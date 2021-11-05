package handler

import (
	"DnDApi/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTalent(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.Talent
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	categoryId, err := ParsId(c.Param("talent_category_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	id, err = h.service.Talent.Create(id, input, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getTalentByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	talentId, err := ParsId(c.Param("talent_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	talent, err := h.service.Talent.GetTalentById(userId, talentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Talent": talent,
	})
}

type getAllTalentResponse struct {
	Data []model.Talent `json:"data"`
}

func (h *Handler) getAllTalent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	talents, err := h.service.Talent.GetAllTalentByCategoryId(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTalentResponse{
		Data: talents,
	})
}

func (h *Handler) updateTalent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	talentId, err := ParsId(c.Param("talent_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	var input model.Talent
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Talent.UpdateTalent(userId, talentId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "updated",
	})
}

func (h *Handler) deleteTalent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	talentId, err := ParsId(c.Param("talent_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	err = h.service.Talent.DeleteTalent(userId, talentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "delete",
	})
}
