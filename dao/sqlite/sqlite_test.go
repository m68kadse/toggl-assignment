package sqlite

import (
	"context"
	"fmt"
	"testing"

	"github.com/m68kadse/toggl-assignment/dao"
	"github.com/m68kadse/toggl-assignment/dto"
)

func setupTestDB(t *testing.T) (*SQLiteDAO, func()) {
	t.Helper()

	// Open an in-memory SQLite database for testing
	sqldao, err := NewDAO(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize SQLiteDAO: %v", err)
	}

	cleanup := func() {
		sqldao.db.Close()
	}

	return sqldao, cleanup
}

func TestSQLiteDAO(t *testing.T) {
	// Initialize the SQLiteDAO for testing
	sqldao, cleanup := setupTestDB(t)
	defer cleanup()

	// Create a test question
	testQuestion := &dto.Question{
		Body: "Test Question",
		Options: []*dto.Option{
			{Body: "Option A", Correct: true},
			{Body: "Option B", Correct: false},
		},
	}

	t.Run("CreateQuestion", func(t *testing.T) {
		createdQuestion, err := sqldao.CreateQuestion(context.Background(), testQuestion)
		if err != nil {
			t.Errorf("CreateQuestion error: %v", err)
		}
		if createdQuestion == nil {
			t.Error("CreateQuestion did not return a question")
		}
	})

	t.Run("GetQuestionByID", func(t *testing.T) {
		createdQuestion, err := sqldao.CreateQuestion(context.Background(), testQuestion)
		if err != nil {
			t.Errorf("CreateQuestion error: %v", err)
		}

		retrievedQuestion, err := sqldao.GetQuestionByID(context.Background(), createdQuestion.ID)
		if err != nil {
			t.Errorf("GetQuestionByID error: %v", err)
		}
		if retrievedQuestion == nil {
			t.Error("GetQuestionByID did not return a question")
		}
	})

	t.Run("UpdateQuestion", func(t *testing.T) {
		createdQuestion, err := sqldao.CreateQuestion(context.Background(), testQuestion)
		if err != nil {
			t.Errorf("CreateQuestion error: %v", err)
		}

		createdQuestion.Body = "Updated Test Question"
		updatedQuestion, err := sqldao.UpdateQuestion(context.Background(), createdQuestion)
		if err != nil {
			t.Errorf("UpdateQuestion error: %v", err)
		}
		if updatedQuestion == nil {
			t.Error("UpdateQuestion did not return an updated question")
		}
		if updatedQuestion.Body != "Updated Test Question" {
			t.Errorf("UpdateQuestion returned the wrong question body")
		}
	})

	t.Run("DeleteQuestion", func(t *testing.T) {
		createdQuestion, err := sqldao.CreateQuestion(context.Background(), testQuestion)
		if err != nil {
			t.Errorf("CreateQuestion error: %v", err)
		}

		deletedID, err := sqldao.DeleteQuestion(context.Background(), createdQuestion.ID)
		if err != nil {
			t.Errorf("DeleteQuestion error: %v", err)
		}
		if deletedID != createdQuestion.ID {
			t.Errorf("DeleteQuestion returned the wrong deleted ID")
		}

		// Try to retrieve the deleted question, it should not exist
		retrieved, err := sqldao.GetQuestionByID(context.Background(), createdQuestion.ID)
		if retrieved != nil && err != nil {
			t.Error("Deleted question still exists")
		}
	})

	t.Run("GetQuestions", func(t *testing.T) {
		// Create some test questions for pagination
		for i := 1; i <= 10; i++ {
			question := &dto.Question{
				Body:    "Test Question " + fmt.Sprint(i),
				Options: []*dto.Option{{Body: "Option A", Correct: true}},
			}
			_, err := sqldao.CreateQuestion(context.Background(), question)
			if err != nil {
				t.Errorf("CreateQuestion error: %v", err)
			}
		}

		// Test pagination
		paginationParams := dao.PaginationParams{Limit: 5, Offset: 0}
		questions, err := sqldao.GetQuestions(context.Background(), paginationParams)
		if err != nil {
			t.Errorf("GetQuestions error: %v", err)
		}
		if len(questions) != 5 {
			t.Errorf("GetQuestions did not return the expected number of questions, returned %d", len(questions))
		}
	})
}
