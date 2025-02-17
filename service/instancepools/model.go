// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instancepools

import "fmt"

// all definitions in this file are in alphabetical order

type CreateInstancePool struct {
	// Attributes related to pool running on Amazon Web Services. If not
	// specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to pool running on Azure. If not specified at pool
	// creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
	// The fleet related setting to power the instance pool.
	InstancePoolFleetAttributes *InstancePoolFleetAttributes `json:"instance_pool_fleet_attributes,omitempty"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string `json:"instance_pool_name,omitempty"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Custom Docker Image BYOC
	PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
	// A list of preloaded Spark image versions for the pool. Pool-backed
	// clusters started with the preloaded Spark version will start faster. A
	// list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
}

type CreateInstancePoolResponse struct {
	// The ID of the created instance pool.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
}

type DeleteInstancePool struct {
	// The instance pool to be terminated.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
}

type DiskSpec struct {
	// The number of disks launched for each instance: - This feature is only
	// enabled for supported node types. - Users can choose up to the limit of
	// the disks supported by the node type. - For node types with no OS disk,
	// at least one disk must be specified; otherwise, cluster creation will
	// fail.
	//
	// If disks are attached, Databricks will configure Spark to use only the
	// disks for scratch storage, because heterogenously sized scratch devices
	// can lead to inefficient disk utilization. If no disks are attached,
	// Databricks will configure Spark to use instance store disks.
	//
	// Note: If disks are specified, then the Spark configuration
	// `spark.local.dir` will be overridden.
	//
	// Disks will be mounted at: - For AWS: `/ebs0`, `/ebs1`, and etc. - For
	// Azure: `/remote_volume0`, `/remote_volume1`, and etc.
	DiskCount int `json:"disk_count,omitempty"`

	DiskIops int `json:"disk_iops,omitempty"`
	// The size of each disk (in GiB) launched for each instance. Values must
	// fall into the supported range for a particular instance type.
	//
	// For AWS: - General Purpose SSD: 100 - 4096 GiB - Throughput Optimized
	// HDD: 500 - 4096 GiB
	//
	// For Azure: - Premium LRS (SSD): 1 - 1023 GiB - Standard LRS (HDD): 1-
	// 1023 GiB
	DiskSize int `json:"disk_size,omitempty"`

	DiskThroughput int `json:"disk_throughput,omitempty"`
	// The type of disks that will be launched with this cluster.
	DiskType *DiskType `json:"disk_type,omitempty"`
}

type DiskType struct {
	AzureDiskVolumeType DiskTypeAzureDiskVolumeType `json:"azure_disk_volume_type,omitempty"`

	EbsVolumeType DiskTypeEbsVolumeType `json:"ebs_volume_type,omitempty"`
}

type DiskTypeAzureDiskVolumeType string

const DiskTypeAzureDiskVolumeTypePremiumLrs DiskTypeAzureDiskVolumeType = `PREMIUM_LRS`

const DiskTypeAzureDiskVolumeTypeStandardLrs DiskTypeAzureDiskVolumeType = `STANDARD_LRS`

// String representation for [fmt.Print]
func (dtadvt *DiskTypeAzureDiskVolumeType) String() string {
	return string(*dtadvt)
}

// Set raw string value and validate it against allowed values
func (dtadvt *DiskTypeAzureDiskVolumeType) Set(v string) error {
	switch v {
	case `PREMIUM_LRS`, `STANDARD_LRS`:
		*dtadvt = DiskTypeAzureDiskVolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PREMIUM_LRS", "STANDARD_LRS"`, v)
	}
}

// Type always returns DiskTypeAzureDiskVolumeType to satisfy [pflag.Value] interface
func (dtadvt *DiskTypeAzureDiskVolumeType) Type() string {
	return "DiskTypeAzureDiskVolumeType"
}

type DiskTypeEbsVolumeType string

const DiskTypeEbsVolumeTypeGeneralPurposeSsd DiskTypeEbsVolumeType = `GENERAL_PURPOSE_SSD`

const DiskTypeEbsVolumeTypeThroughputOptimizedHdd DiskTypeEbsVolumeType = `THROUGHPUT_OPTIMIZED_HDD`

// String representation for [fmt.Print]
func (dtevt *DiskTypeEbsVolumeType) String() string {
	return string(*dtevt)
}

