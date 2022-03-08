package grpc

//
//import (
//	"context"
//	"github.com/reb-felipe/grpc-from-rest/domain/entity"
//	"github.com/reb-felipe/grpc-from-rest/internal/controller/http/grpc/presenter"
//	grpc2 "github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/grpc"
//	"google.golang.org/protobuf/types/known/emptypb"
//)
//
//func NewClient(users grpc2.UsersClient) *Client {
//	return &Client{users: users}
//}
//
//type Client struct {
//	users grpc2.UsersClient
//}
//
//func (c *Client) CreateUser(ctx context.Context, name string, coordinates []float64) (*entity.User, error) {
//	user, err := c.users.Create(ctx, &presenter.UserInput{
//		Name:        name,
//		Coordinates: coordinates,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return user.ToEntity(), nil
//}
//
//func (c *Client) UpdateUser(ctx context.Context, userID string, name string, coordinates []float64) (*entity.User, error) {
//	user, err := c.users.Update(ctx, &presenter.UpdateUserInput{
//		Id: userID,
//		Data: &presenter.UserInput{
//			Name:        name,
//			Coordinates: coordinates,
//		},
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return user.ToEntity(), nil
//}
//
//func (c *Client) ListUsers(ctx context.Context) ([]*entity.User, error) {
//	result, err := c.users.List(ctx, new(emptypb.Empty))
//	if err != nil {
//		return nil, err
//	}
//
//	users := make([]*entity.User, len(result.Users))
//	for i, v := range result.Users {
//		users[i] = v.ToEntity()
//	}
//
//	return users, nil
//}
//
//func (c *Client) DeleteUser(ctx context.Context, userID string) error {
//	_, err := c.users.Delete(ctx, &presenter.IDInput{Id: userID})
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (c *Client) GetUser(ctx context.Context, id string) (*entity.User, error) {
//	user, err := c.users.Get(ctx, &presenter.IDInput{Id: id})
//	if err != nil {
//		return nil, err
//	}
//
//	return user.ToEntity(), nil
//}
