package out

type PenerimaanBarangRequest struct {
	TrxInPK      int64     `json:"trx_in_pk"`
	TrxInNo      string    `json:"trx_in_no"`
	WhsIdf       int64     `json:"whs_idf"`
	WarehouseName string   `json:"warehouse_name"`
	TrxInDate    string `json:"trx_in_date"`
	TrxInSuppIdf int64     `json:"trx_in_supp_idf"`
	SupplierName string    `json:"supplier_name"`
	TrxInNotes   string    `json:"trx_in_notes"`
	PenerimaanDetail []PenerimaanBarangDetailRequest `json:"penerimaan_detail"`
}

type PenerimaanBarangDetailRequest struct {
	TrxInDPK          int64 `json:"trx_in_dpk"`
	TrxInIDF          int64 `json:"trx_in_idf"`
	TrxInDProductIdf  int64 `json:"trx_in_d_product_idf"`
	ProductName       string `json:"product_name"`
	TrxInDQtyDus      int64 `json:"trx_in_d_qty_dus"`
	TrxInDQtyPcs      int64 `json:"trx_in_d_qty_pcs"`
}