// Set raw string value and validate it against allowed values
func (dtevt *DiskTypeEbsVolumeType) Set(v string) error {
	switch v {
	case `GENERAL_PURPOSE_SSD`, `THROUGHPUT_OPTIMIZED_HDD`:
		*dtevt = DiskTypeEbsVolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GENERAL_PURPOSE_SSD", "THROUGHPUT_OPTIMIZED_HDD"`, v)
	}
}

// Type always returns DiskTypeEbsVolumeType to satisfy [pflag.Value] interface
func (dtevt *DiskTypeEbsVolumeType) Type() string {
	return "DiskTypeEbsVolumeType"
}

type DockerBasicAuth struct {
	Password string `json:"password,omitempty"`

	Username string `json:"username,omitempty"`
}

type DockerImage struct {
	BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty"`
	// URL of the docker image.
	Url string `json:"url,omitempty"`
}

type EditInstancePool struct {
	// Attributes related to pool running on Amazon Web Services. If not
	// specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to pool running on Azure. If not specified at pool
	// creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows up to 45 custom tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Defines the specification of the disks that will be attached to all Spark
	// containers.
	DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`

	InstancePoolId string `json:"instance_pool_id"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string `json:"instance_pool_name,omitempty"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle ones. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Custom Docker image BYOC
	PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
	// A list of preloaded Spark image versions for the pool. Pool-backed
	// clusters started with the preloaded Spark version will start faster. A
	// list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
}

type FleetLaunchTemplateOverride struct {
	// User-assigned preferred availability zone. It will adjust to the default
	// zone of the worker environment if the preferred zone does not exist in
	// the subnet.
	AvailabilityZone string `json:"availability_zone"`

	InstanceType string `json:"instance_type"`
	// The maximum price per unit hour that you are willing to pay for a Spot
	// Instance.
	MaxPrice float64 `json:"max_price,omitempty"`
	// The priority for the launch template override. If AllocationStrategy is
	// set to prioritized, EC2 Fleet uses priority to determine which launch
	// template override or to use first in fulfilling On-Demand capacity. The
	// highest priority is launched first. Valid values are whole numbers
	// starting at 0. The lower the number, the higher the priority. If no
	// number is set, the launch template override has the lowest priority.
	Priority float64 `json:"priority,omitempty"`
}

type FleetOnDemandOption struct {
	// Only lowest-price and prioritized are allowed
	AllocationStrategy FleetOnDemandOptionAllocationStrategy `json:"allocation_strategy,omitempty"`
	// The maximum amount per hour for On-Demand Instances that you're willing
	// to pay.
	MaxTotalPrice float64 `json:"max_total_price,omitempty"`
	// If you specify use-capacity-reservations-first, the fleet uses unused
	// Capacity Reservations to fulfill On-Demand capacity up to the target
	// On-Demand capacity. If multiple instance pools have unused Capacity
	// Reservations, the On-Demand allocation strategy (lowest-price or
	// prioritized) is applied. If the number of unused Capacity Reservations is
	// less than the On-Demand target capacity, the remaining On-Demand target
	// capacity is launched according to the On-Demand allocation strategy
	// (lowest-price or prioritized).
	UseCapacityReservationsFirst bool `json:"use_capacity_reservations_first,omitempty"`
}

// Only lowest-price and prioritized are allowed
type FleetOnDemandOptionAllocationStrategy string

const FleetOnDemandOptionAllocationStrategyCapacityOptimized FleetOnDemandOptionAllocationStrategy = `CAPACITY_OPTIMIZED`

const FleetOnDemandOptionAllocationStrategyDiversified FleetOnDemandOptionAllocationStrategy = `DIVERSIFIED`

const FleetOnDemandOptionAllocationStrategyLowestPrice FleetOnDemandOptionAllocationStrategy = `LOWEST_PRICE`

const FleetOnDemandOptionAllocationStrategyPrioritized FleetOnDemandOptionAllocationStrategy = `PRIORITIZED`

// String representation for [fmt.Print]
func (fodoas *FleetOnDemandOptionAllocationStrategy) String() string {
	return string(*fodoas)
}

// Set raw string value and validate it against allowed values
func (fodoas *FleetOnDemandOptionAllocationStrategy) Set(v string) error {
	switch v {
	case `CAPACITY_OPTIMIZED`, `DIVERSIFIED`, `LOWEST_PRICE`, `PRIORITIZED`:
		*fodoas = FleetOnDemandOptionAllocationStrategy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAPACITY_OPTIMIZED", "DIVERSIFIED", "LOWEST_PRICE", "PRIORITIZED"`, v)
	}
}

