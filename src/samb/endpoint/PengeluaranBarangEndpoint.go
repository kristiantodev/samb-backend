package endpoint

import (
	"net/http"
	"samb/service/PengeluaranBarangService"
)

func PengeluaranBarangEndpointWithoutId(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		PengeluaranBarangService.PengeluaranBarangInsert(response, request)
		break
	case "GET":
		PengeluaranBarangService.GetPengeluaranBarangList(response, request)
		break
	}
}