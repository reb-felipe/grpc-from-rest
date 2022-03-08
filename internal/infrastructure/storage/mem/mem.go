package mem

import (
	"context"
	"errors"
	"github.com/reb-felipe/grpc-from-rest/domain/entity"
	"sync"
	"time"
)

var initialUsers = map[string]*entity.User{
	"xyz": {
		ID:        "xyz",
		Name:      "John Doe",
		CreatedAt: time.Now(),
		Coordinates: []float64{
			12.35, 15.231,
		},
	},
	"xpto": {
		ID:        "xpto",
		Name:      "Tyler Durden",
		CreatedAt: time.Now().UTC(),
	},
}

func NewUsers() *Users {
	return &Users{
		users: initialUsers,
		m:     new(sync.Mutex),
	}
}

type Users struct {
	users map[string]*entity.User
	m     *sync.Mutex
}

func (u *Users) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	u.m.Lock()
	defer u.m.Unlock()
	u.users[user.ID] = user

	return user, nil
}

func (u *Users) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	u.m.Lock()
	defer u.m.Unlock()
	v, ok := u.users[user.ID]
	if !ok {
		return nil, errors.New("user not found")
	}

	user.CreatedAt = v.CreatedAt
	u.users[user.ID] = user
	return user, nil
}

func (u *Users) ListUsers(ctx context.Context) ([]*entity.User, error) {
	u.m.Lock()
	defer u.m.Unlock()
	users := make([]*entity.User, len(u.users))
	c := 0
	for _, v := range u.users {
		users[c] = v
		c++
	}

	return users, nil
}

func (u *Users) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	u.m.Lock()
	defer u.m.Unlock()
	user, ok := u.users[userID]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (u *Users) DeleteUser(ctx context.Context, userID string) error {
	u.m.Lock()
	defer u.m.Unlock()
	delete(u.users, userID)
	return nil
}
