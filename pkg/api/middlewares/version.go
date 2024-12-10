// 2024 Daytona Platforms Inc. Tüm hakları saklıdır.
// SPDX-License-Identifier: Apache-2.0

package middlewares

import (
	"github.com/gin-gonic/gin"
)

const SUNUCU_SURUM_BASLIGI = "X-Server-Version"

func SurumOrtaKatmani(surum string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add(SUNUCU_SURUM_BASLIGI, surum)
	}
}
