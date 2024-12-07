package dao

import "database/sql"

type warehouseDAO struct {
	AbstractDAO
}

var WarehouseDAO = warehouseDAO{}.New()

func (input warehouseDAO) New() (output warehouseDAO) {
	output.TableName = "masterwarehouse"
	output.FileName = "WarehouseDAO.go"
	return
}

func (input warehouseDAO) GetDetailWarehouse(db *sql.DB, id int64) (resultId int64, err error) {
	query := "SELECT WhsPK " +
		"FROM " + input.TableName + " WHERE WhsPK = ?"

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