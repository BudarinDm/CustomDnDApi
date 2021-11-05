package handler

import (
	"DnDApi/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createClass(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.Class
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	collectionId, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	id, err = h.service.Class.Create(id, input, collectionId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getClassByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	classId, err := ParsId(c.Param("class_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	class, err := h.service.Class.GetClassById(userId, classId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Class": class,
	})
}

type getAllClassResponse struct {
	Data []model.Class `json:"data"`
}

func (h *Handler) getAllClass(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	classes, err := h.service.Class.GetAllClassByCollectionId(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllClassResponse{
		Data: classes,
	})
}

func (h *Handler) updateClass(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	classId, err := ParsId(c.Param("class_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	var input model.Class
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Class.UpdateClass(userId, classId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "updated",
	})
}

func (h *Handler) deleteClass(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	classId, err := ParsId(c.Param("class_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	err = h.service.Class.DeleteClass(userId, classId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "delete",
	})
}
