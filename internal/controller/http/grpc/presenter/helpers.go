package presenter

import (
	"github.com/reb-felipe/grpc-from-rest/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserInput) ToEntity() *entity.User {
	return &entity.User{
		Name:        u.Name,
		Coordinates: u.Coordinates,
	}
}

func (u *UpdateUserInput) ToEntity() *entity.User {
	user := u.Data.ToEntity()
	user.ID = u.Id
	return user
}

func UserFromEntity(u *entity.User) *User {
	user := &User{
		Id:          u.ID,
		Name:        u.Name,
		CreatedAt:   timestamppb.New(u.CreatedAt),
		Coordinates: u.Coordinates,
	}

	if u.UpdatedAt != nil {
		user.UpdatedAt = timestamppb.New(*u.UpdatedAt)
	}

	return user
}

func (u *User) ToEntity() *entity.User {
	user := &entity.User{
		ID:          u.Id,
		Name:        u.Name,
		CreatedAt:   u.CreatedAt.AsTime(),
		Coordinates: u.Coordinates,
	}

	if u.UpdatedAt != nil {
		t := u.UpdatedAt.AsTime()
		user.UpdatedAt = &t
	}

	return user
}
