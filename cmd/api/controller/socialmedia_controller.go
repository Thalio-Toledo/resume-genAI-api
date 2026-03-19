package controller

import (
	"net/http"
	"resume-genAI-api/cmd/api/model"
	"resume-genAI-api/cmd/api/useCase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	useCase *useCase.SocialMediaUseCase
}

func NewSocialMediaController(uc *useCase.SocialMediaUseCase) *SocialMediaController {
	return &SocialMediaController{useCase: uc}
}

func (ctrl *SocialMediaController) RegisterRoutes(r *gin.Engine) {
	sm := r.Group("/socialmedias/")
	{
		sm.GET("", ctrl.GetAll)
		sm.GET("/:id", ctrl.FindByID)
		sm.POST("", ctrl.Create)
		sm.PUT("/:id", ctrl.Update)
		sm.DELETE("/:id", ctrl.Delete)
	}
}

// GetAll godoc
//
//	@Summary	Lista todas as redes sociais
//	@Tags		socialmedias
//	@Produce	json
//	@Success	200	{array}	model.SocialMedia
//	@Router		/socialmedia/ [get]
func (ctrl *SocialMediaController) GetAll(c *gin.Context) {
	sm, err := ctrl.useCase.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sm)
}

// FindByID godoc
//
//	@Summary	Busca rede social por ID
//	@Tags		socialmedias
//	@Produce	json
//	@Param		id	path		int	true	"ID da rede social"
//	@Success	200	{object}	model.SocialMedia
//	@Failure	404	{object}	model.ErrorResponse
//	@Router		/socialmedias/{id} [get]
func (ctrl *SocialMediaController) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sm, err := ctrl.useCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, sm)
}

// Create godoc
//
//	@Summary	Cria uma nova rede social
//	@Tags		socialmedias
//	@Accept		json
//	@Produce	json
//	@Param		socialmedia	body		model.SocialMedia	true	"Rede social a ser criada"
//	@Success	201			{object}	model.SocialMedia
//	@Failure	400			{object}	model.ErrorResponse
//	@Router		/socialmedia/ [post]
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
//
//	@Summary	Atualiza uma rede social
//	@Tags		socialmedias
//	@Accept		json
//	@Produce	json
//	@Param		id			path		int					true	"ID da rede social"
//	@Param		socialmedia	body		model.SocialMedia	true	"Rede social atualizada"
//	@Success	200			{object}	model.SocialMedia
//	@Failure	400			{object}	model.ErrorResponse
//	@Failure	404			{object}	model.ErrorResponse
//	@Router		/socialmedias/{id} [put]
func (ctrl *SocialMediaController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var sm model.SocialMedia
	if err := c.ShouldBindJSON(&sm); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	sm.SocialMediaId = id
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
//
//	@Summary	Remove uma rede social
//	@Tags		socialmedias
//	@Produce	json
//	@Param		id	path		int	true	"ID da rede social"
//	@Success	200	{object}	map[string]interface{}
//	@Failure	404	{object}	model.ErrorResponse
//	@Router		/socialmedias/{id} [delete]
func (ctrl *SocialMediaController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	success, err := ctrl.useCase.Delete(id)
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
