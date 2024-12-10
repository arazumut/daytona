// Daytona Platforms Inc. 2024
// SPDX-License-Identifier: Apache-2.0

package profiledata

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/profiledata"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// ProfilVerisiniGetir godoc
//
//	@Tags			profil
//	@Summary		Profil verisini getir
//	@Description	Profil verisini getir
//	@Accept			json
//	@Success		200 {object} profiledata.ProfileData
//	@Router			/profil [get]
//
//	@id				ProfilVerisiniGetir
func ProfilVerisiniGetir(ctx *gin.Context) {
	server := server.GetInstance(nil)
	profileData, err := server.ProfileDataService.Get()
	if err != nil {
		if profiledata.IsProfileDataNotFound(err) {
			ctx.JSON(200, &profiledata.ProfileData{})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("profil verisi alınamadı: %w", err))
		return
	}

	ctx.JSON(200, profileData)
}

// ProfilVerisiniAyarla godoc
//
//	@Tags			profil
//	@Summary		Profil verisini ayarla
//	@Description	Profil verisini ayarla
//	@Accept			json
//	@Param			profileData	body	profiledata.ProfileData	true	"Profil verisi"
//	@Success		201
//	@Router			/profil [put]
//
//	@id				ProfilVerisiniAyarla
func ProfilVerisiniAyarla(ctx *gin.Context) {
	var req profiledata.ProfileData
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("geçersiz istek gövdesi: %w", err))
		return
	}

	server := server.GetInstance(nil)
	err = server.ProfileDataService.Save(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("profil verisi kaydedilemedi: %w", err))
		return
	}

	ctx.Status(201)
}

// ProfilVerisiniSil godoc
//
//	@Tags			profil
//	@Summary		Profil verisini sil
//	@Description	Profil verisini sil
//	@Success		204
//	@Router			/profil [delete]
//
//	@id				ProfilVerisiniSil
func ProfilVerisiniSil(ctx *gin.Context) {
	server := server.GetInstance(nil)
	err := server.ProfileDataService.Delete()
	if err != nil {
		if profiledata.IsProfileDataNotFound(err) {
			ctx.Status(204)
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("profil verisi silinemedi: %w", err))
		return
	}

	ctx.Status(204)
}
