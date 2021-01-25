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

// StartApplication invokes the edas.StartApplication API synchronously
func (client *Client) StartApplication(request *StartApplicationRequest) (response *StartApplicationResponse, err error) {
	response = CreateStartApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// StartApplicationWithChan invokes the edas.StartApplication API asynchronously
func (client *Client) StartApplicationWithChan(request *StartApplicationRequest) (<-chan *StartApplicationResponse, <-chan error) {
	responseChan := make(chan *StartApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartApplication(request)
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

// StartApplicationWithCallback invokes the edas.StartApplication API asynchronously
func (client *Client) StartApplicationWithCallback(request *StartApplicationRequest, callback func(response *StartApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartApplicationResponse
		var err error
		defer close(result)
		response, err = client.StartApplication(request)
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

// StartApplicationRequest is the request struct for api StartApplication
type StartApplicationRequest struct {
	*requests.RoaRequest
	AppId   string `position:"Query" name:"AppId"`
	EccInfo string `position:"Query" name:"EccInfo"`
}

// StartApplicationResponse is the response struct for api StartApplication
type StartApplicationResponse struct {
	*responses.BaseResponse
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateStartApplicationRequest creates a request to invoke StartApplication API
func CreateStartApplicationRequest() (request *StartApplicationRequest) {
	request = &StartApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "StartApplication", "/pop/v5/changeorder/co_start", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartApplicationResponse creates a response to parse from StartApplication response
func CreateStartApplicationResponse() (response *StartApplicationResponse) {
	response = &StartApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
