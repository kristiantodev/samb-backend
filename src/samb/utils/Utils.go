package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"samb/dto/in"
	"strconv"
)

func StructToJSON(input interface{}) (output string) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		output = ""
		return
	}
	output = string(b)
	return
}

func readBody(request *http.Request) (output string) {
	byteBody, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		return ""
	}
	return string(byteBody)
}

func PenerimaanBarangBody(request *http.Request) (bodyRequest in.PenerimaanBarangRequest) {
	jsonString := readBody(request)
	json.Unmarshal([]byte(jsonString), &bodyRequest)
	return
}

func PengeluaranBarangBody(request *http.Request) (bodyRequest in.PengeluaranBarangRequest) {
	jsonString := readBody(request)
	json.Unmarshal([]byte(jsonString), &bodyRequest)
	return
}

func ReadParam(request *http.Request) (id int64, err error) {
	strId, ok := mux.Vars(request)["Id"]
	idParam, errConvert := strconv.Atoi(strId)
	id = int64(idParam)

	if !ok || errConvert != nil {
		err = errConvert
		return
	}
	return
}