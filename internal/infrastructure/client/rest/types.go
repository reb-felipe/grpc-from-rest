package rest

import (
	"github.com/reb-felipe/grpc-from-rest/domain/entity"
	"time"
)

type CreateOrUpdateUserPayload struct {
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type ListUserResponse struct {
	Result []User `json:"result"`
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
