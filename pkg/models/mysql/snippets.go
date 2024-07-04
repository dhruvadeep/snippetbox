package mysql

import (
	"database/sql"
	// import the Snippet struct from the models package
	"dhruvadeep.dev/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Insert new data into the database
func (m *SnippetModel) Insert( title, content, expires string) (int, error) {
		return 0, nil
}

// Return a specific snippet based on its ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
