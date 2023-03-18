package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func New(connection string) (*DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	queries := []string{
		`
        CREATE TABLE IF NOT EXISTS sessions(
            id SERIAL PRIMARY KEY,
            name VARCHAR (255) UNIQUE NOT NULL,
            created_at TIMESTAMP
        );
        `,
		`
        CREATE UNIQUE INDEX IF NOT EXISTS sessions_name ON
            sessions(name)
        `,
		`
        CREATE TABLE IF NOT EXISTS surveys(
            id SERIAL PRIMARY KEY,
            session_id SERIAL NOT NULL,
            access_token TEXT NOT NULL,

            presurvey_assessment REAL CHECK(presurvey_assessment >= 0),
            r_star_formation REAL CHECK(r_star_formation > 0),
            fp_planetary_systems REAL CHECK(fp_planetary_systems > 0),
            ne_habitable_planets REAL CHECK(ne_habitable_planets > 0),
            fl_life_emergence REAL CHECK(fl_life_emergence > 0),
            fi_intelligence_emergence REAL CHECK(fi_intelligence_emergence > 0),
            fc_technology_emergence REAL CHECK(fc_technology_emergence > 0),
            l_lifespan REAL CHECK(l_lifespan > 0),
            n_civilizations REAL CHECK(n_civilizations >= 0),
            postsurvey_learn_gain BOOL,
            postsurvey_reason TEXT,

            created_at TIMESTAMP,
            updated_at TIMESTAMP,
            FOREIGN KEY(session_id) REFERENCES sessions(id) ON DELETE CASCADE
        );
		`,
		`
        CREATE UNIQUE INDEX IF NOT EXISTS surveys_access_token ON
            surveys(session_id, access_token)
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
