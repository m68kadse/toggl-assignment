package sqlite

import (
	"context"
	"database/sql"
	"sort"

	"github.com/m68kadse/toggl-assignment/dao"
	"github.com/m68kadse/toggl-assignment/dto"
)

func sortQuestionsByID(questions []*dto.Question) {
	sort.Slice(questions, func(i, j int) bool {
		return questions[i].ID < questions[j].ID
	})
}

func (dao *SQLiteDAO) GetQuestions(ctx context.Context, params dao.PaginationParams) ([]*dto.Question, error) {
	query := `
		SELECT q.id, q.body, o.id, o.body, o.correct
		FROM (SELECT * FROM question
			ORDER BY id
			LIMIT ? OFFSET ?) AS q
		LEFT JOIN "option" AS o ON o.fk_question = q.id
		ORDER BY q.id, o.id
	`

	rows, err := dao.db.QueryContext(ctx, query, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questionsMap = make(map[int64]*dto.Question)

	for rows.Next() {
		var (
			qID, oID     sql.NullInt64
			qBody, oBody sql.NullString
			correct      sql.NullInt64
		)

		err := rows.Scan(&qID, &qBody, &oID, &oBody, &correct)
		if err != nil {
			return nil, err
		}

		question, exists := questionsMap[qID.Int64]
		if !exists {
			question = &dto.Question{
				ID:      qID.Int64,
				Body:    qBody.String,
				Options: make([]*dto.Option, 0),
			}
			questionsMap[qID.Int64] = question
		}

		if oID.Valid {
			// Create and append the option to the question
			option := &dto.Option{
				ID:      oID.Int64,
				Body:    oBody.String,
				Correct: correct.Int64 == 1,
			}
			question.Options = append(question.Options, option)
		}
	}

	// Convert the map of questions to a slice
	questions := make([]*dto.Question, 0, len(questionsMap))
	for _, question := range questionsMap {
		questions = append(questions, question)
	}

	sortQuestionsByID(questions)

	return questions, nil
}

func (dao *SQLiteDAO) GetQuestionByID(ctx context.Context, id int64) (*dto.Question, error) {
	query := `
		SELECT q.id, q.body, o.id, o.body, o.correct
		FROM question AS q
		LEFT JOIN option AS o ON o.fk_question = q.id
		WHERE q.id = ?
		ORDER BY o.id
	`

	var question *dto.Question
	optionsMap := make(map[int64]*dto.Option)

	rows, err := dao.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			qID, oID     sql.NullInt64
			qBody, oBody sql.NullString
			correct      sql.NullInt64
		)

		err := rows.Scan(&qID, &qBody, &oID, &oBody, &correct)
		if err != nil {
			return nil, err
		}

		if question == nil {
			question = &dto.Question{
				ID:      qID.Int64,
				Body:    qBody.String,
				Options: make([]*dto.Option, 0),
			}
		}

		if oID.Valid {
			// Create and append the option to the question
			option := &dto.Option{
				ID:      oID.Int64,
				Body:    oBody.String,
				Correct: correct.Int64 == 1,
			}
			optionsMap[oID.Int64] = option
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Convert the map of options to a slice and assign to the question
	for _, option := range optionsMap {
		question.Options = append(question.Options, option)
	}

	return question, nil
}

func (dao *SQLiteDAO) CreateQuestion(ctx context.Context, question *dto.Question) (*dto.Question, error) {
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx, `
		INSERT INTO question (body) VALUES (?)`, question.Body)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	for _, option := range question.Options {
		_, err = tx.ExecContext(ctx, `
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

func (dao *SQLiteDAO) UpdateQuestion(ctx context.Context, q *dto.Question) (*dto.Question, error) {
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// update question
	_, err = tx.ExecContext(ctx, `
		UPDATE question SET body = ? WHERE id = ?`, q.Body, q.ID)
	if err != nil {
		return nil, err
	}

	// delete old options
	_, err = tx.ExecContext(ctx, `
		DELETE FROM "option" WHERE fk_question = ?`, q.ID)
	if err != nil {
		return nil, err
	}

	// insert updated options
	for _, option := range q.Options {
		_, err = tx.ExecContext(ctx, `
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

func (dao *SQLiteDAO) DeleteQuestion(ctx context.Context, id int64) (int64, error) {
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// delete associated options
	_, err = tx.Exec(`
		DELETE FROM "option" WHERE fk_question = ?
	`, id)
	if err != nil {
		return 0, err
	}

	// delete question
	_, err = tx.Exec(`
		DELETE FROM question WHERE id = ?
	`, id)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}
