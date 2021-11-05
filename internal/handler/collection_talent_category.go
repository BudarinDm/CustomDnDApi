package handler

import (
	"DnDApi/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTalentCategory(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.TalentCategory
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	collectionId, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	id, err = h.service.TalentCategory.Create(id, input, collectionId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getTalentCategoryByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	talentCategoryId, err := ParsId(c.Param("talent_category_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	talentCategory, err := h.service.TalentCategory.GetTalentCategoryById(userId, talentCategoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"TalentCategory": talentCategory,
	})
}

type getAllTalentCategoryResponse struct {
	Data []model.TalentCategory `json:"data"`
}

func (h *Handler) getAllTalentCategory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	talentCategories, err := h.service.TalentCategory.GetAllTalentCategoryByCollectionId(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTalentCategoryResponse{
		Data: talentCategories,
	})
}

func (h *Handler) updateTalentCategory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	talentCategoryId, err := ParsId(c.Param("talent_category_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	var input model.TalentCategory
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.TalentCategory.UpdateTalentCategory(userId, talentCategoryId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "updated",
	})
}

func (h *Handler) deleteTalentCategory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	talentCategoryId, err := ParsId(c.Param("talent_category_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	err = h.service.TalentCategory.DeleteTalentCategory(userId, talentCategoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "delete",
	})
}
