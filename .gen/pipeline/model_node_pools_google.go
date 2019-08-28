/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.29.0-dev.1
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type NodePoolsGoogle struct {
	Autoscaling bool `json:"autoscaling,omitempty"`
	Preemptible bool `json:"preemptible,omitempty"`
	Count int32 `json:"count"`
	MinCount int32 `json:"minCount,omitempty"`
	MaxCount int32 `json:"maxCount,omitempty"`
	InstanceType string `json:"instanceType"`
	Labels map[string]string `json:"labels,omitempty"`
}
