package model

import "myapp/dataStore/postgres"

const queryInsertUser = `INSERT INTO student (stdid, firstname, lastname, email) VALUES ($1, $2, $3, $4);`

type Student struct {
	StdId     int64  `json:"stdId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
	return err
}
