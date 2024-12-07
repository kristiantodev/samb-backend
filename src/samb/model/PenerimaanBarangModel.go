package model

import "database/sql"

type PenerimaanBarangModel struct {
	TrxInPK        sql.NullInt64
	TrxInNo        sql.NullString
	WhsIdf         sql.NullInt64
	WarehouseName  sql.NullString
	TrxInDate      sql.NullString
	TrxInSuppIdf   sql.NullInt64
	SupplierName   sql.NullString
	TrxInNotes     sql.NullString
}

type PenerimaanBarangDetailModel struct {
	TrxInDPK         sql.NullInt64
	TrxInIDF         sql.NullInt64
	TrxInDProductIdf sql.NullInt64
	ProductName      sql.NullString
	TrxInDQtyDus     sql.NullInt64
	TrxInDQtyPcs     sql.NullInt64
}