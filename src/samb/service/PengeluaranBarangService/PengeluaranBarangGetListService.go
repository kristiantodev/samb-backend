package PengeluaranBarangService

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

func GetPengeluaranBarangList(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> GetPengeluaranBarangGetListService.go On ", now.Format("2006-01-02 15:04:05"))
	db := confiq.Connect()

	datas, err := dao.PengeluaranBarangDAO.GetPengeluaranBarangList(db)

	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	db.Close()
	out.ResponseOut(response, convertRepoToDTO(datas), true, constanta.CodeSuccessResponse, constanta.SuccessGetData)
	return nil
}

func convertRepoToDTO(datas []model.PengeluaranBarangModel) (output []out.PengeluaranBarangRequest) {
	for i:=0;i<len(datas);i++ {
		db := confiq.Connect()
		var detailList []out.PengeluaranBarangDetailRequest
		detail, _ := dao.PengeluaranBarangDetailDAO.GetPengeluaranBarangDetailList(db, datas[i].TrxOutPK.Int64)
		for u:=0; u<len(detail);u++  {
			detailList = append(detailList, out.PengeluaranBarangDetailRequest{
				TrxOutDPK     : detail[u].TrxOutDPK.Int64,
				TrxOutIDF     : detail[u].TrxOutIDF.Int64,
				TrxOutDProductIdf       : detail[u].TrxOutDProductIdf.Int64,
				ProductName : detail[u].ProductName.String,
				TrxOutDQtyDus    : detail[u].TrxOutDQtyDus.Int64,
				TrxOutDQtyPcs : detail[u].TrxOutDQtyPcs.Int64,
			})
		}
		db.Close()
		output = append(output, out.PengeluaranBarangRequest{
			TrxOutPK     : datas[i].TrxOutPK.Int64,
			TrxOutNo     : datas[i].TrxOutNo.String,
			WhsIdf       : datas[i].WhsIdf.Int64,
			WarehouseName : datas[i].WarehouseName.String,
			TrxOutDate    : datas[i].TrxOutDate.String,
			TrxOutSuppIdf : datas[i].TrxOutSuppIdf.Int64,
			SupplierName : datas[i].SupplierName.String,
			TrxOutNotes   : datas[i].TrxOutNotes.String,
			PengeluaranDetail: detailList,
		})
	}
	return output
}
