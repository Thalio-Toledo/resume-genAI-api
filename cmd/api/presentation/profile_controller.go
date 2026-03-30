package presentation

import (
	"net/http"
	"strconv"

	"resume-genAI-api/cmd/api/application"
	"resume-genAI-api/cmd/api/domain"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	useCase *application.ProfileUseCase
}

func NewProfileController(uc *application.ProfileUseCase) *ProfileController {
	return &ProfileController{useCase: uc}
}

func (ctrl *ProfileController) RegisterRoutes(r *gin.Engine) {
	profiles := r.Group("/profiles/")
	{
		profiles.GET("", ctrl.Get)
		profiles.GET(":id", ctrl.FindByID)
		profiles.POST("", ctrl.Create)
		profiles.PUT("", ctrl.Update)
		profiles.DELETE(":id", ctrl.Delete)
	}
}

// Get godoc
//
//	@Summary		Lista todos os perfis
//	@Description	Retorna todos os perfis cadastrados
//	@Tags			profiles
//	@Produce		json
//	@Success		200	{array}	model.ProfileDTO
//	@Router			/profiles/ [get]
func (ctrl *ProfileController) Get(c *gin.Context) {
	profiles, err := ctrl.useCase.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profiles)
}

// FindByID godoc
//
//	@Summary		Busca perfil por ID
//	@Description	Retorna um perfil pelo ID
//	@Tags			profiles
//	@Produce		json
//	@Param			id	path		int	true	"ID do perfil"
//	@Success		200	{object}	model.ProfileDTO
//	@Failure		400	{object}	model.ErrorResponse
//	@Failure		404	{object}	model.ErrorResponse
//	@Router			/profiles/{id} [get]
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
//
//	@Summary		Cria um novo perfil
//	@Description	Cria um novo perfil com os dados informados
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			profile	body		model.ProfileDTO	true	"Perfil a ser criado"
//	@Success		201		{object}	model.ProfileDTO
//	@Failure		400		{object}	model.ErrorResponse
//	@Router			/profiles/ [post]
func (ctrl *ProfileController) Create(c *gin.Context) {
	var profile domain.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.useCase.Create(&profile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, profile)
}

// Update godoc
//
//	@Summary		Atualiza um perfil
//	@Description	Atualiza um perfil existente pelo ID
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"ID do perfil"
//	@Param			profile	body		model.ProfileDTO	true	"Perfil atualizado"
//	@Success		200		{object}	model.ProfileDTO
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		404		{object}	model.ErrorResponse
//	@Router			/profiles/{id} [put]
func (ctrl *ProfileController) Update(c *gin.Context) {
	var profile domain.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.Update(&profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (ctrl *ProfileController) AddCertification(c *gin.Context) {
	var certification domain.Certification
	if err := c.ShouldBindJSON(&certification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.AddCertification(certification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, certification)
}

func (ctrl *ProfileController) AddEducation(c *gin.Context) {
	var education domain.Education
	if err := c.ShouldBindJSON(&education); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.AddEducation(education)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, education)
}

func (ctrl *ProfileController) AddExperience(c *gin.Context) {
	var experience domain.Experience
	if err := c.ShouldBindJSON(&experience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.AddExperience(experience)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, experience)
}

func (ctrl *ProfileController) AddLanguage(c *gin.Context) {
	var language domain.Language
	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.AddLanguage(language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, language)
}

func (ctrl *ProfileController) AddProject(c *gin.Context) {
	var project domain.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.AddProject(project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (ctrl *ProfileController) AddSkill(c *gin.Context) {
	var skill domain.Skill
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.AddSkill(skill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skill)
}

func (ctrl *ProfileController) AddSocialMedia(c *gin.Context) {
	var socialMedia domain.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.AddSocialMedia(socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

// Delete godoc
//
//	@Summary		Remove um perfil
//	@Description	Remove um perfil pelo ID
//	@Tags			profiles
//	@Produce		json
//	@Param			id	path		int	true	"ID do perfil"
//	@Success		200	{object}	map[string]interface{}
//	@Failure		400	{object}	model.ErrorResponse
//	@Failure		404	{object}	model.ErrorResponse
//	@Router			/profiles/{id} [delete]
func (ctrl *ProfileController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = ctrl.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted"})
}
