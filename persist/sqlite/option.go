package sqlite

import "github.com/m68kadse/toggl-assignment/dto"

func (db *SQLiteDB) GetOptionByID(id int64) *dto.Option {
	stmt, err := db.PrepareStmt(
		`SELECT id, body
		 FROM question
		 WHERE id=?`)

}
