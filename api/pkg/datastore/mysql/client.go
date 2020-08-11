package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/goguardian/fullstack-code-challenge/api/pkg/datastore"
	"github.com/pkg/errors"
)

// Config represents the configuration for a datastore client
type Config struct {
	DatabaseAddress string
	ReadTimeout     time.Duration
}

// New creates and returns a new datastore client.
func New(conf *Config) (datastore.Client, error) {
	if conf.DatabaseAddress == "" {
		return nil, errors.New("invalid database address")
	}
	if conf.ReadTimeout <= 0 {
		return nil, errors.New("invalid read timeout")
	}

	mysqlClient, err := sql.Open("mysql", conf.DatabaseAddress)
	if err != nil {
		return nil, errors.Wrap(err, "error creating MySQL client")
	}

	return &client{
		mysqlClient: mysqlClient,
		readTimeout: conf.ReadTimeout,
	}, nil
}

type client struct {
	mysqlClient *sql.DB
	readTimeout time.Duration
}

func (c *client) ListClassroomsAndStudents(ctx context.Context) ([]datastore.Classroom, error) {

	timeoutCtx, cancel := context.WithTimeout(ctx, c.readTimeout)
	defer cancel()

	query := `SELECT
		classrooms.id AS classroomID,
		classroom_students.id AS studentID,
		classrooms.name AS classroomName,
		classroom_students.name AS studentName
		FROM classrooms LEFT JOIN classroom_students ON classrooms.id = classroom_students.classroom_id`

	rows, err := c.mysqlClient.QueryContext(
		timeoutCtx,
		query,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error listing classrooms and students")
	}
	defer rows.Close()

	classroomsMap := make(map[int32]string)
	classroomStudentsMap := make(map[int32][]datastore.Student)
	for rows.Next() {
		var classroomID int32
		var studentID int32
		var classroomName string
		var studentName string

		if err := rows.Scan(&classroomID, &studentID, &classroomName, &studentName); err != nil {
			return nil, errors.Wrap(err, "error reading classrooms and students")
		}

		classroomsMap[classroomID] = classroomName

		if students, found := classroomStudentsMap[classroomID]; found {
			classroomStudentsMap[classroomID] = append(students, datastore.Student{
				ID:   uint64(studentID),
				Name: studentName,
			})
		} else {
			classroomStudentsMap[classroomID] = []datastore.Student{
				datastore.Student{
					ID:   uint64(studentID),
					Name: studentName,
				},
			}
		}
	}

	var classrooms []datastore.Classroom
	for classroomID, students := range classroomStudentsMap {
		classrooms = append(classrooms, datastore.Classroom{
			ID:       uint32(classroomID),
			Name:     classroomsMap[classroomID],
			Students: students,
		})
	}

	return classrooms, nil
}
