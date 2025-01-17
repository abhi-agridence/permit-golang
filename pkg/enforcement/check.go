package enforcement

import (
	"bytes"
	"encoding/json"
	"github.com/permitio/permit-golang/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	// client does not need to be initialized each request (thread safe). This will cause a huge latency and not good for concurrent requests.
	client *http.Client = &http.Client{
		Timeout: DefaultTimeout * time.Second,
	}
)

type Action string

type CheckRequest struct {
	User     User              `json:"user"`
	Action   Action            `json:"action"`
	Resource Resource          `json:"resource"`
	Context  map[string]string `json:"context"`
}
type CheckResponse struct {
	allow string
}

func NewCheckRequest(user User, action Action, resource Resource, context map[string]string) *CheckRequest {
	return &CheckRequest{
		User:     user,
		Action:   action,
		Resource: resource,
		Context:  context,
	}
}

func (e *PermitEnforcer) Check(user User, action Action, resource Resource, additionalContext ...map[string]string) (bool, error) {
	const (
		reqMethod           = "POST"
		reqContentTypeKey   = "Content-Type"
		reqContentTypeValue = "application/json"
		reqAuthKey          = "Authorization"
	)
	var checkRes map[string]interface{}
	reqAuthValue := "Bearer " + e.config.GetToken()
	reqEndpoint := e.config.GetPdpUrl() + "/allowed"

	if additionalContext == nil {
		additionalContext = make([]map[string]string, 0)
		additionalContext = append(additionalContext, make(map[string]string))
	}

	checkReq := NewCheckRequest(user, action, resource, additionalContext[0])
	jsonCheckReq, err := json.Marshal(checkReq)
	if err != nil {
		permitError := errors.NewPermitUnexpectedError(err)
		e.logger.Error("error marshalling Permit.Check() request", zap.Error(permitError))
		return false, permitError
	}
	reqBody := bytes.NewBuffer(jsonCheckReq)
	httpRequest, err := http.NewRequest(reqMethod, reqEndpoint, reqBody)
	if err != nil {
		permitError := errors.NewPermitUnexpectedError(err)
		e.logger.Error("error creating Permit.Check() request", zap.Error(permitError))
		return false, permitError
	}
	httpRequest.Header.Set(reqContentTypeKey, reqContentTypeValue)
	httpRequest.Header.Set(reqAuthKey, reqAuthValue)
	res, err := client.Do(httpRequest)
	if err != nil {
		permitError := errors.NewPermitUnexpectedError(err)
		e.logger.Error("error sending Permit.Check() request to PDP", zap.Error(permitError))
		return false, permitError
	}

	err = json.NewDecoder(res.Body).Decode(&checkRes)
	if err != nil {
		permitError := errors.NewPermitUnexpectedError(err)
		e.logger.Error("error decoding Permit.Check() response from PDP", zap.Error(permitError))
		return false, permitError
	}
	stringRes, err := json.Marshal(checkRes)
	if err != nil {
		permitError := errors.NewPermitUnexpectedError(err)
		e.logger.Error("error marshalling Permit.Check() response from PDP", zap.Error(permitError))
		return false, permitError
	}
	err = errors.HttpErrorHandle(err, res)
	if err != nil {
		e.logger.Error(string(stringRes), zap.Error(err))
		return false, err
	}
	if _, found := checkRes[AllowKey]; found {
		return checkRes[AllowKey].(bool), nil
	}

	permitError := errors.NewPermitUnexpectedError(err)
	e.logger.Error("error parsing Permit.Check() response from PDP", zap.Error(permitError))
	return false, permitError
}
