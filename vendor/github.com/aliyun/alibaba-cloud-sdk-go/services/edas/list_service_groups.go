package edas

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

// ListServiceGroups invokes the edas.ListServiceGroups API synchronously
// api document: https://help.aliyun.com/api/edas/listservicegroups.html
func (client *Client) ListServiceGroups(request *ListServiceGroupsRequest) (response *ListServiceGroupsResponse, err error) {
	response = CreateListServiceGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListServiceGroupsWithChan invokes the edas.ListServiceGroups API asynchronously
// api document: https://help.aliyun.com/api/edas/listservicegroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListServiceGroupsWithChan(request *ListServiceGroupsRequest) (<-chan *ListServiceGroupsResponse, <-chan error) {
	responseChan := make(chan *ListServiceGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListServiceGroups(request)
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

// ListServiceGroupsWithCallback invokes the edas.ListServiceGroups API asynchronously
// api document: https://help.aliyun.com/api/edas/listservicegroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListServiceGroupsWithCallback(request *ListServiceGroupsRequest, callback func(response *ListServiceGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListServiceGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListServiceGroups(request)
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

// ListServiceGroupsRequest is the request struct for api ListServiceGroups
type ListServiceGroupsRequest struct {
	*requests.RoaRequest
}

// ListServiceGroupsResponse is the response struct for api ListServiceGroups
type ListServiceGroupsResponse struct {
	*responses.BaseResponse
	Code              int               `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	ServiceGroupsList ServiceGroupsList `json:"ServiceGroupsList" xml:"ServiceGroupsList"`
}

// CreateListServiceGroupsRequest creates a request to invoke ListServiceGroups API
func CreateListServiceGroupsRequest() (request *ListServiceGroupsRequest) {
	request = &ListServiceGroupsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListServiceGroups", "/pop/v5/service/serviceGroups", "edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListServiceGroupsResponse creates a response to parse from ListServiceGroups response
func CreateListServiceGroupsResponse() (response *ListServiceGroupsResponse) {
	response = &ListServiceGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
