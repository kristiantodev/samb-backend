package dao

import (
	"database/sql"
	"fmt"
	"samb/model"
)

type pengeluaranBarangDAO struct {
	AbstractDAO
}

var PengeluaranBarangDAO = pengeluaranBarangDAO{}.New()

func (input pengeluaranBarangDAO) New() (output pengeluaranBarangDAO) {
	output.TableName = "transaksipengeluaranbarangheader"
	output.FileName = "PengeluaranBarangDAO.go"
	return
}

func (input pengeluaranBarangDAO) InsertPengeluaranBarang(db *sql.DB, inputStruct model.PengeluaranBarangModel) (lastId int64, err error) {
	query := "INSERT INTO " + input.TableName + " (" +
		" TrxOutNo, WhsIdf, TrxOutDate, TrxOutSuppIdf, TrxOutNotes)  " +
		"VALUES (?,?,?,?,?) "
	params := []interface{}{
		inputStruct.TrxOutNo.String,
		inputStruct.WhsIdf.Int64,
		inputStruct.TrxOutDate.String,
		inputStruct.TrxOutSuppIdf.Int64,
		inputStruct.TrxOutNotes.String,
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

func (input pengeluaranBarangDAO) GetPengeluaranBarangList(db *sql.DB) (results []model.PengeluaranBarangModel, err error) {
	query := "SELECT p.TrxOutPK, p.TrxOutNo, p.WhsIdf, p.TrxOutDate, p.TrxOutSuppIdf, p.TrxOutNotes, w.WhsName, s.SupplierName " +
		" FROM " + input.TableName + " p" +
		" LEFT JOIN " + WarehouseDAO.TableName +" w ON w.WhsPK = p.WhsIdf " +
		" LEFT JOIN " + SupplierDAO.TableName +" s ON s.SupplierPK = p.TrxOutSuppIdf "

	var params []interface{}

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.PengeluaranBarangModel
		if err := rows.Scan(&data.TrxOutPK,
			&data.TrxOutNo,
			&data.WhsIdf,
			&data.TrxOutDate,
			&data.TrxOutSuppIdf,
			&data.TrxOutNotes,
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