//
// (C) Copyright 2022 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

package ucx

import (
	"context"

	"github.com/daos-stack/daos/src/control/common"
	"github.com/daos-stack/daos/src/control/lib/dlopen"
	"github.com/daos-stack/daos/src/control/lib/hardware"
	"github.com/daos-stack/daos/src/control/logging"
)

// NewProvider creates a new UCX data provider.
func NewProvider(log logging.Logger) *Provider {
	return &Provider{
		log: log,
	}
}

// Provider provides information from UCX's API.
type Provider struct {
	log logging.Logger
}

// GetFabricInterfaces harvests the collection of fabric interfaces from UCX.
func (p *Provider) GetFabricInterfaces(ctx context.Context) (*hardware.FabricInterfaceSet, error) {
	uctHdl, err := openUCT()
	if err != nil {
		return nil, err
	}
	defer uctHdl.Close()

	components, cleanupComp, err := getUCTComponents(uctHdl)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cleanupComp(); err != nil {
			p.log.Errorf("error cleaning up UCT components: %s", err.Error())
		}
	}()

	fis := hardware.NewFabricInterfaceSet()

	for _, comp := range components {
		// TBD: Should we limit components by what we presume to support? I.e. "tcp", "ib"

		mdResources, err := getMDResourceNames(uctHdl, comp)
		if err != nil {
			p.log.Error(err.Error())
			continue
		}

		cfg, cleanupCfg, err := getComponentMDConfig(uctHdl, comp)
		if err != nil {
			p.log.Error(err.Error())
			continue
		}
		defer func(name string) {
			if err := cleanupCfg(); err != nil {
				p.log.Errorf(err.Error())
			}
		}(comp.name)

		for _, mdName := range mdResources {
			if err := p.addFabricDevicesFromMD(uctHdl, comp, mdName, cfg, fis); err != nil {
				p.log.Errorf(err.Error())
			}
		}
	}

	return fis, nil
}

func (p *Provider) addFabricDevicesFromMD(uctHdl *dlopen.LibHandle, comp *uctComponent,
	mdName string, cfg *uctMDConfig, fis *hardware.FabricInterfaceSet) error {
	md, cleanupMD, err := openMDResource(uctHdl, comp, mdName, cfg)
	if err != nil {
		return err
	}
	defer func() {
		if err := cleanupMD(); err != nil {
			p.log.Errorf(err.Error())
		}
	}()

	tlDevs, err := getMDTransportDevices(uctHdl, md)
	if err != nil {
		return err
	}

	for _, dev := range tlDevs {
		if !dev.isNetwork() {
			continue
		}

		fis.Update(&hardware.FabricInterface{
			Name:      dev.device,
			Providers: common.NewStringSet("ucx+" + dev.transport),
		})
	}

	return nil
}
