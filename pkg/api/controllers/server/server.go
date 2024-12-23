// Daytona Platforms Inc. 2024
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// GetConfig 			godoc
//
//	@Tags			server
//	@Summary		Sunucu yapılandırmasını al
//	@Description	Sunucu yapılandırmasını al
//	@Produce		json
//	@Success		200	{object}	ServerConfig
//	@Router			/server/config [get]
//
//	@id				GetConfig
func GetConfig(ctx *gin.Context) {
	config, err := server.GetConfig()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("yapılandırma alınamadı: %w", err))
		return
	}

	ctx.JSON(200, config)
}

// SetConfig 			godoc
//
//	@Tags			server
//	@Summary		Sunucu yapılandırmasını ayarla
//	@Description	Sunucu yapılandırmasını ayarla
//	@Accept			json
//	@Produce		json
//	@Param			config	body		ServerConfig	true	"Sunucu yapılandırması"
//	@Success		200		{object}	ServerConfig
//	@Router			/server/config [post]
//
//	@id				SetConfig
func SetConfig(ctx *gin.Context) {
	var c server.Config
	err := ctx.BindJSON(&c)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("geçersiz istek gövdesi: %w", err))
		return
	}

	err = server.Save(c)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("yapılandırma kaydedilemedi: %w", err))
		return
	}

	ctx.JSON(200, c)
}

// GenerateNetworkKey 		godoc
//
//	@Tags			server
//	@Summary		Yeni bir kimlik doğrulama anahtarı oluştur
//	@Description	Yeni bir kimlik doğrulama anahtarı oluştur
//	@Produce		json
//	@Success		200	{object}	NetworkKey
//	@Router			/server/network-key [post]
//
//	@id				GenerateNetworkKey
func GenerateNetworkKey(ctx *gin.Context) {
	s := server.GetInstance(nil)

	authKey, err := s.TailscaleServer.CreateAuthKey()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("ağ anahtarı oluşturulamadı: %w", err))
		return
	}

	ctx.JSON(200, &server.NetworkKey{Key: authKey})
}
