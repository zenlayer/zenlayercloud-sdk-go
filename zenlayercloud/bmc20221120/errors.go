/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2023 All Rights Reserved.
 */
package bmc

const (
	// This product has a unique error code

	// Invalid hostname length. The specified hostname exceeds the maximum length.
	INVALID_PARAMETER_HOSTNAME_EXCEED = "Invalid.Parameter.Hostname.Exceed"

	// Invalid hostname format.
	INVALID_PARAMETER_HOSTNAME_MALFORMED = "Invalid.Parameter.Hostname.Malformed"

	// Invalid instance name length. The specified instance name exceeds the maximum length.
	INVALID_PARAMETER_INSTANCE_NAME_EXCEED = "Invalid.Parameter.Instance.Name.Exceed"

	// Network billing method not supported in current zone.
	OPERATION_FILED_INTERNET_CHARGE_TYPE_NOT_SUPPORT = "Operation.Filed.Internet.Charge.Type.Not.Support"

	// Image not found.
	INVALID_IMAGE_NOT_FOUND = "Invalid.Image.Not.Found"

	// Zone does not exist.
	INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"

	// Instance model not found.
	INVALID_INSTANCE_TYPE_NOT_FOUND = "Invalid.Instance.Type.Not.Found"

	// Custom partition are only supported for specific operating system.
	INVALID_PARTITION_IMAGE_NOT_SET = "Invalid.Partition.Image.Not.Set"

	// You are trying to create more instances than your remaining quota allows.
	OPERATION_DENIED_INSTANCE_QUOTA_EXCEED = "Operation.Denied.Instance.Quota.Exceed"

	// Bandwidth exceeds the maximum limit.
	INVALID_BANDWIDTH_VALUE_EXCEED_MAXIMUM = "Invalid.Bandwidth.Value.Exceed.MaxImum"

	// Invalid password.
	INVALID_PARAMETER_VALUE_PASSWORD_MALFORMED = "Invalid.Parameter.Value.Password.Malformed"

	// Password login and SSH key login cannot be specified at the same time.
	INVALID_PARAMETER_INSTANCE_LOGIN_CONFLICT = "Invalid.Parameter.Instance.Login.Conflict"

	// Invalid SSH key format. Usually start with rsa.
	INVALID_PARAMETER_SSH_KEY_MALFORMED = "Invalid.Parameter.Ssh.Key.Malformed"

	// Custom RAID and rapid RAID cannot be specified at the same time.
	INVALID_RAID_CONFIG_FAST_CUSTOM_CONFLICT = "Invalid.Raid.Config.Fast.Custom.Conflict"

	// RAID level not supported for current instance model.
	INVALID_INSTANCE_TYPE_RAID_NOT_SUPPORT = "Invalid.Instance.Type.Raid.Not.Support"

	// Public NIC name and private NIC name cannot be the same.
	INVALID_PARAMETER_NIC_NAME_CONFLICT = "Invalid.Parameter.Nic.Name.Conflict"

	// Invalid public or private NIC name.
	INVALID_PARAMETER_NIC_NAME_MALFORMED = "Invalid.Parameter.Nic.Name.Malformed"

	// Invalid partition size. Does not reach the specified capacity.
	INVALID_PARTITION_SIZE_NOT_FULL = "Invalid.Partition.Size.Not.Full"

	// Duplicated file path or drive letters of the partition.
	INVALID_PARTITION_DUPLICATE_FILE_PATH = "Invalid.Partition.Duplicate.File.Path"

	// File path (drive letter) is needed.
	// "c" must be included for Windows, and "/" must be included for Linux.
	INVALID_PARTITION_MISSING_REQUIRED_FILE_PATH = "Invalid.Partition.Missing.Required.File.Path"

	// Invalid file type or path format.
	INVALID_PARTITION_MALFORMED = "Invalid.Partition.Malformed"

	// Invalid file type.
	INVALID_PARTITION_NO_TYPE = "Invalid.Partition.No.Type"

	// Drive letters must be filled in alphabetical order, such as CDEFG for Windows.
	INVALID_PARTITION_INVALID_ORDER = "Invalid.Partition.Invalid.Order"

	// Disk quantity does not meet the RAID level requirement in custom RAID configuration.
	INVALID_PARAMETER_VALUE_RAID_DISK_MISMATCH = "Invalid.Parameter.Value.Raid.Disk.Mismatch"

	// Serial numbers of disks must be filled in order, such as [1,2,3], in custom RAID configuration.
	INVALID_PARAMETER_VALUE_RAID_DISK_DISORDER = "Invalid.Parameter.Value.Raid.Disk.Disorder"

	// Duplicated disk serial numbers in custom RAID configuration.
	INVALID_PARAMETER_VALUE_RAID_DISK_SEQUENCE_DUPLICATE = "Invalid.Parameter.Value.Raid.Disk.Sequence.Duplicate"

	// Disk serial number exceeds disk quantity in custom RAID configuration.
	INVALID_PARAMETER_VALUE_RAID_DISK_SEQUENCE_RANGE = "Invalid.Parameter.Value.Raid.Disk.Sequence.Range"

	// Resources not for sale.
	RESOURCE_INSUFFICIENT_PRODUCT_SOLD_OUT = "Resource.Insufficient.Product.Sold.out"

	// Pricing model not supported in current zone.
	OPERATION_DENIED_CHARGE_TYPE_NOT_SUPPORT = "Operation.Denied.Charge.Type.Not.Support"

	// Instance models sold out.
	RESOURCES_SOLDOUT_INSTANCE_TYPE = "Resource.Soldout.Instance.type"

	// Pricing model not supported. If necessary, please contact Zenlayer Support.
	INVALID_CHARGE_TYPE_NOT_SUPPORT = "Invalid.Charge.Type.Not.Support"

	// Subnet not found.
	INVALID_SUBNET_NOT_FOUND = "Invalid.Subnet.Not.Found"

	// Insufficient private IPs in subnet to assign to created instances.
	INVALID_SUBNET_PRIVATE_IP_INSUFFICIENT = "Invalid.Subnet.Private.Ip.Insufficient"

	// Subnet not in the current zone.
	INVALID_SUBNET_ZONE_MISMATCH = "Invalid.Subnet.Zone.Mismatch"

	// Invalid SSH key. Duplicated values.
	INVALID_PARAMETER_SSH_KEY_DUPLICATE = "Invalid.Parameter.Ssh.Key.Duplicate"

	// Instance does not exist.
	INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"

	// Operations on instances in recycle bin are not supported.
	OPERATION_DENIED_INSTANCE_RECYCLED = "Operation.Denied.Instance.Recycled"

	// Only operations on instances in STOPPED status are supported.
	OPERATION_DENIED_INSTANCE_NOT_STOPPED = "Operation.Denied.Instance.Not.Stopped"

	// Only operations on instances in RUNNING status are supported.
	OPERATION_DENIED_INSTANCE_NOT_RUNNING = "Operation.Denied.Instance.Not.Running"

	// Only operations on instances in STOPPED or RUNNING status are supported.
	OPERATION_DENIED_INSTANCE_STATUS = "Operation.Denied.Instance.Status"

	// Operations on instances in CREATING status are not supported.
	OPERATION_DENIED_INSTANCE_CREATING = "Operation.Denied.Instance.Creating"

	// Operations on instances in INSTALLING status are not supported.
	OPERATION_DENIED_INSTANCE_STATUS_INSTALLING = "Operation.Denied.Instance.Status.Installing"

	// Operations on subscription instances are not supported.
	OPERATION_DENIED_INSTANCE_SUBSCRIPTION = "Operation.Denied.Instance.Subscription"

	// Operations failed. Resources in commitment period.
	OPERATION_DENIED_POSTPAID_PROMISE = "Operation.Denied.Postpaid.Promise"

	// Only operations on instances in recycle bin are supported.
	OPERATION_DENIED_INSTANCE_NOT_RECYCLED = "Operation.Denied.Instance.Not.Recycled"

	// Resource group does not exist.
	OPERATION_FAILED_RESOURCE_GROUP_NOT_FOUND = "Operation.Failed.Resource.Group.Not.Found"

	// Resources do not exist.
	OPERATION_FAILED_RESOURCE_NOT_FOUND = "Operation.Failed.Resource.Not.Found"

	// Instance model not for sale in the zone.
	INVALID_INSTANCE_TYPE_ZONE_NO_SELL = "Invalid.Instance.Type.Zone.No.Sell"

	// Public network billing method not supported in the zone.
	INVALID_INSTANCE_BANDWIDTH_ZONE_NO_SELL = "Invalid.Instance.Bandwidth.Zone.No.Sell"

	// Bandwidth exceeds upper limit.
	INVALID_PARAMETER_BANDWIDTH_EXCEED = "Invalid.Parameter.Bandwidth.Exceed"

	// Bandwidth is being modified.
	OPERATION_FAILED_INSTANCE_BANDWIDTH_PROCESSING = "Operation.Failed.Instance.Bandwidth.Processing"

	// InternetChargeType is not ByBandwidth.
	OPERATION_DENIED_INTERNET_CHARGE_TYPE_NOT_SUPPORT = "Operation.Denied.Internet.Charge.Type.Not.Support"

	// No bandwidth downgrade order exists.
	OPERATION_DENIED_DOWNGRADE_NOT_EXIST = "Operation.Denied.Downgrade.Not.Exist"

	// Traffic package is being modified.
	OPERATION_DENIED_INSTANCE_TRAFFIC_PACKAGE_PROCESSING = "Operation.Denied.Instance.Traffic.Package.Processing"

	// Traffic package parameter validation error.
	INVALID_PARAMETER_TRAFFIC_PACKAGE_ERROR = "Invalid.Parameter.Traffic.Package.Error"

	// No traffic package exists.
	OPERATION_FILED_INSTANCE_NOT_EXIST_TRAFFIC_PACKAGE = "Operation.Filed.Instance.Not.Exist.Traffic.Package"

	// Scheduled downgrade order exists.
	OPERATION_FAILED_INSTANCE_EXIST_PLAN_TRAFFIC_PACKAGE = "Operation.Failed.Instance.Exist.Plan.Traffic.Package"

	// Traffic package parameter needs to be greater than or equal to the default value.
	INVALID_PARAMETER_TRAFFIC_PACKAGE_LESS = "Invalid.Parameter.Traffic.Package.Less"

	// Traffic package parameter exceeds upper limit.
	INVALID_PARAMETER_TRAFFIC_PACKAGE_EXCEED = "Invalid.Parameter.Traffic.Package.Exceed"

	// Elastic IPs not available for sale in zone.
	INVALID_EIP_TYPE_ZONE_NO_SELL = "Invalid.Eip.Type.Zone.No.Sell"

	// The parameter is empty.
	MISSING_PARAMETER = "Missing.Parameter"

	// Instances of ESXi operating system are not supported.
	INVALID_ESXI_NOT_SUPPORT = "Invalid.Esxi.Not.Support"

	// Elastic IP does not exist.
	INVALID_EIP_NOT_FOUND = "Invalid.Eip.Not.Found"

	// Operations on a subscription elastic IP are not supported.
	OPERATION_DENIED_EIP_SUBSCRIPTION = "Operation.Denied.Eip.Subscription"

	// Operations on an elastic IP in recycle bin are not supported.
	OPERATION_DENIED_EIP_RECYCLED = "Operation.Denied.Eip.Recycled"

	// Elastic IP not in recycle bin.
	OPERATION_DENIED_EIP_NOT_RECYCLED = "Operation.Denied.Eip.Not.Recycled"

	// Elastic IP and instance should be in the same zone.
	OPERATION_DENIED_EIP_ZONE_NOT_SAME = "Operation.Denied.Eip.Zone.Not.Same"

	// Elastic IP is in recycle bin.
	FAILED_OPERATION_FOR_RECYCLE_RESOURCE = "Failed.Operation.For.Recycle.Resource"

	// Elastic IP status not supported.
	OPERATION_DENIED_EIP_STATUS_NOT_AVAILABLE = "Operation.Denied.Eip.Status.Not.Available"

	// Instance of ESXi operating system cannot bind an elastic IP.
	OPERATION_DENIED_EIP_ESXI_INSTANCE_ASSIGN = "Operation.Denied.Eip.Esxi.Instance.Assign"

	// Number of elastic IPs bound to an instance exceeds the limit.
	OPERATION_DENIED_EIP_INSTANCE_EXCEED_LIMIT = "Operation.Denied.Eip.Instance.Exceed.Limit"

	// Elastic IP status not supported.
	OPERATION_DENIED_EIP_STATUS_NOT_SUPPORT = "Operation.Denied.Eip.Status.Not.Support"

	// DDoS protected IPs not available for sale in zone.
	INVALID_DDOS_IP_TYPE_ZONE_NO_SELL = "Invalid.Ddos.Ip.Type.Zone.No.Sell"

	// DDoS protected IP does not exist.
	INVALID_DDOS_IP_NOT_FOUND = "Invalid.Ddos.Ip.Not.Found"

	// Operations on a subscription DDoS protected IP are not supported.
	OPERATION_DENIED_DDOS_IP_SUBSCRIPTION = "Operation.Denied.Ddos.Ip.Subscription"

	// Operations on a DDoS protected IP in recycle bin are not supported.
	OPERATION_DENIED_DDOS_IP_RECYCLED = "Operation.Denied.Ddos.Ip.Recycled"

	// DDoS protected IP not in recycle bin.
	OPERATION_DENIED_DDOS_IP_NOT_RECYCLED = "Operation.Denied.Ddos.Ip.Not.Recycled"

	// DDoS protected IP and instance should be in the same zone.
	OPERATION_DENIED_DDOS_IP_ZONE_NOT_SAME = "Operation.Denied.Ddos.Ip.Zone.Not.Same"

	// DDoS protected IP status not supported.
	OPERATION_DENIED_DDOS_IP_STATUS_NOT_AVAILABLE = "Operation.Denied.Ddos.Ip.Status.Not.Available"

	// Instance of ESXi operating system cannot bind a DDoS protected IP.
	OPERATION_DENIED_DDOS_IP_ESXI_INSTANCE_ASSIGN = "Operation.Denied.Ddos.Ip.Esxi.Instance.Assign"

	// DDoS protected IP status not supported.
	OPERATION_DENIED_DDOS_IP_STATUS_NOT_SUPPORT = "Operation.Denied.Ddos.Ip.Status.Not.Support"

	// Availability region not found.
	INVALID_VPC_REGION_NOT_FOUND = "Invalid.Vpc.Region.Not.Found"

	// Quantity of VPCs exceeds the limit of the availability region.
	OPERATION_DENIED_VPC_REGION_EXCEED_LIMIT = "Operation.Denied.Vpc.Region.Exceed.Limit"

	// Invalid netmask of CIDR block.
	INVALID_PARAMETER_VPC_LAN_IP_NETMASK = "Invalid.Parameter.Vpc.Lan.Ip.Netmask"

	// Invalid CIDR block.
	INVALID_PARAMETER_VPC_CIDR_BLOCK = "Invalid.Parameter.Vpc.Cidr.Block"

	// Only CIDR block with private IP range is supported.
	INVALID_PARAMETER_VPC_LAN_IP = "Invalid.Parameter.Vpc.Lan.Ip"

	// Only operations on VPCs in the status of AVAILABLE or CREATE_FAILED are supported.
	OPERATION_DENIED_VPC_STATUS_NOT_SUPPORT = "Operation.Denied.Vpc.Status.Not.Support"

	// Operations on VPCs with subnets in it are not supported.
	OPERATION_DENIED_VPC_EXIST_SUBNET_ASSIGN = "Operation.Denied.Vpc.Exist.Subnet.Assign"

	// Operations on VPCs with instances in it are not supported.
	OPERATION_DENIED_VPC_EXIST_INSTANCE = "Operation.Denied.Vpc.Exist.Instance"

	// Subnet not supported in the zone.
	OPERATION_DENIED_ZONE_NOT_SUPPORT_SUBNET = "Operation.Denied.Zone.Not.Support.Subnet"

	// VPC not found.
	INVALID_VPC_NOT_FOUND = "Invalid.Vpc.Not.Found"

	// Invalid availability region ID of VPC.
	OPERATION_DENIED_NO_AVAILABLE_VPC_REGION = "Operation.Denied.No.Available.Vpc.Region"

	// Zone does not belong to the availability region of VPC.
	OPERATION_DENIED_ZONE_NOT_BELONG_VPC = "Operation.Denied.Zone.Not.Belong.Vpc"

	// Invalid CIDR block.
	INVALID_PARAMETER_SUBNET_CIDR_BLOCK = "Invalid.Parameter.Subnet.Cidr.Block"

	// Invalid netmask of CIDR block.
	INVALID_PARAMETER_SUBNET_LAN_IP_NETMASK = "Invalid.Parameter.Subnet.Lan.Ip.Netmask"

	// Only CIDR block with private IP range is supported.
	INVALID_PARAMETER_SUBNET_LAN_IP = "Invalid.Parameter.Subnet.Lan.Ip"

	// Quantity of subnets exceeds the limit of the zone.
	OPERATION_DENIED_SUBNET_EXCEED_LIMIT = "Operation.Denied.Subnet.Exceed.Limit"

	// Quantity of subnets exceeds the limit of the availability region of VPC.
	OPERATION_DENIED_VPC_ZONE_SUBNET_EXCEED_LIMIT = "Operation.Denied.Vpc.Zone.Subnet.Exceed.Limit"

	// Overlapped CIDR blocks of subnets in VPC.
	INVALID_PARAMETER_VPC_SUBNET_OVER_LAP = "Invalid.Parameter.Vpc.Subnet.Over.Lan"

	// CIDR block of subnet not in the IP range of VPC.
	INVALID_PARAMETER_SUBNET_CIDR_NOT_BELONG_VPC = "Invalid.Parameter.Subnet.Cidr.Not.Belong.Vpc"

	// Operations on subnets with instances in it are not supported.
	OPERATION_DENIED_SUBNET_EXIST_INSTANCE = "Operation.Denied.Subnet.Exist.Instance"

	// Subnet status not supported.
	OPERATION_DENIED_SUBNET_STATUS_NOT_SUPPORT = "Operation.Denied.Subnet.Status.Not.Support"

	// Instance is being removed.
	OPERATION_DENIED_SUBNET_ASSOCIATING_INSTANCE = "Operation.Denied.Subnet.Associating.Instance"

	// Operations on adding the instance into a subnet not supported.
	OPERATION_DENIED_INSTANCE_NOT_SUPPORT_SUBNET = "Operation.Denied.Instance.Not.Support.Subnet"

	// Instances can only be added into the subnet in the same zone.
	OPERATION_DENIED_DIFFERENT_ZONE = "Operation.Denied.Different.Zone"

	// Duplicated private IPs assigned to instances.
	INVALID_PARAMETER_DUPLICATE_LAN_IP = "Invalid.Parameter.Duplicate.Lan.Ip"

	// Private IP ended with .1 cannot be assigned.
	INVALID_PARAMETER_LAN_IP_NOT_SUPPORT = "Invalid.Parameter.Lan.Ip.Not.Support"

	// Unavailable private IP. This has already been assigned to an instance.
	OPERATION_DENIED_IP_ASSOCIATED_INSTANCE = "Operation.Denied.Ip.Associated.Instance"

	// Duplicated instances in the subnet.
	OPERATION_DENIED_SUBNET_NOT_REPEAT_INSTANCE = "Operation.Denied.Subnet.Not.Repeat.Instance"

	// Subnet is in other VPC.
	OPERATION_DENIED_SUBNET_ASSOCIATED_OTHER_VPC = "Operation.Denied.Subnet.Associated.Other.Vpc"

	// Unavailable netmask.
	OPERATION_DENIED_UNAVAILABLE_NETMASK = "Operation.Denied.Unavailable.Netmask"

	// Netmask out of stock.
	OPERATION_DENIED_NETMASK_OUT_OF_STOCK = "Operation.Denied.Netmask.Out.Of.Stock"

	// Renewal not supported for this type of CIDR block.
	OPERATION_DENIED_CIDRBLOCK_TYPE_NOT_SUPPORTED = "Operation.Denied.CidrBlock.Type.Not.Support"

	// Return on prepaid CIDR blocks not supported.
	OPERATION_DENIED_CIDRBLOCK_SUBSCRIPTION = "Operation.Denied.CidrBlock.Subscription"

	// Return on CIDR blocks with instances assigned not supported.
	OPERATION_DENIED_INSTANCE_EXIST = "Operation.Denied.Instance.Exist"

	// Only operations on IPv4 CIDR blocks in recycle bin are supported.
	OPERATION_DENIED_CIDRBLOCK_NOT_RECYCLED = "Operation.Denied.CidrBlock.Not.Recycled"

	// CIDR block not found.
	INVALID_CIDRBLOCK_NOT_FOUND = "Invalid.CidrBlock.Not.Found"

	// Operations on CIDR block in recycle bin are not supported.
	OPERATION_DENIED_CIDRBLOCK_RECYCLED = "Operation.Denied.CidrBlock.Recycled"

	// Number of IPs assigned to an instance exceeds the limit.
	OPERATION_DENIED_CIDRBLOCK_IP_COUNT_REACH_LIMIT = "Operation.Denied.CidrBlock.Ip.Count.Reach.Limit"

	// IP to be unassigned does not belong to the CIDR block.
	OPERATION_DENIED_IP_NOT_BELONG = "Operation.Denied.Ip.Not.Belong"

	// IP status not supported.
	OPERATION_DENIED_INVALID_STATUS = "Operation.Denied.Invalid.Status"
)
