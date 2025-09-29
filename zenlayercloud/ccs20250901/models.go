package ccs

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// DescribeKeyPairsRequest 查询密钥对列表的请求参数。
type DescribeKeyPairsRequest struct {
    *common.BaseRequest

    // KeyIds 根据密钥对ID列表进行筛选。
    KeyIds []string `json:"keyIds,omitempty"`

    // KeyName 根据密钥对名称进行筛选。该字段支持模糊搜索。
    KeyName *string `json:"keyName,omitempty"`

    // PageSize 返回的分页大小。默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。默认为1。
    PageNum *int `json:"pageNum,omitempty"`

}

// DescribeKeyPairsResponseParams 描述了密钥对的基本信息的响应信息。
type DescribeKeyPairsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 密钥对结果集。
    DataSet []*KeyPair `json:"dataSet,omitempty"`

}

// KeyPair 描述了密钥对的基本信息。
type KeyPair struct {

    // KeyId 密钥对ID。示例值：key-xxxxxxxx。
    KeyId *string `json:"keyId,omitempty"`

    // KeyName 密钥对名称。
    KeyName *string `json:"keyName,omitempty"`

    // PublicKey 密钥对的公钥内容。OpenSSH 格式。示例值：ssh-rsa XXXXXXXXXXXX key_xxx。
    PublicKey *string `json:"publicKey,omitempty"`

    // KeyDescription 密钥对描述信息。
    KeyDescription *string `json:"keyDescription,omitempty"`

    // CreateTime 创建时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

}

type DescribeKeyPairsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeKeyPairsResponseParams `json:"response,omitempty"`

}

// ImportKeyPairRequest 导入密钥对的请求参数。
type ImportKeyPairRequest struct {
    *common.BaseRequest

    // KeyName 密钥对名称。可由数字，大小写字母和-组成，长度不超过32个字符，不能和账号下其他的密钥对重名。示例值：my_key, my-key。
    KeyName *string `json:"keyName,omitempty"`

    // KeyDescription 密钥对描述信息。不超过255个字符。
    KeyDescription *string `json:"keyDescription,omitempty"`

    // PublicKey 密钥对的公钥内容，OpenSSH 格式。公钥内容至多传5个公钥，通过分隔。示例值：ssh-rsa XXXXXXXXXXXX key。
    PublicKey *string `json:"publicKey,omitempty"`

}

// ImportKeyPairResponseParams 导入密钥对的响应信息。
type ImportKeyPairResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // KeyId 密钥对ID。
    KeyId *string `json:"keyId,omitempty"`

}

type ImportKeyPairResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *ImportKeyPairResponseParams `json:"response,omitempty"`

}

// ModifyKeyPairAttributeRequest 修改一个密钥对的属性的请求参数。
type ModifyKeyPairAttributeRequest struct {
    *common.BaseRequest

    // KeyId 密钥对ID。示例值：key-xxxxxxxx。
    KeyId *string `json:"keyId,omitempty"`

    // KeyDescription 密钥对描述信息。不超过255个字符。
    KeyDescription *string `json:"keyDescription,omitempty"`

}

type ModifyKeyPairAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteKeyPairsRequest 删除一个或多个密钥对的响应结果。
type DeleteKeyPairsRequest struct {
    *common.BaseRequest

    // KeyIds 一个或多个待操作的密钥对ID。每次请求批量密钥对的上限为100。
    KeyIds []string `json:"keyIds,omitempty"`

}

type DeleteKeyPairsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

