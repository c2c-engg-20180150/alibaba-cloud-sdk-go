package ons

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

// OnsMqttQueryClientByClientId invokes the ons.OnsMqttQueryClientByClientId API synchronously
// api document: https://help.aliyun.com/api/ons/onsmqttqueryclientbyclientid.html
func (client *Client) OnsMqttQueryClientByClientId(request *OnsMqttQueryClientByClientIdRequest) (response *OnsMqttQueryClientByClientIdResponse, err error) {
	response = CreateOnsMqttQueryClientByClientIdResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMqttQueryClientByClientIdWithChan invokes the ons.OnsMqttQueryClientByClientId API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmqttqueryclientbyclientid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMqttQueryClientByClientIdWithChan(request *OnsMqttQueryClientByClientIdRequest) (<-chan *OnsMqttQueryClientByClientIdResponse, <-chan error) {
	responseChan := make(chan *OnsMqttQueryClientByClientIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMqttQueryClientByClientId(request)
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

// OnsMqttQueryClientByClientIdWithCallback invokes the ons.OnsMqttQueryClientByClientId API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmqttqueryclientbyclientid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMqttQueryClientByClientIdWithCallback(request *OnsMqttQueryClientByClientIdRequest, callback func(response *OnsMqttQueryClientByClientIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMqttQueryClientByClientIdResponse
		var err error
		defer close(result)
		response, err = client.OnsMqttQueryClientByClientId(request)
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

// OnsMqttQueryClientByClientIdRequest is the request struct for api OnsMqttQueryClientByClientId
type OnsMqttQueryClientByClientIdRequest struct {
	*requests.RpcRequest
	PreventCache requests.Integer `position:"Query" name:"PreventCache"`
	ClientId     string           `position:"Query" name:"ClientId"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
}

// OnsMqttQueryClientByClientIdResponse is the response struct for api OnsMqttQueryClientByClientId
type OnsMqttQueryClientByClientIdResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	HelpUrl          string           `json:"HelpUrl" xml:"HelpUrl"`
	MqttClientInfoDo MqttClientInfoDo `json:"MqttClientInfoDo" xml:"MqttClientInfoDo"`
}

// CreateOnsMqttQueryClientByClientIdRequest creates a request to invoke OnsMqttQueryClientByClientId API
func CreateOnsMqttQueryClientByClientIdRequest() (request *OnsMqttQueryClientByClientIdRequest) {
	request = &OnsMqttQueryClientByClientIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMqttQueryClientByClientId", "ons", "openAPI")
	return
}

// CreateOnsMqttQueryClientByClientIdResponse creates a response to parse from OnsMqttQueryClientByClientId response
func CreateOnsMqttQueryClientByClientIdResponse() (response *OnsMqttQueryClientByClientIdResponse) {
	response = &OnsMqttQueryClientByClientIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
