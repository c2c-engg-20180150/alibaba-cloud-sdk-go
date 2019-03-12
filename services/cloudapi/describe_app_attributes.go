package cloudapi

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

// DescribeAppAttributes invokes the cloudapi.DescribeAppAttributes API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeappattributes.html
func (client *Client) DescribeAppAttributes(request *DescribeAppAttributesRequest) (response *DescribeAppAttributesResponse, err error) {
	response = CreateDescribeAppAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppAttributesWithChan invokes the cloudapi.DescribeAppAttributes API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeappattributes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppAttributesWithChan(request *DescribeAppAttributesRequest) (<-chan *DescribeAppAttributesResponse, <-chan error) {
	responseChan := make(chan *DescribeAppAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAppAttributes(request)
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

// DescribeAppAttributesWithCallback invokes the cloudapi.DescribeAppAttributes API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeappattributes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppAttributesWithCallback(request *DescribeAppAttributesRequest, callback func(response *DescribeAppAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppAttributesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAppAttributes(request)
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

// DescribeAppAttributesRequest is the request struct for api DescribeAppAttributes
type DescribeAppAttributesRequest struct {
	*requests.RpcRequest
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	AppId         requests.Integer `position:"Query" name:"AppId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeAppAttributesResponse is the response struct for api DescribeAppAttributes
type DescribeAppAttributesResponse struct {
	*responses.BaseResponse
	RequestId  string                      `json:"RequestId" xml:"RequestId"`
	TotalCount int                         `json:"TotalCount" xml:"TotalCount"`
	PageSize   int                         `json:"PageSize" xml:"PageSize"`
	PageNumber int                         `json:"PageNumber" xml:"PageNumber"`
	Apps       AppsInDescribeAppAttributes `json:"Apps" xml:"Apps"`
}

// CreateDescribeAppAttributesRequest creates a request to invoke DescribeAppAttributes API
func CreateDescribeAppAttributesRequest() (request *DescribeAppAttributesRequest) {
	request = &DescribeAppAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAppAttributes", "apigateway", "openAPI")
	return
}

// CreateDescribeAppAttributesResponse creates a response to parse from DescribeAppAttributes response
func CreateDescribeAppAttributesResponse() (response *DescribeAppAttributesResponse) {
	response = &DescribeAppAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
