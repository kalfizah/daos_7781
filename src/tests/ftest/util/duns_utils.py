#!/usr/bin/env python3
"""
  (C) Copyright 2018-2022 Intel Corporation.

  SPDX-License-Identifier: BSD-2-Clause-Patent
"""

def uns_path(pool, cont, path=None):
    """Format a DAOS UNS path of the form daos://pool/cont/[path]

    Args:
        pool (str/obj): pool string or object containing identifier, label, or uuid.
        cont (str/obj): cont string or object containing identifier, label, or uuid.
        path (str, optional): path relative to container root

    Return:
        str: the DAOS UNS path
    """
    def _get_id(obj):
        for attr_name in ('identifier', 'label', 'uuid'):
            try:
                return getattr(obj, attr_name)
            except AttributeError:
                pass
        return obj

    return '/'.join([
        'daos:/',
        _get_id(pool),
        _get_id(cont),
        path.lstrip('/') if path else ""])
