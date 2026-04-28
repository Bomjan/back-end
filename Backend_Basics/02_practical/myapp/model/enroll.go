package model

import "myapp/dataStore/postgres"

const queryEnrollStud = `INSERT INTO enroll (std_id, course_id, date_enrolled) VALUES ($1, $2, $3) RETURNING std_id ;`
const queryGetEnroll = `SELECT std_id, course_id, date_enrolled FROM enroll WHERE std_id=$1 and course_id=$2;`

type Enroll struct {
	StdId        int64  `json:"stdid"`
	CourseId     int64  `json:"courseid"`
	DateEnrolled string `json:"date"`
}

func (e *Enroll) EnrollStud() error {
	row := postgres.Db.QueryRow(queryEnrollStud, e.StdId, e.CourseId, e.DateEnrolled)
	err := row.Scan(&e.StdId)
	return err
}

func (e *Enroll) Get() error {
	return postgres.Db.QueryRow(queryGetEnroll, e.StdId, e.CourseId).Scan(&e.StdId, &e.CourseId, &e.DateEnrolled)
}
