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

// UnsubscribeBillToOSS invokes the bssopenapi.UnsubscribeBillToOSS API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/unsubscribebilltooss.html
func (client *Client) UnsubscribeBillToOSS(request *UnsubscribeBillToOSSRequest) (response *UnsubscribeBillToOSSResponse, err error) {
	response = CreateUnsubscribeBillToOSSResponse()
	err = client.DoAction(request, response)
	return
}

// UnsubscribeBillToOSSWithChan invokes the bssopenapi.UnsubscribeBillToOSS API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/unsubscribebilltooss.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnsubscribeBillToOSSWithChan(request *UnsubscribeBillToOSSRequest) (<-chan *UnsubscribeBillToOSSResponse, <-chan error) {
	responseChan := make(chan *UnsubscribeBillToOSSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnsubscribeBillToOSS(request)
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

// UnsubscribeBillToOSSWithCallback invokes the bssopenapi.UnsubscribeBillToOSS API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/unsubscribebilltooss.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnsubscribeBillToOSSWithCallback(request *UnsubscribeBillToOSSRequest, callback func(response *UnsubscribeBillToOSSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnsubscribeBillToOSSResponse
		var err error
		defer close(result)
		response, err = client.UnsubscribeBillToOSS(request)
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

// UnsubscribeBillToOSSRequest is the request struct for api UnsubscribeBillToOSS
type UnsubscribeBillToOSSRequest struct {
	*requests.RpcRequest
	SubscribeType           string `position:"Query" name:"SubscribeType"`
	MultAccountRelSubscribe string `position:"Query" name:"MultAccountRelSubscribe"`
}

// UnsubscribeBillToOSSResponse is the response struct for api UnsubscribeBillToOSS
type UnsubscribeBillToOSSResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUnsubscribeBillToOSSRequest creates a request to invoke UnsubscribeBillToOSS API
func CreateUnsubscribeBillToOSSRequest() (request *UnsubscribeBillToOSSRequest) {
	request = &UnsubscribeBillToOSSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "UnsubscribeBillToOSS", "BssOpenApi", "openAPI")
	return
}

// CreateUnsubscribeBillToOSSResponse creates a response to parse from UnsubscribeBillToOSS response
func CreateUnsubscribeBillToOSSResponse() (response *UnsubscribeBillToOSSResponse) {
	response = &UnsubscribeBillToOSSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
