package v0alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// The data source resource is a reflection of the individual datasource instances
// that are exposed in the groups: {datasource}.datasource.grafana.app
// The status is updated periodically.
// The name is the plugin id
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DataSourceApiServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The display name
	Title string `json:"title"`

	// Describe the plugin
	Description string `json:"description,omitempty"`

	// The group + preferred version
	GroupVersion string `json:"groupVersion"`

	// Possible alternative plugin IDs
	AliasIDs []string `json:"aliasIDs,omitempty"`
}

// List of datasource plugins
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DataSourceApiServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DataSourceApiServer `json:"items,omitempty"`
}
