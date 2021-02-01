package grpc

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	m "github.com/yuriimakohon/RunecharmsCRUD/api/models"
	"github.com/yuriimakohon/RunecharmsCRUD/pkg/api"
	"google.golang.org/grpc"
	"log"
)

// gRPC implementation of storage.Storage
type Storage struct {
	cli api.CharmCRUDServiceClient
}

func New() *Storage {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Storage{api.NewCharmCRUDServiceClient(conn)}
}

func (s *Storage) Add(c m.Charm) (m.Charm, error) {
	req := &api.EntityRequest{
		Entity: &api.Charm{
			Rune:  c.Rune,
			God:   c.God,
			Power: c.Power,
		},
	}

	resp, err := s.cli.Add(context.Background(), req)
	if err != nil {
		return m.Charm{}, err
	}

	entity := resp.GetEntities()[0]

	return m.Charm{
		Id:    entity.GetId(),
		Rune:  entity.GetRune(),
		God:   entity.GetGod(),
		Power: entity.GetPower(),
	}, nil
}

func (s *Storage) GetAll() ([]m.Charm, error) {
	resp, err := s.cli.GetAll(context.Background(), &empty.Empty{})
	if err != nil {
		return []m.Charm{}, err
	}

	lenResp, err := s.cli.Len(context.Background(), &empty.Empty{})
	if err != nil {
		return []m.Charm{}, err
	}

	slice := make([]m.Charm, 0, lenResp.Value)
	for _, c := range resp.Entities {
		slice = append(
			slice,
			m.Charm{
				Id:    c.Id,
				Rune:  c.Rune,
				God:   c.God,
				Power: c.Power,
			},
		)
	}
	return slice, nil
}

func (s *Storage) Get(id int32) (m.Charm, error) {
	fmt.Println("Implement Get")
	return m.Charm{}, nil
	//for _, c := range s.Charms {
	//	if c.Id == id {
	//		return c, nil
	//	}
	//}
	//return m.Charm{}, storage.ErrNotFound
}

func (s *Storage) Delete(id int32) (m.Charm, error) {
	fmt.Println("Implement Delete")
	return m.Charm{}, nil
	//for idx, c := range s.Charms {
	//	if c.Id == id {
	//		s.Charms = append(s.Charms[:idx], s.Charms[idx+1:]...)
	//		return c, nil
	//	}
	//}
	//return m.Charm{}, storage.ErrNotFound
}

func (s *Storage) Update(id int32, u m.Charm) (m.Charm, error) {
	fmt.Println("Implement Update")
	return m.Charm{}, nil
	//for idx, c := range s.Charms {
	//	if c.Id == id {
	//		u.Id = id
	//		s.Charms[idx] = u
	//		return s.Charms[idx], nil
	//	}
	//}
	//return m.Charm{}, storage.ErrNotFound
}

func (s *Storage) Len() (int, error) {
	resp, err := s.cli.Len(context.Background(), &empty.Empty{})
	if err != nil {
		return 0, err
	}
	return int(resp.Value), nil
}
