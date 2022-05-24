// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TcprouteAction struct {
	/* Optional. The destination services to which traffic should be forwarded. At least one destination service is required. */
	// +optional
	Destinations []TcprouteDestinations `json:"destinations,omitempty"`

	/* Optional. If true, Router will use the destination IP and port of the original connection as the destination of the request. Default is false. */
	// +optional
	OriginalDestination *bool `json:"originalDestination,omitempty"`
}

type TcprouteDestinations struct {
	/*  */
	ServiceRef v1alpha1.ResourceRef `json:"serviceRef"`

	/* Optional. Specifies the proportion of requests forwarded to the backend referenced by the serviceName field. This is computed as: weight/Sum(weights in this destination list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. If only one serviceName is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weights are specified for any one service name, they need to be specified for all of them. If weights are unspecified for all services, then, traffic is distributed in equal proportions to all of them. */
	// +optional
	Weight *int `json:"weight,omitempty"`
}

type TcprouteMatches struct {
	/* Required. Must be specified in the CIDR range format. A CIDR range consists of an IP Address and a prefix length to construct the subnet mask. By default, the prefix length is 32 (i.e. matches a single IP address). Only IPV4 addresses are supported. Examples: “10.0.0.1” - matches against this exact IP address. “10.0.0.0/8" - matches against any IP address within the 10.0.0.0 subnet and 255.255.255.0 mask. "0.0.0.0/0" - matches against any IP address'. */
	Address string `json:"address"`

	/* Required. Specifies the destination port to match against. */
	Port string `json:"port"`
}

type TcprouteRules struct {
	/* Required. The detailed rule defining how to route matched traffic. */
	Action TcprouteAction `json:"action"`

	/* Optional. RouteMatch defines the predicate used to match requests to a given action. Multiple match types are “OR”ed for evaluation. If no routeMatch field is specified, this rule will unconditionally match traffic. */
	// +optional
	Matches []TcprouteMatches `json:"matches,omitempty"`
}

type NetworkServicesTCPRouteSpec struct {
	/* Optional. A free-text description of the resource. Max length 1024 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/*  */
	// +optional
	Gateways []v1alpha1.ResourceRef `json:"gateways,omitempty"`

	/* Optional. Set of label tags associated with the TcpRoute resource. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/*  */
	// +optional
	Meshes []v1alpha1.ResourceRef `json:"meshes,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Optional. Routers define a list of routers this TcpRoute should be served by. Each router reference should match the pattern: `projects/* /locations/global/routers/` The attached Router should be of a type PROXY */
	// +optional
	Routers []string `json:"routers,omitempty"`

	/* Required. Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match. */
	Rules []TcprouteRules `json:"rules"`
}

type NetworkServicesTCPRouteStatus struct {
	/* Conditions represent the latest available observations of the
	   NetworkServicesTCPRoute's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The timestamp when the resource was created. */
	CreateTime string `json:"createTime,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. Server-defined URL of this resource */
	SelfLink string `json:"selfLink,omitempty"`
	/* Output only. The timestamp when the resource was updated. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesTCPRoute is the Schema for the networkservices API
// +k8s:openapi-gen=true
type NetworkServicesTCPRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesTCPRouteSpec   `json:"spec,omitempty"`
	Status NetworkServicesTCPRouteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesTCPRouteList contains a list of NetworkServicesTCPRoute
type NetworkServicesTCPRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesTCPRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesTCPRoute{}, &NetworkServicesTCPRouteList{})
}
