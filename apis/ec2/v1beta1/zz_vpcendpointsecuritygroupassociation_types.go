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

type VPCEndpointSecurityGroupAssociationObservation struct {

	// The ID of the association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPCEndpointSecurityGroupAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether this association should replace the association with the VPC's default security group that is created when no security groups are specified during VPC endpoint creation. At most 1 association per-VPC endpoint should be configured with replace_default_association = true.
	// +kubebuilder:validation:Optional
	ReplaceDefaultAssociation *bool `json:"replaceDefaultAssociation,omitempty" tf:"replace_default_association,omitempty"`

	// The ID of the security group to be associated with the VPC endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecurityGroup in ec2 to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in ec2 to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// The ID of the VPC endpoint with which the security group will be associated.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCEndpoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// Reference to a VPCEndpoint in ec2 to populate vpcEndpointId.
	// +kubebuilder:validation:Optional
	VPCEndpointIDRef *v1.Reference `json:"vpcEndpointIdRef,omitempty" tf:"-"`

	// Selector for a VPCEndpoint in ec2 to populate vpcEndpointId.
	// +kubebuilder:validation:Optional
	VPCEndpointIDSelector *v1.Selector `json:"vpcEndpointIdSelector,omitempty" tf:"-"`
}

// VPCEndpointSecurityGroupAssociationSpec defines the desired state of VPCEndpointSecurityGroupAssociation
type VPCEndpointSecurityGroupAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointSecurityGroupAssociationParameters `json:"forProvider"`
}

// VPCEndpointSecurityGroupAssociationStatus defines the observed state of VPCEndpointSecurityGroupAssociation.
type VPCEndpointSecurityGroupAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointSecurityGroupAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointSecurityGroupAssociation is the Schema for the VPCEndpointSecurityGroupAssociations API. Provides a resource to create an association between a VPC endpoint and a security group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpointSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCEndpointSecurityGroupAssociationSpec   `json:"spec"`
	Status            VPCEndpointSecurityGroupAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointSecurityGroupAssociationList contains a list of VPCEndpointSecurityGroupAssociations
type VPCEndpointSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointSecurityGroupAssociation `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpointSecurityGroupAssociation_Kind             = "VPCEndpointSecurityGroupAssociation"
	VPCEndpointSecurityGroupAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpointSecurityGroupAssociation_Kind}.String()
	VPCEndpointSecurityGroupAssociation_KindAPIVersion   = VPCEndpointSecurityGroupAssociation_Kind + "." + CRDGroupVersion.String()
	VPCEndpointSecurityGroupAssociation_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpointSecurityGroupAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpointSecurityGroupAssociation{}, &VPCEndpointSecurityGroupAssociationList{})
}
