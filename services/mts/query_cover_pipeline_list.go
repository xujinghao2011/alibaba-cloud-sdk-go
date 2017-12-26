package mts

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

func (client *Client) QueryCoverPipelineList(request *QueryCoverPipelineListRequest) (response *QueryCoverPipelineListResponse, err error) {
	response = CreateQueryCoverPipelineListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryCoverPipelineListWithChan(request *QueryCoverPipelineListRequest) (<-chan *QueryCoverPipelineListResponse, <-chan error) {
	responseChan := make(chan *QueryCoverPipelineListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCoverPipelineList(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) QueryCoverPipelineListWithCallback(request *QueryCoverPipelineListRequest, callback func(response *QueryCoverPipelineListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCoverPipelineListResponse
		var err error
		defer close(result)
		response, err = client.QueryCoverPipelineList(request)
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

type QueryCoverPipelineListRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string           `position:"Query" name:"PipelineIds"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryCoverPipelineListResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	NonExistIds struct {
		String []string `json:"String" xml:"String"`
	} `json:"NonExistIds" xml:"NonExistIds"`
	PipelineList struct {
		Pipeline []struct {
			Id           string `json:"Id" xml:"Id"`
			Name         string `json:"Name" xml:"Name"`
			State        string `json:"State" xml:"State"`
			Priority     string `json:"Priority" xml:"Priority"`
			Role         string `json:"Role" xml:"Role"`
			NotifyConfig struct {
				Topic string `json:"Topic" xml:"Topic"`
				Queue string `json:"Queue" xml:"Queue"`
			} `json:"NotifyConfig" xml:"NotifyConfig"`
		} `json:"Pipeline" xml:"Pipeline"`
	} `json:"PipelineList" xml:"PipelineList"`
}

func CreateQueryCoverPipelineListRequest() (request *QueryCoverPipelineListRequest) {
	request = &QueryCoverPipelineListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryCoverPipelineList", "", "")
	return
}

func CreateQueryCoverPipelineListResponse() (response *QueryCoverPipelineListResponse) {
	response = &QueryCoverPipelineListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}