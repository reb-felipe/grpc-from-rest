package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/reb-felipe/grpc-from-rest/domain/service"
	"github.com/reb-felipe/grpc-from-rest/internal/controller/http/rest/presenter"
	"net/http"
)

func NewUsers(users *service.Users) *UsersController {
	return &UsersController{users: users}
}

type UsersController struct {
	users *service.Users
}

func (r *UsersController) CreateUser(ctx *gin.Context) {
	var request presenter.CreateOrUpdateUserRequest
	if err := ctx.BindJSON(&request); err != nil {
		return
	}

	user, err := r.users.CreateUser(ctx, request.ToEntity())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.ErrorMessage{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, presenter.UserFromEntity(user))
}

func (r *UsersController) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("userID")
	var request presenter.CreateOrUpdateUserRequest
	if err := ctx.BindJSON(&request); err != nil {
		return
	}

	update := request.ToEntity()
	update.ID = userID

	update, err := r.users.UpdateUser(ctx, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.ErrorMessage{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, presenter.UserFromEntity(update))
}

func (r *UsersController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("userID")

	err := r.users.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.ErrorMessage{Message: err.Error()})
		return
	}
}

func (r *UsersController) ListUsers(ctx *gin.Context) {
	users, err := r.users.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.ErrorMessage{Message: err.Error()})
	}

	usersResponse := make([]*presenter.User, len(users))
	for i, v := range users {
		usersResponse[i] = presenter.UserFromEntity(v)
	}

	ctx.JSON(http.StatusOK, presenter.ListUserReponse{
		Results: usersResponse,
	})
}

func (r *UsersController) GetUser(ctx *gin.Context) {
	userID := ctx.Param("userID")

	user, err := r.users.GetUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.ErrorMessage{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, presenter.UserFromEntity(user))
}
