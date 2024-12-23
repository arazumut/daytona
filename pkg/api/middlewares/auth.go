// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package middlewares

import (
	"errors"
	"strings"

	"github.com/daytonaio/daytona/pkg/apikey"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware, gelen isteklerin yetkilendirilmesini sağlar.
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			ctx.AbortWithError(401, errors.New("yetkisiz"))
			return
		}

		token := ExtractToken(bearerToken)
		if token == "" {
			ctx.AbortWithError(401, errors.New("yetkisiz"))
			return
		}

		server := server.GetInstance(nil)

		if !server.ApiKeyService.IsValidApiKey(token) {
			ctx.AbortWithError(401, errors.New("yetkisiz"))
			return
		}

		apiKeyType := apikey.ApiKeyTypeClient

		if server.ApiKeyService.IsWorkspaceApiKey(token) {
			apiKeyType = apikey.ApiKeyTypeWorkspace
		} else if server.ApiKeyService.IsProjectApiKey(token) {
			apiKeyType = apikey.ApiKeyTypeProject
		}

		ctx.Set("apiKeyType", apiKeyType)
		ctx.Next()
	}
}

// ExtractToken, Bearer token'ını Authorization başlığından çıkarır.
func ExtractToken(bearerToken string) string {
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return ""
	}

	return strings.TrimPrefix(bearerToken, "Bearer ")
}
