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
		profiles.POST("add-certification", ctrl.AddCertification)
		profiles.PUT("update-certification", ctrl.UpdateCertification)
		profiles.POST("add-education", ctrl.AddEducation)
		profiles.PUT("update-education", ctrl.UpdateEducation)
		profiles.POST("add-experience", ctrl.AddExperience)
		profiles.PUT("update-experience", ctrl.UpdateExperience)
		profiles.POST("add-language", ctrl.AddLanguage)
		profiles.PUT("update-language", ctrl.UpdateLanguage)
		profiles.POST("add-project", ctrl.AddProject)
		profiles.PUT("update-project", ctrl.UpdateProject)
		profiles.POST("add-skill", ctrl.AddSkill)
		profiles.PUT("update-skill", ctrl.UpdateSkill)
		profiles.POST("add-social-media", ctrl.AddSocialMedia)
		profiles.PUT("update-social-media", ctrl.UpdateSocialMedia)
	}
}

// Get godoc
//
//	@Summary		Lista todos os perfis
//	@Description	Retorna todos os perfis cadastrados
//	@Tags			profiles
//	@Produce		json
//	@Success		200	{array}	domain.ProfileDTO
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
//	@Success		200	{object}	domain.ProfileDTO
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
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
//	@Param			profile	body		domain.ProfileDTO	true	"Perfil a ser criado"
//	@Success		201		{object}	domain.ProfileDTO
//	@Failure		400		{object}	domain.ErrorResponse
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
//	@Description	Atualiza um perfil existente
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			profile	body		domain.ProfileDTO	true	"Perfil atualizado"
//	@Success		200		{object}	domain.ProfileDTO
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		404		{object}	domain.ErrorResponse
//	@Router			/profiles/ [put]
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

// AddCertification godoc
//
//	@Summary		Adiciona uma certificação ao perfil
//	@Description	Adiciona uma nova certificação a um perfil existente
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			certification	body		domain.Certification	true	"Certificação a ser adicionada"
//	@Success		200		{object}	domain.Certification
//	@Failure		400		{object}	domain.ErrorResponse
//	@Router			/profiles/add-certification [post]
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

// AddEducation godoc
//
//	@Summary		Adiciona educação ao perfil
//	@Description	Adiciona uma nova formação educacional a um perfil existente
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			education	body		domain.Education	true	"Educação a ser adicionada"
//	@Success		200		{object}	domain.Education
//	@Failure		400		{object}	domain.ErrorResponse
//	@Router			/profiles/add-education [post]
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

// AddExperience godoc
//
//	@Summary		Adiciona experiência ao perfil
//	@Description	Adiciona uma nova experiência profissional a um perfil existente
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			experience	body		domain.Experience	true	"Experiência a ser adicionada"
//	@Success		200		{object}	domain.Experience
//	@Failure		400		{object}	domain.ErrorResponse
//	@Router			/profiles/add-experience [post]
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

// AddLanguage godoc
//
//	@Summary		Adiciona idioma ao perfil
//	@Description	Adiciona um novo idioma a um perfil existente
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			language	body		domain.Language	true	"Idioma a ser adicionado"
//	@Success		200		{object}	domain.Language
//	@Failure		400		{object}	domain.ErrorResponse
//	@Router			/profiles/add-language [post]
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

// AddProject godoc
//
//	@Summary		Adiciona projeto ao perfil
//	@Description	Adiciona um novo projeto a um perfil existente
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			project	body		domain.Project	true	"Projeto a ser adicionado"
//	@Success		200		{object}	domain.Project
//	@Failure		400		{object}	domain.ErrorResponse
//	@Router			/profiles/add-project [post]
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

// AddSkill godoc
//
//	@Summary		Adiciona habilidade ao perfil
//	@Description	Adiciona uma nova habilidade a um perfil existente
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			skill	body		domain.Skill	true	"Habilidade a ser adicionada"
//	@Success		200		{object}	domain.Skill
//	@Failure		400		{object}	domain.ErrorResponse
//	@Router			/profiles/add-skill [post]
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

