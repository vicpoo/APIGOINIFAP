// CreateRecomendacionNutricionalController.go
package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/application"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type CreateRecomendacionNutricionalController struct {
	createUseCase *application.CreateRecomendacionNutricionalUseCase
	fileUploader  *FileUploader
}

func NewCreateRecomendacionNutricionalController(createUseCase *application.CreateRecomendacionNutricionalUseCase) *CreateRecomendacionNutricionalController {
	return &CreateRecomendacionNutricionalController{
		createUseCase: createUseCase,
		fileUploader:  NewFileUploader(),
	}
}

func (ctrl *CreateRecomendacionNutricionalController) Run(c *gin.Context) {
	// Parsear form-data
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil { // 32 MB max
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error al procesar formulario",
			"error":   err.Error(),
		})
		return
	}

	// Obtener campos del form
	municipioIDStr := c.Request.FormValue("municipio_id_FK")
	userIDStr := c.Request.FormValue("user_id_FK")

	// Convertir IDs
	var municipioID, userID int32
	if _, err := fmt.Sscanf(municipioIDStr, "%d", &municipioID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "municipio_id_FK debe ser un número",
			"error":   err.Error(),
		})
		return
	}

	if _, err := fmt.Sscanf(userIDStr, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user_id_FK debe ser un número",
			"error":   err.Error(),
		})
		return
	}

	// Subir archivo PDF
	nombrePDF, rutaPDF, err := ctrl.fileUploader.UploadPDF(c, "archivo_pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error al subir archivo PDF",
			"error":   err.Error(),
		})
		return
	}

	recomendacion := entities.NewRecomendacionNutricional(
		municipioID,
		nombrePDF,
		rutaPDF,
		userID,
	)

	createdRecomendacion, err := ctrl.createUseCase.Run(recomendacion)
	if err != nil {
		// Si falla, eliminar el archivo subido
		ctrl.fileUploader.DeleteFile("pdfs", nombrePDF)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la recomendación nutricional",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdRecomendacion)
}