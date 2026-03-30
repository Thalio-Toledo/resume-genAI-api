// O import abaixo referencia o pacote docs gerado pelo swag, localizado em cmd/api/docs
//
//	@title			Resume GenAI API
//	@version		1.0
//	@description	API para gerenciamento de perfis e currículos com IA.
//	@host			localhost:8080
//	@BasePath		/
package main

import (
	"context"
	"log"
	"resume-genAI-api/cmd/api/application"
	"resume-genAI-api/cmd/api/database"
	"resume-genAI-api/cmd/api/infrastructure"
	"resume-genAI-api/cmd/api/middleware"
	"resume-genAI-api/cmd/api/presentation"
	"time"

	// Importa o pacote docs gerado pelo swag para registrar a documentação Swagger

	_ "resume-genAI-api/cmd/api/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// skills, err := ai.Generate(`Requisitos da vaga Hard skills:
	// 				C#, .NET 6/7, ASP.NET Core e gRPC;
	// 				RabbitMQ e Kafka;
	// 				SQL Server, PostgreSQL e Redis;
	// 				Angular 15+, TypeScript e RxJS;
	// 				Material Design e Bootstrap;
	// 				Azure, Docker e Kubernetes;
	// 				Swagger/OpenAPI, Postman e Git;
	// 				GitHub Actions;
	// 				Experiência com IA e produtos financeiros.`)
	//ai.GenerateEmbedding("C#")

	//fmt.Println(strings.Join(skills, ", "))

	// Usar gin.New() para controle total
	r := gin.New()

	// Desabilitar redirects automáticos de trailing slash (que removem headers CORS)
	r.RedirectTrailingSlash = false

	// Aplicar middleware CORS PRIMEIRO
	r.Use(middleware.CORSMiddleware())

	// Adicionar middlewares padrão após CORS
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Banco indisponível:", err)
	}

	log.Println("🚀 API iniciada e banco conectado")

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "pong",
		})
	})

	// Handler global para OPTIONS em qualquer rota
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})

	// Dependências
	profileCommandRepo := infrastructure.NewProfileCommandRepository(db)
	profileQueryRepo := infrastructure.NewProfileQueryRepository(db)
	profileUC := application.NewProfileUseCase(profileCommandRepo, profileQueryRepo)
	profileCtrl := presentation.NewProfileController(profileUC)

	// Dependências ResumeController
	resumeUC := application.NewGenerateResumeUseCase(profileUC)
	resumeCtrl := presentation.NewResumeController(resumeUC)

	// Rotas
	profileCtrl.RegisterRoutes(r)
	resumeCtrl.RegisterRoutes(r)

	log.Println("SkillMatch AI running on :8080")
	r.Run(":8080")
}
