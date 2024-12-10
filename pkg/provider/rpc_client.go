// 2024 Daytona Platforms Inc. Telif Hakkı
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"net/rpc"

	"github.com/daytonaio/daytona/pkg/provider/util"
	"github.com/daytonaio/daytona/pkg/workspace"
	"github.com/daytonaio/daytona/pkg/workspace/project"
)

// ProviderRPCClient, RPC istemcisi ile iletişim kurmak için kullanılır
type ProviderRPCClient struct {
	client *rpc.Client
}

// Initialize, sağlayıcıyı başlatır
func (m *ProviderRPCClient) Initialize(req InitializeProviderRequest) (*util.Empty, error) {
	err := m.client.Call("Plugin.Initialize", &req, new(util.Empty))
	return new(util.Empty), err
}

// GetInfo, sağlayıcı bilgilerini alır
func (m *ProviderRPCClient) GetInfo() (ProviderInfo, error) {
	var resp ProviderInfo
	err := m.client.Call("Plugin.GetInfo", new(interface{}), &resp)
	return resp, err
}

// GetTargetManifest, hedef manifestosunu alır
func (m *ProviderRPCClient) GetTargetManifest() (*ProviderTargetManifest, error) {
	var resp ProviderTargetManifest
	err := m.client.Call("Plugin.GetTargetManifest", new(interface{}), &resp)
	return &resp, err
}

// GetPresetTargets, ön ayarlı hedefleri alır
func (m *ProviderRPCClient) GetPresetTargets() (*[]ProviderTarget, error) {
	var resp []ProviderTarget
	err := m.client.Call("Plugin.GetPresetTargets", new(interface{}), &resp)
	return &resp, err
}

// CreateWorkspace, bir çalışma alanı oluşturur
func (m *ProviderRPCClient) CreateWorkspace(workspaceReq *WorkspaceRequest) (*util.Empty, error) {
	err := m.client.Call("Plugin.CreateWorkspace", workspaceReq, new(util.Empty))
	return new(util.Empty), err
}

// StartWorkspace, bir çalışma alanını başlatır
func (m *ProviderRPCClient) StartWorkspace(workspaceReq *WorkspaceRequest) (*util.Empty, error) {
	err := m.client.Call("Plugin.StartWorkspace", workspaceReq, new(util.Empty))
	return new(util.Empty), err
}

// StopWorkspace, bir çalışma alanını durdurur
func (m *ProviderRPCClient) StopWorkspace(workspaceReq *WorkspaceRequest) (*util.Empty, error) {
	err := m.client.Call("Plugin.StopWorkspace", workspaceReq, new(util.Empty))
	return new(util.Empty), err
}

// DestroyWorkspace, bir çalışma alanını yok eder
func (m *ProviderRPCClient) DestroyWorkspace(workspaceReq *WorkspaceRequest) (*util.Empty, error) {
	err := m.client.Call("Plugin.DestroyWorkspace", workspaceReq, new(util.Empty))
	return new(util.Empty), err
}

// GetWorkspaceInfo, çalışma alanı bilgilerini alır
func (m *ProviderRPCClient) GetWorkspaceInfo(workspaceReq *WorkspaceRequest) (*workspace.WorkspaceInfo, error) {
	var response workspace.WorkspaceInfo
	err := m.client.Call("Plugin.GetWorkspaceInfo", workspaceReq, &response)
	return &response, err
}

// CreateProject, bir proje oluşturur
func (m *ProviderRPCClient) CreateProject(projectReq *ProjectRequest) (*util.Empty, error) {
	err := m.client.Call("Plugin.CreateProject", projectReq, new(util.Empty))
	return new(util.Empty), err
}

// StartProject, bir projeyi başlatır
func (m *ProviderRPCClient) StartProject(projectReq *ProjectRequest) (*util.Empty, error) {
	err := m.client.Call("Plugin.StartProject", projectReq, new(util.Empty))
	return new(util.Empty), err
}

// StopProject, bir projeyi durdurur
func (m *ProviderRPCClient) StopProject(projectReq *ProjectRequest) (*util.Empty, error) {
	err := m.client.Call("Plugin.StopProject", projectReq, new(util.Empty))
	return new(util.Empty), err
}

// DestroyProject, bir projeyi yok eder
func (m *ProviderRPCClient) DestroyProject(projectReq *ProjectRequest) (*util.Empty, error) {
	err := m.client.Call("Plugin.DestroyProject", projectReq, new(util.Empty))
	return new(util.Empty), err
}

// GetProjectInfo, proje bilgilerini alır
func (m *ProviderRPCClient) GetProjectInfo(projectReq *ProjectRequest) (*project.ProjectInfo, error) {
	var resp project.ProjectInfo
	err := m.client.Call("Plugin.GetProjectInfo", projectReq, &resp)
	return &resp, err
}
