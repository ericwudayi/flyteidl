/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Defines a pod spec and additional pod metadata that is created when a task is executed.
type CoreK8sPod struct {
	// Contains additional metadata for building a kubernetes pod.
	Metadata *CoreK8sObjectMetadata `json:"metadata,omitempty"`
	PodSpec *ProtobufStruct `json:"pod_spec,omitempty"`
	DataConfig *CoreDataLoadingConfig `json:"data_config,omitempty"`
}
