package postgres

import (
	"github.com/louisbranch/drake"
	"github.com/pkg/errors"
)

func (db *DB) CreateSession(s *drake.Session) error {
	q := `INSERT INTO sessions (name, created_at) VALUES ($1, $2) RETURNING id`

	err := db.QueryRow(q, s.Name, s.CreatedAt).Scan(&s.ID)
	if err != nil {
		return errors.Wrap(err, "create session")
	}

	return nil
}

func (db *DB) FindSession(name string) (drake.Session, error) {
	q := "SELECT id, name, created_at FROM sessions where name = $1"

	s := drake.Session{}

	err := db.QueryRow(q, name).Scan(&s.ID, &s.Name, &s.CreatedAt)

	if err != nil {
		return s, errors.Wrap(err, "find session")
	}

	return s, nil
}

func (db *DB) FindSessions() ([]drake.Session, error) {
	var sessions []drake.Session

	query := `SELECT id, name, created_at FROM sessions ORDER BY created_at DESC
    LIMIT 10`

	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "query sessions")
	}
	defer rows.Close()

	for rows.Next() {
		s := drake.Session{}
		err = rows.Scan(&s.ID, &s.Name, &s.CreatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "scan sessions")
		}
		sessions = append(sessions, s)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "find sessions")
	}
	return sessions, nil
}
