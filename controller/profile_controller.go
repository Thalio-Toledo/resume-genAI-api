package controller

import (
	"net/http"
	"strconv"

	"resume-genAI-api/model"
	"resume-genAI-api/useCase"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	useCase *useCase.ProfileUseCase
}

func NewProfileController(uc *useCase.ProfileUseCase) *ProfileController {
	return &ProfileController{useCase: uc}
}

func (ctrl *ProfileController) RegisterRoutes(r *gin.Engine) {
	profiles := r.Group("/profiles")
	{
		profiles.GET("/", ctrl.Get)
		profiles.GET(":id", ctrl.FindByID)
		profiles.POST("/", ctrl.Create)
		profiles.PUT(":id", ctrl.Update)
		profiles.DELETE(":id", ctrl.Delete)
	}
}

// Get godoc
// @Summary Lista todos os perfis
// @Description Retorna todos os perfis cadastrados
// @Tags profiles
// @Produce json
// @Success 200 {array} model.Profile
// @Router /profiles/ [get]
func (ctrl *ProfileController) Get(c *gin.Context) {
	profiles, err := ctrl.useCase.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profiles)
}

// FindByID godoc
// @Summary Busca perfil por ID
// @Description Retorna um perfil pelo ID
// @Tags profiles
// @Produce json
// @Param id path int true "ID do perfil"
// @Success 200 {object} model.Profile
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /profiles/{id} [get]
func (ctrl *ProfileController) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	profile, err := ctrl.useCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// Create godoc
// @Summary Cria um novo perfil
// @Description Cria um novo perfil com os dados informados
// @Tags profiles
// @Accept json
// @Produce json
// @Param profile body model.Profile true "Perfil a ser criado"
// @Success 201 {object} model.Profile
// @Failure 400 {object} model.ErrorResponse
// @Router /profiles/ [post]
func (ctrl *ProfileController) Create(c *gin.Context) {
	var profile model.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := ctrl.useCase.Create(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	profile.ProfileId = id
	c.JSON(http.StatusCreated, profile)
}

// Update godoc
// @Summary Atualiza um perfil
// @Description Atualiza um perfil existente pelo ID
// @Tags profiles
// @Accept json
// @Produce json
// @Param id path int true "ID do perfil"
// @Param profile body model.Profile true "Perfil atualizado"
// @Success 200 {object} model.Profile
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /profiles/{id} [put]
func (ctrl *ProfileController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var profile model.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	profile.ProfileId = id
	success, err := ctrl.useCase.Update(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// Delete godoc
// @Summary Remove um perfil
// @Description Remove um perfil pelo ID
// @Tags profiles
// @Produce json
// @Param id path int true "ID do perfil"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /profiles/{id} [delete]
func (ctrl *ProfileController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	success, err := ctrl.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted"})
}
