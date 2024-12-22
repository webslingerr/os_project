package handler

import (
	"app/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var createPosrt models.CreatePost

	err := c.ShouldBindJSON(&createPosrt)
	if err != nil {
		h.handlerResponse(c, "create post", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Post().Create(context.Background(), &createPosrt)
	if err != nil {
		h.handlerResponse(c, "storage create post", http.StatusInternalServerError, err.Error())
		return
	}

	post, err := h.storages.Post().GetById(context.Background(), &models.PostPrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "storage get by id post", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create post", http.StatusCreated, post)
}

func (h *Handler) GetByIdPost(c *gin.Context) {
	id := c.Param("id")

	post, err := h.storages.Post().GetById(context.Background(), &models.PostPrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "Storage get by id post", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get by id post", http.StatusOK, post)
}

func (h *Handler) GetListPost(c *gin.Context) {
	var getListPost models.GetListPostRequest

	err := c.ShouldBindJSON(&getListPost)
	if err != nil {
		h.handlerResponse(c, "get list post", http.StatusBadRequest, err.Error())
		return
	}

	if getListPost.Limit == 0 {
		getListPost.Limit = 10
	}

	posts, err := h.storages.Post().GetList(context.Background(), &models.GetListPostRequest{
		Offset:         getListPost.Offset,
		Limit:          getListPost.Limit,
		Search:         getListPost.Search,
		Status:         getListPost.Status,
		Region:         getListPost.Region,
		RealEstateType: getListPost.RealEstateType,
		UserId:         getListPost.UserId,
	})
	if err != nil {
		h.handlerResponse(c, "Storage get list post", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get list post", http.StatusOK, posts)
}

func (h *Handler) UpdatePost(c *gin.Context) {
	var updatePost models.UpdatePost

	id := c.Param("id")

	err := c.ShouldBindJSON(&updatePost)
	if err != nil {
		h.handlerResponse(c, "Update post", http.StatusBadRequest, err.Error())
		return
	}
	updatePost.ID = id

	rowsAffected, err := h.storages.Post().Update(context.Background(), &updatePost)
	if err != nil {
		h.handlerResponse(c, "Storage update post", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage update post", http.StatusBadRequest, "no rows affected")
		return
	}

	post, err := h.storages.Post().GetById(context.Background(), &models.PostPrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "Storage get by id post", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Update post", http.StatusOK, post)
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	var updateStatus models.UpdateStatus

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateStatus)
	if err != nil {
		h.handlerResponse(c, "Update status", http.StatusBadRequest, err.Error())
		return
	}
	updateStatus.ID = id

	err = h.storages.Post().UpdateStatus(context.Background(), &updateStatus)
	if err != nil {
		h.handlerResponse(c, "Storage update status", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Update Status", http.StatusOK, "")
}

func (h *Handler) DeletePost(c *gin.Context) {
	id := c.Param("id")

	rowsAffected, err := h.storages.Post().Delete(context.Background(), &models.PostPrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "Storage delete post", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage delete post", http.StatusBadRequest, "no rows affected")
		return
	}

	h.handlerResponse(c, "Delete post", http.StatusNoContent, "Deleted Successfully")
}
