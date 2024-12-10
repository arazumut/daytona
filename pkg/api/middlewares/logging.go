// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

// LoggingMiddleware, API isteklerini ve yanıtlarını loglamak için bir ara katman fonksiyonudur.
func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()               // İstek başlangıç zamanını al
		ctx.Next()                            // İstek işlenmeye devam etsin
		endTime := time.Now()                 // İstek bitiş zamanını al
		latencyTime := endTime.Sub(startTime) // İstek süresini hesapla
		reqMethod := ctx.Request.Method       // İstek metodunu al (GET, POST, vb.)
		reqUri := ctx.Request.RequestURI      // İstek URI'sini al
		statusCode := ctx.Writer.Status()     // Yanıt durum kodunu al

		if len(ctx.Errors) > 0 { // Eğer hata varsa
			log.WithFields(log.Fields{
				"method":  reqMethod,
				"URI":     reqUri,
				"status":  statusCode,
				"latency": latencyTime,
				"error":   ctx.Errors.String(),
			}).Error("API HATASI") // Hata mesajını logla
			ctx.JSON(statusCode, gin.H{"error": ctx.Errors[0].Err.Error()}) // Hata mesajını JSON olarak döndür
		} else { // Eğer hata yoksa
			log.WithFields(log.Fields{
				"method":  reqMethod,
				"URI":     reqUri,
				"status":  statusCode,
				"latency": latencyTime,
			}).Info("API İSTEĞİ") // İstek bilgilerini logla
		}

		ctx.Next() // İstek işlenmeye devam etsin
	}
}
