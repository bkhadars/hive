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

// DeleteDomainGroup invokes the alidns.DeleteDomainGroup API synchronously
func (client *Client) DeleteDomainGroup(request *DeleteDomainGroupRequest) (response *DeleteDomainGroupResponse, err error) {
	response = CreateDeleteDomainGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDomainGroupWithChan invokes the alidns.DeleteDomainGroup API asynchronously
func (client *Client) DeleteDomainGroupWithChan(request *DeleteDomainGroupRequest) (<-chan *DeleteDomainGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteDomainGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDomainGroup(request)
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

// DeleteDomainGroupWithCallback invokes the alidns.DeleteDomainGroup API asynchronously
func (client *Client) DeleteDomainGroupWithCallback(request *DeleteDomainGroupRequest, callback func(response *DeleteDomainGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDomainGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteDomainGroup(request)
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

// DeleteDomainGroupRequest is the request struct for api DeleteDomainGroup
type DeleteDomainGroupRequest struct {
	*requests.RpcRequest
	GroupId      string `position:"Query" name:"GroupId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DeleteDomainGroupResponse is the response struct for api DeleteDomainGroup
type DeleteDomainGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	GroupName string `json:"GroupName" xml:"GroupName"`
}

// CreateDeleteDomainGroupRequest creates a request to invoke DeleteDomainGroup API
func CreateDeleteDomainGroupRequest() (request *DeleteDomainGroupRequest) {
	request = &DeleteDomainGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DeleteDomainGroup", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDomainGroupResponse creates a response to parse from DeleteDomainGroup response
func CreateDeleteDomainGroupResponse() (response *DeleteDomainGroupResponse) {
	response = &DeleteDomainGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
