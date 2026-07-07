package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response interface {
	ParseErrorResponse(body []byte) error
}

type BaseResponse struct {
}

type ErrorResponse struct {
	RequestId string `json:"requestId"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

func (r *BaseResponse) ParseErrorResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json body: %s, error: %s", body, err)
		return NewZenlayerCloudSdkError(JsonParseError, msg, "")
	}
	return NewZenlayerCloudSdkError(resp.Code, resp.Message, resp.RequestId)
}

func ParseFromHttpResponse(hr *http.Response, response Response) (err error) {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body, error %s", err)
		return NewZenlayerCloudSdkError(IoError, msg, "")
	}

	if hr.StatusCode != 200 {
		if hr.StatusCode == 403 && hr.Header.Get("cf-mitigated") == "challenge" {
			return NewZenlayerCloudSdkError(SecurityChallengeError,
				"Request was intercepted by a security challenge (HTTP 403). This is a network-layer block, not an API error. Contact support if it persists.",
				"")
		}
		if hr.StatusCode == 451 {
			return NewZenlayerCloudSdkError(RequestBlockedError,
				"Request was blocked by a security policy (HTTP 451). Contact support to investigate.",
				"")
		}
		return response.ParseErrorResponse(body)
	}

	//todo log.Printf("[DEBUG] Response body=%s", body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json body: %s, error: %s", body, err)
		return NewZenlayerCloudSdkError(JsonParseError, msg, "")
	}
	return
}
