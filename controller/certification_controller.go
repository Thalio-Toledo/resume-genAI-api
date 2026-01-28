package controller

import (
	"net/http"
	"resume-genAI-api/model"
	"resume-genAI-api/useCase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CertificationController struct {
	useCase *useCase.CertificationUseCase
}

func NewCertificationController(uc *useCase.CertificationUseCase) *CertificationController {
	return &CertificationController{useCase: uc}
}

func (ctrl *CertificationController) RegisterRoutes(r *gin.Engine) {
	certs := r.Group("/certifications")
	{
		certs.GET("/", ctrl.GetAll)
		certs.GET(":id", ctrl.FindByID)
		certs.POST("/", ctrl.Create)
		certs.PUT(":id", ctrl.Update)
		certs.DELETE(":id", ctrl.Delete)
	}
}

// GetAll godoc
// @Summary Lista todas as certificações
// @Tags certifications
// @Produce json
// @Success 200 {array} model.Certification
// @Router /certifications/ [get]
func (ctrl *CertificationController) GetAll(c *gin.Context) {
	certs, err := ctrl.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, certs)
}

// FindByID godoc
// @Summary Busca certificação por ID
// @Tags certifications
// @Produce json
// @Param id path string true "ID da certificação"
// @Success 200 {object} model.Certification
// @Failure 404 {object} model.ErrorResponse
// @Router /certifications/{id} [get]
func (ctrl *CertificationController) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	cert, err := ctrl.useCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, cert)
}

// Create godoc
// @Summary Cria uma nova certificação
// @Tags certifications
// @Accept json
// @Produce json
// @Param certification body model.Certification true "Certificação a ser criada"
// @Success 201 {object} model.Certification
// @Failure 400 {object} model.ErrorResponse
// @Router /certifications/ [post]
func (ctrl *CertificationController) Create(c *gin.Context) {
	var cert model.Certification
	if err := c.ShouldBindJSON(&cert); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	_, err := ctrl.useCase.Create(cert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cert)
}

// Update godoc
// @Summary Atualiza uma certificação
// @Tags certifications
// @Accept json
// @Produce json
// @Param id path string true "ID da certificação"
// @Param certification body model.Certification true "Certificação atualizada"
// @Success 200 {object} model.Certification
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /certifications/{id} [put]
func (ctrl *CertificationController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var cert model.Certification
	if err := c.ShouldBindJSON(&cert); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	cert.Certification_Id = id
	success, err := ctrl.useCase.Update(cert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Certification not found"})
		return
	}
	c.JSON(http.StatusOK, cert)
}

// Delete godoc
// @Summary Remove uma certificação
// @Tags certifications
// @Produce json
// @Param id path string true "ID da certificação"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} model.ErrorResponse
// @Router /certifications/{id} [delete]
func (ctrl *CertificationController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	success, err := ctrl.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Certification not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Certification deleted"})
}
