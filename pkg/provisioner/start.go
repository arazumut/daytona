// 2024 Daytona Platforms Inc. Telif Hakkı
// SPDX-License-Identifier: Apache-2.0

package provisioner

import (
	"github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/workspace"
	"github.com/daytonaio/daytona/pkg/workspace/project"
)

// Workspace'i başlatır
func (p *Provisioner) StartWorkspace(workspace *workspace.Workspace, target *provider.ProviderTarget) error {
	targetProvider, err := p.providerManager.GetProvider(target.ProviderInfo.Name)
	if err != nil {
		return err
	}

	_, err = (*targetProvider).StartWorkspace(&provider.WorkspaceRequest{
		TargetOptions: target.Options,
		Workspace:     workspace,
	})

	return err
}

// Projeyi başlatır
func (p *Provisioner) StartProject(proj *project.Project, target *provider.ProviderTarget) error {
	targetProvider, err := p.providerManager.GetProvider(target.ProviderInfo.Name)
	if err != nil {
		return err
	}

	_, err = (*targetProvider).StartProject(&provider.ProjectRequest{
		TargetOptions: target.Options,
		Project:       proj,
	})

	return err
}
