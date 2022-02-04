//
// (C) Copyright 2022 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

package ucx

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/daos-stack/daos/src/control/common"
	"github.com/daos-stack/daos/src/control/logging"
)

func TestUCX_Provider_GetFabricInterfaces_Integrated(t *testing.T) {
	cleanup, err := Load()
	if err != nil {
		t.Skipf("ucx not installed (%s)", err.Error())
	}
	defer cleanup()

	// Can't mock the underlying UCX calls, but we can make sure it doesn't crash or
	// error on the normal happy path.

	log, buf := logging.NewTestLogger(t.Name())
	defer common.ShowBufferOnFailure(t, buf)

	p := NewProvider(log)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := p.GetFabricInterfaces(ctx)

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Printf("FabricInterfaceSet:\n%s\n", result)
}
