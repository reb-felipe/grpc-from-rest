package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/reb-felipe/grpc-from-rest/domain/service"
	"net/http"
)

type Rest struct {
	users service.Users
}

func (r *Rest) CreateUser(ctx *gin.Context) {
	var request CreateOrUpdateUserRequest
	if err := ctx.BindJSON(&request); err != nil {
		return
	}

	user, err := r.users.CreateUser(ctx, request.ToEntity())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, userFromEntity(user))
}

func (r *Rest) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("userID")
	var request CreateOrUpdateUserRequest
	if err := ctx.BindJSON(&request); err != nil {
		return
	}

	update := request.ToEntity()
	update.ID = userID

	update, err := r.users.UpdateUser(ctx, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, userFromEntity(update))
}

func (r *Rest) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("userID")

	err := r.users.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}
}

func (r *Rest) ListUsers(ctx *gin.Context) {
	users, err := r.users.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	usersResponse := make([]*User, len(users))
	for i, v := range users {
		usersResponse[i] = userFromEntity(v)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"result": usersResponse,
	})
}

func (r *Rest) GetUser(ctx *gin.Context) {
	userID := ctx.Param("userID")

	user, err := r.users.GetUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, userFromEntity(user))
}
