package controller

import (
	"net/http"
	"resume-genAI-api/model"
	"resume-genAI-api/useCase"

	"github.com/gin-gonic/gin"
)

type ExperienceController struct {
	useCase *useCase.ExperienceUseCase
}

func NewExperienceController(uc *useCase.ExperienceUseCase) *ExperienceController {
	return &ExperienceController{useCase: uc}
}

func (ctrl *ExperienceController) RegisterRoutes(r *gin.Engine) {
	exp := r.Group("/experiences")
	{
		exp.GET("/", ctrl.GetAll)
		exp.GET(":id", ctrl.FindByID)
		exp.POST("/", ctrl.Create)
		exp.PUT(":id", ctrl.Update)
		exp.DELETE(":id", ctrl.Delete)
	}
}

// GetAll godoc
// @Summary Lista todas as experiências
// @Tags experiences
// @Produce json
// @Success 200 {array} model.Experience
// @Router /experiences/ [get]
func (ctrl *ExperienceController) GetAll(c *gin.Context) {
	exp := ctrl.useCase.GetAll()
	c.JSON(http.StatusOK, exp)
}

// FindByID godoc
// @Summary Busca experiência por ID
// @Tags experiences
// @Produce json
// @Param id path string true "ID da experiência"
// @Success 200 {object} model.Experience
// @Failure 404 {object} model.ErrorResponse
// @Router /experiences/{id} [get]
func (ctrl *ExperienceController) FindByID(c *gin.Context) {
	id := c.Param("id")
	exp, err := ctrl.useCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, exp)
}

// Create godoc
// @Summary Cria uma nova experiência
// @Tags experiences
// @Accept json
// @Produce json
// @Param experience body model.Experience true "Experiência a ser criada"
// @Success 201 {object} model.Experience
// @Failure 400 {object} model.ErrorResponse
// @Router /experiences/ [post]
func (ctrl *ExperienceController) Create(c *gin.Context) {
	var exp model.Experience
	if err := c.ShouldBindJSON(&exp); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	_, err := ctrl.useCase.Create(exp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, exp)
}

// Update godoc
// @Summary Atualiza uma experiência
// @Tags experiences
// @Accept json
// @Produce json
// @Param id path string true "ID da experiência"
// @Param experience body model.Experience true "Experiência atualizada"
// @Success 200 {object} model.Experience
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /experiences/{id} [put]
func (ctrl *ExperienceController) Update(c *gin.Context) {
	id := c.Param("id")
	var exp model.Experience
	if err := c.ShouldBindJSON(&exp); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	exp.ID = id
	success, err := ctrl.useCase.Update(exp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Experience not found"})
		return
	}
	c.JSON(http.StatusOK, exp)
}

// Delete godoc
// @Summary Remove uma experiência
// @Tags experiences
// @Produce json
// @Param id path string true "ID da experiência"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} model.ErrorResponse
// @Router /experiences/{id} [delete]
func (ctrl *ExperienceController) Delete(c *gin.Context) {
	id := c.Param("id")
	success, err := ctrl.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Experience not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Experience deleted"})
}
