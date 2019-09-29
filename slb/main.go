package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

func NewClient(regionId, accessKeyId, accessKeySecret string) *slb.Client {
	client, err := slb.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		panic(fmt.Sprintf("[SLB]: create client failed, %s", err))
	}
	return client
}

func GetSlbInfo(client *slb.Client) []slb.LoadBalancer {
	request := slb.CreateDescribeLoadBalancersRequest()
	response, err := client.DescribeLoadBalancers(request)
	if err != nil {
		panic("[SLB]: can not get slb info")
	}
	return response.LoadBalancers.LoadBalancer
}

func GetSlbByIp(client *slb.Client, ip string, lbs []slb.LoadBalancer) {
	for _, s := range lbs {
		if strings.Contains(s.Address, ip) {
			fmt.Println(s.LoadBalancerId, s.Address, s.LoadBalancerName)
			break
		}
	}
}

func GetSlbByName(client *slb.Client, name string, lbs []slb.LoadBalancer) {
	for _, s := range lbs {
		sName := strings.ToLower(s.LoadBalancerName)
		if strings.Contains(sName, strings.ToLower(name)) {
			fmt.Println(s.LoadBalancerId, s.Address, s.LoadBalancerName)
		}
	}
}

func GetSlbDetailById(client *slb.Client, id string) {
	request := slb.CreateDescribeLoadBalancerAttributeRequest()
	request.LoadBalancerId = id
	response, err := client.DescribeLoadBalancerAttribute(request)
	if err != nil {
		panic(fmt.Sprintf("[SLB]: can not get %s slb info", id))
	}

	fmt.Println(response.LoadBalancerId, response.Address, response.LoadBalancerName)
	for _, listener := range response.ListenerPortsAndProtocol.ListenerPortAndProtocol {
		fmt.Println(listener.ListenerPort, listener.ListenerProtocol, listener.Description)
	}
}

func main() {
	accessKeyId := flag.String("access-key-id", "", "aliyun access key id")
	accessKeySecret := flag.String("access-key-secret", "", "aliyun access key secret")
	regionId := flag.String("region-id", "cn-hangzhou", "aliyun region id")
	queryType := flag.String("query-type", "", "ip|name|id")
	queryValue := flag.String("query-value", "", "query value")
	flag.Parse()

	if *accessKeyId == "" {
		panic("access-key-id is required")
	}
	if *accessKeySecret == "" {
		panic("access-key-secret is required")
	}
	if *queryType == "" {
		panic("query-type is required")
	}
	if *queryValue == "" {
		panic("query-value is required")
	}

	client := NewClient(*regionId, *accessKeyId, *accessKeySecret)

	if *queryType == "name" {
		lbs := GetSlbInfo(client)
		GetSlbByName(client, *queryValue, lbs)
	} else if *queryType == "ip" {
		lbs := GetSlbInfo(client)
		GetSlbByIp(client, *queryValue, lbs)
	} else if *queryType == "id" {
		GetSlbDetailById(client, *queryValue)
	}
}
