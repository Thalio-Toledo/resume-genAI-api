package presentation

import (
	"net/http"
	"resume-genAI-api/cmd/api/application"
	"resume-genAI-api/cmd/api/dto"

	"github.com/gin-gonic/gin"
)

type ResumeController struct {
	resumeUseCase *application.GenerateResumeUseCase
}

func NewResumeController(uc *application.GenerateResumeUseCase) *ResumeController {
	return &ResumeController{resumeUseCase: uc}
}

func (ctrl *ResumeController) RegisterRoutes(r *gin.Engine) {
	profiles := r.Group("/resume/")
	{
		profiles.POST("/generate", ctrl.Generate)
	}
}

// Generate godoc
//
//	@Summary		Gera um novo currículo baseado na descrição de vaga
//	@Description	Gera um novo currículo usando LLMs com base na descrição da vaga
//	@Tags			resume
//	@Accept			json
//	@Produce		json
//	@Param			roleDescription	body		dto.RoleDescription	true	"Descrição da vaga"
//	@Success		201				{object}	domain.ProfileDTO
//	@Failure		400				{object}	domain.ErrorResponse
//	@Router			/resume/generate [post]
func (ctrl *ResumeController) Generate(c *gin.Context) {

	var job_description dto.RoleDescription
	if err := c.ShouldBindJSON(&job_description); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resume, err := ctrl.resumeUseCase.Generate(job_description)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resume)
}
