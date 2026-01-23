package controller

import (
	"net/http"
	"resume-genAI-api/model"
	"resume-genAI-api/useCase"

	"github.com/gin-gonic/gin"
)

type ResumeController struct {
	useCase *useCase.ResumeUseCase
}

func NewResumeController(uc *useCase.ResumeUseCase) *ResumeController {
	return &ResumeController{useCase: uc}
}

func (ctrl *ResumeController) RegisterRoutes(r *gin.Engine) {
	resumes := r.Group("/resumes")
	{
		resumes.GET("/", ctrl.GetAll)
		resumes.GET(":id", ctrl.FindByID)
		resumes.POST("/", ctrl.Create)
		resumes.PUT(":id", ctrl.Update)
		resumes.DELETE(":id", ctrl.Delete)
	}
}

// GetAll godoc
// @Summary Lista todos os currículos
// @Tags resumes
// @Produce json
// @Success 200 {array} model.Resume
// @Router /resumes/ [get]
func (ctrl *ResumeController) GetAll(c *gin.Context) {
	resumes := ctrl.useCase.GetAll()
	c.JSON(http.StatusOK, resumes)
}

// FindByID godoc
// @Summary Busca currículo por ID
// @Tags resumes
// @Produce json
// @Param id path string true "ID do currículo"
// @Success 200 {object} model.Resume
// @Failure 404 {object} model.ErrorResponse
// @Router /resumes/{id} [get]
func (ctrl *ResumeController) FindByID(c *gin.Context) {
	id := c.Param("id")
	resume, err := ctrl.useCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resume)
}

// Create godoc
// @Summary Cria um novo currículo
// @Tags resumes
// @Accept json
// @Produce json
// @Param resume body model.Resume true "Currículo a ser criado"
// @Success 201 {object} model.Resume
// @Failure 400 {object} model.ErrorResponse
// @Router /resumes/ [post]
func (ctrl *ResumeController) Create(c *gin.Context) {
	var resume model.Resume
	if err := c.ShouldBindJSON(&resume); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	_, err := ctrl.useCase.Create(resume)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resume)
}

// Update godoc
// @Summary Atualiza um currículo
// @Tags resumes
// @Accept json
// @Produce json
// @Param id path string true "ID do currículo"
// @Param resume body model.Resume true "Currículo atualizado"
// @Success 200 {object} model.Resume
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /resumes/{id} [put]
func (ctrl *ResumeController) Update(c *gin.Context) {
	id := c.Param("id")
	var resume model.Resume
	if err := c.ShouldBindJSON(&resume); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	resume.ID = id
	success, err := ctrl.useCase.Update(resume)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Resume not found"})
		return
	}
	c.JSON(http.StatusOK, resume)
}

// Delete godoc
// @Summary Remove um currículo
// @Tags resumes
// @Produce json
// @Param id path string true "ID do currículo"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} model.ErrorResponse
// @Router /resumes/{id} [delete]
func (ctrl *ResumeController) Delete(c *gin.Context) {
	id := c.Param("id")
	success, err := ctrl.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Resume not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Resume deleted"})
}
