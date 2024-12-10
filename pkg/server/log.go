// 2024 Daytona Platforms Inc. Tüm hakları saklıdır.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"io"
	"os"

	frp_log "github.com/fatedier/frp/pkg/util/log"
	log "github.com/sirupsen/logrus"
)

type logFormatter struct {
	textFormatter *log.TextFormatter
	file          *os.File
}

func (f *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	formatted, err := f.textFormatter.Format(entry)
	if err != nil {
		return nil, err
	}

	if f.file != nil {
		// Dosyaya yaz
		_, err = f.file.Write(formatted)
		if err != nil {
			return nil, err
		}
	}

	return formatted, nil
}

func (s *Server) initLogs() error {
	dosyaYolu := s.config.LogFilePath

	dosya, err := os.OpenFile(dosyaYolu, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	logFormatter := &logFormatter{
		textFormatter: &log.TextFormatter{
			ForceColors: true,
		},
		file: dosya,
	}

	log.SetFormatter(logFormatter)

	frpLogSeviyesi := "error"
	if os.Getenv("FRP_LOG_LEVEL") != "" {
		frpLogSeviyesi = os.Getenv("FRP_LOG_LEVEL")
	}

	frpCikti := dosyaYolu
	if os.Getenv("FRP_LOG_OUTPUT") != "" {
		frpCikti = os.Getenv("FRP_LOG_OUTPUT")
	}

	frp_log.InitLogger(frpCikti, frpLogSeviyesi, 0, false)

	return nil
}

func (s *Server) GetLogReader() (io.Reader, error) {
	dosya, err := os.Open(s.config.LogFilePath)
	if err != nil {
		return nil, err
	}

	return dosya, nil
}
