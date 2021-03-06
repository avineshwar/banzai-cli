# CreateEksPropertiesEks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | [optional] 
**LogTypes** | **[]string** |  | [optional] 
**NodePools** | [**map[string]EksNodePool**](EKSNodePool.md) |  | 
**Vpc** | [**EksVpc**](EKSVpc.md) |  | [optional] 
**RouteTableId** | **string** | Id of the RouteTable of the VPC to be used by subnets. This is used only when subnets are created into existing VPC. | [optional] 
**Subnets** | [**[]EksSubnet**](EKSSubnet.md) | Subnets for EKS master and worker nodes. All worker nodes will be launched in the same subnet (the first subnet in the list - which may not coincide with first subnet in the cluster create request payload as the deserialization may change the order) unless a subnet is specified for the workers that belong to a node pool at node pool level. | [optional] 
**Iam** | [**EksIam**](EKSIam.md) |  | [optional] 
**ApiServerAccessPoints** | **[]string** | List of access point types for the API server; public and private are the only valid values | [optional] [default to ["public"]]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


