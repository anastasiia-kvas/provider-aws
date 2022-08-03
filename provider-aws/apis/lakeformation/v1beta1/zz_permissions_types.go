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

type DataLocationObservation struct {
}

type DataLocationParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/lakeformation/v1beta1.Resource
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`
}

type DatabaseObservation struct {
}

type DatabaseParameters struct {

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/glue/v1beta1.CatalogDatabase
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
}

type PermissionsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PermissionsParameters struct {

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// +kubebuilder:validation:Optional
	CatalogResource *bool `json:"catalogResource,omitempty" tf:"catalog_resource,omitempty"`

	// +kubebuilder:validation:Optional
	DataLocation []DataLocationParameters `json:"dataLocation,omitempty" tf:"data_location,omitempty"`

	// +kubebuilder:validation:Optional
	Database []DatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// +kubebuilder:validation:Required
	Permissions []*string `json:"permissions" tf:"permissions,omitempty"`

	// +kubebuilder:validation:Optional
	PermissionsWithGrantOption []*string `json:"permissionsWithGrantOption,omitempty" tf:"permissions_with_grant_option,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// +kubebuilder:validation:Optional
	PrincipalRef *v1.Reference `json:"principalRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrincipalSelector *v1.Selector `json:"principalSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Table []TableParameters `json:"table,omitempty" tf:"table,omitempty"`

	// +kubebuilder:validation:Optional
	TableWithColumns []TableWithColumnsParameters `json:"tableWithColumns,omitempty" tf:"table_with_columns,omitempty"`
}

type TableObservation struct {
}

type TableParameters struct {

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Wildcard *bool `json:"wildcard,omitempty" tf:"wildcard,omitempty"`
}

type TableWithColumnsObservation struct {
}

type TableWithColumnsParameters struct {

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// +kubebuilder:validation:Optional
	ColumnNames []*string `json:"columnNames,omitempty" tf:"column_names,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/glue/v1beta1.CatalogTable
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("databaseName",false)
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ExcludedColumnNames []*string `json:"excludedColumnNames,omitempty" tf:"excluded_column_names,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/glue/v1beta1.CatalogTable
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Wildcard *bool `json:"wildcard,omitempty" tf:"wildcard,omitempty"`
}

// PermissionsSpec defines the desired state of Permissions
type PermissionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionsParameters `json:"forProvider"`
}

// PermissionsStatus defines the observed state of Permissions.
type PermissionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Permissions is the Schema for the Permissionss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Permissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionsSpec   `json:"spec"`
	Status            PermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionsList contains a list of Permissionss
type PermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Permissions `json:"items"`
}

// Repository type metadata.
var (
	Permissions_Kind             = "Permissions"
	Permissions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Permissions_Kind}.String()
	Permissions_KindAPIVersion   = Permissions_Kind + "." + CRDGroupVersion.String()
	Permissions_GroupVersionKind = CRDGroupVersion.WithKind(Permissions_Kind)
)

func init() {
	SchemeBuilder.Register(&Permissions{}, &PermissionsList{})
}
