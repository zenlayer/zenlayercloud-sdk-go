/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2023 All Rights Reserved.
 */
package vm

const (
	// This product has a unique error code

	// Instance does not exist.
	INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"

	// Shut down the instance before creating an image.
	UNSUPPORTED_OPERATION_INSTANCE_STATE_STARTING = "Unsupported.Operation.Instance.State.Starting"

	// Quantity of images exceeds the upper limit.
	LIMIT_EXCEEDED_IMAGE_QUOTA = "Limit.Exceeded.Image.Quota"

	// Operations not supported in current zone.
	UNSUPPORTED_OPERATION_ZONE_NOT_SUPPORT = "Unsupported.Operation.Zone.Not.Support"

	// Operations on the image with current status are not supported.
	UNSUPPORTED_OPERATION_DISK_UNAVAILABLE = "Unsupported.Operation.Disk.Unavailable"

	// System disk of the instance is being mirrored.
	UNSUPPORTED_OPERATION_DISK_MAKING_IMAGE = "Unsupported.Operation.Disk.Making.Image"

	// Zone not found.
	INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"

	// Zone not found.
	INVALID_REGION_NOT_FOUND = "Invalid.Region.Not.Found"

	// Image not found.
	INVALID_IMAGE_NOT_FOUND = "Invalid.Image.Not.Found"

	// Operations on the image with current status are not supported.
	INVALID_IMAGE_STATUS = "Invalid.Image.Status"

	// Operations on the security group with current status are not supported.
	OPERATION_DENIED_SECURITY_GROUP_STATUS_NOT_AVAILABLE = "Operation.Denied.Security.Group.Status.Not.Available"

	// Security group not found.
	INVALID_SECURITY_GROUP_NOT_FOUND = "Invalid.Security.Group.Not.Found"

	// Operations on the default security groups are not supported.
	OPERATION_DENIED_DEFAULT_SECURITY_GROUP_NOT_SUPPORT = "Operation.Denied.Default.Security.Group.Not.Support"

	// Security group rule ID is not allowed to be passed in.
	INVALID_PARAMETER_SECURITY_GROUP_RULE_ID_NOT_ALLOW = "Invalid.Parameter.Security.Group.Rule.Id.Not.Allow"

	// Invalid value of policy.
	INVALID_PARAMETER_SECURITY_GROUP_POLICY = "Invalid.Parameter.Security.Group.Policy"

	// Invalid value of portRange.
	INVALID_PARAMETER_SECURITY_GROUP_PORT_RANGE = "Invalid.Parameter.Security.Group.Port.Range"

	// Invalid value of priority.
	INVALID_PARAMETER_SECURITY_GROUP_PRIORITY = "Invalid.Parameter.Security.Group.Priority"

	// Invalid value of source cidrIp.
	INVALID_PARAMETER_SECURITY_GROUP_SOURCE_CIDR_IP = "Invalid.Parameter.Security.Group.Source.Cidr.Ip"

	// Duplicated security group rules.
	OPERATION_DENIED_SECURITY_GROUP_EXIST_REPEAT_RULE = "Operation.Denied.Security.Group.Exist.Repeat.Rule"

	// Rule quantity exceeds limit.
	OPERATION_DENIED_SECURITY_GROUP_RULE_EXCEED_LIMIT = "Operation.Denied.Security.Group.Rule.Exceed.Limit"

	// Security group quantity exceeds limit.
	OPERATION_DENIED_SECURITY_GROUP_TEAM_EXCEED_LIMIT = "Operation.Denied.Security.Group.Team.Exceed.Limit"

	// Operations on the security group with current status are not supported.
	OPERATION_DENIED_SECURITY_GROUP_STATUS_NOT_SUPPORT = "Operation.Denied.Security.Group.Status.Not.Support"

	// Operations on security groups with instances in it are not supported.
	OPERATION_DENIED_SECURITY_GROUP_EXIST_INSTANCE = "Operation.Denied.Security.Group.Exist.Instance"

	// Instance is being created.
	OPERATION_DENIED_SECURITY_GROUP_EXIST_PRE_PRODUCT_INSTANCE = "Operation.Denied.Security.Group.Exist.Pre.Product.Instance"
)
