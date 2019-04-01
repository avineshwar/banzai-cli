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

type Resource struct {
	Capacity string `json:"capacity,omitempty"`
	Allocatable string `json:"allocatable,omitempty"`
	Limit string `json:"limit,omitempty"`
	Request string `json:"request,omitempty"`
}
