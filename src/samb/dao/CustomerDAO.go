package dao

import "database/sql"

type customerDAO struct {
	AbstractDAO
}

var CustomerDAO = customerDAO{}.New()

func (input customerDAO) New() (output customerDAO) {
	output.TableName = "mastercustomer"
	output.FileName = "CustomerDAO.go"
	return
}

func (input customerDAO) GetDetailCustomer(db *sql.DB, id int64) (resultId int64, err error) {
	query := "SELECT CustomerPK " +
		"FROM " + input.TableName + " WHERE CustomerPK = ?"

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