// Type always returns FleetOnDemandOptionAllocationStrategy to satisfy [pflag.Value] interface
func (fodoas *FleetOnDemandOptionAllocationStrategy) Type() string {
	return "FleetOnDemandOptionAllocationStrategy"
}

type FleetSpotOption struct {
	// lowest-price | diversified | capacity-optimized
	AllocationStrategy FleetSpotOptionAllocationStrategy `json:"allocation_strategy,omitempty"`
	// The number of Spot pools across which to allocate your target Spot
	// capacity. Valid only when Spot Allocation Strategy is set to
	// lowest-price. EC2 Fleet selects the cheapest Spot pools and evenly
	// allocates your target Spot capacity across the number of Spot pools that
	// you specify.
	InstancePoolsToUseCount int `json:"instance_pools_to_use_count,omitempty"`
	// The maximum amount per hour for Spot Instances that you're willing to
	// pay.
	MaxTotalPrice float64 `json:"max_total_price,omitempty"`
}

// lowest-price | diversified | capacity-optimized
type FleetSpotOptionAllocationStrategy string

const FleetSpotOptionAllocationStrategyCapacityOptimized FleetSpotOptionAllocationStrategy = `CAPACITY_OPTIMIZED`

const FleetSpotOptionAllocationStrategyDiversified FleetSpotOptionAllocationStrategy = `DIVERSIFIED`

const FleetSpotOptionAllocationStrategyLowestPrice FleetSpotOptionAllocationStrategy = `LOWEST_PRICE`

const FleetSpotOptionAllocationStrategyPrioritized FleetSpotOptionAllocationStrategy = `PRIORITIZED`

// String representation for [fmt.Print]
func (fsoas *FleetSpotOptionAllocationStrategy) String() string {
	return string(*fsoas)
}

// Set raw string value and validate it against allowed values
func (fsoas *FleetSpotOptionAllocationStrategy) Set(v string) error {
	switch v {
	case `CAPACITY_OPTIMIZED`, `DIVERSIFIED`, `LOWEST_PRICE`, `PRIORITIZED`:
		*fsoas = FleetSpotOptionAllocationStrategy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAPACITY_OPTIMIZED", "DIVERSIFIED", "LOWEST_PRICE", "PRIORITIZED"`, v)
	}
}

// Type always returns FleetSpotOptionAllocationStrategy to satisfy [pflag.Value] interface
func (fsoas *FleetSpotOptionAllocationStrategy) Type() string {
	return "FleetSpotOptionAllocationStrategy"
}

// Get instance pool information
type Get struct {
	// The canonical unique identifier for the instance pool.
	InstancePoolId string `json:"-" url:"instance_pool_id,omitempty"`
}

type GetInstancePool struct {
	// Attributes related to pool running on Amazon Web Services. If not
	// specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to pool running on Azure. If not specified at pool
	// creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Tags that are added by Databricks regardless of any `custom_tags`,
	// including:
	//
	// - Vendor: Databricks
	//
	// - InstancePoolCreator: <user_id_of_creator>
	//
	// - InstancePoolName: <name_of_pool>
	//
	// - InstancePoolId: <id_of_pool>
	DefaultTags map[string]string `json:"default_tags,omitempty"`
	// Defines the specification of the disks that will be attached to all Spark
	// containers.
	DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
	// Autoscaling Local Storage: When enabled, the instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
	// Canonical unique identifier for the pool.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// Pool name requested by the user. Name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string `json:"instance_pool_name,omitempty"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle ones. Clusters that require
	// further instance provision would fail during upsize requests.
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Custom Docker image BYOC
	PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
	// A list of preloaded Spark image versions for the pool. Pool-backed
	// clusters started with the preloaded Spark version will start faster. A
	// list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
	// Current state of the instance pool.
	State GetInstancePoolState `json:"state,omitempty"`
	// Usage statistics about the instance pool.
	Stats *InstancePoolStats `json:"stats,omitempty"`
	// Status of failed pending instances in the pool.
	Status *InstancePoolStatus `json:"status,omitempty"`
}

// Current state of the instance pool.
type GetInstancePoolState string

const GetInstancePoolStateActive GetInstancePoolState = `ACTIVE`

const GetInstancePoolStateDeleted GetInstancePoolState = `DELETED`

