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

// DescribeChanges invokes the webplus.DescribeChanges API synchronously
// api document: https://help.aliyun.com/api/webplus/describechanges.html
func (client *Client) DescribeChanges(request *DescribeChangesRequest) (response *DescribeChangesResponse, err error) {
	response = CreateDescribeChangesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeChangesWithChan invokes the webplus.DescribeChanges API asynchronously
// api document: https://help.aliyun.com/api/webplus/describechanges.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeChangesWithChan(request *DescribeChangesRequest) (<-chan *DescribeChangesResponse, <-chan error) {
	responseChan := make(chan *DescribeChangesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeChanges(request)
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

// DescribeChangesWithCallback invokes the webplus.DescribeChanges API asynchronously
// api document: https://help.aliyun.com/api/webplus/describechanges.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeChangesWithCallback(request *DescribeChangesRequest, callback func(response *DescribeChangesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeChangesResponse
		var err error
		defer close(result)
		response, err = client.DescribeChanges(request)
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

// DescribeChangesRequest is the request struct for api DescribeChanges
type DescribeChangesRequest struct {
	*requests.RoaRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	EnvId      string           `position:"Query" name:"EnvId"`
	ActionName string           `position:"Query" name:"ActionName"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeChangesResponse is the response struct for api DescribeChanges
type DescribeChangesResponse struct {
	*responses.BaseResponse
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	PageNumber int     `json:"PageNumber" xml:"PageNumber"`
	PageSize   int     `json:"PageSize" xml:"PageSize"`
	TotalCount int     `json:"TotalCount" xml:"TotalCount"`
	Changes    Changes `json:"Changes" xml:"Changes"`
}

// CreateDescribeChangesRequest creates a request to invoke DescribeChanges API
func CreateDescribeChangesRequest() (request *DescribeChangesRequest) {
	request = &DescribeChangesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DescribeChanges", "/pop/v1/wam/change", "webx", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeChangesResponse creates a response to parse from DescribeChanges response
func CreateDescribeChangesResponse() (response *DescribeChangesResponse) {
	response = &DescribeChangesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
