package alidns

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteSubDomainRecords invokes the alidns.DeleteSubDomainRecords API synchronously
// api document: https://help.aliyun.com/api/alidns/deletesubdomainrecords.html
func (client *Client) DeleteSubDomainRecords(request *DeleteSubDomainRecordsRequest) (response *DeleteSubDomainRecordsResponse, err error) {
	response = CreateDeleteSubDomainRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSubDomainRecordsWithChan invokes the alidns.DeleteSubDomainRecords API asynchronously
// api document: https://help.aliyun.com/api/alidns/deletesubdomainrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSubDomainRecordsWithChan(request *DeleteSubDomainRecordsRequest) (<-chan *DeleteSubDomainRecordsResponse, <-chan error) {
	responseChan := make(chan *DeleteSubDomainRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSubDomainRecords(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteSubDomainRecordsWithCallback invokes the alidns.DeleteSubDomainRecords API asynchronously
// api document: https://help.aliyun.com/api/alidns/deletesubdomainrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSubDomainRecordsWithCallback(request *DeleteSubDomainRecordsRequest, callback func(response *DeleteSubDomainRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSubDomainRecordsResponse
		var err error
		defer close(result)
		response, err = client.DeleteSubDomainRecords(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteSubDomainRecordsRequest is the request struct for api DeleteSubDomainRecords
type DeleteSubDomainRecordsRequest struct {
	*requests.RpcRequest
	RR           string `position:"Query" name:"RR"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	Type         string `position:"Query" name:"Type"`
}

// DeleteSubDomainRecordsResponse is the response struct for api DeleteSubDomainRecords
type DeleteSubDomainRecordsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	RR         string `json:"RR" xml:"RR"`
	TotalCount string `json:"TotalCount" xml:"TotalCount"`
}

// CreateDeleteSubDomainRecordsRequest creates a request to invoke DeleteSubDomainRecords API
func CreateDeleteSubDomainRecordsRequest() (request *DeleteSubDomainRecordsRequest) {
	request = &DeleteSubDomainRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DeleteSubDomainRecords", "alidns", "openAPI")
	return
}

// CreateDeleteSubDomainRecordsResponse creates a response to parse from DeleteSubDomainRecords response
func CreateDeleteSubDomainRecordsResponse() (response *DeleteSubDomainRecordsResponse) {
	response = &DeleteSubDomainRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