const GetInstancePoolStateStopped GetInstancePoolState = `STOPPED`

// String representation for [fmt.Print]
func (gips *GetInstancePoolState) String() string {
	return string(*gips)
}

// Set raw string value and validate it against allowed values
func (gips *GetInstancePoolState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DELETED`, `STOPPED`:
		*gips = GetInstancePoolState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DELETED", "STOPPED"`, v)
	}
}

// Type always returns GetInstancePoolState to satisfy [pflag.Value] interface
func (gips *GetInstancePoolState) Type() string {
	return "GetInstancePoolState"
}

type InstancePoolAndStats struct {
	// Attributes related to pool running on Amazon Web Services. If not
	// specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to pool running on Azure. If not specified at pool
	// creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows up to 45 custom tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Tags that are added by Databricks regardless of any `custom_tags`,
	// including:
	//
	// - Vendor: Databricks
	//
	// - InstancePoolCreator: <user_id_of_creator>
	//
	// - InstancePoolName: <name_of_pool>
	//
	// - InstancePoolId: <id_of_pool>
	DefaultTags map[string]string `json:"default_tags,omitempty"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
	// Autoscaling Local Storage: When enabled, the instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
	// Canonical unique identifier for the pool.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string `json:"instance_pool_name,omitempty"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle ones. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Custom Docker image BYOC
	PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
	// A list of preloaded Spark image versions for the pool. Pool-backed
	// clusters started with the preloaded Spark version will start faster. A
	// list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
	// Current state of the instance pool.
	State InstancePoolAndStatsState `json:"state,omitempty"`
	// Usage statistics about the instance pool.
	Stats *InstancePoolStats `json:"stats,omitempty"`
	// Status of failed pending instances in the pool.
	Status *InstancePoolStatus `json:"status,omitempty"`
}

// Current state of the instance pool.
type InstancePoolAndStatsState string

const InstancePoolAndStatsStateActive InstancePoolAndStatsState = `ACTIVE`

const InstancePoolAndStatsStateDeleted InstancePoolAndStatsState = `DELETED`

const InstancePoolAndStatsStateStopped InstancePoolAndStatsState = `STOPPED`

// String representation for [fmt.Print]
func (ipass *InstancePoolAndStatsState) String() string {
	return string(*ipass)
}

// Set raw string value and validate it against allowed values
func (ipass *InstancePoolAndStatsState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DELETED`, `STOPPED`:
		*ipass = InstancePoolAndStatsState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DELETED", "STOPPED"`, v)
	}
}

// Type always returns InstancePoolAndStatsState to satisfy [pflag.Value] interface
func (ipass *InstancePoolAndStatsState) Type() string {
	return "InstancePoolAndStatsState"
}

type InstancePoolAwsAttributes struct {
	// Availability type used for the spot nodes.
	//
	// The default value is defined by
	// InstancePoolConf.instancePoolDefaultAwsAvailability
	Availability InstancePoolAwsAttributesAvailability `json:"availability,omitempty"`
	// Calculates the bid price for AWS spot instances, as a percentage of the
	// corresponding instance type's on-demand price. For example, if this field
	// is set to 50, and the cluster needs a new `r3.xlarge` spot instance, then
	// the bid price is half of the price of on-demand `r3.xlarge` instances.
	// Similarly, if this field is set to 200, the bid price is twice the price
	// of on-demand `r3.xlarge` instances. If not specified, the default value
	// is 100. When spot instances are requested for this cluster, only spot
	// instances whose bid price percentage matches this field will be
	// considered. Note that, for safety, we enforce this field to be no more
	// than 10000.
	//
	// The default value and documentation here should be kept consistent with
	// CommonConf.defaultSpotBidPricePercent and
	// CommonConf.maxSpotBidPricePercent.
	SpotBidPricePercent int `json:"spot_bid_price_percent,omitempty"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. The list of available zones as well as the default value
	// can be found by using the `List Zones`_ method.
	ZoneId string `json:"zone_id,omitempty"`
}

// Availability type used for the spot nodes.
//
// The default value is defined by
// InstancePoolConf.instancePoolDefaultAwsAvailability
type InstancePoolAwsAttributesAvailability string

const InstancePoolAwsAttributesAvailabilityOnDemand InstancePoolAwsAttributesAvailability = `ON_DEMAND`

const InstancePoolAwsAttributesAvailabilitySpot InstancePoolAwsAttributesAvailability = `SPOT`

const InstancePoolAwsAttributesAvailabilitySpotWithFallback InstancePoolAwsAttributesAvailability = `SPOT_WITH_FALLBACK`

// String representation for [fmt.Print]
func (ipaaa *InstancePoolAwsAttributesAvailability) String() string {
	return string(*ipaaa)
}

// Set raw string value and validate it against allowed values
func (ipaaa *InstancePoolAwsAttributesAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND`, `SPOT`, `SPOT_WITH_FALLBACK`:
		*ipaaa = InstancePoolAwsAttributesAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND", "SPOT", "SPOT_WITH_FALLBACK"`, v)
	}
}

