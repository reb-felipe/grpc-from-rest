package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/reb-felipe/grpc-from-rest/domain/entity"
	"os"
	"strconv"
	"strings"
	"time"
)

func NewApp(t *Terminal, client UsersClient) *App {
	return &App{
		terminal: t,
		client:   client,
	}
}

type App struct {
	terminal *Terminal
	client   UsersClient
}

func (a *App) Run() {
	a.terminal.Clear()

	for {
		a.WriteMainOptions()
		opt, err := a.ParseMenuOptions()
		if err != nil {
			a.terminal.Clear()
			a.terminal.Write("unrecognized operation\n")
			continue
		}

		switch opt {
		case 0:
			os.Exit(0)
		case 1:
			a.CreateUser()
		case 2:
			a.GetUser()
		case 3:
			a.UpdateUser()
		case 4:
			a.DeleteUser()
		case 5:
			a.ListUsers()
		default:
			a.terminal.Clear()
			a.terminal.Write("unrecognized operation\n")
			continue
		}
	}
}

func (a *App) WriteMainOptions() {
	options := "1 - Create User\n2 - Get User\n3 - Update User\n4 - Delete User\n5 - List User\n0 - Exit\n"
	a.terminal.Write(options)
}

func (a *App) ParseMenuOptions() (int, error) {
	input, err := a.terminal.WaitInput(">")
	if err != nil {
		return -1, err
	}

	i, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}

	if i > 5 {
		return -1, errors.New("invalid option")
	}

	return i, nil
}

type UserData struct {
	Name        string
	Coordinates []float64
}

func (a *App) DeleteUser() {
	a.terminal.Clear()
	a.terminal.WriteLn("Please write the attributes of the user")
	var userID string
	for userID == "" {
		input, err := a.terminal.WaitInput("User ID >")
		if err != nil {
			a.terminal.WriteLn("error reading id")
			continue
		}

		if len(strings.Trim(input, "")) == 0 {
			a.terminal.WriteLn("invalid user ID")
			continue
		}

		userID = input
	}

	err := a.client.DeleteUser(context.Background(), userID)
	if err != nil {
		a.terminal.WriteLn("Error deleting user: " + err.Error())
		a.EnterAndClear()
		return
	}
	a.terminal.Clear()
	a.terminal.WriteLn("User deleted with success")
	a.EnterAndClear()
}

func (a *App) GetUser() {
	a.terminal.Clear()
	a.terminal.WriteLn("Please write the attributes of the user")
	var userID string
	for userID == "" {
		input, err := a.terminal.WaitInput("User ID >")
		if err != nil {
			a.terminal.WriteLn("error reading id")
			continue
		}

		if len(strings.Trim(input, "")) == 0 {
			a.terminal.WriteLn("invalid user ID")
			continue
		}

		userID = input
	}

	user, err := a.client.GetUser(context.Background(), userID)
	if err != nil {
		a.terminal.WriteLn("Error deleting user: " + err.Error())
		a.EnterAndClear()
		return
	}
	a.terminal.Clear()
	a.WriteUser(user)
	a.terminal.Write("\n")
	a.EnterAndClear()
}

func (a *App) ListUsers() {
	a.terminal.Clear()
	closeFn := a.terminal.SetLoading("Listing users")
	users, err := a.client.ListUsers(context.Background())
	closeFn()
	if err != nil {
		a.terminal.WriteLn("Error listing users: " + err.Error())
		a.EnterAndClear()
		return
	}
	a.terminal.Clear()
	for _, v := range users {
		a.WriteUser(v)
		a.terminal.Write("\n\n")
	}
	a.EnterAndClear()
}

func (a *App) UpdateUser() {
	a.terminal.Clear()
	a.terminal.WriteLn("Please write the attributes of the user\n")
	var userID string
	for userID == "" {
		input, err := a.terminal.WaitInput("User ID >")
		if err != nil {
			a.terminal.WriteLn("error reading id")
			continue
		}

		if len(strings.Trim(input, "")) == 0 {
			a.terminal.WriteLn("invalid user ID")
			continue
		}

		userID = input
	}

	userData := a.ReadUserData()
	a.terminal.Clear()
	closeFn := a.terminal.SetLoading("Updating user")
	user, err := a.client.UpdateUser(context.Background(), userID, userData.Name, userData.Coordinates)
	closeFn()
	if err != nil {
		a.terminal.WriteLn("Error updating user: " + err.Error())
		a.EnterAndClear()
		return
	}
	a.terminal.Clear()
	a.terminal.WriteLn("User updated")
	a.WriteUser(user)
	a.terminal.Write("\n\n")
	a.EnterAndClear()
}

func (a *App) CreateUser() {
	a.terminal.Clear()
	a.terminal.WriteLn("Please write the attributes of the new user\n")
	userData := a.ReadUserData()
	a.terminal.Clear()
	closeFn := a.terminal.SetLoading("Creating user")
	user, err := a.client.CreateUser(context.Background(), userData.Name, userData.Coordinates)
	closeFn()
	if err != nil {
		a.terminal.WriteLn("Error creating user: " + err.Error())
		a.EnterAndClear()
		return
	}
	a.terminal.Clear()
	a.terminal.WriteLn("User created\n")
	a.WriteUser(user)
	a.terminal.Write("\n\n")
	a.EnterAndClear()
}

func (a *App) WriteUser(user *entity.User) {
	a.terminal.WriteF("USER %s\n", user.ID)
	a.terminal.WriteF("\n\tName: %s", user.Name)
	if len(user.Coordinates) == 2 {
		a.terminal.WriteF("\n\tLat: %f", user.Coordinates[0])
		a.terminal.WriteF("\n\tLng: %f", user.Coordinates[1])
	} else {
		a.terminal.Write("\n\tCoordinates: invalid coordinates")
	}
	a.terminal.WriteF("\n\tCreatedAt: %s", user.CreatedAt.Format(time.RFC3339))
	if user.UpdatedAt != nil {
		a.terminal.WriteF("\n\tUpdatedAt: %s", user.UpdatedAt.Format(time.RFC3339))
	}
}

func (a *App) EnterAndClear() {
	_, _ = a.terminal.WaitInput("Press enter to continue")
	a.terminal.Clear()
}

func (a *App) ReadUserData() *UserData {
	prefixList := []string{"Name >", "Lat (float64) >", "Lng (float64) >"}
	step := 0
	hasError := false
	var name string
	coordinates := make([]float64, 2)
	for step < 3 {
		if hasError {
			a.terminal.Write("\ninvalid argument")
		}

		input, err := a.terminal.WaitInput(fmt.Sprintf("%s", prefixList[step]))
		if err != nil && err.Error() != "expected newline" {
			hasError = true
			continue
		}

		switch step {
		case 0:
			t := strings.Trim(input, "")
			if len(t) == 0 {
				hasError = true
				continue
			}
			name = t
			hasError = false
		case 1:
			f, err := strconv.ParseFloat(input, 64)
			if err != nil {
				hasError = true
				continue
			}
			coordinates[0] = f
			hasError = false
		case 2:
			f, err := strconv.ParseFloat(input, 64)
			if err != nil {
				hasError = true
				continue
			}
			coordinates[1] = f
			hasError = false
		}

		step++
	}

	return &UserData{
		Name:        name,
		Coordinates: coordinates,
	}
}
