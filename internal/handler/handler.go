package handler

import (
	"DnDApi/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth", h.userIdentity)
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		collection := api.Group("/collection")
		{
			collection.POST("/", h.createCollection)
			collection.GET("/type/:type", h.getAllCollectionTheType)
			collection.GET("/:id", h.getCollection)
			collection.PUT("/:id", h.updateCollection)
			collection.DELETE("/:id", h.deleteCollection)

			weapon := collection.Group(":id/weapon")
			{
				weapon.POST("/", h.createWeapon)
				weapon.GET("/", h.getAllWeapon)
				weapon.GET("/:weapon_id", h.getWeaponByID)
				weapon.PUT("/:weapon_id", h.updateWeapon)
				weapon.DELETE("/:weapon_id", h.deleteWeapon)
			}

			spell := collection.Group(":id/spell")
			{
				spell.POST("/", h.createSpell)
				spell.GET("/", h.getAllSpell)
				spell.GET("/:spell_id", h.getSpellByID)
				spell.PUT("/:spell_id", h.updateSpell)
				spell.DELETE("/:spell_id", h.deleteSpell)
			}

			armor := collection.Group(":id/armor")
			{
				armor.POST("/", h.createArmor)
				armor.GET("/", h.getAllArmor)
				armor.GET("/:armor_id", h.getArmorByID)
				armor.PUT("/:armor_id", h.updateArmor)
				armor.DELETE("/:armor_id", h.deleteArmor)
			}

			talentCategory := collection.Group(":id/tcategory")
			{
				talentCategory.POST("/", h.createTalentCategory)
				talentCategory.GET("/", h.getAllTalentCategory)
				talentCategory.GET("/:talent_category_id", h.getTalentCategoryByID)
				talentCategory.PUT("/:talent_category_id", h.updateTalentCategory)
				talentCategory.DELETE("/:talent_category_id", h.deleteTalentCategory)

				talent := talentCategory.Group(":talent_category_id/talent")
				{
					talent.POST("/", h.createTalent)
					talent.GET("/", h.getAllTalent)
					talent.GET("/:talent_id", h.getTalentByID)
					talent.PUT("/:talent_id", h.updateTalent)
					talent.DELETE("/:talent_id", h.deleteTalent)
				}
			}

			race := collection.Group(":id/race")
			{
				race.POST("/", h.createRace)
				race.GET("/", h.getAllRace)
				race.GET("/:race_id", h.getRaceByID)
				race.PUT("/:race_id", h.updateRace)
				race.DELETE("/:race_id", h.deleteRace)
			}

			class := collection.Group(":id/class")
			{
				class.POST("/", h.createClass)
				class.GET("/", h.getAllClass)
				class.GET("/:class_id", h.getClassByID)
				class.PUT("/:class_id", h.updateClass)
				class.DELETE("/:class_id", h.deleteClass)
			}
		}
	}

	return router
}
