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

type ProvisionedConcurrencyConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProvisionedConcurrencyConfigParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/lambda/v1beta1.Alias
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("functionName",false)
	// +kubebuilder:validation:Optional
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// +kubebuilder:validation:Optional
	FunctionNameRef *v1.Reference `json:"functionNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FunctionNameSelector *v1.Selector `json:"functionNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ProvisionedConcurrentExecutions *float64 `json:"provisionedConcurrentExecutions" tf:"provisioned_concurrent_executions,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/lambda/v1beta1.Alias
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// +kubebuilder:validation:Optional
	QualifierRef *v1.Reference `json:"qualifierRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	QualifierSelector *v1.Selector `json:"qualifierSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ProvisionedConcurrencyConfigSpec defines the desired state of ProvisionedConcurrencyConfig
type ProvisionedConcurrencyConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProvisionedConcurrencyConfigParameters `json:"forProvider"`
}

// ProvisionedConcurrencyConfigStatus defines the observed state of ProvisionedConcurrencyConfig.
type ProvisionedConcurrencyConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProvisionedConcurrencyConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProvisionedConcurrencyConfig is the Schema for the ProvisionedConcurrencyConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ProvisionedConcurrencyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProvisionedConcurrencyConfigSpec   `json:"spec"`
	Status            ProvisionedConcurrencyConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProvisionedConcurrencyConfigList contains a list of ProvisionedConcurrencyConfigs
type ProvisionedConcurrencyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProvisionedConcurrencyConfig `json:"items"`
}

// Repository type metadata.
var (
	ProvisionedConcurrencyConfig_Kind             = "ProvisionedConcurrencyConfig"
	ProvisionedConcurrencyConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProvisionedConcurrencyConfig_Kind}.String()
	ProvisionedConcurrencyConfig_KindAPIVersion   = ProvisionedConcurrencyConfig_Kind + "." + CRDGroupVersion.String()
	ProvisionedConcurrencyConfig_GroupVersionKind = CRDGroupVersion.WithKind(ProvisionedConcurrencyConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&ProvisionedConcurrencyConfig{}, &ProvisionedConcurrencyConfigList{})
}
