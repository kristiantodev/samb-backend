package dao

import "database/sql"

type productDAO struct {
	AbstractDAO
}

var ProductDAO = productDAO{}.New()

func (input productDAO) New() (output productDAO) {
	output.TableName = "masterproduct"
	output.FileName = "ProductDAO.go"
	return
}

func (input productDAO) GetDetailProduct(db *sql.DB, id int64) (resultId int64, err error) {
	query := "SELECT ProductPK " +
		"FROM " + input.TableName + " WHERE ProductPK = ?"

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