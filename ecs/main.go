package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func NewClient(regionId, accessKeyId, accessKeySecret string) *ecs.Client {
	client, err := ecs.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		panic(fmt.Sprintf("[ECS]: create client failed, %s", err))
	}
	return client
}

func GetEcsByIp(client *ecs.Client, ip string, vpc bool) {
	request := ecs.CreateDescribeInstancesRequest()
	if vpc {
		request.PrivateIpAddresses = fmt.Sprintf(`["%s"]`, ip)
	} else {
		request.InnerIpAddresses = fmt.Sprintf(`["%s"]`, ip)
	}

	response, err := client.DescribeInstances(request)
	if err != nil {
		panic(fmt.Sprintf("[ECS]: can not find %s instance, %s", ip, err))
	}

	for _, instance := range response.Instances.Instance {
		if vpc {
			fmt.Println(instance.InstanceId, instance.InstanceName, instance.VpcAttributes.PrivateIpAddress.IpAddress, instance.Status)
		} else {
			fmt.Println(instance.InstanceId, instance.InstanceName, instance.InnerIpAddress.IpAddress, instance.Status)
		}
	}
}

func GetStatusById(client *ecs.Client, id string) string {
	request := ecs.CreateDescribeInstanceAttributeRequest()
	request.InstanceId = id
	response, err := client.DescribeInstanceAttribute(request)
	if err != nil {
		panic(fmt.Sprintf("[ECS]: can not find %s instance, %s", id, err))
	}
	return response.Status
}

func RebootEcsById(client *ecs.Client, id string, force bool) {
	request := ecs.CreateRebootInstanceRequest()
	request.InstanceId = id
	request.ForceStop = requests.NewBoolean(force)
	_, err := client.RebootInstance(request)
	if err != nil {
		panic(fmt.Sprintf("[ECS]: can not reboot %s instance, %s", id, err))
	}
	for {
		status := GetStatusById(client, id)
		fmt.Println(fmt.Sprintf("%s status: %s", id, status))
		if status == "Running" {
			break
		} else {
			time.Sleep(time.Second * 5)
		}
	}
}

func main() {
	accessKeyId := flag.String("access-key-id", "", "aliyun access key id")
	accessKeySecret := flag.String("access-key-secret", "", "aliyun access key secret")
	regionId := flag.String("region-id", "cn-hangzhou", "aliyun region id")
	ip := flag.String("query-by-ip", "", "query ecs by ip")
	action := flag.String("action", "query", "query or reboot")
	id := flag.String("instance-id", "", "ecs instance id")
	vpc := flag.Bool("vpc", false, "is vpc network")
	force := flag.Bool("force-reboot", false, "is force reboot")
	flag.Parse()

	if *accessKeyId == "" {
		panic("access-key-id is required")
	}
	if *accessKeySecret == "" {
		panic("access-key-secret is required")
	}

	client := NewClient(*regionId, *accessKeyId, *accessKeySecret)

	if *action == "query" {
		if *ip == "" {
			panic("query-by-ip is required")
		} else {
			GetEcsByIp(client, *ip, *vpc)
		}
	} else if *action == "reboot" {
		if *id == "" {
			panic("instance-id is required")
		} else {
			RebootEcsById(client, *id, *force)
		}
	}
}
