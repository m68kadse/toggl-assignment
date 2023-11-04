package sqlite

import "github.com/m68kadse/toggl-assignment/dto"

func (db *SQLiteDB) GetQuestionByID(id int64) (*dto.Question, error) {
	stmt, err := db.PrepareStmt(`
		SELECT q.id, q.body, o.id, o.body, o.correct
		FROM question AS q
		LEFT JOIN "option" AS o ON o.fk_question = q.id
		WHERE q.id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var question *dto.Question
	optionsMap := make(map[int64]*dto.Option)

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			qID, oID     int64
			qBody, oBody string
			correct      int
		)

		err := rows.Scan(&qID, &qBody, &oID, &oBody, &correct)
		if err != nil {
			return nil, err
		}

		if question == nil {
			question = &dto.Question{
				ID:      qID,
				Body:    qBody,
				Options: make([]*dto.Option, 0),
			}
		}

		option, exists := optionsMap[oID]
		if !exists {
			option = &dto.Option{
				ID:      oID,
				Body:    oBody,
				Correct: correct == 1,
			}
			optionsMap[oID] = option
		}

		question.Options = append(question.Options, option)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return question, nil
}

func (db *SQLiteDB) CreateQuestion(question *dto.Question) (*dto.Question, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(`
		INSERT INTO question (body) VALUES (?)`, question.Body)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	for _, option := range question.Options {
		_, err = tx.Exec(`
			INSERT INTO "option" (fk_question, body, correct) VALUES (?, ?, ?)`,
			lastInsertID, option.Body, option.Correct)
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	question.ID = lastInsertID

	return question, nil
}

func (db *SQLiteDB) UpdateQuestion(q *dto.Question) (*dto.Question, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// update question
	_, err = tx.Exec(`
		UPDATE question SET body = ? WHERE id = ?`, q.Body, q.ID)
	if err != nil {
		return nil, err
	}

	// delete old options
	_, err = tx.Exec(`
		DELETE FROM "option" WHERE fk_question = ?`, q.ID)
	if err != nil {
		return nil, err
	}

	// insert updated options
	for _, option := range q.Options {
		_, err = tx.Exec(`
			INSERT INTO "option" (fk_question, body, correct) VALUES (?, ?, ?)`, q.ID, option.Body, option.Correct)
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return q, nil
}

	// commit the transaction
func (db *SQLiteDB) DeleteQuestion(q *dto.Question) (*dto.Question, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() /

	// delete associated options
	_, err = tx.Exec(`
		DELETE FROM "option" WHERE fk_question = ?
	`, q.ID)
	if err != nil {
		return nil, err
	}

	// delete question
	_, err = tx.Exec(`
		DELETE FROM question WHERE id = ?
	`, q.ID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return q, nil
}
