package v1alpha1

import (
	toolchainv1alpha1 "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:openapi-gen=true
type ToolchainSpec struct {
	// +optional
	Pipelines OpenShiftPipelines `json:"openshiftPipelines,omitempty"`

	// +optional
	CRW CodeReadyWorkspaces `json:"codereadyWorkspaces,omitempty"`
}

//specDescriptors:
//- description: Specify if CodeReady Workspaces should be installed
//  displayName: Install CodeReady Workspaces
//  path: codereadyWorkspaces.install
//  x-descriptors:
//    - 'urn:alm:descriptor:com.tectonic.ui:fieldGroup:CodeReadyWorkspaces'
//    - 'urn:alm:descriptor:com.tectonic.ui:booleanSwitch'
//- description: Specify if OpenShift Pipelines should be installed
//  displayName: Install OpenShift Pipelines
//  path: openshiftPipelines.install
//  x-descriptors:
//    - 'urn:alm:descriptor:com.tectonic.ui:booleanSwitch'
//statusDescriptors:
//- description: CodeReady Workspaces Dashboard URL
//  displayName: CodeReady Workspaces Dashboard
//  path: cheDashboardURL
//  x-descriptors:
//    - 'urn:alm:descriptor:com.tectonic.ui:fieldGroup:CodeReadyWorkspaces'
//    - 'urn:alm:descriptor:com.tectonic.ui:text'

type OpenShiftPipelines struct {

	// Specify if OpenShift Pipelines should be installed
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.displayName="Install OpenShift Pipelines"
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.x-descriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	// +optional
	Install bool `json:"install"`
}

type CodeReadyWorkspaces struct {

	// Specify if CodeReady Workspaces should be installed
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.displayName="Install CodeReady Workspaces"
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.x-descriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch,urn:alm:descriptor:com.tectonic.ui:fieldGroup:CodeReadyWorkspaces"
	// +optional
	Install bool `json:"install"`

	// CodeReady Workspaces namespace
	// +optional
	Namespace string `json:"namespace"`

	// TLS
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.displayName="TLS"
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.x-descriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch,urn:alm:descriptor:com.tectonic.ui:fieldGroup:CodeReadyWorkspaces"
	// +optional
	TLS bool `json:"tls"`
}

// +k8s:openapi-gen=true
type ToolchainStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current Toolchain conditions
	// Supported condition types:
	// Ready
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType
	Conditions []toolchainv1alpha1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// CodeReady Workspaces Dashboard URL
	// +operator-sdk:gen-csv:customresourcedefinitions.statusDescriptors=true
	// +operator-sdk:gen-csv:customresourcedefinitions.statusDescriptors.displayName="CodeReady Workspaces Dashboard"
	// +operator-sdk:gen-csv:customresourcedefinitions.statusDescriptors.x-descriptors="urn:alm:descriptor:org.w3:link,urn:alm:descriptor:com.tectonic.ui:fieldGroup:CodeReadyWorkspaces"
	// +optional
	CheDashboardURL string `json:"cheDashboardURL,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Toolchain is the Schema for the toolchains API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=toolchains,scope=Namespaced
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].reason"
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Toolchain"
type Toolchain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ToolchainSpec   `json:"spec,omitempty"`
	Status ToolchainStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToolchainList contains a list of Toolchian
type ToolchainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Toolchain `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Toolchain{}, &ToolchainList{})
}
