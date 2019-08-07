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

// DescribeConfigSettings invokes the webplus.DescribeConfigSettings API synchronously
// api document: https://help.aliyun.com/api/webplus/describeconfigsettings.html
func (client *Client) DescribeConfigSettings(request *DescribeConfigSettingsRequest) (response *DescribeConfigSettingsResponse, err error) {
	response = CreateDescribeConfigSettingsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConfigSettingsWithChan invokes the webplus.DescribeConfigSettings API asynchronously
// api document: https://help.aliyun.com/api/webplus/describeconfigsettings.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConfigSettingsWithChan(request *DescribeConfigSettingsRequest) (<-chan *DescribeConfigSettingsResponse, <-chan error) {
	responseChan := make(chan *DescribeConfigSettingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConfigSettings(request)
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

// DescribeConfigSettingsWithCallback invokes the webplus.DescribeConfigSettings API asynchronously
// api document: https://help.aliyun.com/api/webplus/describeconfigsettings.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConfigSettingsWithCallback(request *DescribeConfigSettingsRequest, callback func(response *DescribeConfigSettingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConfigSettingsResponse
		var err error
		defer close(result)
		response, err = client.DescribeConfigSettings(request)
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

// DescribeConfigSettingsRequest is the request struct for api DescribeConfigSettings
type DescribeConfigSettingsRequest struct {
	*requests.RoaRequest
	EnvId      string `position:"Query" name:"EnvId"`
	TemplateId string `position:"Query" name:"TemplateId"`
}

// DescribeConfigSettingsResponse is the response struct for api DescribeConfigSettings
type DescribeConfigSettingsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	ConfigSettings ConfigSettings `json:"ConfigSettings" xml:"ConfigSettings"`
}

// CreateDescribeConfigSettingsRequest creates a request to invoke DescribeConfigSettings API
func CreateDescribeConfigSettingsRequest() (request *DescribeConfigSettingsRequest) {
	request = &DescribeConfigSettingsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DescribeConfigSettings", "/pop/v1/wam/config/configSetting", "webx", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeConfigSettingsResponse creates a response to parse from DescribeConfigSettings response
func CreateDescribeConfigSettingsResponse() (response *DescribeConfigSettingsResponse) {
	response = &DescribeConfigSettingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
