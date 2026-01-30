package mocks

import (
	"database/sql"

	"snippetbox.senor-crab.com/internal/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Info(id int) (*models.User, error) {
	stmt := `SELECT name, email, created FROM users WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)
	u := &models.User{}
	err := row.Scan(&u.Name, &u.Email, &u.Created)
	if err != nil {
		return nil, err
	}

	return u, nil
}
