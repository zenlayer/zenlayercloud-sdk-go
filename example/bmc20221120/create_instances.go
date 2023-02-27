package main

import (
	"encoding/json"
	"fmt"
	bmc "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/bmc20221120"
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
	"os"
)

func main() {
	// create custom config
	config := common.NewConfig()
	// open debug
	config.Debug = common.Bool(true)
	client, _ := bmc.NewClient(config, os.Getenv("ZENLAYERCLOUD_SECRET_KEY_ID"), os.Getenv("ZENLAYERCLOUD_SECRET_KEY_PASSWORD"))
	request := bmc.NewCreateInstancesRequest()
	request.InstanceTypeId = "S8I"
	request.InstanceChargeType = "PREPAID"
	request.InstanceChargePrepaid = &bmc.ChargePrepaid{
		Period: 1,
	}
	request.InternetChargeType = "ByBandwidth"
	request.ZoneId = "SEL-A"

	response, err := client.CreateInstances(request)
	if err != nil {
		fmt.Printf("An API error has returned: %s", err)
	}
	b, _ := json.Marshal(response.Response)
	fmt.Printf("%s", b)
}
