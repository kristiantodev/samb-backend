package model

import "database/sql"

type PengeluaranBarangModel struct {
	TrxOutPK          sql.NullInt64
	TrxOutNo          sql.NullString
	WhsIdf            sql.NullInt64
	WarehouseName     sql.NullString
	TrxOutDate        sql.NullString
	TrxOutSuppIdf     sql.NullInt64
	SupplierName      sql.NullString
	TrxOutNotes       sql.NullString
}

type PengeluaranBarangDetailModel struct {
	TrxOutDPK        sql.NullInt64
	TrxOutIDF        sql.NullInt64
	TrxOutDProductIdf sql.NullInt64
	ProductName      sql.NullString
	TrxOutDQtyDus    sql.NullInt64
	TrxOutDQtyPcs    sql.NullInt64
}