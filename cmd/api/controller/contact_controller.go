package controller

import (
	"net/http"
	"resume-genAI-api/cmd/api/model"
	"resume-genAI-api/cmd/api/useCase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
	useCase *useCase.ContactUseCase
}

func NewContactController(uc *useCase.ContactUseCase) *ContactController {
	return &ContactController{useCase: uc}
}

func (ctrl *ContactController) RegisterRoutes(r *gin.Engine) {
	contacts := r.Group("/contacts/")
	{
		contacts.GET("", ctrl.GetAll)
		contacts.GET(":id", ctrl.FindByID)
		contacts.POST("", ctrl.Create)
		contacts.PUT(":id", ctrl.Update)
		contacts.DELETE(":id", ctrl.Delete)
	}
}

// GetAll godoc
// @Summary Lista todos os contatos
// @Tags contacts
// @Produce json
// @Success 200 {array} model.Contact
// @Router /contacts/ [get]
func (ctrl *ContactController) GetAll(c *gin.Context) {
	contacts, err := ctrl.useCase.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contacts)
}

// FindByID godoc
// @Summary Busca contato por ID
// @Tags contacts
// @Produce json
// @Param id path int true "ID do contato"
// @Success 200 {object} model.Contact
// @Failure 404 {object} model.ErrorResponse
// @Router /contacts/{id} [get]
func (ctrl *ContactController) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	contact, err := ctrl.useCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, contact)
}

// Create godoc
// @Summary Cria um novo contato
// @Tags contacts
// @Accept json
// @Produce json
// @Param contact body model.Contact true "Contato a ser criado"
// @Success 201 {object} model.Contact
// @Failure 400 {object} model.ErrorResponse
// @Router /contacts/ [post]
func (ctrl *ContactController) Create(c *gin.Context) {
	var contact model.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	_, err := ctrl.useCase.Create(contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, contact)
}

// Update godoc
// @Summary Atualiza um contato
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path int true "ID do contato"
// @Param contact body model.Contact true "Contato atualizado"
// @Success 200 {object} model.Contact
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /contacts/{id} [put]
func (ctrl *ContactController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var contact model.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	contact.ContactId = id
	success, err := ctrl.useCase.Update(contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Contact not found"})
		return
	}
	c.JSON(http.StatusOK, contact)
}

// Delete godoc
// @Summary Remove um contato
// @Tags contacts
// @Produce json
// @Param id path int true "ID do contato"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} model.ErrorResponse
// @Router /contacts/{id} [delete]
func (ctrl *ContactController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	success, err := ctrl.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Contact not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted"})
}
