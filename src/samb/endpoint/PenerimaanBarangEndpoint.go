package endpoint

import (
	"net/http"
	"samb/service/PenerimaanBarangService"
)

func PenerimaanBarangEndpointWithoutId(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		PenerimaanBarangService.PenerimaanBarangInsert(response, request)
		break
	case "GET":
		PenerimaanBarangService.GetPenerimaanBarangList(response, request)
		break
	}
}
