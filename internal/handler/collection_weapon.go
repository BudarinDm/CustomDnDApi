package handler

import (
	"DnDApi/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createWeapon(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.Weapon
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	collectionId, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	id, err = h.service.Weapon.Create(id, input, collectionId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getWeaponByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	weaponId, err := ParsId(c.Param("weapon_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	weapon, err := h.service.Weapon.GetWeaponById(userId, weaponId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Weapon": weapon,
	})
}

type getAllWeaponResponse struct {
	Data []model.Weapon `json:"data"`
}

func (h *Handler) getAllWeapon(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	weapons, err := h.service.Weapon.GetAllWeaponByCollectionId(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWeaponResponse{
		Data: weapons,
	})
}

func (h *Handler) updateWeapon(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	weaponId, err := ParsId(c.Param("weapon_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	var input model.Weapon
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Weapon.UpdateWeapon(userId, weaponId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "updated",
	})
}

func (h *Handler) deleteWeapon(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	weaponId, err := ParsId(c.Param("weapon_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	err = h.service.Weapon.DeleteWeapon(userId, weaponId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "delete",
	})
}
