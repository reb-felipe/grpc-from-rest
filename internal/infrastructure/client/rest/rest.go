package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/reb-felipe/grpc-from-rest/domain/entity"
	"io/ioutil"
	"net/http"
)

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

type Client struct {
	addr string
}

func (c *Client) CreateUser(ctx context.Context, name string, coordinates []float64) (*entity.User, error) {
	b, err := json.Marshal(CreateOrUpdateUserPayload{
		Name:        name,
		Coordinates: coordinates,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s", c.addr), bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(b, &user); err != nil {
		return nil, err
	}

	return user.ToEntity(), nil
}

func (c *Client) UpdateUser(ctx context.Context, userID string, name string, coordinates []float64) (*entity.User, error) {
	b, err := json.Marshal(CreateOrUpdateUserPayload{
		Name:        name,
		Coordinates: coordinates,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", c.addr, userID), bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(b, &user); err != nil {
		return nil, err
	}

	return user.ToEntity(), nil
}

func (c *Client) ListUsers(ctx context.Context) ([]*entity.User, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s", c.addr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ListUserResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, err
	}

	users := make([]*entity.User, len(response.Result))
	for i, v := range response.Result {
		users[i] = v.ToEntity()
	}

	return users, nil
}

func (c *Client) DeleteUser(ctx context.Context, userID string) error {
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", c.addr, userID), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var errMsg ErrorMessage
	if err := json.Unmarshal(b, &errMsg); err != nil {
		return err
	}

	return errors.New(errMsg.Message)
}

func (c *Client) GetUser(ctx context.Context, id string) (*entity.User, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.addr, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(b, &user); err != nil {
		return nil, err
	}

	return user.ToEntity(), nil
}
