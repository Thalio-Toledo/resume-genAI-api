// O import abaixo referencia o pacote docs gerado pelo swag, localizado em cmd/api/docs
// @title Resume GenAI API
// @version 1.0
// @description API para gerenciamento de perfis e currículos com IA.
// @host localhost:8080
// @BasePath /
package main

import (
	"log"
	"resume-genAI-api/controller"
	"resume-genAI-api/repository"
	"resume-genAI-api/useCase"

	// Importa o pacote docs gerado pelo swag para registrar a documentação Swagger
	_ "resume-genAI-api/cmd/api/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "pong",
		})
	})

	// Dependências
	profileRepo := repository.NewProfileRepository()
	profileUC := useCase.NewProfileUseCase(profileRepo)
	profileCtrl := controller.NewProfileController(profileUC)

	certificationRepo := repository.NewCertificationRepository()
	certificationUC := useCase.NewCertificationUseCase(certificationRepo)
	certificationCtrl := controller.NewCertificationController(certificationUC)

	contactRepo := repository.NewContactRepository()
	contactUC := useCase.NewContactUseCase(contactRepo)
	contactCtrl := controller.NewContactController(contactUC)

	educationRepo := repository.NewEducationRepository()
	educationUC := useCase.NewEducationUseCase(educationRepo)
	educationCtrl := controller.NewEducationController(educationUC)

	experienceRepo := repository.NewExperienceRepository()
	experienceUC := useCase.NewExperienceUseCase(experienceRepo)
	experienceCtrl := controller.NewExperienceController(experienceUC)

	languageRepo := repository.NewLanguageRepository()
	languageUC := useCase.NewLanguageUseCase(languageRepo)
	languageCtrl := controller.NewLanguageController(languageUC)

	projectRepo := repository.NewProjectRepository()
	projectUC := useCase.NewProjectUseCase(projectRepo)
	projectCtrl := controller.NewProjectController(projectUC)

	resumeRepo := repository.NewResumeRepository()
	resumeUC := useCase.NewResumeUseCase(resumeRepo)
	resumeCtrl := controller.NewResumeController(resumeUC)

	skillRepo := repository.NewSkillRepository()
	skillUC := useCase.NewSkillUseCase(skillRepo)
	skillCtrl := controller.NewSkillController(skillUC)

	socialMediaRepo := repository.NewSocialMediaRepository()
	socialMediaUC := useCase.NewSocialMediaUseCase(socialMediaRepo)
	socialMediaCtrl := controller.NewSocialMediaController(socialMediaUC)

	// Rotas
	profileCtrl.RegisterRoutes(r)
	certificationCtrl.RegisterRoutes(r)
	contactCtrl.RegisterRoutes(r)
	educationCtrl.RegisterRoutes(r)
	experienceCtrl.RegisterRoutes(r)
	languageCtrl.RegisterRoutes(r)
	projectCtrl.RegisterRoutes(r)
	resumeCtrl.RegisterRoutes(r)
	skillCtrl.RegisterRoutes(r)
	socialMediaCtrl.RegisterRoutes(r)

	log.Println("SkillMatch AI running on :8080")
	r.Run(":8080")
}
