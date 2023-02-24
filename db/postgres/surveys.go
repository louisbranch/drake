package postgres

import (
	"github.com/louisbranch/drake"
	"github.com/pkg/errors"
)

func (db *DB) CreateSurvey(s *drake.Survey) error {
	q := `INSERT into surveys (
        session_id, access_token, created_at, updated_at)
        values ($1, $2, $3, $4) returning id`

	err := db.QueryRow(q, s.SessionID, s.AccessToken, s.CreatedAt,
		s.UpdatedAt).Scan(&s.ID)
	if err != nil {
		return errors.Wrap(err, "create survey")
	}

	return nil
}

func (db *DB) UpdateSurvey(s *drake.Survey) error {
	_, err := db.Exec(`UPDATE surveys SET
    presurvey_assessment=$1,
    r_star_formation=$2,
    fp_planetary_systems=$3,
    ne_habitable_planets=$4,
    fl_life_emergence=$5,
    fi_intelligence_emergence=$6,
    fc_technology_emergence=$7,
    l_lifespan=$8,
    n_civilizations=$9,
    postsurvey_learn_gain=$10,
    postsurvey_reason=$11,
    created_at=$12,
    updated_at=$13
	where id = $14`, s.PresurveyAssessment,
		s.R, s.Fp, s.Ne, s.Fl, s.Fi, s.Fc, s.L, s.N,
		s.PostsurveyLearnGain, s.PostsurveyReason,
		s.CreatedAt, s.UpdatedAt, s.ID)

	return err
}

func (db *DB) FindSurvey(sessionID string, token string) (drake.Survey, error) {
	q := `SELECT id,
    presurvey_assessment,
    r_star_formation,
    fp_planetary_systems,
    ne_habitable_planets,
    fl_life_emergence,
    fi_intelligence_emergence,
    fc_technology_emergence,
    l_lifespan,
    n_civilizations,
    postsurvey_learn_gain,
    postsurvey_reason,
    created_at,
    updated_at
    FROM surveys where session_id = $1 AND access_token = $2`

	s := drake.Survey{
		SessionID:   sessionID,
		AccessToken: token,
	}

	err := db.QueryRow(q, s.SessionID, token).Scan(&s.ID,
		&s.PresurveyAssessment,
		&s.R, &s.Fp, &s.Ne, &s.Fl, &s.Fi, &s.Fc, &s.L, &s.N,
		&s.PostsurveyLearnGain, &s.PostsurveyReason,
		&s.CreatedAt, &s.UpdatedAt)

	if err != nil {
		return s, errors.Wrap(err, "find survey")
	}

	return s, nil
}

func (db *DB) FindSurveys(sessionID string) ([]drake.Survey, error) {
	query := `SELECT id,
    presurvey_assessment,
    r_star_formation,
    fp_planetary_systems,
    ne_habitable_planets,
    fl_life_emergence,
    fi_intelligence_emergence,
    fc_technology_emergence,
    l_lifespan,
    n_civilizations,
    postsurvey_learn_gain,
    postsurvey_reason,
    created_at,
    updated_at
    FROM surveys where session_id = $1`

	var surveys []drake.Survey

	rows, err := db.Query(query, sessionID)
	if err != nil {
		return nil, errors.Wrap(err, "query surveys")
	}
	defer rows.Close()

	for rows.Next() {
		s := drake.Survey{}
		err = rows.Scan(&s.ID,
			&s.PresurveyAssessment,
			&s.R, &s.Fp, &s.Ne, &s.Fl, &s.Fi, &s.Fc, &s.L, &s.N,
			&s.PostsurveyLearnGain, &s.PostsurveyReason,
			&s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "scan surveys")
		}
		surveys = append(surveys, s)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "find surveys")
	}
	return surveys, nil
}
