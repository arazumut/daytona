// Daytona Platforms Inc. 2024
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"fmt"
	"time"

	"github.com/daytonaio/daytona/pkg/telemetry"
	log "github.com/sirupsen/logrus"
)

func (s *Server) Temizle(ctx context.Context, zorla bool) []error {
	log.SetLevel(log.PanicLevel)

	telemetriAktif := telemetry.TelemetryEnabled(ctx)
	telemetriOzellikleri := map[string]interface{}{
		"zorla":     zorla,
		"sunucu_id": s.Id,
	}

	if telemetriAktif {
		err := s.TelemetryService.TrackServerEvent(telemetry.ServerEventPurgeStarted, telemetry.ClientId(ctx), telemetriOzellikleri)
		if err != nil {
			log.Trace(err)
		}
	}

	fmt.Println("Tüm çalışma alanları siliniyor...")

	err := server.Start()
	if err != nil {
		s.temizlemeHatasiniIzle(ctx, zorla, err)
		return []error{err}
	}

	calismaAlanlari, err := s.WorkspaceService.ListWorkspaces(ctx, false)
	if err != nil {
		s.temizlemeHatasiniIzle(ctx, zorla, err)
		if !zorla {
			return []error{err}
		}
	}

	if err == nil {
		for _, calismaAlani := range calismaAlanlari {
			err := s.WorkspaceService.RemoveWorkspace(ctx, calismaAlani.Id)
			if err != nil {
				s.temizlemeHatasiniIzle(ctx, zorla, err)
				if !zorla {
					return []error{err}
				} else {
					fmt.Printf("%s silinemedi: %v\n", calismaAlani.Name, err)
				}
			} else {
				fmt.Printf("Çalışma alanı %s silindi\n", calismaAlani.Name)
			}
		}
	} else {
		fmt.Printf("Çalışma alanları listelenemedi: %v\n", err)
	}

	fmt.Println("Sağlayıcılar temizleniyor...")
	err = s.ProviderManager.Purge()
	if err != nil {
		s.temizlemeHatasiniIzle(ctx, zorla, err)
		if !zorla {
			return []error{err}
		} else {
			fmt.Printf("Sağlayıcılar temizlenemedi: %v\n", err)
		}
	}

	fmt.Println("Yapılar temizleniyor...")
	hatalar := s.BuildService.MarkForDeletion(nil, zorla)
	if len(hatalar) > 0 {
		s.temizlemeHatasiniIzle(ctx, zorla, hatalar[0])
		if !zorla {
			return hatalar
		} else {
			fmt.Printf("Yapılar silinmek üzere işaretlenemedi: %v\n", hatalar[0])
		}
	}

	err = s.BuildService.AwaitEmptyList(time.Minute)
	if err != nil {
		s.temizlemeHatasiniIzle(ctx, zorla, err)
		if !zorla {
			return []error{err}
		} else {
			fmt.Printf("Boş yapı listesi beklenemedi: %v\n", err)
		}
	}

	if telemetriAktif {
		err := s.TelemetryService.TrackServerEvent(telemetry.ServerEventPurgeCompleted, telemetry.ClientId(ctx), telemetriOzellikleri)
		if err != nil {
			log.Trace(err)
		}
	}

	return nil
}

func (s *Server) temizlemeHatasiniIzle(ctx context.Context, zorla bool, err error) {
	telemetriAktif := telemetry.TelemetryEnabled(ctx)
	telemetriOzellikleri := map[string]interface{}{
		"sunucu_id": s.Id,
		"zorla":     zorla,
		"hata":      err.Error(),
	}

	if telemetriAktif {
		err := s.TelemetryService.TrackServerEvent(telemetry.ServerEventPurgeError, telemetry.ClientId(ctx), telemetriOzellikleri)
		if err != nil {
			log.Trace(err)
		}
	}
}
