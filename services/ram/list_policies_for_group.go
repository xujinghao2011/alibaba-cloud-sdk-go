package ram

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

// ListPoliciesForGroup invokes the ram.ListPoliciesForGroup API synchronously
// api document: https://help.aliyun.com/api/ram/listpoliciesforgroup.html
func (client *Client) ListPoliciesForGroup(request *ListPoliciesForGroupRequest) (response *ListPoliciesForGroupResponse, err error) {
	response = CreateListPoliciesForGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListPoliciesForGroupWithChan invokes the ram.ListPoliciesForGroup API asynchronously
// api document: https://help.aliyun.com/api/ram/listpoliciesforgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPoliciesForGroupWithChan(request *ListPoliciesForGroupRequest) (<-chan *ListPoliciesForGroupResponse, <-chan error) {
	responseChan := make(chan *ListPoliciesForGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPoliciesForGroup(request)
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

// ListPoliciesForGroupWithCallback invokes the ram.ListPoliciesForGroup API asynchronously
// api document: https://help.aliyun.com/api/ram/listpoliciesforgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPoliciesForGroupWithCallback(request *ListPoliciesForGroupRequest, callback func(response *ListPoliciesForGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPoliciesForGroupResponse
		var err error
		defer close(result)
		response, err = client.ListPoliciesForGroup(request)
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

// ListPoliciesForGroupRequest is the request struct for api ListPoliciesForGroup
type ListPoliciesForGroupRequest struct {
	*requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
}

// ListPoliciesForGroupResponse is the response struct for api ListPoliciesForGroup
type ListPoliciesForGroupResponse struct {
	*responses.BaseResponse
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	Policies  PoliciesInListPoliciesForGroup `json:"Policies" xml:"Policies"`
}

// CreateListPoliciesForGroupRequest creates a request to invoke ListPoliciesForGroup API
func CreateListPoliciesForGroupRequest() (request *ListPoliciesForGroupRequest) {
	request = &ListPoliciesForGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ListPoliciesForGroup", "ram", "openAPI")
	return
}

// CreateListPoliciesForGroupResponse creates a response to parse from ListPoliciesForGroup response
func CreateListPoliciesForGroupResponse() (response *ListPoliciesForGroupResponse) {
	response = &ListPoliciesForGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
