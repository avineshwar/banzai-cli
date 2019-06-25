/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.21.2
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type PkeOnAzureNodePool struct {
	Name string `json:"name"`
	Roles []string `json:"roles"`
	Labels map[string]string `json:"labels,omitempty"`
	Subnet PkeOnAzureNodePoolSubnet `json:"subnet,omitempty"`
	Zones []string `json:"zones,omitempty"`
	Autoscaling bool `json:"autoscaling,omitempty"`
	MinCount int32 `json:"minCount,omitempty"`
	MaxCount int32 `json:"maxCount,omitempty"`
	Count int32 `json:"count,omitempty"`
	InstanceType string `json:"instanceType"`
}