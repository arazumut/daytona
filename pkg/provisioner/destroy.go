// 2024 Daytona Platforms Inc. Tüm hakları saklıdır.
// SPDX-License-Identifier: Apache-2.0

package provisioner

import (
	"github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/workspace"
	"github.com/daytonaio/daytona/pkg/workspace/project"
)

// DestroyWorkspace, belirtilen çalışma alanını yok eder.
func (p *Provisioner) DestroyWorkspace(workspace *workspace.Workspace, target *provider.ProviderTarget) error {
	// Hedef sağlayıcıyı al
	targetProvider, err := p.providerManager.GetProvider(target.ProviderInfo.Name)
	if err != nil {
		return err
	}

	// Çalışma alanını yok et
	_, err = (*targetProvider).DestroyWorkspace(&provider.WorkspaceRequest{
		TargetOptions: target.Options,
		Workspace:     workspace,
	})

	return err
}

// DestroyProject, belirtilen projeyi yok eder.
func (p *Provisioner) DestroyProject(proj *project.Project, target *provider.ProviderTarget) error {
	// Hedef sağlayıcıyı al
	targetProvider, err := p.providerManager.GetProvider(target.ProviderInfo.Name)
	if err != nil {
		return err
	}

	// Projeyi yok et
	_, err = (*targetProvider).DestroyProject(&provider.ProjectRequest{
		TargetOptions: target.Options,
		Project:       proj,
	})

	return err
}
