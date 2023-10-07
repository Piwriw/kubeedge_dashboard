/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// DeploymentSpec is the specification of the desired behavior of the Deployment.
type ExtensionsV1beta1DeploymentSpec struct {

	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	MinReadySeconds int32 `json:"minReadySeconds,omitempty"`

	// Indicates that the deployment is paused and will not be processed by the deployment controller.
	Paused bool `json:"paused,omitempty"`

	// The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. This is not set by default.
	ProgressDeadlineSeconds int32 `json:"progressDeadlineSeconds,omitempty"`

	// Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	Replicas int32 `json:"replicas,omitempty"`

	// The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified.
	RevisionHistoryLimit int32 `json:"revisionHistoryLimit,omitempty"`

	// DEPRECATED. The config this deployment is rolling back to. Will be cleared after rollback is done.
	RollbackTo *ExtensionsV1beta1RollbackConfig `json:"rollbackTo,omitempty"`

	// Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment.
	Selector *V1LabelSelector `json:"selector,omitempty"`

	// The deployment strategy to use to replace existing pods with new ones.
	Strategy *ExtensionsV1beta1DeploymentStrategy `json:"strategy,omitempty"`

	// Template describes the pods that will be created.
	Template *V1PodTemplateSpec `json:"template"`
}
