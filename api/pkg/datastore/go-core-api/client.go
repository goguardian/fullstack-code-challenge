package gocoreapi

import (
	"context"
	"database/sql"
	"time"

	"github.com/goguardian/Development/services/examples/go-grpc/pkg/datastore"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
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

	// Create sql client here
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
	mysqlClient pb.GoCoreAPIClient
	readTimeout time.Duration
}

func (c *client) ListClassroomsAndStudents(ctx context.Context, classroomIDs []uint32) ([]datastore.Classroom, error) {

	classrooms := []datastore.Classroom{}

	timeoutCtx, cancel := context.WithTimeout(ctx, c.readTimeout)
	defer cancel()

	query := `SELECT
		classrooms.id AS classroomID,
		classroom_students.id AS studentID,
		classrooms.name AS classroomName,
		classroom_students.name AS studentName
		FROM classrooms LEFT JOIN classroom_students ON classrooms.id = classroom_students.classroom_id`

	if len(classroomIDs) > 0 {
		query = fmt.Sprintf("%s WHERE classrooms.id IN (?)")
	}

	int32ClassroomIDs := make([]int32, len(classroomIDs))
	for i, classroomID := range classroomIDs {
		int32ClassroomIDs[i] = int32(classroomID)
	}

	rows, err := c.mysqlClient.QueryContext(
		timeoutCtx,
		classroomIDs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error listing classrooms and students")
	}
	defer rows.Close()

	classroomsMap = make(map[int32]string)
	classroomStudentsMap = make(map[int32][]datastore.Student)
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
			students = append(students, datastore.Student{
				ID: studentID,
				Name: studentName,
			})
		} else {
			classroomStudentsMap[classroomID] := make[]datastore.Student{
				datastore.Student{
					ID: studentID,
					Name: studentName,
				}
			}
		}
	}

	var classrooms []datastore.Classroom
	for classroomID, students := range classroomStudents {
		classrooms = append(classrooms, datastore.Classroom{
			ID:   uint32(classroomID),
			Name: classroomsMap[classroomID],
			Students: students
		})
	}

	return classrooms, nil
}
