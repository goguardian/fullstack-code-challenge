package service

import (
	"context"
	"errors"
	"log"
	"net/http"

	gw "github.com/goguardian/fullstack-code-challenge/proto/gen/go/fullstack_code_challenge/v1"
)

func (s *service) GetClassroomsAndStudents(ctx context.Context, req *pb.GetClassroomsAndStudentsRequest) (*pb.GetClassroomsAndStudentsResponse, error) {
	classrooms, err := s.datastoreClient.ListClassroomsAndStudents(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	res := &pb.GetClassroomsAndStudentsResponse{
		Classrooms: make(map[uint32]*pb.Classroom, len(classrooms)),
	}

	for _, classroom := range classrooms {
		students := make([]*gw.Student, len(classroom.Students))
		for i, student := range classroom.Students {
			students[i] = &gw.Student{
				Id:   uint32(student.ID),
				Name: student.Name,
			}
		}
		res.Classrooms[classroom.ID] = &gw.Classroom{
			Id:       classroom.ID,
			Name:     classroom.Name,
			Students: students,
		}
	}

	return res, nil
}
