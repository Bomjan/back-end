package model

import (
	"myapp/dataStore/postgres"
)

// * Student Model

// queryInsertUser is the SQL statement for inserting a new student record.
// Using parameterized queries ($1, $2, ...) prevents SQL injection - important for security.
const queryInsertUser = `INSERT INTO student (stdid, firstname, lastname, email) VALUES ($1, $2, $3, $4);`
const queryGetUser = `SELECT stdid, firstname, lastname, email FROM student WHERE stdid=$1;`
const queryUpdate = `UPDATE student SET stdid=$1, firstname=$2, lastname=$3, email=$4 WHERE stdid=$5 RETURNING stdid`
const queryDeleteUser = `DELETE FROM student WHERE stdid=$1 RETURNING stdid;`
const queryGetAllStudents = `SELECT * FROM student;`

// Student represents the student table structure.
// The json tags define how the fields appear in API requests/responses.
// StdId is the primary key - must be unique for each student.
type Student struct {
	StdId     int64  `json:"stdId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// Create inserts the student record into the database.
// This is where the actual data persistence happens - without this, the data never leaves the request.
// Returns an error if the insert fails (e.g., duplicate ID, constraint violation).
func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
	return err
}

func (s *Student) Read() error {
	return postgres.Db.QueryRow(queryGetUser, s.StdId).Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
}

func (s *Student) Update(oldID int64) error {
	return postgres.Db.QueryRow(queryUpdate, s.StdId, s.FirstName, s.LastName, s.Email, oldID).Scan(&s.StdId)
}

func (s *Student) Delete() error {
	if err := postgres.Db.QueryRow(queryDeleteUser, s.StdId).Scan(&s.StdId); err != nil {
		return err
	} else {
		return nil
	}
}

func GetAllStudents() ([]Student, error) {
	rows, getErr := postgres.Db.Query(queryGetAllStudents)

	if getErr != nil {
		return nil, getErr
	}

	// create a slice of students from rows
	students := []Student{}
	for rows.Next() {
		var s Student
		dbErr := rows.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
		if dbErr != nil {
			return nil, dbErr
		}
		students = append(students, s)
	}
	rows.Close()
	return students, nil
}

// * Course Model
const queryCreateCourse = `INSERT INTO course (cid, coursename) VALUES ($1, $2);`
const queryDeleteCourse = `DELETE FROM course WHERE cid=$1 RETURNING cid;`
const queryUpdateCourse = `UPDATE course SET cid=$1, coursename=$2 WHERE cid=$3 RETURNING cid;`
const queryGetCourse = `SELECT cid, coursename FROM course WHERE cid=$1;`
const queryGetAllCourses = `SELECT * FROM course;`

type Course struct {
	CourseID   int64  `json:"courseId"`
	CourseName string `json:"courseName"`
}

func (c *Course) Create() error {
	_, err := postgres.Db.Exec(queryCreateCourse, c.CourseID, c.CourseName)

	return err
}

func (c *Course) Delete() error {
	return postgres.Db.QueryRow(queryDeleteCourse, c.CourseID).Scan(&c.CourseID)
}

func (c *Course) Read() error {
	return postgres.Db.QueryRow(queryGetCourse, c.CourseID).Scan(&c.CourseID, &c.CourseName)
}
func (c *Course) Update(oldId int64) error {
	return postgres.Db.QueryRow(queryUpdateCourse, c.CourseID, c.CourseName, oldId).Scan(&c.CourseID)
}

func GetAllCourses() ([]Course, error) {
	rows, getErr := postgres.Db.Query(queryGetAllCourses)
	if getErr != nil {
		return nil, getErr
	}
	courses := []Course{}

	for rows.Next() {
		var c Course
		if dbErr := rows.Scan(&c.CourseID, &c.CourseName); dbErr != nil {
			return nil, dbErr
		}
		courses = append(courses, c)
	}
	rows.Close()
	return courses, nil
}

// * Signup

const querysignup = `INSERT INTO admin (firstname, lastname, email, password) VALUES ($1, $2, $3, $4);`

type Signup struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (s *Signup) Create() error {
	_, err := postgres.Db.Exec(querysignup, s.FirstName, s.LastName, s.Email, s.Password)

	return err
}
