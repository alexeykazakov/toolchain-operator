package v1alpha1

import (
	toolchainv1alpha1 "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ToolchainClusterSpec defines the desired state of ToolchainCluster
// +k8s:openapi-gen=true
type ToolchainClusterSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// namespace where Toolchain will be installed
	TargetNamespace string `json:"targetNamespace"`
}

// ToolchainClusterStatus defines the observed state of ToolchainCluster
// +k8s:openapi-gen=true
type ToolchainClusterStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current ToolchainCluster conditions
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []toolchainv1alpha1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToolchainCluster is the Schema for the toolchainclusters API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Target Namespace",type="string",JSONPath=".spec.targetNamespace"
type ToolchainCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ToolchainClusterSpec   `json:"spec,omitempty"`
	Status ToolchainClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToolchainClusterList contains a list of ToolchainCluster
type ToolchainClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ToolchainCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ToolchainCluster{}, &ToolchainClusterList{})
}
