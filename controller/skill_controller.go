package controller

import (
	"net/http"
	"resume-genAI-api/model"
	"resume-genAI-api/useCase"

	"github.com/gin-gonic/gin"
)

type SkillController struct {
	useCase *useCase.SkillUseCase
}

func NewSkillController(uc *useCase.SkillUseCase) *SkillController {
	return &SkillController{useCase: uc}
}

func (ctrl *SkillController) RegisterRoutes(r *gin.Engine) {
	skills := r.Group("/skills")
	{
		skills.GET("/", ctrl.GetAll)
		skills.GET(":id", ctrl.FindByID)
		skills.POST("/", ctrl.Create)
		skills.PUT(":id", ctrl.Update)
		skills.DELETE(":id", ctrl.Delete)
	}
}

// GetAll godoc
// @Summary Lista todas as skills
// @Tags skills
// @Produce json
// @Success 200 {array} model.Skill
// @Router /skills/ [get]
func (ctrl *SkillController) GetAll(c *gin.Context) {
	skills := ctrl.useCase.GetAll()
	c.JSON(http.StatusOK, skills)
}

// FindByID godoc
// @Summary Busca skill por ID
// @Tags skills
// @Produce json
// @Param id path string true "ID da skill"
// @Success 200 {object} model.Skill
// @Failure 404 {object} model.ErrorResponse
// @Router /skills/{id} [get]
func (ctrl *SkillController) FindByID(c *gin.Context) {
	id := c.Param("id")
	skill, err := ctrl.useCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, skill)
}

// Create godoc
// @Summary Cria uma nova skill
// @Tags skills
// @Accept json
// @Produce json
// @Param skill body model.Skill true "Skill a ser criada"
// @Success 201 {object} model.Skill
// @Failure 400 {object} model.ErrorResponse
// @Router /skills/ [post]
func (ctrl *SkillController) Create(c *gin.Context) {
	var skill model.Skill
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	_, err := ctrl.useCase.Create(skill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, skill)
}

// Update godoc
// @Summary Atualiza uma skill
// @Tags skills
// @Accept json
// @Produce json
// @Param id path string true "ID da skill"
// @Param skill body model.Skill true "Skill atualizada"
// @Success 200 {object} model.Skill
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /skills/{id} [put]
func (ctrl *SkillController) Update(c *gin.Context) {
	id := c.Param("id")
	var skill model.Skill
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	skill.ID = id
	success, err := ctrl.useCase.Update(skill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Skill not found"})
		return
	}
	c.JSON(http.StatusOK, skill)
}

// Delete godoc
// @Summary Remove uma skill
// @Tags skills
// @Produce json
// @Param id path string true "ID da skill"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} model.ErrorResponse
// @Router /skills/{id} [delete]
func (ctrl *SkillController) Delete(c *gin.Context) {
	id := c.Param("id")
	success, err := ctrl.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Skill not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Skill deleted"})
}
