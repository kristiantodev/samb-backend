package PengeluaranBarangService

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

func PengeluaranBarangInsert(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> PengeluaranBarangInsert.go On ", now.Format("2006-01-02 15:04:05"))


	body := utils.PengeluaranBarangBody(request)
	bodyRepo := pengeluaranBarangRepository(body)
	db := confiq.Connect()

	if bodyRepo.TrxOutNo.String == ""{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "TrxOutNo tidak boleh kosong")
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

	idSpl, err := dao.SupplierDAO.GetDetailSupplier(db, bodyRepo.TrxOutSuppIdf.Int64)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if idSpl == 0 {
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "TrxOutSuppIdf tidak diketahui")
		return
	}

	for i:=0; i< len(body.PengeluaranDetail);i++ {
		idPrd, err2 := dao.ProductDAO.GetDetailProduct(db, body.PengeluaranDetail[i].TrxOutDProductIdf)
		if err2 != nil{
			err = err2
			out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
			return
		}

		if idPrd == 0 {
			out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "TrxOutDProductIdf tidak diketahui")
			return
		}
	}

	lastId, err := dao.PengeluaranBarangDAO.InsertPengeluaranBarang(db, bodyRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	for i:=0; i< len(body.PengeluaranDetail);i++  {
		detailRepo := pengeluaranBarangDetailRepository(body.PengeluaranDetail[i])
		detailRepo.TrxOutIDF.Int64 = lastId
		dao.PengeluaranBarangDetailDAO.InsertPengeluaranBarangDetail(db, detailRepo)
	}

	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessAddData)
	return nil
}

func pengeluaranBarangRepository(body in.PengeluaranBarangRequest) model.PengeluaranBarangModel  {
	return model.PengeluaranBarangModel{
		TrxOutNo: sql.NullString{String: body.TrxOutNo},
		WhsIdf:  sql.NullInt64{Int64: body.WhsIdf},
		TrxOutDate: sql.NullString{String: body.TrxOutDate},
		TrxOutSuppIdf : sql.NullInt64{Int64: body.TrxOutSuppIdf},
		TrxOutNotes: sql.NullString{String: body.TrxOutNotes},
	}
}

func pengeluaranBarangDetailRepository(body in.PengeluaranBarangDetailRequest) model.PengeluaranBarangDetailModel  {
	return model.PengeluaranBarangDetailModel{
		TrxOutIDF: sql.NullInt64{Int64: body.TrxOutIDF},
		TrxOutDProductIdf:  sql.NullInt64{Int64: body.TrxOutDProductIdf},
		TrxOutDQtyDus: sql.NullInt64{Int64: body.TrxOutDQtyDus},
		TrxOutDQtyPcs: sql.NullInt64{Int64: body.TrxOutDQtyPcs},
	}
}
