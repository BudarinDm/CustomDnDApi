package handler

import (
	"DnDApi/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createSpell(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.Spell
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	collectionId, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	id, err = h.service.Spell.Create(id, input, collectionId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getSpellByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	spellId, err := ParsId(c.Param("weapon_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	spell, err := h.service.Spell.GetSpellById(userId, spellId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Spell": spell,
	})
}

type getAllSpellResponse struct {
	Data []model.Spell `json:"data"`
}

func (h *Handler) getAllSpell(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	spells, err := h.service.Spell.GetAllSpellByCollectionId(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllSpellResponse{
		Data: spells,
	})
}

func (h *Handler) updateSpell(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	spellId, err := ParsId(c.Param("spell_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	var input model.Spell
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Spell.UpdateSpell(userId, spellId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "updated",
	})
}

func (h *Handler) deleteSpell(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	spellId, err := ParsId(c.Param("spell_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	err = h.service.Spell.DeleteSpell(userId, spellId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "delete",
	})
}