// Type always returns InstancePoolAwsAttributesAvailability to satisfy [pflag.Value] interface
func (ipaaa *InstancePoolAwsAttributesAvailability) Type() string {
	return "InstancePoolAwsAttributesAvailability"
}

type InstancePoolAzureAttributes struct {
	// Shows the Availability type used for the spot nodes.
	//
	// The default value is defined by
	// InstancePoolConf.instancePoolDefaultAzureAvailability
	Availability InstancePoolAzureAttributesAvailability `json:"availability,omitempty"`
	// The default value and documentation here should be kept consistent with
	// CommonConf.defaultSpotBidMaxPrice.
	SpotBidMaxPrice float64 `json:"spot_bid_max_price,omitempty"`
}

// Shows the Availability type used for the spot nodes.
//
// The default value is defined by
// InstancePoolConf.instancePoolDefaultAzureAvailability
type InstancePoolAzureAttributesAvailability string

const InstancePoolAzureAttributesAvailabilityOnDemandAzure InstancePoolAzureAttributesAvailability = `ON_DEMAND_AZURE`

const InstancePoolAzureAttributesAvailabilitySpotAzure InstancePoolAzureAttributesAvailability = `SPOT_AZURE`

const InstancePoolAzureAttributesAvailabilitySpotWithFallbackAzure InstancePoolAzureAttributesAvailability = `SPOT_WITH_FALLBACK_AZURE`

// String representation for [fmt.Print]
func (ipaaa *InstancePoolAzureAttributesAvailability) String() string {
	return string(*ipaaa)
}

// Set raw string value and validate it against allowed values
func (ipaaa *InstancePoolAzureAttributesAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND_AZURE`, `SPOT_AZURE`, `SPOT_WITH_FALLBACK_AZURE`:
		*ipaaa = InstancePoolAzureAttributesAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND_AZURE", "SPOT_AZURE", "SPOT_WITH_FALLBACK_AZURE"`, v)
	}
}

// Type always returns InstancePoolAzureAttributesAvailability to satisfy [pflag.Value] interface
func (ipaaa *InstancePoolAzureAttributesAvailability) Type() string {
	return "InstancePoolAzureAttributesAvailability"
}

type InstancePoolFleetAttributes struct {
	FleetOnDemandOption *FleetOnDemandOption `json:"fleet_on_demand_option,omitempty"`

	FleetSpotOption *FleetSpotOption `json:"fleet_spot_option,omitempty"`

	LaunchTemplateOverrides []FleetLaunchTemplateOverride `json:"launch_template_overrides,omitempty"`
}

type InstancePoolStats struct {
	// Number of active instances in the pool that are NOT part of a cluster.
	IdleCount int `json:"idle_count,omitempty"`
	// Number of pending instances in the pool that are NOT part of a cluster.
	PendingIdleCount int `json:"pending_idle_count,omitempty"`
	// Number of pending instances in the pool that are part of a cluster.
	PendingUsedCount int `json:"pending_used_count,omitempty"`
	// Number of active instances in the pool that are part of a cluster.
	UsedCount int `json:"used_count,omitempty"`
}

type InstancePoolStatus struct {
	// List of error messages for the failed pending instances. The
	// pending_instance_errors follows FIFO with maximum length of the min_idle
	// of the pool. The pending_instance_errors is emptied once the number of
	// exiting available instances reaches the min_idle of the pool.
	PendingInstanceErrors []PendingInstanceError `json:"pending_instance_errors,omitempty"`
}

type ListInstancePools struct {
	InstancePools []InstancePoolAndStats `json:"instance_pools,omitempty"`
}

type PendingInstanceError struct {
	InstanceId string `json:"instance_id,omitempty"`

	Message string `json:"message,omitempty"`
}
