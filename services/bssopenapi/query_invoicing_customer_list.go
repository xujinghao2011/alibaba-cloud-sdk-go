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

// QueryInvoicingCustomerList invokes the bssopenapi.QueryInvoicingCustomerList API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryinvoicingcustomerlist.html
func (client *Client) QueryInvoicingCustomerList(request *QueryInvoicingCustomerListRequest) (response *QueryInvoicingCustomerListResponse, err error) {
	response = CreateQueryInvoicingCustomerListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryInvoicingCustomerListWithChan invokes the bssopenapi.QueryInvoicingCustomerList API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryinvoicingcustomerlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryInvoicingCustomerListWithChan(request *QueryInvoicingCustomerListRequest) (<-chan *QueryInvoicingCustomerListResponse, <-chan error) {
	responseChan := make(chan *QueryInvoicingCustomerListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInvoicingCustomerList(request)
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

// QueryInvoicingCustomerListWithCallback invokes the bssopenapi.QueryInvoicingCustomerList API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryinvoicingcustomerlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryInvoicingCustomerListWithCallback(request *QueryInvoicingCustomerListRequest, callback func(response *QueryInvoicingCustomerListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInvoicingCustomerListResponse
		var err error
		defer close(result)
		response, err = client.QueryInvoicingCustomerList(request)
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

// QueryInvoicingCustomerListRequest is the request struct for api QueryInvoicingCustomerList
type QueryInvoicingCustomerListRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryInvoicingCustomerListResponse is the response struct for api QueryInvoicingCustomerList
type QueryInvoicingCustomerListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryInvoicingCustomerListRequest creates a request to invoke QueryInvoicingCustomerList API
func CreateQueryInvoicingCustomerListRequest() (request *QueryInvoicingCustomerListRequest) {
	request = &QueryInvoicingCustomerListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryInvoicingCustomerList", "BssOpenApi", "openAPI")
	return
}

// CreateQueryInvoicingCustomerListResponse creates a response to parse from QueryInvoicingCustomerList response
func CreateQueryInvoicingCustomerListResponse() (response *QueryInvoicingCustomerListResponse) {
	response = &QueryInvoicingCustomerListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
