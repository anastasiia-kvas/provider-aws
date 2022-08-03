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

type BucketReplicationConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketReplicationConfigurationParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Rule []BucketReplicationConfigurationRuleParameters `json:"rule" tf:"rule,omitempty"`

	// +kubebuilder:validation:Optional
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`
}

type BucketReplicationConfigurationRuleFilterObservation struct {
}

type BucketReplicationConfigurationRuleFilterParameters struct {

	// +kubebuilder:validation:Optional
	And []FilterAndParameters `json:"and,omitempty" tf:"and,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Tag []FilterTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BucketReplicationConfigurationRuleObservation struct {
}

type BucketReplicationConfigurationRuleParameters struct {

	// +kubebuilder:validation:Optional
	DeleteMarkerReplication []DeleteMarkerReplicationParameters `json:"deleteMarkerReplication,omitempty" tf:"delete_marker_replication,omitempty"`

	// +kubebuilder:validation:Required
	Destination []RuleDestinationParameters `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	ExistingObjectReplication []ExistingObjectReplicationParameters `json:"existingObjectReplication,omitempty" tf:"existing_object_replication,omitempty"`

	// +kubebuilder:validation:Optional
	Filter []BucketReplicationConfigurationRuleFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	SourceSelectionCriteria []RuleSourceSelectionCriteriaParameters `json:"sourceSelectionCriteria,omitempty" tf:"source_selection_criteria,omitempty"`

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type DeleteMarkerReplicationObservation struct {
}

type DeleteMarkerReplicationParameters struct {

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type DestinationAccessControlTranslationObservation struct {
}

type DestinationAccessControlTranslationParameters struct {

	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`
}

type DestinationMetricsObservation struct {
}

type DestinationMetricsParameters struct {

	// +kubebuilder:validation:Optional
	EventThreshold []EventThresholdParameters `json:"eventThreshold,omitempty" tf:"event_threshold,omitempty"`

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type DestinationReplicationTimeObservation struct {
}

type DestinationReplicationTimeParameters struct {

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`

	// +kubebuilder:validation:Required
	Time []TimeParameters `json:"time" tf:"time,omitempty"`
}

type EncryptionConfigurationObservation struct {
}

type EncryptionConfigurationParameters struct {

	// +kubebuilder:validation:Required
	ReplicaKMSKeyID *string `json:"replicaKmsKeyId" tf:"replica_kms_key_id,omitempty"`
}

type EventThresholdObservation struct {
}

type EventThresholdParameters struct {

	// +kubebuilder:validation:Required
	Minutes *float64 `json:"minutes" tf:"minutes,omitempty"`
}

type ExistingObjectReplicationObservation struct {
}

type ExistingObjectReplicationParameters struct {

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type FilterAndObservation struct {
}

type FilterAndParameters struct {

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FilterTagObservation struct {
}

type FilterTagParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ReplicaModificationsObservation struct {
}

type ReplicaModificationsParameters struct {

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type RuleDestinationObservation struct {
}

type RuleDestinationParameters struct {

	// +kubebuilder:validation:Optional
	AccessControlTranslation []DestinationAccessControlTranslationParameters `json:"accessControlTranslation,omitempty" tf:"access_control_translation,omitempty"`

	// +kubebuilder:validation:Optional
	Account *string `json:"account,omitempty" tf:"account,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Metrics []DestinationMetricsParameters `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// +kubebuilder:validation:Optional
	ReplicationTime []DestinationReplicationTimeParameters `json:"replicationTime,omitempty" tf:"replication_time,omitempty"`

	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleSourceSelectionCriteriaObservation struct {
}

type RuleSourceSelectionCriteriaParameters struct {

	// +kubebuilder:validation:Optional
	ReplicaModifications []ReplicaModificationsParameters `json:"replicaModifications,omitempty" tf:"replica_modifications,omitempty"`

	// +kubebuilder:validation:Optional
	SseKMSEncryptedObjects []SourceSelectionCriteriaSseKMSEncryptedObjectsParameters `json:"sseKmsEncryptedObjects,omitempty" tf:"sse_kms_encrypted_objects,omitempty"`
}

type SourceSelectionCriteriaSseKMSEncryptedObjectsObservation struct {
}

type SourceSelectionCriteriaSseKMSEncryptedObjectsParameters struct {

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type TimeObservation struct {
}

type TimeParameters struct {

	// +kubebuilder:validation:Required
	Minutes *float64 `json:"minutes" tf:"minutes,omitempty"`
}

// BucketReplicationConfigurationSpec defines the desired state of BucketReplicationConfiguration
type BucketReplicationConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketReplicationConfigurationParameters `json:"forProvider"`
}

// BucketReplicationConfigurationStatus defines the observed state of BucketReplicationConfiguration.
type BucketReplicationConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketReplicationConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketReplicationConfiguration is the Schema for the BucketReplicationConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketReplicationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketReplicationConfigurationSpec   `json:"spec"`
	Status            BucketReplicationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketReplicationConfigurationList contains a list of BucketReplicationConfigurations
type BucketReplicationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketReplicationConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketReplicationConfiguration_Kind             = "BucketReplicationConfiguration"
	BucketReplicationConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketReplicationConfiguration_Kind}.String()
	BucketReplicationConfiguration_KindAPIVersion   = BucketReplicationConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketReplicationConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketReplicationConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketReplicationConfiguration{}, &BucketReplicationConfigurationList{})
}
