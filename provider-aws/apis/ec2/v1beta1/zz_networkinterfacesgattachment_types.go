/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkInterfaceSgAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NetworkInterfaceSgAttachmentParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("primaryNetworkInterfaceId",true)
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`
}

// NetworkInterfaceSgAttachmentSpec defines the desired state of NetworkInterfaceSgAttachment
type NetworkInterfaceSgAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceSgAttachmentParameters `json:"forProvider"`
}

// NetworkInterfaceSgAttachmentStatus defines the observed state of NetworkInterfaceSgAttachment.
type NetworkInterfaceSgAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceSgAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceSgAttachment is the Schema for the NetworkInterfaceSgAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NetworkInterfaceSgAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSgAttachmentSpec   `json:"spec"`
	Status            NetworkInterfaceSgAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceSgAttachmentList contains a list of NetworkInterfaceSgAttachments
type NetworkInterfaceSgAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceSgAttachment `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterfaceSgAttachment_Kind             = "NetworkInterfaceSgAttachment"
	NetworkInterfaceSgAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterfaceSgAttachment_Kind}.String()
	NetworkInterfaceSgAttachment_KindAPIVersion   = NetworkInterfaceSgAttachment_Kind + "." + CRDGroupVersion.String()
	NetworkInterfaceSgAttachment_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterfaceSgAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceSgAttachment{}, &NetworkInterfaceSgAttachmentList{})
}
