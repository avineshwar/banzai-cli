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

type PkeClusterReadinessResponseMaster struct {
	// true when the node has been reported to be ready
	Ready bool `json:"ready,omitempty"`
}
