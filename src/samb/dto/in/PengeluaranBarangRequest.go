package in

type PengeluaranBarangRequest struct {
	TrxOutPK      int64       `json:"trx_out_pk"`
	TrxOutNo      string    `json:"trx_out_no"`
	WhsIdf        int64       `json:"whs_idf"`
	TrxOutDate    string `json:"trx_out_date"`
	TrxOutSuppIdf int64       `json:"trx_out_supp_idf"`
	TrxOutNotes   string    `json:"trx_out_notes"`
	PengeluaranDetail []PengeluaranBarangDetailRequest `json:"pengeluaran_detail"`
}

type PengeluaranBarangDetailRequest struct {
	TrxOutDPK        int64 `json:"trx_out_dpk"`
	TrxOutIDF        int64 `json:"trx_out_idf"`
	TrxOutDProductIdf int64 `json:"trx_out_d_product_idf"`
	TrxOutDQtyDus    int64 `json:"trx_out_d_qty_dus"`
	TrxOutDQtyPcs    int64 `json:"trx_out_d_qty_pcs"`
}
