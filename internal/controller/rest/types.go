package rest

import (
	"github.com/reb-felipe/grpc-from-rest/domain/entity"
	"time"
)

type CreateOrUpdateUserRequest struct {
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates"`
}

func (c *CreateOrUpdateUserRequest) ToEntity() *entity.User {
	return &entity.User{
		ID:          "",
		Name:        c.Name,
		CreatedAt:   time.Time{},
		Coordinates: c.Coordinates,
		UpdatedAt:   nil,
	}
}

func userFromEntity(u *entity.User) *User {
	return &User{
		ID:          u.ID,
		Name:        u.Name,
		CreatedAt:   u.CreatedAt,
		Coordinates: u.Coordinates,
		UpdatedAt:   u.UpdatedAt,
	}
}

type User struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	CreatedAt   time.Time  `json:"created_at"`
	Coordinates []float64  `json:"coordinates"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:          u.ID,
		Name:        u.Name,
		CreatedAt:   u.CreatedAt,
		Coordinates: u.Coordinates,
		UpdatedAt:   u.UpdatedAt,
	}
}
