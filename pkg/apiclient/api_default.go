/*
Daytona Sunucu API'si

Daytona Sunucu API'si

API sürümü: v0.0.0-dev
*/

// Bu kod OpenAPI Generator (https://openapi-generator.tech) tarafından oluşturulmuştur; LÜTFEN DÜZENLEMEYİN.

package apiclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// DefaultAPIService varsayılan API servisi
type DefaultAPIService service

// ApiHealthCheckRequest sağlık kontrolü isteği
type ApiHealthCheckRequest struct {
	ctx        context.Context
	ApiService *DefaultAPIService
}

// Execute isteği çalıştırır
func (r ApiHealthCheckRequest) Execute() (map[string]string, *http.Response, error) {
	return r.ApiService.HealthCheckExecute(r)
}

/*
HealthCheck sağlık kontrolü

Sağlık kontrolü

	@param ctx context.Context - kimlik doğrulama, loglama, iptal, son tarihler, izleme vb. için kullanılır. http.Request veya context.Background()'dan geçirilir.
	@return ApiHealthCheckRequest
*/
func (a *DefaultAPIService) HealthCheck(ctx context.Context) ApiHealthCheckRequest {
	return ApiHealthCheckRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// HealthCheckExecute isteği çalıştırır
//
//	@return map[string]string
func (a *DefaultAPIService) HealthCheckExecute(r ApiHealthCheckRequest) (map[string]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.HealthCheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/health"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// İçerik Türü başlığını belirlemek için
	localVarHTTPContentTypes := []string{}

	// İçerik Türü başlığını ayarla
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// Kabul başlığını belirlemek için
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// Kabul başlığını ayarla
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Anahtarı Kimlik Doğrulama
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
