// Daytona Platforms Inc. 2024 Telif Hakkı
// SPDX-License-Identifier: Apache-2.0

package controllers

import (
	"net/http"
	"regexp"
	"strconv"
)

// Hata mesajından HTTP durum kodu ve mesajını alır
func HataMesajindanHTTPDurumKoduVeMesajiAl(err error) (int, string, error) {
	re := regexp.MustCompile(`durum kodu: (\d{3}) hata: (.+)`)
	eslesme := re.FindStringSubmatch(err.Error())
	if len(eslesme) > 2 {
		// eşleşen stringi bir tamsayıya (durum kodu) dönüştür
		durumKodu, cevirimHatasi := strconv.Atoi(eslesme[1])
		if cevirimHatasi == nil {
			hataMesaji := eslesme[2]
			return durumKodu, hataMesaji, nil
		}
	}

	return http.StatusInternalServerError, "", err
}
