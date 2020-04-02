package csb

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

// FindServiceStatisticalData invokes the csb.FindServiceStatisticalData API synchronously
// api document: https://help.aliyun.com/api/csb/findservicestatisticaldata.html
func (client *Client) FindServiceStatisticalData(request *FindServiceStatisticalDataRequest) (response *FindServiceStatisticalDataResponse, err error) {
	response = CreateFindServiceStatisticalDataResponse()
	err = client.DoAction(request, response)
	return
}

// FindServiceStatisticalDataWithChan invokes the csb.FindServiceStatisticalData API asynchronously
// api document: https://help.aliyun.com/api/csb/findservicestatisticaldata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindServiceStatisticalDataWithChan(request *FindServiceStatisticalDataRequest) (<-chan *FindServiceStatisticalDataResponse, <-chan error) {
	responseChan := make(chan *FindServiceStatisticalDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindServiceStatisticalData(request)
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

// FindServiceStatisticalDataWithCallback invokes the csb.FindServiceStatisticalData API asynchronously
// api document: https://help.aliyun.com/api/csb/findservicestatisticaldata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindServiceStatisticalDataWithCallback(request *FindServiceStatisticalDataRequest, callback func(response *FindServiceStatisticalDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindServiceStatisticalDataResponse
		var err error
		defer close(result)
		response, err = client.FindServiceStatisticalData(request)
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

// FindServiceStatisticalDataRequest is the request struct for api FindServiceStatisticalData
type FindServiceStatisticalDataRequest struct {
	*requests.RpcRequest
	CsbId       requests.Integer `position:"Query" name:"CsbId"`
	EndTime     requests.Integer `position:"Query" name:"EndTime"`
	StartTime   requests.Integer `position:"Query" name:"StartTime"`
	ServiceName string           `position:"Query" name:"ServiceName"`
}

// FindServiceStatisticalDataResponse is the response struct for api FindServiceStatisticalData
type FindServiceStatisticalDataResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFindServiceStatisticalDataRequest creates a request to invoke FindServiceStatisticalData API
func CreateFindServiceStatisticalDataRequest() (request *FindServiceStatisticalDataRequest) {
	request = &FindServiceStatisticalDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "FindServiceStatisticalData", "", "")
	return
}

// CreateFindServiceStatisticalDataResponse creates a response to parse from FindServiceStatisticalData response
func CreateFindServiceStatisticalDataResponse() (response *FindServiceStatisticalDataResponse) {
	response = &FindServiceStatisticalDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
