package handler

import (
	"app/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.User().Create(context.Background(), &createUser)
	if err != nil {
		h.handlerResponse(c, "storage create user", http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.storages.User().GetById(context.Background(), &models.UserPrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "storage get by id user", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create user", http.StatusCreated, user)
}

func (h *Handler) GetByIdUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.storages.User().GetById(context.Background(), &models.UserPrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "Storage get by id user", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get by id user", http.StatusOK, user)
}

func (h *Handler) GetListUser(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Get list user", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Get list user", http.StatusBadRequest, "invalid limit")
		return
	}

	users, err := h.storages.User().GetList(context.Background(), &models.GetListUserRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "Storage get list user", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get list user", http.StatusOK, users)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var updateUser models.UpdateUser

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		h.handlerResponse(c, "Update user", http.StatusBadRequest, err.Error())
		return
	}
	updateUser.ID = id

	rowsAffected, err := h.storages.User().Update(context.Background(), &updateUser)
	if err != nil {
		h.handlerResponse(c, "Storage update user", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage update user", http.StatusBadRequest, "no rows affected")
		return
	}

	user, err := h.storages.User().GetById(context.Background(), &models.UserPrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "Storage get by id user", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Update user", http.StatusOK, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	rowsAffected, err := h.storages.User().Delete(context.Background(), &models.UserPrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "Storage delete category", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage delete user", http.StatusBadRequest, "no rows affected")
		return
	}

	h.handlerResponse(c, "Delete user", http.StatusNoContent, "Deleted Successfully")
}
