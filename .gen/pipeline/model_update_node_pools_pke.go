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

type UpdateNodePoolsPke struct {
	// Instance type for the nodes in the node pool. This field is ignored when existing node pool is updated as instance type can not be changed.
	InstanceType string `json:"instanceType,omitempty"`
	// The spot price for a node in the node pool. Provide \"\" or \"0\" for on-demand nodes.
	SpotPrice string `json:"spotPrice,omitempty"`
	// Whether to enable Kubernetes cluster autoscaler for this node pool.
	Autoscaling bool `json:"autoscaling,omitempty"`
	// If cluster autoscaler is enabled for this node pool it sets the minimum node count the cluster autoscaler can downscale the node pool to.
	MinCount int32 `json:"minCount,omitempty"`
	// If cluster autoscaler is enabled for this node pool it sets the maximum node count the cluster autoscaler can upscale the node pool to.
	MaxCount int32 `json:"maxCount,omitempty"`
	// If cluster autoscaler is not enabled this specifies the desired ndoe count in the node pool. If cluster autoscaler is enabled this specifies the initial node count in the ndoe pool.
	Count int32 `json:"count,omitempty"`
	// The subnet to create the node pool into. If this field is omitted than the subnet from the cluster level network configuration is used.
	Subnets []string `json:"subnets,omitempty"`
}
