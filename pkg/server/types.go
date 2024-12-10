// 2024 Daytona Platforms Inc. Tüm Hakları Saklıdır.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"net/http"
)

// TailscaleServer arayüzü, Tailscale sunucusuyla ilgili işlemleri tanımlar.
type TailscaleServer interface {
	Connect() error
	CreateAuthKey() (string, error)
	CreateUser() error
	HTTPClient() *http.Client
	Start(errChan chan error) error
	Stop() error
	Purge() error
}

// ILocalContainerRegistry arayüzü, yerel konteyner kayıt defteri işlemlerini tanımlar.
type ILocalContainerRegistry interface {
	Start() error
	Stop() error
	Purge() error
}

// FRPSConfig, FRP sunucusu yapılandırmasını temsil eder.
type FRPSConfig struct {
	Domain   string `json:"domain" validate:"required"`
	Port     uint32 `json:"port" validate:"required"`
	Protocol string `json:"protocol" validate:"required"`
} // @name FRPSConfig

// NetworkKey, ağ anahtarını temsil eder.
type NetworkKey struct {
	Key string `json:"key" validate:"required"`
} // @name NetworkKey

// Config, sunucu yapılandırmasını temsil eder.
type Config struct {
	ProvidersDir              string      `json:"providersDir" validate:"required"`
	RegistryUrl               string      `json:"registryUrl" validate:"required"`
	Id                        string      `json:"id" validate:"required"`
	ServerDownloadUrl         string      `json:"serverDownloadUrl" validate:"required"`
	Frps                      *FRPSConfig `json:"frps,omitempty" validate:"optional"`
	ApiPort                   uint32      `json:"apiPort" validate:"required"`
	HeadscalePort             uint32      `json:"headscalePort" validate:"required"`
	BinariesPath              string      `json:"binariesPath" validate:"required"`
	LogFilePath               string      `json:"logFilePath" validate:"required"`
	DefaultProjectImage       string      `json:"defaultProjectImage" validate:"required"`
	DefaultProjectUser        string      `json:"defaultProjectUser" validate:"required"`
	BuilderImage              string      `json:"builderImage" validate:"required"`
	LocalBuilderRegistryPort  uint32      `json:"localBuilderRegistryPort" validate:"required"`
	LocalBuilderRegistryImage string      `json:"localBuilderRegistryImage" validate:"required"`
	BuilderRegistryServer     string      `json:"builderRegistryServer" validate:"required"`
	BuildImageNamespace       string      `json:"buildImageNamespace" validate:"optional"`
	SamplesIndexUrl           string      `json:"samplesIndexUrl" validate:"optional"`
} // @name ServerConfig
