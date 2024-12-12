快速开始[(中文)](./README-CN.md)

--- 

# Introduction

Welcome to Zenlayer Cloud API Software Developer Kit (SDK). SDK is a supporting tool for Zenlayer Cloud API. It
currently supports Bare Metal Instance, Elastic IP, DDoS Protected IP and other products. More cloud services will be
supported for SDK.

# Requirements

1. You must use Go 1.9.x or later.
2. A Zenlayer Cloud account is created and an Access Key ID and an Access Key Password are created.
   See [Generate an API Access Key](https://docs.console.zenlayer.com/welcome/platform/team-management/generate-an-api-access-key)
   for more details.

# Installation

Use go get to install SDK：

```shell
$ go get -u gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go
```

# Quick Examples

Take DescribeInstances as an example.

```go
package main

import (
	"fmt"
	bmc "gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go/zenlayercloud/bmc20221120"
	"gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go/zenlayercloud/common"
	"os"
)

func main() {
	client, _ := bmc.NewClientWithSecretKey(os.Getenv("ZENLAYERCLOUD_SECRET_KEY_ID"), os.Getenv("ZENLAYERCLOUD_SECRET_KEY_PASSWORD"))
	request := bmc.NewDescribeInstancesRequest()
	request.PageSize = 1
	request.PageNum = 100
	response, err := client.DescribeInstances(request)
	if _, ok := err.(*common.ZenlayerCloudSdkError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	fmt.Printf("%v\n", response)
}
```

# Other Configurations

Before creating client, if needed, you can specify some other configuration by setting `common.NewConfig`

```go
conf := common.NewConfig()
```

The configurations are as follows:

## Request Timeout

The Zenlayer Cloud SDK for GO has default request timeout, and please do not modify the default value if necessary. You
can check the latest default timeout value in the code. Unit: second. For example, the current default timeout value is
30 seconds.

```go
conf.Timeout = 30
```

You can set `DEBUG=on` to enable the debug mode. The debug mode will print more detailed logs (including request and
response data), which can be enabled when you need to troubleshoot errors in detail. By default, the debug mode is
disabled. You can also set config.Debug = Bool(true) to enable the debug mode as follows:

Default value: `false`

```go
conf.Debug = Bool(true)
```

## Request Retries

If a request fails due to network error, it may be desirable to retry the request. By default, the request retries are
disabled.  
You can enable request reties on all API interfaces or just on a specified one.

```go
package main

import (
	bmc "gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go/zenlayercloud/bmc20221120"
	"gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go/zenlayercloud/common"
	"os"
)

func main() {
	// Open retry for all api invocation and set up retry for 3 times
	config := common.NewConfig()
	config.AutoRetry = true
	config.MaxRetryTime = 3

	client, _ := bmc.NewClient(config, os.Getenv("ZENLAYERCLOUD_SECRET_KEY_ID"), os.Getenv("ZENLAYERCLOUD_SECRET_KEY_PASSWORD"))
	request := bmc.NewDescribeInstanceTypesRequest()
	// Specify retry config for DescribeInstanceTypes
	request.SetAutoRetries(true) // if autoretries not set to true，the retry config will inherent client config
	request.SetMaxAttempts(2)    // specify retry times to 2 replace the client config 3
	response, err = client.DescribeInstanceTypes(request)
}

```

---
快速开始[(中文)](./README-CN.md)
