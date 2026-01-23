package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"resume-genAI-api/model"
	"resume-genAI-api/useCase"
)

type SocialMediaController struct {
	useCase *useCase.SocialMediaUseCase
}

func NewSocialMediaController(uc *useCase.SocialMediaUseCase) *SocialMediaController {
	return &SocialMediaController{useCase: uc}
}

func (ctrl *SocialMediaController) RegisterRoutes(r *gin.Engine) {
	sm := r.Group("/socialmedias")
	{
		sm.GET("/", ctrl.GetAll)
		sm.GET(":handle", ctrl.FindByHandle)
		sm.POST("/", ctrl.Create)
		sm.PUT(":handle", ctrl.Update)
		sm.DELETE(":handle", ctrl.Delete)
	}
}

// GetAll godoc
// @Summary Lista todas as redes sociais
// @Tags socialmedias
// @Produce json
// @Success 200 {array} model.SocialMedia
// @Router /socialmedias/ [get]
func (ctrl *SocialMediaController) GetAll(c *gin.Context) {
	sm := ctrl.useCase.GetAll()
	c.JSON(http.StatusOK, sm)
}

// FindByHandle godoc
// @Summary Busca rede social por handle
// @Tags socialmedias
// @Produce json
// @Param handle path string true "Handle da rede social"
// @Success 200 {object} model.SocialMedia
// @Failure 404 {object} model.ErrorResponse
// @Router /socialmedias/{handle} [get]
func (ctrl *SocialMediaController) FindByHandle(c *gin.Context) {
	handle := c.Param("handle")
	sm, err := ctrl.useCase.FindByHandle(handle)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, sm)
}

// Create godoc
// @Summary Cria uma nova rede social
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param socialmedia body model.SocialMedia true "Rede social a ser criada"
// @Success 201 {object} model.SocialMedia
// @Failure 400 {object} model.ErrorResponse
// @Router /socialmedias/ [post]
func (ctrl *SocialMediaController) Create(c *gin.Context) {
	var sm model.SocialMedia
	if err := c.ShouldBindJSON(&sm); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	_, err := ctrl.useCase.Create(sm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sm)
}

// Update godoc
// @Summary Atualiza uma rede social
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param handle path string true "Handle da rede social"
// @Param socialmedia body model.SocialMedia true "Rede social atualizada"
// @Success 200 {object} model.SocialMedia
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /socialmedias/{handle} [put]
func (ctrl *SocialMediaController) Update(c *gin.Context) {
	handle := c.Param("handle")
	var sm model.SocialMedia
	if err := c.ShouldBindJSON(&sm); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	sm.Handle = handle
	success, err := ctrl.useCase.Update(sm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "SocialMedia not found"})
		return
	}
	c.JSON(http.StatusOK, sm)
}

// Delete godoc
// @Summary Remove uma rede social
// @Tags socialmedias
// @Produce json
// @Param handle path string true "Handle da rede social"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} model.ErrorResponse
// @Router /socialmedias/{handle} [delete]
func (ctrl *SocialMediaController) Delete(c *gin.Context) {
	handle := c.Param("handle")
	success, err := ctrl.useCase.Delete(handle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "SocialMedia not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "SocialMedia deleted"})
}
