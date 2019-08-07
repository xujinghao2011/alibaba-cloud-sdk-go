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

// DeletePkgVersion invokes the webplus.DeletePkgVersion API synchronously
// api document: https://help.aliyun.com/api/webplus/deletepkgversion.html
func (client *Client) DeletePkgVersion(request *DeletePkgVersionRequest) (response *DeletePkgVersionResponse, err error) {
	response = CreateDeletePkgVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePkgVersionWithChan invokes the webplus.DeletePkgVersion API asynchronously
// api document: https://help.aliyun.com/api/webplus/deletepkgversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePkgVersionWithChan(request *DeletePkgVersionRequest) (<-chan *DeletePkgVersionResponse, <-chan error) {
	responseChan := make(chan *DeletePkgVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePkgVersion(request)
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

// DeletePkgVersionWithCallback invokes the webplus.DeletePkgVersion API asynchronously
// api document: https://help.aliyun.com/api/webplus/deletepkgversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePkgVersionWithCallback(request *DeletePkgVersionRequest, callback func(response *DeletePkgVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePkgVersionResponse
		var err error
		defer close(result)
		response, err = client.DeletePkgVersion(request)
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

// DeletePkgVersionRequest is the request struct for api DeletePkgVersion
type DeletePkgVersionRequest struct {
	*requests.RoaRequest
	PkgVersionId string `position:"Query" name:"PkgVersionId"`
}

// DeletePkgVersionResponse is the response struct for api DeletePkgVersion
type DeletePkgVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeletePkgVersionRequest creates a request to invoke DeletePkgVersion API
func CreateDeletePkgVersionRequest() (request *DeletePkgVersionRequest) {
	request = &DeletePkgVersionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DeletePkgVersion", "/pop/v1/wam/pkgVersion", "webx", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeletePkgVersionResponse creates a response to parse from DeletePkgVersion response
func CreateDeletePkgVersionResponse() (response *DeletePkgVersionResponse) {
	response = &DeletePkgVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
