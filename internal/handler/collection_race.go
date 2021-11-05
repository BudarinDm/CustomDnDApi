package handler

import (
	"DnDApi/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createRace(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.Race
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	collectionId, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	id, err = h.service.Race.Create(id, input, collectionId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getRaceByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	raceId, err := ParsId(c.Param("race_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	race, err := h.service.Race.GetRaceById(userId, raceId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Race": race,
	})
}

type getAllRaceResponse struct {
	Data []model.Race `json:"data"`
}

func (h *Handler) getAllRace(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := ParsId(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	races, err := h.service.Race.GetAllRaceByCollectionId(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllRaceResponse{
		Data: races,
	})
}

func (h *Handler) updateRace(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	raceId, err := ParsId(c.Param("race_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	var input model.Race
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Race.UpdateRace(userId, raceId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "updated",
	})
}

func (h *Handler) deleteRace(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	raceId, err := ParsId(c.Param("race_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "parse id error")
		return
	}

	err = h.service.Race.DeleteRace(userId, raceId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "delete",
	})
}
