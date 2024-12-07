package out

import (
	"encoding/json"
	"net/http"
	"samb/utils"
)

type APIResponse struct {
	Samb SambMessage `json:"samb"`
}

type SambMessage struct {
	Success bool     `json:"success"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Content interface{} `json:"content"`
}

func (ar APIResponse) String() string {
	return utils.StructToJSON(ar)
}

func ResponseOut(response http.ResponseWriter, data interface{}, success bool, code int, message string){
	response.Header().Set("Content-type", "application/json")
	var apiResponse APIResponse
	apiResponse.Samb.Success = success
	apiResponse.Samb.Code = code
	apiResponse.Samb.Message = message
	apiResponse.Samb.Content = data
	response.WriteHeader(code)
	json.NewEncoder(response).Encode(apiResponse)
}