package sqlite

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/louisbranch/drake"
)

func (db *DB) CreateSession(s *drake.Session) error {
	q := `INSERT into sessions (name, created_at) values (?, ?);`

	res, err := db.Exec(q, s.Name, s.CreatedAt)

	if err != nil {
		return errors.Wrap(err, "create session")
	}

	id, err := res.LastInsertId()

	if err != nil {
		return errors.Wrap(err, "retrieve last session id")
	}

	s.ID = strconv.FormatInt(id, 10)

	return nil
}

func (db *DB) FindSessions() ([]drake.Session, error) {
	var sessions []drake.Session

	query := `SELECT sessions.id, name, COUNT(surveys.id) AS participants,
    sessions.created_at FROM sessions
    LEFT JOIN surveys ON surveys.session_id=sessions.id
    GROUP BY sessions.id ORDER BY sessions.created_at DESC LIMIT 10
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "query sessions")
	}
	defer rows.Close()

	for rows.Next() {
		s := drake.Session{}
		err = rows.Scan(&s.ID, &s.Name, &s.Participants, &s.CreatedAt)
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

func (db *DB) FindSession(name string) (drake.Session, error) {
	q := "SELECT id, name, created_at FROM sessions where name = ?"

	s := drake.Session{}

	err := db.QueryRow(q, name).Scan(&s.ID, &s.Name, &s.CreatedAt)

	if err != nil {
		return s, errors.Wrap(err, "find session")
	}

	return s, nil
}
