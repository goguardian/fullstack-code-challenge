package service

import (
	"context"
	"errors"
	"net/http"
	gw "proto/fullstack-code-challenge"
)

func (s *service) GetClassroomsAndStudents(ctx context.Context, req *gw.GetClassroomsAndStudentsRequest) (*gw.GetClassroomsAndStudentsResponse, error) {
	products, err := s.datastoreClient.ListClassroomsAndStudents(ctx, nil)
	if err != nil {
		log.Error(err)
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	res := &gw.GetClassroomsAndStudentsResponse{
		Classrooms: make(map[uint64]*gw.Classroom, len(classrooms)),
	}

	for _, classroom := range classrooms {
		res.Classrooms[classroom.ID] = &gw.Classroom{
			Id:       classroom.ID,
			Name:     classroom.Name,
			Students: classroom.Students,
		}
	}

	return res, nil
}
