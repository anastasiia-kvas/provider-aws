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

type AutoscalingGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LoadBalancers []*string `json:"loadBalancers,omitempty" tf:"load_balancers,omitempty"`

	TargetGroupArns []*string `json:"targetGroupArns,omitempty" tf:"target_group_arns,omitempty"`
}

type AutoscalingGroupParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityRebalance *bool `json:"capacityRebalance,omitempty" tf:"capacity_rebalance,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultCooldown *float64 `json:"defaultCooldown,omitempty" tf:"default_cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	DesiredCapacity *float64 `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	EnabledMetrics []*string `json:"enabledMetrics,omitempty" tf:"enabled_metrics,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDeleteWarmPool *bool `json:"forceDeleteWarmPool,omitempty" tf:"force_delete_warm_pool,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckGracePeriod *float64 `json:"healthCheckGracePeriod,omitempty" tf:"health_check_grace_period,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// +kubebuilder:validation:Optional
	InitialLifecycleHook []InitialLifecycleHookParameters `json:"initialLifecycleHook,omitempty" tf:"initial_lifecycle_hook,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceRefresh []InstanceRefreshParameters `json:"instanceRefresh,omitempty" tf:"instance_refresh,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchConfiguration *string `json:"launchConfiguration,omitempty" tf:"launch_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// +kubebuilder:validation:Optional
	MaxInstanceLifetime *float64 `json:"maxInstanceLifetime,omitempty" tf:"max_instance_lifetime,omitempty"`

	// +kubebuilder:validation:Required
	MaxSize *float64 `json:"maxSize" tf:"max_size,omitempty"`

	// +kubebuilder:validation:Optional
	MetricsGranularity *string `json:"metricsGranularity,omitempty" tf:"metrics_granularity,omitempty"`

	// +kubebuilder:validation:Optional
	MinELBCapacity *float64 `json:"minElbCapacity,omitempty" tf:"min_elb_capacity,omitempty"`

	// +kubebuilder:validation:Required
	MinSize *float64 `json:"minSize" tf:"min_size,omitempty"`

	// +kubebuilder:validation:Optional
	MixedInstancesPolicy []MixedInstancesPolicyParameters `json:"mixedInstancesPolicy,omitempty" tf:"mixed_instances_policy,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.PlacementGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PlacementGroup *string `json:"placementGroup,omitempty" tf:"placement_group,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementGroupRef *v1.Reference `json:"placementGroupRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PlacementGroupSelector *v1.Selector `json:"placementGroupSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProtectFromScaleIn *bool `json:"protectFromScaleIn,omitempty" tf:"protect_from_scale_in,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ServiceLinkedRoleArn *string `json:"serviceLinkedRoleArn,omitempty" tf:"service_linked_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceLinkedRoleArnRef *v1.Reference `json:"serviceLinkedRoleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceLinkedRoleArnSelector *v1.Selector `json:"serviceLinkedRoleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SuspendedProcesses []*string `json:"suspendedProcesses,omitempty" tf:"suspended_processes,omitempty"`

	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TerminationPolicies []*string `json:"terminationPolicies,omitempty" tf:"termination_policies,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	VPCZoneIdentifier []*string `json:"vpcZoneIdentifier,omitempty" tf:"vpc_zone_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	VPCZoneIdentifierRefs []v1.Reference `json:"vpcZoneIdentifierRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCZoneIdentifierSelector *v1.Selector `json:"vpcZoneIdentifierSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	WaitForCapacityTimeout *string `json:"waitForCapacityTimeout,omitempty" tf:"wait_for_capacity_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	WaitForELBCapacity *float64 `json:"waitForElbCapacity,omitempty" tf:"wait_for_elb_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	WarmPool []WarmPoolParameters `json:"warmPool,omitempty" tf:"warm_pool,omitempty"`
}

type InitialLifecycleHookObservation struct {
}

