package dao

import (
	"database/sql"
	"fmt"
	"samb/model"
)

type penerimaanBarangDAO struct {
	AbstractDAO
}

var PenerimaanBarangDAO = penerimaanBarangDAO{}.New()

func (input penerimaanBarangDAO) New() (output penerimaanBarangDAO) {
	output.TableName = "transaksipenerimaanbarangheader"
	output.FileName = "PenerimaanBarangDAO.go"
	return
}

func (input penerimaanBarangDAO) InsertPenerimaanBarang(db *sql.DB, inputStruct model.PenerimaanBarangModel) (lastId int64,err error) {
	query := "INSERT INTO " + input.TableName + " (" +
		" TrxInNo, WhsIdf, TrxInDate, TrxInSuppIdf, TrxInNotes)  " +
		"VALUES (?,?,?,?,?) "
	params := []interface{}{
		inputStruct.TrxInNo.String,
		inputStruct.WhsIdf.Int64,
		inputStruct.TrxInDate.String,
		inputStruct.TrxInSuppIdf.Int64,
		inputStruct.TrxInNotes.String,
	}

	result, err := db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (input penerimaanBarangDAO) GetPenerimaanBarangList(db *sql.DB) (results []model.PenerimaanBarangModel, err error) {
	query := "SELECT p.TrxInPK, p.TrxInNo, p.WhsIdf, p.TrxInDate, p.TrxInSuppIdf, p.TrxInNotes, w.WhsName, s.SupplierName " +
		" FROM " + input.TableName + " p" +
		" LEFT JOIN " + WarehouseDAO.TableName +" w ON w.WhsPK = p.WhsIdf " +
		" LEFT JOIN " + SupplierDAO.TableName +" s ON s.SupplierPK = p.TrxInSuppIdf "

	var params []interface{}

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.PenerimaanBarangModel
		if err := rows.Scan(&data.TrxInPK,
			&data.TrxInNo,
			&data.WhsIdf,
			&data.TrxInDate,
			&data.TrxInSuppIdf,
			&data.TrxInNotes,
			&data.WarehouseName,
			&data.SupplierName); err != nil {
			return nil, err
		}
		results = append(results, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}