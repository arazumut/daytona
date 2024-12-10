// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"net/rpc"

	"github.com/daytonaio/daytona/pkg/provider/util"
	"github.com/daytonaio/daytona/pkg/workspace"
	"github.com/daytonaio/daytona/pkg/workspace/project"
	"github.com/hashicorp/go-plugin"
)

// Sağlayıcı arayüzü
type Provider interface {
	Initialize(InitializeProviderRequest) (*util.Empty, error)
	GetInfo() (ProviderInfo, error)

	GetTargetManifest() (*ProviderTargetManifest, error)
	GetPresetTargets() (*[]ProviderTarget, error)

	CreateWorkspace(*WorkspaceRequest) (*util.Empty, error)
	StartWorkspace(*WorkspaceRequest) (*util.Empty, error)
	StopWorkspace(*WorkspaceRequest) (*util.Empty, error)
	DestroyWorkspace(*WorkspaceRequest) (*util.Empty, error)
	GetWorkspaceInfo(*WorkspaceRequest) (*workspace.WorkspaceInfo, error)

	CreateProject(*ProjectRequest) (*util.Empty, error)
	StartProject(*ProjectRequest) (*util.Empty, error)
	StopProject(*ProjectRequest) (*util.Empty, error)
	DestroyProject(*ProjectRequest) (*util.Empty, error)
	GetProjectInfo(*ProjectRequest) (*project.ProjectInfo, error)
}

// Sağlayıcı Eklentisi
type ProviderPlugin struct {
	Impl Provider
}

// Sağlayıcı RPC Sunucusu
func (p *ProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ProviderRPCServer{Impl: p.Impl}, nil
}

// Sağlayıcı RPC İstemcisi
func (p *ProviderPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ProviderRPCClient{client: c}, nil
}
