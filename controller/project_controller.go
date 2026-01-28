package controller

import (
	"net/http"
	"resume-genAI-api/model"
	"resume-genAI-api/useCase"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	useCase *useCase.ProjectUseCase
}

func NewProjectController(uc *useCase.ProjectUseCase) *ProjectController {
	return &ProjectController{useCase: uc}
}

func (ctrl *ProjectController) RegisterRoutes(r *gin.Engine) {
	projs := r.Group("/projects")
	{
		projs.GET("/", ctrl.GetAll)
		projs.GET(":id", ctrl.FindByID)
		projs.POST("/", ctrl.Create)
		projs.PUT(":id", ctrl.Update)
		projs.DELETE(":id", ctrl.Delete)
	}
}

// GetAll godoc
// @Summary Lista todos os projetos
// @Tags projects
// @Produce json
// @Success 200 {array} model.Project
// @Router /projects/ [get]
func (ctrl *ProjectController) GetAll(c *gin.Context) {
	projs, err := ctrl.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projs)
}

// FindByID godoc
// @Summary Busca projeto por ID
// @Tags projects
// @Produce json
// @Param id path string true "ID do projeto"
// @Success 200 {object} model.Project
// @Failure 404 {object} model.ErrorResponse
// @Router /projects/{id} [get]
func (ctrl *ProjectController) FindByID(c *gin.Context) {
	id := c.Param("id")
	proj, err := ctrl.useCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, proj)
}

// Create godoc
// @Summary Cria um novo projeto
// @Tags projects
// @Accept json
// @Produce json
// @Param project body model.Project true "Projeto a ser criado"
// @Success 201 {object} model.Project
// @Failure 400 {object} model.ErrorResponse
// @Router /projects/ [post]
func (ctrl *ProjectController) Create(c *gin.Context) {
	var proj model.Project
	if err := c.ShouldBindJSON(&proj); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	_, err := ctrl.useCase.Create(proj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, proj)
}

// Update godoc
// @Summary Atualiza um projeto
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "ID do projeto"
// @Param project body model.Project true "Projeto atualizado"
// @Success 200 {object} model.Project
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /projects/{id} [put]
func (ctrl *ProjectController) Update(c *gin.Context) {
	id := c.Param("id")
	var proj model.Project
	if err := c.ShouldBindJSON(&proj); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	proj.ProjectId = id
	success, err := ctrl.useCase.Update(proj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Project not found"})
		return
	}
	c.JSON(http.StatusOK, proj)
}

// Delete godoc
// @Summary Remove um projeto
// @Tags projects
// @Produce json
// @Param id path string true "ID do projeto"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} model.ErrorResponse
// @Router /projects/{id} [delete]
func (ctrl *ProjectController) Delete(c *gin.Context) {
	id := c.Param("id")
	success, err := ctrl.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Project not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
