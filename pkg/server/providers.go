// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/daytonaio/daytona/pkg/provider/manager"
	log "github.com/sirupsen/logrus"
)

func (s *Server) varsayılanSaglayicilariIndir() error {
	manifest, err := s.ProviderManager.GetProvidersManifest()
	if err != nil {
		return err
	}

	varsayılanSaglayicilar := manifest.GetDefaultProviders()

	log.Info("Varsayılan sağlayıcılar indiriliyor")
	for saglayiciAdi, saglayici := range varsayılanSaglayicilar {
		kilitDosyaYolu := filepath.Join(s.config.ProvidersDir, saglayiciAdi, manager.INITIAL_SETUP_LOCK_FILE_NAME)

		_, err := os.Stat(kilitDosyaYolu)
		if err == nil {
			continue
		}

		_, err = s.ProviderManager.DownloadProvider(context.Background(), saglayici.DownloadUrls, saglayiciAdi)
		if err != nil {
			if !manager.IsProviderAlreadyDownloaded(err, saglayiciAdi) {
				log.Error(err)
			}
			continue
		}
	}

	log.Info("Varsayılan sağlayıcılar indirildi")

	return nil
}

func (s *Server) saglayicilariKaydet() error {
	log.Info("Sağlayıcılar kaydediliyor")

	manifest, err := s.ProviderManager.GetProvidersManifest()
	if err != nil {
		return err
	}

	dizinIcerikleri, err := os.ReadDir(s.config.ProvidersDir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("Sağlayıcı bulunamadı")
			return nil
		}
		return err
	}

	for _, icerik := range dizinIcerikleri {
		if icerik.IsDir() {
			saglayiciDizini := filepath.Join(s.config.ProvidersDir, icerik.Name())

			pluginYolu, err := s.getPluginPath(saglayiciDizini)
			if err != nil {
				if !manager.IsNoPluginFound(err, saglayiciDizini) {
					log.Error(err)
				}
				continue
			}

			err = s.ProviderManager.RegisterProvider(pluginYolu, false)
			if err != nil {
				log.Error(err)
				continue
			}

			// İlk kurulum kilidini oluştur
			kilitDosyaYolu := filepath.Join(s.config.ProvidersDir, icerik.Name(), manager.INITIAL_SETUP_LOCK_FILE_NAME)

			_, err = os.Stat(kilitDosyaYolu)
			if err != nil {
				dosya, err := os.Create(kilitDosyaYolu)
				if err != nil {
					return err
				}
				defer dosya.Close()
			}

			// Güncellemeleri kontrol et
			saglayici, err := s.ProviderManager.GetProvider(icerik.Name())
			if err != nil {
				log.Error(err)
				continue
			}

			bilgi, err := (*saglayici).GetInfo()
			if err != nil {
				log.Error(err)
				continue
			}

			if manifest.HasUpdateAvailable(bilgi.Name, bilgi.Version) {
				log.Infof("%s için güncelleme mevcut. `daytona provider update` ile güncelleyebilirsiniz.", bilgi.Name)
			}
		}
	}

	log.Info("Sağlayıcılar kaydedildi")

	return nil
}

func (s *Server) getPluginPath(dizin string) (string, error) {
	dosyalar, err := os.ReadDir(dizin)
	if err != nil {
		return "", err
	}

	for _, dosya := range dosyalar {
		if !dosya.IsDir() && dosya.Name() != manager.INITIAL_SETUP_LOCK_FILE_NAME {
			return filepath.Join(dizin, dosya.Name()), nil
		}
	}

	return "", errors.New(dizin + " dizininde eklenti bulunamadı")
}
