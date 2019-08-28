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

// A collection of whitelist items to match a policy evaluation against.
type Whitelist struct {
	Id string `json:"id"`
	Name string `json:"name,omitempty"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
	Items []WhitelistItem `json:"items,omitempty"`
}
