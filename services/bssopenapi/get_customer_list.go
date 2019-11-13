package bssopenapi

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

// GetCustomerList invokes the bssopenapi.GetCustomerList API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/getcustomerlist.html
func (client *Client) GetCustomerList(request *GetCustomerListRequest) (response *GetCustomerListResponse, err error) {
	response = CreateGetCustomerListResponse()
	err = client.DoAction(request, response)
	return
}

// GetCustomerListWithChan invokes the bssopenapi.GetCustomerList API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/getcustomerlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCustomerListWithChan(request *GetCustomerListRequest) (<-chan *GetCustomerListResponse, <-chan error) {
	responseChan := make(chan *GetCustomerListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCustomerList(request)
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

// GetCustomerListWithCallback invokes the bssopenapi.GetCustomerList API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/getcustomerlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCustomerListWithCallback(request *GetCustomerListRequest, callback func(response *GetCustomerListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCustomerListResponse
		var err error
		defer close(result)
		response, err = client.GetCustomerList(request)
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

// GetCustomerListRequest is the request struct for api GetCustomerList
type GetCustomerListRequest struct {
	*requests.RpcRequest
}

// GetCustomerListResponse is the response struct for api GetCustomerList
type GetCustomerListResponse struct {
	*responses.BaseResponse
	RequestId string                `json:"RequestId" xml:"RequestId"`
	Success   bool                  `json:"Success" xml:"Success"`
	Code      string                `json:"Code" xml:"Code"`
	Message   string                `json:"Message" xml:"Message"`
	Data      DataInGetCustomerList `json:"Data" xml:"Data"`
}

// CreateGetCustomerListRequest creates a request to invoke GetCustomerList API
func CreateGetCustomerListRequest() (request *GetCustomerListRequest) {
	request = &GetCustomerListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "GetCustomerList", "BssOpenApi", "openAPI")
	return
}

// CreateGetCustomerListResponse creates a response to parse from GetCustomerList response
func CreateGetCustomerListResponse() (response *GetCustomerListResponse) {
	response = &GetCustomerListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
