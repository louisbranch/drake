package sqlite

import (
	"database/sql"
	"errors"

	"github.com/louisbranch/drake"
	sqlite3 "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func init() {
	sql.Register("sqlite3_with_fk",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				_, err := conn.Exec("PRAGMA foreign_keys = ON", nil)
				return err
			},
		})
}

func New(path string) (*DB, error) {
	db, err := sql.Open("sqlite3_with_fk", path)
	if err != nil {
		return nil, err
	}

	queries := []string{
		`
        CREATE TABLE IF NOT EXISTS sessions(
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL UNIQUE CHECK(name <> ''),
            created_at DATETIME
        );
        `,
		`
        CREATE TABLE IF NOT EXISTS surveys(
            id INTEGER PRIMARY KEY,
            session_id INTEGER NOT NULL,

            presurvey_guess REAL CHECK(presurvey_guess > 0),
            r_star_formation REAL CHECK(r_star_formation > 0),
            fp_planetary_systems REAL CHECK(fp_planetary_systems > 0),
            ne_habitable_planets REAL CHECK(ne_habitable_planets > 0),
            fl_life_emergence REAL CHECK(fl_life_emergence > 0),
            fi_intelligence_emergence REAL CHECK(fi_intelligence_emergence > 0),
            fc_technology_emergence REAL CHECK(fc_technology_emergence > 0),
            l_lifespan REAL CHECK(l_lifespan > 0),
            n_civilizations REAL CHECK(n_civilizations > 0),

            created_at DATETIME,
            updated_at DATETIME,
            FOREIGN KEY(session_id) REFERENCES sessions(id) ON DELETE CASCADE
        );
		`,
	}

	for _, q := range queries {
		_, err = db.Exec(q)

		if err != nil {
			return nil, err
		}
	}

	return &DB{db}, nil
}

func (db *DB) CreateSession(session *drake.Session) error {
	return errors.New("not implemented")
}
