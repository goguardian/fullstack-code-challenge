package datastore

import (
	"context"
)

// Client describes the interface for interacting with a datastore client.
type Client interface {
	// ListClassroomsAndStudents returns classrooms matching the classroomIDs parameter. If the
	// provided argument is empty or nil, it returns all classrooms.
	ListClassroomsAndStudents(ctx context.Context) ([]Classroom, error)
}

// Student represents a GoGuardian student.
type Student struct {
	ID          uint64
	Name        string
	ClassroomID uint32
}

// Classroom represents a GoGuardian classroom.
type Classroom struct {
	ID       uint32
	Name     string
	Students []Student
}
