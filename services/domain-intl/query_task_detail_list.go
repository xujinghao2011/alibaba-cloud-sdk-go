package domain_intl

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

func (client *Client) QueryTaskDetailList(request *QueryTaskDetailListRequest) (response *QueryTaskDetailListResponse, err error) {
	response = CreateQueryTaskDetailListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryTaskDetailListWithChan(request *QueryTaskDetailListRequest) (<-chan *QueryTaskDetailListResponse, <-chan error) {
	responseChan := make(chan *QueryTaskDetailListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTaskDetailList(request)
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

func (client *Client) QueryTaskDetailListWithCallback(request *QueryTaskDetailListRequest, callback func(response *QueryTaskDetailListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTaskDetailListResponse
		var err error
		defer close(result)
		response, err = client.QueryTaskDetailList(request)
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

type QueryTaskDetailListRequest struct {
	*requests.RpcRequest
	TaskStatus   requests.Integer `position:"Query" name:"TaskStatus"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	TaskNo       string           `position:"Query" name:"TaskNo"`
	DomainName   string           `position:"Query" name:"DomainName"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	PageNum      requests.Integer `position:"Query" name:"PageNum"`
}

type QueryTaskDetailListResponse struct {
	*responses.BaseResponse
	RequestId      string           `json:"RequestId" xml:"RequestId"`
	TotalItemNum   requests.Integer `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum requests.Integer `json:"CurrentPageNum" xml:"CurrentPageNum"`
	TotalPageNum   requests.Integer `json:"TotalPageNum" xml:"TotalPageNum"`
	PageSize       requests.Integer `json:"PageSize" xml:"PageSize"`
	PrePage        requests.Boolean `json:"PrePage" xml:"PrePage"`
	NextPage       requests.Boolean `json:"NextPage" xml:"NextPage"`
	Data           struct {
		TaskDetail []struct {
			TaskNo       string           `json:"TaskNo" xml:"TaskNo"`
			TaskDetailNo string           `json:"TaskDetailNo" xml:"TaskDetailNo"`
			TaskType     string           `json:"TaskType" xml:"TaskType"`
			InstanceId   string           `json:"InstanceId" xml:"InstanceId"`
			DomainName   string           `json:"DomainName" xml:"DomainName"`
			TaskStatus   string           `json:"TaskStatus" xml:"TaskStatus"`
			UpdateTime   string           `json:"UpdateTime" xml:"UpdateTime"`
			CreateTime   string           `json:"CreateTime" xml:"CreateTime"`
			TryCount     requests.Integer `json:"TryCount" xml:"TryCount"`
			ErrorMsg     string           `json:"ErrorMsg" xml:"ErrorMsg"`
		} `json:"TaskDetail" xml:"TaskDetail"`
	} `json:"Data" xml:"Data"`
}

func CreateQueryTaskDetailListRequest() (request *QueryTaskDetailListRequest) {
	request = &QueryTaskDetailListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryTaskDetailList", "", "")
	return
}

func CreateQueryTaskDetailListResponse() (response *QueryTaskDetailListResponse) {
	response = &QueryTaskDetailListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
