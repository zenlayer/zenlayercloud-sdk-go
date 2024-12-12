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
