package zos

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// CreateCommandRequest 创建命令的请求参数。
type CreateCommandRequest struct {
    *common.BaseRequest

    // Name 命令名称。长度不能超过64个字符。
    Name *string `json:"name,omitempty"`

    // Content 命令内容。长度不能超过4096个字符。
    Content *string `json:"content,omitempty"`

    // Type 命令类型。
    Type *string `json:"type,omitempty"`

    // Description 命令描述信息。最长不超过255个字符。
    Description *string `json:"description,omitempty"`

    // ResourceGroupId 命令关联的资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Tags 命令绑定的标签信息。
    Tags *TagAssociation `json:"tags,omitempty"`

}

// TagAssociation 描述创建资源时同时绑定的标签对的信息。
type TagAssociation struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

// Tag 描述一个标签键值对的信息。
type Tag struct {

    // Key 标签键。长度限制：1～64个字符。
    Key *string `json:"key,omitempty"`

    // Value 标签值。长度限制：1～64个字符。
    Value *string `json:"value,omitempty"`

}

// CreateCommandResponseParams 创建命令的响应结果。
type CreateCommandResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CommandId 创建的命令ID。
    CommandId *string `json:"commandId,omitempty"`

}

type CreateCommandResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateCommandResponseParams `json:"response,omitempty"`

}

// DescribeCommandsRequest 
type DescribeCommandsRequest struct {
    *common.BaseRequest

    // CommandIds 根据命令ID进行筛选。最长不超过100个。
    CommandIds []string `json:"commandIds,omitempty"`

    // Name 根据命令名称进行筛选。该字段支持模糊搜索。
    Name *string `json:"name,omitempty"`

    // Types 根据命令类型进行筛选。
    Types []string `json:"types,omitempty"`

    // PageSize 返回的分页大小，默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数，默认为1。
    PageNum *int `json:"pageNum,omitempty"`

    // TagKeys 根据标签键进行搜索。 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。 最长不得超过20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribeCommandsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataSet 命令列表。
    DataSet []*Command `json:"dataSet,omitempty"`

    // TotalCount 匹配条件的命令总数。
    TotalCount *int `json:"totalCount,omitempty"`

}

// Command 描述命令的基本信息。
type Command struct {

    // CommandId 命令ID。
    CommandId *string `json:"commandId,omitempty"`

    // Name 命令的名称。
    Name *string `json:"name,omitempty"`

    // Description 命令的描述信息。
    Description *string `json:"description,omitempty"`

    // Content 命令的内容。
    Content *string `json:"content,omitempty"`

    // Type 命令的类型。
    Type *string `json:"type,omitempty"`

    // CreateTime 命令的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // LatestInvocationTime 最近执行的时间。
    LatestInvocationTime *string `json:"latestInvocationTime,omitempty"`

    // Tags 指令的标签。
    Tags *Tags `json:"tags,omitempty"`

    // ResourceGroup 指令所属的资源组。
    ResourceGroup *ResourceGroupInfo `json:"resourceGroup,omitempty"`

}

// Tags 描述资源关联的标签信息。
type Tags struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

// ResourceGroupInfo 描述资源所在资源组的相关信息，包括资源组名称和ID。
type ResourceGroupInfo struct {

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

}

type DescribeCommandsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCommandsResponseParams `json:"response,omitempty"`

}

// ModifyCommandRequest 修改命令的请求参数。
type ModifyCommandRequest struct {
    *common.BaseRequest

    // CommandId 需要修改的命令ID。
    CommandId *string `json:"commandId,omitempty"`

    // Name 命令名称。长度不能超过64个字符。
    Name *string `json:"name,omitempty"`

    // Content 命令内容。长度不能超过4096个字符。
    Content *string `json:"content,omitempty"`

    // Description 命令的描述信息。长度不能超过255个字符。
    Description *string `json:"description,omitempty"`

}

type ModifyCommandResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteCommandRequest 删除命令的请求。
type DeleteCommandRequest struct {
    *common.BaseRequest

    // CommandId 需要删除的命令ID。
    CommandId *string `json:"commandId,omitempty"`

}

type DeleteCommandResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type InvokeCommandRequest struct {
    *common.BaseRequest

    // CommandId 需要执行的命令ID。
    CommandId *string `json:"commandId,omitempty"`

    // InstanceIds 执行的实例ID列表。目前只支持ZEC实例。最多不得超过200个实例。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // Timeout 每台实例执行命令的超时时间。单位秒。
    Timeout *int `json:"timeout,omitempty"`

}

// InvokeCommandResponseParams 执行命令的响应信息。
type InvokeCommandResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InvocationId 命令执行记录编号。
    InvocationId *string `json:"invocationId,omitempty"`

}

type InvokeCommandResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InvokeCommandResponseParams `json:"response,omitempty"`

}

// DescribeCommandInvocationsRequest 查询执行记录列表的请求信息。
type DescribeCommandInvocationsRequest struct {
    *common.BaseRequest

    // CommandIds 根据命令ID进行筛选。
    CommandIds []string `json:"commandIds,omitempty"`

    // InvocationIds 根据执行记录ID筛选。
    InvocationIds []string `json:"invocationIds,omitempty"`

    // InstanceIds 根据执行的实例ID筛选。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // InvocationStatuses 根据执行的总状态筛选。
    InvocationStatuses []string `json:"invocationStatuses,omitempty"`

    // PageSize 返回的分页大小，默认为20，最大为10000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数，默认为1。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribeCommandInvocationsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataSet 命令执行记录列表。
    DataSet []*CommandInvocation `json:"dataSet,omitempty"`

    // TotalCount 匹配条件的命令执行记录总数。
    TotalCount *int `json:"totalCount,omitempty"`

}

// CommandInvocation 描述命令执行记录的信息。
type CommandInvocation struct {

    // InvocationId 执行ID。
    InvocationId *string `json:"invocationId,omitempty"`

    // CommandId 命令的ID。
    CommandId *string `json:"commandId,omitempty"`

    // CommandName 命令的名称。
    CommandName *string `json:"commandName,omitempty"`

    // CommandType 命令的类型。
    CommandType *string `json:"commandType,omitempty"`

    // Content 执行的内容。
    Content *string `json:"content,omitempty"`

    // Timeout 超时时间。
    Timeout *int `json:"timeout,omitempty"`

    // CreateTime 命令执行的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // InvocationStatus 命令的执行的总执行状态。
    InvocationStatus *string `json:"invocationStatus,omitempty"`

    // InvocationInstances 命令执行实例列表。
    InvocationInstances []*InvocationInstance `json:"invocationInstances,omitempty"`

}

// InvocationInstance 描述执行目标实例的信息。
type InvocationInstance struct {

    // InstanceId 执行ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 实例的名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // InstanceExist 实例是否存在。
    InstanceExist *bool `json:"instanceExist,omitempty"`

    // TaskStatus 实例执行任务的状态。
    TaskStatus *string `json:"taskStatus,omitempty"`

    // ExitCode 命令进程的退出代码。
    ExitCode *int `json:"exitCode,omitempty"`

    // ErrorCode 命令执行失败原因的代号。
    ErrorCode *string `json:"errorCode,omitempty"`

    // Output 命令执行的标准输出(stdout)。
    Output *string `json:"output,omitempty"`

    // OutputError 命令执行的标准错误输出(stderr)。
    OutputError *string `json:"outputError,omitempty"`

    // ExecStartTime 命令执行的开始时间。
    ExecStartTime *string `json:"execStartTime,omitempty"`

    // ExecEndTime 命令执行的结束时间。
    ExecEndTime *string `json:"execEndTime,omitempty"`

}

type DescribeCommandInvocationsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCommandInvocationsResponseParams `json:"response,omitempty"`

}

