package dao

import "database/sql"

type supplierDAO struct {
	AbstractDAO
}

var SupplierDAO = supplierDAO{}.New()

func (input supplierDAO) New() (output supplierDAO) {
	output.TableName = "mastersupplier"
	output.FileName = "SupplierDAO.go"
	return
}

func (input supplierDAO) GetDetailSupplier(db *sql.DB, id int64) (resultId int64, err error) {
	query := "SELECT SupplierPK " +
		"FROM " + input.TableName + " WHERE SupplierPK = ?"

	row := db.QueryRow(query, id)

	err = row.Scan(&resultId)
	if err != nil {
		if err == sql.ErrNoRows {
			return resultId, nil
		}
		return resultId, err
	}
	return resultId, nil
}