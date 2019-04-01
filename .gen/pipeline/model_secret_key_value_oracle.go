/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.15.4
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type SecretKeyValueOracle struct {
	UserOcid string `json:"user_ocid"`
	TenancyOcid string `json:"tenancy_ocid"`
	ApiKey string `json:"api_key"`
	ApiKeyFingerprint string `json:"api_key_fingerprint"`
	Region string `json:"region"`
	ClientId string `json:"client_id,omitempty"`
}
