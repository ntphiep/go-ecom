package user

import (
	"database/sql"

	"github.com/ntphiep/go-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {

	}
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)
	err := rows.Scan(
		&u.ID, 
		&u.FirstName, 
		&u.LastName, 
		&u.Email, 
		&u.Password, 
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}