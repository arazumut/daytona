// 2024 Daytona Platforms Inc. Tüm hakları saklıdır.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"net/http"

	"github.com/daytonaio/daytona/pkg/sample"
)

// FetchSamples sunucudan örnekleri alır
func (s *Server) FetchSamples() ([]sample.Sample, *http.Response, error) {
	if s.config.SamplesIndexUrl == "" {
		return []sample.Sample{}, nil, nil
	}

	return sample.FetchSamples(s.config.SamplesIndexUrl)
}
