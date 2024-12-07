package dao

import (
	"database/sql"
	"fmt"
	"samb/model"
)

type pengeluaranBarangDetailDAO struct {
	AbstractDAO
}

var PengeluaranBarangDetailDAO = pengeluaranBarangDetailDAO{}.New()

func (input pengeluaranBarangDetailDAO) New() (output pengeluaranBarangDetailDAO) {
	output.TableName = "transaksipengeluaranbarangdetail"
	output.FileName = "PengeluaranBarangDetailDAO.go"
	return
}

func (input pengeluaranBarangDetailDAO) InsertPengeluaranBarangDetail(db *sql.DB, inputStruct model.PengeluaranBarangDetailModel) (err error) {
	query := "INSERT INTO " + input.TableName + " (" +
		" TrxOutIDF, TrxOutDProductIdf, TrxOutDQtyDus, TrxOutDQtyPcs)  " +
		" VALUES (?,?,?,?) "
	params := []interface{}{
		inputStruct.TrxOutIDF.Int64,
		inputStruct.TrxOutDProductIdf.Int64,
		inputStruct.TrxOutDQtyDus.Int64,
		inputStruct.TrxOutDQtyPcs.Int64,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

func (input pengeluaranBarangDetailDAO) GetPengeluaranBarangDetailList(db *sql.DB, TrxOutDPK int64) (results []model.PengeluaranBarangDetailModel, err error) {
	query := "SELECT p.TrxOutDPK, p.TrxOutIDF, p.TrxOutDProductIdf, p.TrxOutDQtyDus, p.TrxOutDQtyPcs, pr.ProductName " +
		" FROM " + input.TableName + " p " +
		" LEFT JOIN " + ProductDAO.TableName +" pr ON pr.ProductPK = p.TrxOutDProductIdf " +
	    " WHERE p.TrxOutIDF = " + fmt.Sprintf("%d", TrxOutDPK)

	var params []interface{}

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.PengeluaranBarangDetailModel
		if err := rows.Scan(&data.TrxOutDPK,
			&data.TrxOutIDF,
			&data.TrxOutDProductIdf,
			&data.TrxOutDQtyDus,
			&data.TrxOutDQtyPcs,
			&data.ProductName,); err != nil {
			return nil, err
		}
		results = append(results, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}