// AddSocialMedia godoc
//
//	@Summary		Adiciona rede social ao perfil
//	@Description	Adiciona uma nova rede social a um perfil existente
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			socialMedia	body		domain.SocialMedia	true	"Rede social a ser adicionada"
//	@Success		200		{object}	domain.SocialMedia
//	@Failure		400		{object}	domain.ErrorResponse
//	@Router			/profiles/add-social-media [post]
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
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
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

// UpdateCertification godoc
//
//	@Summary		Atualiza uma certificação do perfil
//	@Description	Atualiza uma certificação existente em um perfil
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			certification	body		domain.Certification	true	"Certificação a ser atualizada"
//	@Success		200		{object}	domain.Certification
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		404		{object}	domain.ErrorResponse
//	@Router			/profiles/update-certification [put]
func (ctrl *ProfileController) UpdateCertification(c *gin.Context) {
	var certification domain.Certification
	if err := c.ShouldBindJSON(&certification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.UpdateCertification(certification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, certification)
}

// UpdateEducation godoc
//
//	@Summary		Atualiza educação do perfil
//	@Description	Atualiza uma formação educacional existente em um perfil
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			education	body		domain.Education	true	"Educação a ser atualizada"
//	@Success		200		{object}	domain.Education
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		404		{object}	domain.ErrorResponse
//	@Router			/profiles/update-education [put]
func (ctrl *ProfileController) UpdateEducation(c *gin.Context) {
	var education domain.Education
	if err := c.ShouldBindJSON(&education); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.UpdateEducation(education)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, education)
}

// UpdateExperience godoc
//
//	@Summary		Atualiza experiência do perfil
//	@Description	Atualiza uma experiência profissional existente em um perfil
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			experience	body		domain.Experience	true	"Experiência a ser atualizada"
//	@Success		200		{object}	domain.Experience
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		404		{object}	domain.ErrorResponse
//	@Router			/profiles/update-experience [put]
func (ctrl *ProfileController) UpdateExperience(c *gin.Context) {
	var experience domain.Experience
	if err := c.ShouldBindJSON(&experience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.UpdateExperience(experience)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, experience)
}

// UpdateLanguage godoc
//
//	@Summary		Atualiza idioma do perfil
//	@Description	Atualiza um idioma existente em um perfil
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			language	body		domain.Language	true	"Idioma a ser atualizado"
//	@Success		200		{object}	domain.Language
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		404		{object}	domain.ErrorResponse
//	@Router			/profiles/update-language [put]
func (ctrl *ProfileController) UpdateLanguage(c *gin.Context) {
	var language domain.Language
	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.UpdateLanguage(language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, language)
}

// UpdateProject godoc
//
//	@Summary		Atualiza projeto do perfil
//	@Description	Atualiza um projeto existente em um perfil
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			project	body		domain.Project	true	"Projeto a ser atualizado"
//	@Success		200		{object}	domain.Project
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		404		{object}	domain.ErrorResponse
//	@Router			/profiles/update-project [put]
func (ctrl *ProfileController) UpdateProject(c *gin.Context) {
	var project domain.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.UpdateProject(project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

// UpdateSkill godoc
//
//	@Summary		Atualiza habilidade do perfil
//	@Description	Atualiza uma habilidade existente em um perfil
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			skill	body		domain.Skill	true	"Habilidade a ser atualizada"
//	@Success		200		{object}	domain.Skill
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		404		{object}	domain.ErrorResponse
//	@Router			/profiles/update-skill [put]
func (ctrl *ProfileController) UpdateSkill(c *gin.Context) {
	var skill domain.Skill
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.UpdateSkill(skill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skill)
}

// UpdateSocialMedia godoc
//
//	@Summary		Atualiza rede social do perfil
//	@Description	Atualiza uma rede social existente em um perfil
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			socialMedia	body		domain.SocialMedia	true	"Rede social a ser atualizada"
//	@Success		200		{object}	domain.SocialMedia
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		404		{object}	domain.ErrorResponse
//	@Router			/profiles/update-social-media [put]
func (ctrl *ProfileController) UpdateSocialMedia(c *gin.Context) {
	var socialMedia domain.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.useCase.UpdateSocialMedia(socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}