type InitialLifecycleHookParameters struct {

	// +kubebuilder:validation:Optional
	DefaultResult *string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`

	// +kubebuilder:validation:Optional
	HeartbeatTimeout *float64 `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`

	// +kubebuilder:validation:Required
	LifecycleTransition *string `json:"lifecycleTransition" tf:"lifecycle_transition,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationMetadata *string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationTargetArn *string `json:"notificationTargetArn,omitempty" tf:"notification_target_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type InstanceRefreshObservation struct {
}

type InstanceRefreshParameters struct {

	// +kubebuilder:validation:Optional
	Preferences []PreferencesParameters `json:"preferences,omitempty" tf:"preferences,omitempty"`

	// +kubebuilder:validation:Required
	Strategy *string `json:"strategy" tf:"strategy,omitempty"`

	// +kubebuilder:validation:Optional
	Triggers []*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type InstanceReusePolicyObservation struct {
}

type InstanceReusePolicyParameters struct {

	// +kubebuilder:validation:Optional
	ReuseOnScaleIn *bool `json:"reuseOnScaleIn,omitempty" tf:"reuse_on_scale_in,omitempty"`
}

type InstancesDistributionObservation struct {
}

type InstancesDistributionParameters struct {

	// +kubebuilder:validation:Optional
	OnDemandAllocationStrategy *string `json:"onDemandAllocationStrategy,omitempty" tf:"on_demand_allocation_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	OnDemandBaseCapacity *float64 `json:"onDemandBaseCapacity,omitempty" tf:"on_demand_base_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	OnDemandPercentageAboveBaseCapacity *float64 `json:"onDemandPercentageAboveBaseCapacity,omitempty" tf:"on_demand_percentage_above_base_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	SpotAllocationStrategy *string `json:"spotAllocationStrategy,omitempty" tf:"spot_allocation_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	SpotInstancePools *float64 `json:"spotInstancePools,omitempty" tf:"spot_instance_pools,omitempty"`

	// +kubebuilder:validation:Optional
	SpotMaxPrice *string `json:"spotMaxPrice,omitempty" tf:"spot_max_price,omitempty"`
}

type LaunchTemplateObservation struct {
}

type LaunchTemplateParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.LaunchTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.LaunchTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("latestVersion",true)
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// +kubebuilder:validation:Optional
	VersionRef *v1.Reference `json:"versionRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VersionSelector *v1.Selector `json:"versionSelector,omitempty" tf:"-"`
}

type LaunchTemplateSpecificationObservation struct {
}

type LaunchTemplateSpecificationParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.LaunchTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplateIDRef *v1.Reference `json:"launchTemplateIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LaunchTemplateIDSelector *v1.Selector `json:"launchTemplateIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MixedInstancesPolicyLaunchTemplateObservation struct {
}

type MixedInstancesPolicyLaunchTemplateParameters struct {

	// +kubebuilder:validation:Required
	LaunchTemplateSpecification []LaunchTemplateSpecificationParameters `json:"launchTemplateSpecification" tf:"launch_template_specification,omitempty"`

	// +kubebuilder:validation:Optional
	Override []OverrideParameters `json:"override,omitempty" tf:"override,omitempty"`
}

type MixedInstancesPolicyObservation struct {
}

type MixedInstancesPolicyParameters struct {

	// +kubebuilder:validation:Optional
	InstancesDistribution []InstancesDistributionParameters `json:"instancesDistribution,omitempty" tf:"instances_distribution,omitempty"`

	// +kubebuilder:validation:Required
	LaunchTemplate []MixedInstancesPolicyLaunchTemplateParameters `json:"launchTemplate" tf:"launch_template,omitempty"`
}

type OverrideLaunchTemplateSpecificationObservation struct {
}

type OverrideLaunchTemplateSpecificationParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.LaunchTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplateIDRef *v1.Reference `json:"launchTemplateIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LaunchTemplateIDSelector *v1.Selector `json:"launchTemplateIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type OverrideObservation struct {
}

type OverrideParameters struct {

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplateSpecification []OverrideLaunchTemplateSpecificationParameters `json:"launchTemplateSpecification,omitempty" tf:"launch_template_specification,omitempty"`

	// +kubebuilder:validation:Optional
	WeightedCapacity *string `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type PreferencesObservation struct {
}

type PreferencesParameters struct {

	// +kubebuilder:validation:Optional
	CheckpointDelay *string `json:"checkpointDelay,omitempty" tf:"checkpoint_delay,omitempty"`

	// +kubebuilder:validation:Optional
	CheckpointPercentages []*float64 `json:"checkpointPercentages,omitempty" tf:"checkpoint_percentages,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceWarmup *string `json:"instanceWarmup,omitempty" tf:"instance_warmup,omitempty"`

	// +kubebuilder:validation:Optional
	MinHealthyPercentage *float64 `json:"minHealthyPercentage,omitempty" tf:"min_healthy_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	SkipMatching *bool `json:"skipMatching,omitempty" tf:"skip_matching,omitempty"`
}

type TagObservation struct {
}

type TagParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	PropagateAtLaunch *bool `json:"propagateAtLaunch" tf:"propagate_at_launch,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type WarmPoolObservation struct {
}

type WarmPoolParameters struct {

	// +kubebuilder:validation:Optional
	InstanceReusePolicy []InstanceReusePolicyParameters `json:"instanceReusePolicy,omitempty" tf:"instance_reuse_policy,omitempty"`

	// +kubebuilder:validation:Optional
	MaxGroupPreparedCapacity *float64 `json:"maxGroupPreparedCapacity,omitempty" tf:"max_group_prepared_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// +kubebuilder:validation:Optional
	PoolState *string `json:"poolState,omitempty" tf:"pool_state,omitempty"`
}

// AutoscalingGroupSpec defines the desired state of AutoscalingGroup
type AutoscalingGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoscalingGroupParameters `json:"forProvider"`
}

// AutoscalingGroupStatus defines the observed state of AutoscalingGroup.
type AutoscalingGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoscalingGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingGroup is the Schema for the AutoscalingGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AutoscalingGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingGroupSpec   `json:"spec"`
	Status            AutoscalingGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingGroupList contains a list of AutoscalingGroups
type AutoscalingGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingGroup `json:"items"`
}

// Repository type metadata.
var (
	AutoscalingGroup_Kind             = "AutoscalingGroup"
	AutoscalingGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutoscalingGroup_Kind}.String()
	AutoscalingGroup_KindAPIVersion   = AutoscalingGroup_Kind + "." + CRDGroupVersion.String()
	AutoscalingGroup_GroupVersionKind = CRDGroupVersion.WithKind(AutoscalingGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AutoscalingGroup{}, &AutoscalingGroupList{})
}
