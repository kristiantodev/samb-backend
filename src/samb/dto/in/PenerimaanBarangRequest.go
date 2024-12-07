package in

type PenerimaanBarangRequest struct {
	TrxInPK      int64       `json:"trx_in_pk"`
	TrxInNo      string    `json:"trx_in_no"`
	WhsIdf       int64       `json:"whs_idf"`
	TrxInDate    string `json:"trx_in_date"`
	TrxInSuppIdf int64       `json:"trx_in_supp_idf"`
	TrxInNotes   string    `json:"trx_in_notes"`
	PenerimaanDetail []PenerimaanBarangDetailRequest `json:"penerimaan_detail"`
}

type PenerimaanBarangDetailRequest struct {
	TrxInDPK          int64 `json:"trx_in_dpk"`
	TrxInIDF          int64 `json:"trx_in_idf"`
	TrxInDProductIdf  int64 `json:"trx_in_d_product_idf"`
	TrxInDQtyDus      int64 `json:"trx_in_d_qty_dus"`
	TrxInDQtyPcs      int64 `json:"trx_in_d_qty_pcs"`
}