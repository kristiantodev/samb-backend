package PenerimaanBarangService

import (
	"fmt"
	"net/http"
	"samb/confiq"
	"samb/constanta"
	"samb/dao"
	"samb/dto/out"
	"samb/model"
	"time"
)

func GetPenerimaanBarangList(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> GetPenerimaanBarangGetListService.go On ", now.Format("2006-01-02 15:04:05"))
	db := confiq.Connect()

	datas, err := dao.PenerimaanBarangDAO.GetPenerimaanBarangList(db)

	if err != nil {
		fmt.Println("Error Disini", err.Error())
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	db.Close()
	out.ResponseOut(response, convertRepoToDTO(datas), true, constanta.CodeSuccessResponse, constanta.SuccessGetData)
	return nil
}

func convertRepoToDTO(datas []model.PenerimaanBarangModel) (output []out.PenerimaanBarangRequest) {
	for i:=0;i<len(datas);i++ {
		db := confiq.Connect()
		var detailList []out.PenerimaanBarangDetailRequest
		detail, _ := dao.PenerimaanBarangDetailDAO.GetPenerimaanBarangDetailList(db, datas[i].TrxInPK.Int64)
		for u:=0; u<len(detail);u++  {
			detailList = append(detailList, out.PenerimaanBarangDetailRequest{
				TrxInDPK     : detail[u].TrxInDPK.Int64,
				TrxInIDF     : detail[u].TrxInIDF.Int64,
				TrxInDProductIdf       : detail[u].TrxInDProductIdf.Int64,
				ProductName : detail[u].ProductName.String,
				TrxInDQtyDus    : detail[u].TrxInDQtyDus.Int64,
				TrxInDQtyPcs : detail[u].TrxInDQtyPcs.Int64,
			})
		}
		db.Close()
		output = append(output, out.PenerimaanBarangRequest{
			TrxInPK     : datas[i].TrxInPK.Int64,
			TrxInNo     : datas[i].TrxInNo.String,
			WhsIdf       : datas[i].WhsIdf.Int64,
			WarehouseName : datas[i].WarehouseName.String,
			TrxInDate    : datas[i].TrxInDate.String,
			TrxInSuppIdf : datas[i].TrxInSuppIdf.Int64,
			SupplierName : datas[i].SupplierName.String,
			TrxInNotes   : datas[i].TrxInNotes.String,
			PenerimaanDetail: detailList,
		})
	}
	return output
}
