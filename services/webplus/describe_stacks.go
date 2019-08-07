package webplus

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

// DescribeStacks invokes the webplus.DescribeStacks API synchronously
// api document: https://help.aliyun.com/api/webplus/describestacks.html
func (client *Client) DescribeStacks(request *DescribeStacksRequest) (response *DescribeStacksResponse, err error) {
	response = CreateDescribeStacksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStacksWithChan invokes the webplus.DescribeStacks API asynchronously
// api document: https://help.aliyun.com/api/webplus/describestacks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStacksWithChan(request *DescribeStacksRequest) (<-chan *DescribeStacksResponse, <-chan error) {
	responseChan := make(chan *DescribeStacksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStacks(request)
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

// DescribeStacksWithCallback invokes the webplus.DescribeStacks API asynchronously
// api document: https://help.aliyun.com/api/webplus/describestacks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStacksWithCallback(request *DescribeStacksRequest, callback func(response *DescribeStacksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStacksResponse
		var err error
		defer close(result)
		response, err = client.DescribeStacks(request)
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

// DescribeStacksRequest is the request struct for api DescribeStacks
type DescribeStacksRequest struct {
	*requests.RoaRequest
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CategoryName    string           `position:"Query" name:"CategoryName"`
	RecommendedOnly requests.Boolean `position:"Query" name:"RecommendedOnly"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeStacksResponse is the response struct for api DescribeStacks
type DescribeStacksResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	Stacks     Stacks `json:"Stacks" xml:"Stacks"`
}

// CreateDescribeStacksRequest creates a request to invoke DescribeStacks API
func CreateDescribeStacksRequest() (request *DescribeStacksRequest) {
	request = &DescribeStacksRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DescribeStacks", "/pop/v1/wam/stack", "webx", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeStacksResponse creates a response to parse from DescribeStacks response
func CreateDescribeStacksResponse() (response *DescribeStacksResponse) {
	response = &DescribeStacksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
