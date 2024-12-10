// 2024 Daytona Platforms Inc. Telif Hakkı
// SPDX-License-Identifier: Apache-2.0

package log

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const ZAMAN_ASIMI = 300 * time.Millisecond

var yükseltici = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsYaz(ws *websocket.Conn, c chan []byte, hataKanalı chan error) {
	for {
		hata := ws.WriteMessage(websocket.TextMessage, <-c)
		if hata != nil {
			hataKanalı <- hata
			break
		}
	}
}

func wsJSONYaz(ws *websocket.Conn, c chan interface{}, hataKanalı chan error) {
	for {
		hata := ws.WriteJSON(<-c)
		if hata != nil {
			hataKanalı <- hata
			break
		}
	}
}

// logOku, logReader'dan okur ve websocket'e yazar.
// T, logReader'dan okunacak mesajın türüdür
func logOku[T any](ginCtx *gin.Context, logReader io.Reader, okumaFonksiyonu func(context.Context, io.Reader, bool, chan T, chan error), wsYazmaFonksiyonu func(*websocket.Conn, chan T, chan error)) {
	takipSorgusu := ginCtx.Query("follow")
	takip := takipSorgusu == "true"

	ws, hata := yükseltici.Upgrade(ginCtx.Writer, ginCtx.Request, nil)
	if hata != nil {
		log.Error(hata)
		return
	}

	defer func() {
		kapatmaHatası := websocket.CloseNormalClosure
		if !errors.Is(hata, io.EOF) {
			kapatmaHatası = websocket.CloseInternalServerErr
		}
		hata := ws.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(kapatmaHatası, ""), time.Now().Add(time.Second))
		if hata != nil {
			log.Trace(hata)
		}
		ws.Close()
	}()

	mesajKanalı := make(chan T)
	hataKanalı := make(chan error)
	ctx, iptal := context.WithCancel(ginCtx.Request.Context())

	defer iptal()
	go okumaFonksiyonu(ctx, logReader, takip, mesajKanalı, hataKanalı)
	go wsYazmaFonksiyonu(ws, mesajKanalı, hataKanalı)

	okumaHatası := make(chan error)
	go func() {
		for {
			_, _, hata := ws.ReadMessage()
			okumaHatası <- hata
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case hata = <-hataKanalı:
			if hata != nil {
				if !errors.Is(hata, io.EOF) {
					log.Error(hata)
				}
				iptal()
				return
			}
		case hata := <-okumaHatası:
			if websocket.IsUnexpectedCloseError(hata, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
				log.Error(hata)
			}
			if hata != nil {
				return
			}
		}
	}
}

func SunucuLoguOku(ginCtx *gin.Context) {
	sunucu := server.GetInstance(nil)
	tekrarDeneSorgusu := ginCtx.DefaultQuery("retry", "true")
	tekrarDene := tekrarDeneSorgusu == "true"

	if tekrarDene {
		for {
			okuyucu, hata := sunucu.GetLogReader()
			if hata == nil {
				logOku(ginCtx, okuyucu, util.ReadLog, wsYaz)
				return
			}
			time.Sleep(ZAMAN_ASIMI)
		}
	}

	okuyucu, hata := sunucu.GetLogReader()
	if hata != nil {
		ginCtx.AbortWithError(http.StatusInternalServerError, hata)
		return
	}

	logOku(ginCtx, okuyucu, util.ReadLog, wsYaz)
}

func ÇalışmaAlanıLoguOku(ginCtx *gin.Context) {
	çalışmaAlanıId := ginCtx.Param("workspaceId")
	tekrarDeneSorgusu := ginCtx.DefaultQuery("retry", "true")
	tekrarDene := tekrarDeneSorgusu == "true"

	sunucu := server.GetInstance(nil)

	if tekrarDene {
		for {
			wsLogOkuyucu, hata := sunucu.WorkspaceService.GetWorkspaceLogReader(çalışmaAlanıId)
			if hata == nil {
				logOku(ginCtx, wsLogOkuyucu, util.ReadJSONLog, wsJSONYaz)
				return
			}
			time.Sleep(ZAMAN_ASIMI)
		}
	}

	wsLogOkuyucu, hata := sunucu.WorkspaceService.GetWorkspaceLogReader(çalışmaAlanıId)
	if hata != nil {
		ginCtx.AbortWithError(http.StatusInternalServerError, hata)
		return
	}

	logOku(ginCtx, wsLogOkuyucu, util.ReadJSONLog, wsJSONYaz)
}

func ProjeLoguOku(ginCtx *gin.Context) {
	çalışmaAlanıId := ginCtx.Param("workspaceId")
	projeAdı := ginCtx.Param("projectName")
	tekrarDeneSorgusu := ginCtx.DefaultQuery("retry", "true")
	tekrarDene := tekrarDeneSorgusu == "true"

	sunucu := server.GetInstance(nil)

	if tekrarDene {
		for {
			projeLogOkuyucu, hata := sunucu.WorkspaceService.GetProjectLogReader(çalışmaAlanıId, projeAdı)
			if hata == nil {
				logOku(ginCtx, projeLogOkuyucu, util.ReadJSONLog, wsJSONYaz)
				return
			}
			time.Sleep(ZAMAN_ASIMI)
		}
	}

	projeLogOkuyucu, hata := sunucu.WorkspaceService.GetProjectLogReader(çalışmaAlanıId, projeAdı)
	if hata != nil {
		ginCtx.AbortWithError(http.StatusInternalServerError, hata)
		return
	}

	logOku(ginCtx, projeLogOkuyucu, util.ReadJSONLog, wsJSONYaz)
}

func YapıLoguOku(ginCtx *gin.Context) {
	yapıId := ginCtx.Param("buildId")
	tekrarDeneSorgusu := ginCtx.DefaultQuery("retry", "true")
	tekrarDene := tekrarDeneSorgusu == "true"

	sunucu := server.GetInstance(nil)

	if tekrarDene {
		for {
			yapıLogOkuyucu, hata := sunucu.BuildService.GetBuildLogReader(yapıId)

			if hata == nil {
				logOku(ginCtx, yapıLogOkuyucu, util.ReadJSONLog, wsJSONYaz)
				return
			}
			time.Sleep(ZAMAN_ASIMI)
		}
	}

	yapıLogOkuyucu, hata := sunucu.BuildService.GetBuildLogReader(yapıId)
	if hata != nil {
		ginCtx.AbortWithError(http.StatusInternalServerError, hata)
		return
	}

	logOku(ginCtx, yapıLogOkuyucu, util.ReadJSONLog, wsJSONYaz)
}
