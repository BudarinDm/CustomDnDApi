package handler

import (
	"DnDApi/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createCollection(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.Collection
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err = h.service.Collection.Create(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getCollection(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := ParsId(c.Param("id"))
	if err != nil {
		return
	}

	collection, err := h.service.Collection.GetCollectionById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"collection": collection,
	})
}

type getAllCollectionsTheTypeResponse struct {
	Data []model.Collection `json:"data"`
}

func (h *Handler) getAllCollectionTheType(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	collectionType := c.Param("type")

	collections, err := h.service.Collection.GetAllTheType(userId, collectionType)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCollectionsTheTypeResponse{
		Data: collections,
	})
}

func (h *Handler) updateCollection(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	var input model.Collection
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Collection.UpdateCollection(userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "updated",
	})
}

func (h *Handler) deleteCollection(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	err = h.service.Collection.DeleteCollection(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "delete",
	})
}
