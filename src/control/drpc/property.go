//
// (C) Copyright 2019-2022 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//
package drpc

import "unsafe"
import "math"

/*
#include <daos_prop.h>
#include <daos_pool.h>
#include <daos/object.h>
#include <daos/cont_props.h>
*/
import "C"

const (
	// MaxLabelLength is the maximum length of a label.
	MaxLabelLength = C.DAOS_PROP_LABEL_MAX_LEN
)

const (
	// PoolPropertyLabel is a string that a user can associate with a pool.
	PoolPropertyLabel = C.DAOS_PROP_PO_LABEL
	// PoolPropertyACL is the Access Control List for a pool.
	PoolPropertyACL = C.DAOS_PROP_PO_ACL
	// PoolPropertyReservedSpace is the ratio of space that can be reserved
	// on each target for rebuild purposes.
	PoolPropertyReservedSpace = C.DAOS_PROP_PO_SPACE_RB
	// PoolPropertySelfHealing defines the self-healing behavior of the pool.
	PoolPropertySelfHealing = C.DAOS_PROP_PO_SELF_HEAL
	// PoolPropertySpaceReclaim defines the free space reclamation behavior of the pool.
	PoolPropertySpaceReclaim = C.DAOS_PROP_PO_RECLAIM
	// PoolPropertyOwner is the user who acts as the owner of the pool.
	PoolPropertyOwner = C.DAOS_PROP_PO_OWNER
	// PoolPropertyOwnerGroup is the group that acts as the owner of the pool.
	PoolPropertyOwnerGroup = C.DAOS_PROP_PO_OWNER_GROUP
	// PoolPropertyECCellSize is the EC Cell size.
	PoolPropertyECCellSize = C.DAOS_PROP_PO_EC_CELL_SZ
	// PoolPropertyRedunFac
	PoolPropertyRedunFac = C.DAOS_PROP_PO_REDUN_FAC
	//PoolPropertyECPda is performance domain affinity level of EC object.
	PoolPropertyECPda = C.DAOS_PROP_PO_EC_PDA
	//PoolPropertyRPPda is performance domain affinity level of replicated object.
	PoolPropertyRPPda = C.DAOS_PROP_PO_RP_PDA
	//PoolPropertyPerfDomain is Performance domain
	PoolPropertyPerfDomain = C.DAOS_PROP_PO_PERF_DOMAIN
)

const (
	// PoolSpaceReclaimDisabled sets the PoolPropertySpaceReclaim property to disabled.
	PoolSpaceReclaimDisabled = C.DAOS_RECLAIM_DISABLED
	// PoolSpaceReclaimLazy sets the PoolPropertySpaceReclaim property to lazy.
	PoolSpaceReclaimLazy = C.DAOS_RECLAIM_LAZY
	// PoolSpaceReclaimSnapshot sets the PoolPropertySpaceReclaim property to snapshot.
	PoolSpaceReclaimSnapshot = C.DAOS_RECLAIM_SNAPSHOT
	// PoolSpaceReclaimBatch sets the PoolPropertySpaceReclaim property to batch.
	PoolSpaceReclaimBatch = C.DAOS_RECLAIM_BATCH
	// PoolSpaceReclaimTime sets the PoolPropertySpaceReclaim property to time.
	PoolSpaceReclaimTime = C.DAOS_RECLAIM_TIME
)

const (
	// PoolSelfHealingAutoExclude sets the self-healing strategy to auto-exclude.
	PoolSelfHealingAutoExclude = C.DAOS_SELF_HEAL_AUTO_EXCLUDE
	// PoolSelfHealingAutoRebuild sets the self-healing strategy to auto-rebuild.
	PoolSelfHealingAutoRebuild = C.DAOS_SELF_HEAL_AUTO_REBUILD
)

const (
	// MediaTypeScm is the media type for SCM.
	MediaTypeScm = C.DAOS_MEDIA_SCM
	// MediaTypeNvme is the media type for NVMe.
	MediaTypeNvme = C.DAOS_MEDIA_NVME
)

// LabelIsValid checks a label to verify that it meets length/content
// requirements.
func LabelIsValid(label string) bool {
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))

	return bool(C.daos_label_is_valid(cLabel))
}

func EcCellSizeIsValid(sz uint64) bool {
	if sz > math.MaxUint32 {
		return false
	}
	return bool(C.daos_ec_cs_valid(C.uint32_t(sz)))
}

func EcPdaIsValid(pda uint64) bool {
	if pda > math.MaxUint32 {
		return false
	}
	return bool(C.daos_ec_pda_valid(C.uint32_t(pda)))
}

func RpPdaIsValid(pda uint64) bool {
	if pda > math.MaxUint32 {
		return false
	}
	return bool(C.daos_rp_pda_valid(C.uint32_t(pda)))
}

func PerfDomainIsValid(perfdomain string) bool {
	cPerfDomain := C.CString(perfdomain)
	defer C.free(unsafe.Pointer(cPerfDomain))

	return bool(C.daos_perf_domain_is_valid(cPerfDomain))
}
