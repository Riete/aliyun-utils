package main

import (
	"flag"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

func NewClient(regionId, accessKeyId, accessKeySecret string) *alidns.Client {
	client, err := alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		panic(fmt.Sprintf("[DNS]: create client failed, %s", err))
	}
	return client
}

func QueryDomainARecord(client *alidns.Client, domainName, rr string) {
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = domainName
	request.RRKeyWord = rr
	request.TypeKeyWord = "A"
	response, err := client.DescribeDomainRecords(request)
	if err != nil {
		panic(fmt.Sprintf("[DNS]: query A record failed, %s", err))
	}
	record := response.DomainRecords.Record[0]
	fmt.Println(fmt.Sprintf("%s.%s %s", record.RR, record.DomainName, record.Value))
}

func NewDomainARecord(client *alidns.Client, domainName, value, rr string) {
	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = domainName
	request.Type = "A"
	request.Value = value
	request.RR = rr
	_, err := client.AddDomainRecord(request)
	if err != nil {
		panic(fmt.Sprintf("[DNS]: add A record failed, %s", err))
	}
	QueryDomainARecord(client, domainName, rr)
}

func main() {
	accessKeyId := flag.String("access-key-id", "", "aliyun access key id")
	accessKeySecret := flag.String("access-key-secret", "", "aliyun access key secret")
	regionId := flag.String("region-id", "cn-hangzhou", "aliyun region id")
	action := flag.String("action", "query", "query or new")
	rr := flag.String("record", "", "A record name without domain")
	value := flag.String("value", "", "A record value")
	domainName := flag.String("domain-name", "", "domain name")
	flag.Parse()

	if *rr == "" {
		panic("record is required")
	}
	if *accessKeyId == "" {
		panic("access-key-id is required")
	}
	if *accessKeySecret == "" {
		panic("access-key-secret is required")
	}
	if *domainName == "" {
		panic("domain-name is required")
	}

	client := NewClient(*regionId, *accessKeyId, *accessKeySecret)

	if *action == "query" {
		QueryDomainARecord(client, *domainName, *rr)
	} else if *action == "new" {
		if *value == "" {
			panic("record value is required")
		}
		NewDomainARecord(client, *domainName, *value, *rr)
	}
}
