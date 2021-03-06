/**
 * (C) Copyright 2016-2022 Intel Corporation.
 *
 * SPDX-License-Identifier: BSD-2-Clause-Patent
 */

#include "dfuse_common.h"
#include "dfuse.h"

#define READAHEAD_SIZE (1024 * 1024)

static void
dfuse_cb_read_complete(struct dfuse_event *ev)
{
	if (ev->de_ev.ev_error == 0) {
		size_t len = ev->de_iov.iov_buf_len;

		if (ev->de_len == 0)
			DFUSE_TRA_DEBUG(ev, "Truncated read, (EOF)");
		else if (ev->de_len != len)
			DFUSE_TRA_DEBUG(ev,
					"Truncated read, requested %#zx returned %#zx",
					len, ev->de_len);

		DFUSE_REPLY_BUF(ev, ev->de_req, ev->de_iov.iov_buf, ev->de_len);
	} else {
		DFUSE_REPLY_ERR_RAW(ev, ev->de_req, ev->de_ev.ev_error);
	}
	D_FREE(ev->de_iov.iov_buf);
}

void
dfuse_cb_read(fuse_req_t req, fuse_ino_t ino, size_t len, off_t position,
	      struct fuse_file_info *fi)
{
	struct dfuse_obj_hdl		*oh = (struct dfuse_obj_hdl *)fi->fh;
	struct dfuse_projection_info	*fs_handle = fuse_req_userdata(req);
	const struct fuse_ctx		*fc = fuse_req_ctx(req);
	struct fuse_bufvec		fb = {};
	void				*buff;
	int				rc;
	size_t				buff_len = len;
	bool				skip_read = false;
	bool				readahead = false;
	bool				async = false;
	struct dfuse_event		*ev = NULL;

	D_ALLOC_PTR(ev);
	if (ev == NULL)
		D_GOTO(err, rc = ENOMEM);

	DFUSE_TRA_UP(ev, oh, "event");

	DFUSE_TRA_INFO(ev, "%#zx-%#zx requested pid=%d",
		       position, position + len - 1, fc->pid);

	if (oh->doh_ie->ie_truncated &&
	    position + len < oh->doh_ie->ie_stat.st_size &&
		((oh->doh_ie->ie_start_off == 0 &&
			oh->doh_ie->ie_end_off == 0) ||
			position >= oh->doh_ie->ie_end_off ||
			position + len <= oh->doh_ie->ie_start_off)) {
		off_t pos_ra = position + len + READAHEAD_SIZE;

		DFUSE_TRA_DEBUG(oh, "Returning zeros");
		skip_read = true;

		if (pos_ra <= oh->doh_ie->ie_stat.st_size &&
		    ((oh->doh_ie->ie_start_off == 0 &&
				oh->doh_ie->ie_end_off == 0) ||
				(position >= oh->doh_ie->ie_end_off ||
					pos_ra <= oh->doh_ie->ie_start_off))) {
			readahead = true;
		}
	} else if (oh->doh_caching &&
		len < (1024 * 1024) &&
		oh->doh_ie->ie_stat.st_size > (1024 * 1024)) {
		/* Only do readahead if the requested size is less than 1Mb and
		 * the file size is > 1Mb
		 */

		readahead = true;
	}

	if (readahead) {
		if (!fs_handle->dpi_info->di_wb_cache)
			buff_len += READAHEAD_SIZE;
	} else {
		if (!skip_read) {
			rc = daos_event_init(&ev->de_ev,
					     fs_handle->dpi_eq, NULL);
			if (rc != -DER_SUCCESS)
				D_GOTO(err, rc = daos_der2errno(rc));

			ev->de_req = req;
			ev->de_complete_cb = dfuse_cb_read_complete;
			async = true;
		}
	}

	D_ALLOC(buff, buff_len);
	if (!buff)
		D_GOTO(err, rc = ENOMEM);

	ev->de_sgl.sg_nr = 1;
	d_iov_set(&ev->de_iov, (void *)buff, buff_len);
	ev->de_sgl.sg_iovs = &ev->de_iov;

	if (skip_read) {
		ev->de_len = buff_len;
	} else {
		rc = dfs_read(oh->doh_dfs, oh->doh_obj, &ev->de_sgl, position,
			      &ev->de_len, async ? &ev->de_ev : NULL);
		if (rc != -DER_SUCCESS) {
			DFUSE_REPLY_ERR_RAW(oh, req, rc);
			D_FREE(buff);
			D_FREE(ev);
			return;
		}
	}

	if (async) {
		/* Send a message to the async thread to wake it up and poll
		 * for events
		 */
		sem_post(&fs_handle->dpi_sem);
		return;
	}

	if (ev->de_len <= len) {
		if (ev->de_len == 0)
			DFUSE_TRA_DEBUG(oh, "Truncated read, %#zx-%#zx (EOF)",
					position, position);
		else if (ev->de_len != len)
			DFUSE_TRA_DEBUG(oh, "Truncated read, %#zx-%#zx",
					position, position + ev->de_len - 1);
		DFUSE_REPLY_BUF(oh, req, buff, ev->de_len);
		D_FREE(buff);
		D_FREE(ev);
		return;
	}

	rc = pthread_mutex_trylock(&oh->doh_ie->ie_dfs->dfs_read_mutex);
	if (rc == 0) {
		fb.count = 1;
		fb.buf[0].mem = buff + len;
		fb.buf[0].size = ev->de_len - len;

		DFUSE_TRA_INFO(oh, "%#zx-%#zx was readahead",
			       position + len, position + ev->de_len - 1);

		rc = fuse_lowlevel_notify_store(fs_handle->dpi_info->di_session,
						ino, position + len, &fb, 0);
		if (rc == 0)
			DFUSE_TRA_DEBUG(oh, "notify_store returned %d", rc);
		else
			DFUSE_TRA_INFO(oh, "notify_store returned %d", rc);
		rc = pthread_mutex_unlock(&oh->doh_ie->ie_dfs->dfs_read_mutex);
		if (rc != 0)
			DFUSE_TRA_ERROR(oh, "Mutex unlock failed");
	}

	DFUSE_REPLY_BUF(oh, req, buff, len);
	D_FREE(buff);
	D_FREE(ev);
	return;

err:
	DFUSE_REPLY_ERR_RAW(oh, req, rc);
	D_FREE(ev);
}
