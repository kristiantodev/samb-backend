package PenerimaanBarangService

import (
	"database/sql"
	"fmt"
	"net/http"
	"samb/confiq"
	"samb/constanta"
	"samb/dao"
	"samb/dto/in"
	"samb/dto/out"
	"samb/model"
	"samb/utils"
	"time"
)

func PenerimaanBarangInsert(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> PenerimaanBarangInsert.go On ", now.Format("2006-01-02 15:04:05"))


	body := utils.PenerimaanBarangBody(request)
	bodyRepo := penerimaanBarangRepository(body)
	db := confiq.Connect()

	if bodyRepo.TrxInNo.String == ""{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "TrxInNo tidak boleh kosong")
		return
	}

	idWhs, err := dao.WarehouseDAO.GetDetailWarehouse(db, bodyRepo.WhsIdf.Int64)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if idWhs == 0 {
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "WhsIdf tidak diketahui")
		return
	}

	idSpl, err := dao.SupplierDAO.GetDetailSupplier(db, bodyRepo.TrxInSuppIdf.Int64)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if idSpl == 0 {
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "TrxInSuppIdf tidak diketahui")
		return
	}

	for i:=0; i< len(body.PenerimaanDetail);i++ {
		idPrd, err2 := dao.ProductDAO.GetDetailProduct(db, body.PenerimaanDetail[i].TrxInDProductIdf)
		if err2 != nil{
			err = err2
			out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
			return
		}

		if idPrd == 0 {
			out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "TrxInDProductIdf tidak diketahui")
			return
		}
	}

	lastId, err := dao.PenerimaanBarangDAO.InsertPenerimaanBarang(db, bodyRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	for i:=0; i< len(body.PenerimaanDetail);i++  {
	  detailRepo := penerimaanBarangDetailRepository(body.PenerimaanDetail[i])
	  detailRepo.TrxInIDF.Int64 = lastId
	  dao.PenerimaanBarangDetailDAO.InsertPenerimaanBarangDetail(db, detailRepo)
	}

	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessAddData)
	return nil
}

func penerimaanBarangRepository(body in.PenerimaanBarangRequest) model.PenerimaanBarangModel  {
	return model.PenerimaanBarangModel{
		TrxInNo: sql.NullString{String: body.TrxInNo},
		WhsIdf:  sql.NullInt64{Int64: body.WhsIdf},
        TrxInDate: sql.NullString{String: body.TrxInDate},
		TrxInSuppIdf : sql.NullInt64{Int64: body.TrxInSuppIdf},
		TrxInNotes: sql.NullString{String: body.TrxInNotes},
	}
}

func penerimaanBarangDetailRepository(body in.PenerimaanBarangDetailRequest) model.PenerimaanBarangDetailModel  {
	return model.PenerimaanBarangDetailModel{
		TrxInIDF: sql.NullInt64{Int64: body.TrxInIDF},
		TrxInDProductIdf:  sql.NullInt64{Int64: body.TrxInDProductIdf},
		TrxInDQtyDus: sql.NullInt64{Int64: body.TrxInDQtyDus},
		TrxInDQtyPcs: sql.NullInt64{Int64: body.TrxInDQtyPcs},
	}
}