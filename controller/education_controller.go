package controller

import (
	"net/http"
	"resume-genAI-api/model"
	"resume-genAI-api/useCase"

	"github.com/gin-gonic/gin"
)

type EducationController struct {
	useCase *useCase.EducationUseCase
}

func NewEducationController(uc *useCase.EducationUseCase) *EducationController {
	return &EducationController{useCase: uc}
}

func (ctrl *EducationController) RegisterRoutes(r *gin.Engine) {
	edu := r.Group("/educations")
	{
		edu.GET("/", ctrl.GetAll)
		edu.GET(":id", ctrl.FindByID)
		edu.POST("/", ctrl.Create)
		edu.PUT(":id", ctrl.Update)
		edu.DELETE(":id", ctrl.Delete)
	}
}

// GetAll godoc
// @Summary Lista todas as formações
// @Tags educations
// @Produce json
// @Success 200 {array} model.Education
// @Router /educations/ [get]
func (ctrl *EducationController) GetAll(c *gin.Context) {
	edu := ctrl.useCase.GetAll()
	c.JSON(http.StatusOK, edu)
}

// FindByID godoc
// @Summary Busca formação por ID
// @Tags educations
// @Produce json
// @Param id path string true "ID da formação"
// @Success 200 {object} model.Education
// @Failure 404 {object} model.ErrorResponse
// @Router /educations/{id} [get]
func (ctrl *EducationController) FindByID(c *gin.Context) {
	id := c.Param("id")
	edu, err := ctrl.useCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, edu)
}

// Create godoc
// @Summary Cria uma nova formação
// @Tags educations
// @Accept json
// @Produce json
// @Param education body model.Education true "Formação a ser criada"
// @Success 201 {object} model.Education
// @Failure 400 {object} model.ErrorResponse
// @Router /educations/ [post]
func (ctrl *EducationController) Create(c *gin.Context) {
	var edu model.Education
	if err := c.ShouldBindJSON(&edu); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	_, err := ctrl.useCase.Create(edu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, edu)
}

// Update godoc
// @Summary Atualiza uma formação
// @Tags educations
// @Accept json
// @Produce json
// @Param id path string true "ID da formação"
// @Param education body model.Education true "Formação atualizada"
// @Success 200 {object} model.Education
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /educations/{id} [put]
func (ctrl *EducationController) Update(c *gin.Context) {
	id := c.Param("id")
	var edu model.Education
	if err := c.ShouldBindJSON(&edu); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	edu.EducationId = id
	success, err := ctrl.useCase.Update(edu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Education not found"})
		return
	}
	c.JSON(http.StatusOK, edu)
}

// Delete godoc
// @Summary Remove uma formação
// @Tags educations
// @Produce json
// @Param id path string true "ID da formação"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} model.ErrorResponse
// @Router /educations/{id} [delete]
func (ctrl *EducationController) Delete(c *gin.Context) {
	id := c.Param("id")
	success, err := ctrl.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Education not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Education deleted"})
}
