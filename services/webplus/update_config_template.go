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

// UpdateConfigTemplate invokes the webplus.UpdateConfigTemplate API synchronously
// api document: https://help.aliyun.com/api/webplus/updateconfigtemplate.html
func (client *Client) UpdateConfigTemplate(request *UpdateConfigTemplateRequest) (response *UpdateConfigTemplateResponse, err error) {
	response = CreateUpdateConfigTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateConfigTemplateWithChan invokes the webplus.UpdateConfigTemplate API asynchronously
// api document: https://help.aliyun.com/api/webplus/updateconfigtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateConfigTemplateWithChan(request *UpdateConfigTemplateRequest) (<-chan *UpdateConfigTemplateResponse, <-chan error) {
	responseChan := make(chan *UpdateConfigTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateConfigTemplate(request)
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

// UpdateConfigTemplateWithCallback invokes the webplus.UpdateConfigTemplate API asynchronously
// api document: https://help.aliyun.com/api/webplus/updateconfigtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateConfigTemplateWithCallback(request *UpdateConfigTemplateRequest, callback func(response *UpdateConfigTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateConfigTemplateResponse
		var err error
		defer close(result)
		response, err = client.UpdateConfigTemplate(request)
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

// UpdateConfigTemplateRequest is the request struct for api UpdateConfigTemplate
type UpdateConfigTemplateRequest struct {
	*requests.RoaRequest
	OptionSettings      string `position:"Body" name:"OptionSettings"`
	TemplateId          string `position:"Body" name:"TemplateId"`
	TemplateDescription string `position:"Body" name:"TemplateDescription"`
}

// UpdateConfigTemplateResponse is the response struct for api UpdateConfigTemplate
type UpdateConfigTemplateResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	ConfigTemplate ConfigTemplate `json:"ConfigTemplate" xml:"ConfigTemplate"`
}

// CreateUpdateConfigTemplateRequest creates a request to invoke UpdateConfigTemplate API
func CreateUpdateConfigTemplateRequest() (request *UpdateConfigTemplateRequest) {
	request = &UpdateConfigTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "UpdateConfigTemplate", "/pop/v1/wam/configTemplate", "webx", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateConfigTemplateResponse creates a response to parse from UpdateConfigTemplate response
func CreateUpdateConfigTemplateResponse() (response *UpdateConfigTemplateResponse) {
	response = &UpdateConfigTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
