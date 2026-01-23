package controller

import (
	"net/http"
	"resume-genAI-api/model"
	"resume-genAI-api/useCase"

	"github.com/gin-gonic/gin"
)

type LanguageController struct {
	useCase *useCase.LanguageUseCase
}

func NewLanguageController(uc *useCase.LanguageUseCase) *LanguageController {
	return &LanguageController{useCase: uc}
}

func (ctrl *LanguageController) RegisterRoutes(r *gin.Engine) {
	langs := r.Group("/languages")
	{
		langs.GET("/", ctrl.GetAll)
		langs.GET(":id", ctrl.FindByID)
		langs.POST("/", ctrl.Create)
		langs.PUT(":id", ctrl.Update)
		langs.DELETE(":id", ctrl.Delete)
	}
}

// GetAll godoc
// @Summary Lista todos os idiomas
// @Tags languages
// @Produce json
// @Success 200 {array} model.Language
// @Router /languages/ [get]
func (ctrl *LanguageController) GetAll(c *gin.Context) {
	langs := ctrl.useCase.GetAll()
	c.JSON(http.StatusOK, langs)
}

// FindByID godoc
// @Summary Busca idioma por ID
// @Tags languages
// @Produce json
// @Param id path string true "ID do idioma"
// @Success 200 {object} model.Language
// @Failure 404 {object} model.ErrorResponse
// @Router /languages/{id} [get]
func (ctrl *LanguageController) FindByID(c *gin.Context) {
	id := c.Param("id")
	lang, err := ctrl.useCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, lang)
}

// Create godoc
// @Summary Cria um novo idioma
// @Tags languages
// @Accept json
// @Produce json
// @Param language body model.Language true "Idioma a ser criado"
// @Success 201 {object} model.Language
// @Failure 400 {object} model.ErrorResponse
// @Router /languages/ [post]
func (ctrl *LanguageController) Create(c *gin.Context) {
	var lang model.Language
	if err := c.ShouldBindJSON(&lang); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	_, err := ctrl.useCase.Create(lang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, lang)
}

// Update godoc
// @Summary Atualiza um idioma
// @Tags languages
// @Accept json
// @Produce json
// @Param id path string true "ID do idioma"
// @Param language body model.Language true "Idioma atualizado"
// @Success 200 {object} model.Language
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /languages/{id} [put]
func (ctrl *LanguageController) Update(c *gin.Context) {
	id := c.Param("id")
	var lang model.Language
	if err := c.ShouldBindJSON(&lang); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	lang.ID = id
	success, err := ctrl.useCase.Update(lang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Language not found"})
		return
	}
	c.JSON(http.StatusOK, lang)
}

// Delete godoc
// @Summary Remove um idioma
// @Tags languages
// @Produce json
// @Param id path string true "ID do idioma"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} model.ErrorResponse
// @Router /languages/{id} [delete]
func (ctrl *LanguageController) Delete(c *gin.Context) {
	id := c.Param("id")
	success, err := ctrl.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Language not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Language deleted"})
}
