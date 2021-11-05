package handler

import (
	"DnDApi/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createArmor(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.Armor
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	collectionId, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	id, err = h.service.Armor.Create(id, input, collectionId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getArmorByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	armorId, err := ParsId(c.Param("armor_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	armor, err := h.service.Armor.GetArmorById(userId, armorId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Armor": armor,
	})
}

type getAllArmorResponse struct {
	Data []model.Armor `json:"data"`
}

func (h *Handler) getAllArmor(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	armors, err := h.service.Armor.GetAllArmorByCollectionId(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllArmorResponse{
		Data: armors,
	})
}

func (h *Handler) updateArmor(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	armorId, err := ParsId(c.Param("armor_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	var input model.Armor
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Armor.UpdateArmor(userId, armorId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "updated",
	})
}

func (h *Handler) deleteArmor(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	armorId, err := ParsId(c.Param("armor_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	err = h.service.Armor.DeleteArmor(userId, armorId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "delete",
	})
}
