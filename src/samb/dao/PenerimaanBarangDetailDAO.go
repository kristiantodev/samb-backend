package dao

import (
	"database/sql"
	"fmt"
	"samb/model"
)

type penerimaanBarangDetailDAO struct {
	AbstractDAO
}

var PenerimaanBarangDetailDAO = penerimaanBarangDetailDAO{}.New()

func (input penerimaanBarangDetailDAO) New() (output penerimaanBarangDetailDAO) {
	output.TableName = "transaksipenerimaanbarangdetail"
	output.FileName = "PenerimaanBarangDetailDAO.go"
	return
}

func (input penerimaanBarangDetailDAO) InsertPenerimaanBarangDetail(db *sql.DB, inputStruct model.PenerimaanBarangDetailModel) (err error) {
	query := "INSERT INTO " + input.TableName + " (" +
		" TrxInIDF, TrxInDProductIdf, TrxInDQtyDus, TrxInDQtyPcs)  " +
		"VALUES (?,?,?,?) "
	params := []interface{}{
		inputStruct.TrxInIDF.Int64,
		inputStruct.TrxInDProductIdf.Int64,
		inputStruct.TrxInDQtyDus.Int64,
		inputStruct.TrxInDQtyPcs.Int64,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

func (input penerimaanBarangDetailDAO) GetPenerimaanBarangDetailList(db *sql.DB, TrxInIDF int64) (results []model.PenerimaanBarangDetailModel, err error) {
	query := "SELECT p.TrxInDPK, p.TrxInIDF, p.TrxInDProductIdf, p.TrxInDQtyDus, p.TrxInDQtyPcs, pr.ProductName " +
		" FROM " + input.TableName + " p " +
		" LEFT JOIN " + ProductDAO.TableName +" pr ON pr.ProductPK = p.TrxInDProductIdf " +
		" WHERE p.TrxInIDF = " + fmt.Sprintf("%d", TrxInIDF)

	var params []interface{}

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.PenerimaanBarangDetailModel
		if err := rows.Scan(&data.TrxInDPK,
			&data.TrxInIDF,
			&data.TrxInDProductIdf,
			&data.TrxInDQtyDus,
			&data.TrxInDQtyPcs,
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