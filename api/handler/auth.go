package handler

import (
	"app/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var registerUser models.CreateUser

	err := c.ShouldBindJSON(&registerUser)
	if err != nil {
		h.handlerResponse(c, "register user", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.User().Create(context.Background(), &models.CreateUser{
		Email:    registerUser.Email,
		Fullname: registerUser.Fullname,
		Password: registerUser.Password,
		Address:  registerUser.Address,
		Type:     registerUser.Type,
	})
	if err != nil {
		h.handlerResponse(c, "storage register/create user", http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.storages.User().GetById(context.Background(), &models.UserPrimaryKey{
		ID: id,
	})
	if err != nil {
		h.handlerResponse(c, "storage get by id user inside register", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create user", http.StatusCreated, user)
}

func (h *Handler) LoginUser(c *gin.Context) {
	var loginUser models.Login

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		h.handlerResponse(c, "login user", http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.storages.User().GetByLoginAndPassword(context.Background(), &models.Login{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	})
	if err != nil {
		h.handlerResponse(c, "storage login user", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "ok", http.StatusOK, user)
}
