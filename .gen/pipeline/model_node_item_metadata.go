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

type NodeItemMetadata struct {
	Name string `json:"name,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
	Uid string `json:"uid,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	Labels NodeItemMetadataLabels `json:"labels,omitempty"`
	Annotations NodeItemMetadataAnnotations `json:"annotations,omitempty"`
}
