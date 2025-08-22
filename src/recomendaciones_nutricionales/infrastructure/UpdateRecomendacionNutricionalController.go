// UpdateRecomendacionNutricionalController.go
package infrastructure

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/application"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type UpdateRecomendacionNutricionalController struct {
	updateUseCase *application.UpdateRecomendacionNutricionalUseCase
	fileUploader  *FileUploader
}

func NewUpdateRecomendacionNutricionalController(updateUseCase *application.UpdateRecomendacionNutricionalUseCase) *UpdateRecomendacionNutricionalController {
	return &UpdateRecomendacionNutricionalController{
		updateUseCase: updateUseCase,
		fileUploader:  NewFileUploader(),
	}
}

func (ctrl *UpdateRecomendacionNutricionalController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	// Primero obtener la recomendación existente para manejar el archivo viejo
	existingRecomendacion, err := ctrl.updateUseCase.GetById(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Recomendación no encontrada",
			"error":   err.Error(),
		})
		return
	}

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

	var municipioID, userID int32
	if _, err := fmt.Sscanf(municipioIDStr, "%d", &municipioID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "municipio_id_FK debe ser un número válido",
			"error":   err.Error(),
		})
		return
	}

	if _, err := fmt.Sscanf(userIDStr, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user_id_FK debe ser un número válido",
			"error":   err.Error(),
		})
		return
	}

	var nombrePDF, rutaPDF string
	var oldPDF string

	// Verificar si se subió nuevo archivo
	if c.Request.MultipartForm != nil && c.Request.MultipartForm.File["archivo_pdf"] != nil {
		// Subir nuevo archivo
		nombrePDF, rutaPDF, err = ctrl.fileUploader.UploadPDF(c, "archivo_pdf")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error al subir archivo PDF",
				"error":   err.Error(),
			})
			return
		}
		oldPDF = existingRecomendacion.NombrePDF // Guardar nombre del archivo viejo para eliminarlo después
	} else {
		// Usar archivo existente
		nombrePDF = existingRecomendacion.NombrePDF
		rutaPDF = existingRecomendacion.RutaPDF
	}

	recomendacion := entities.NewRecomendacionNutricional(
		municipioID,
		nombrePDF,
		rutaPDF,
		userID,
	)
	recomendacion.ID = int32(id)
	recomendacion.FechaSubida = existingRecomendacion.FechaSubida // Mantener fecha original

	updatedRecomendacion, err := ctrl.updateUseCase.Run(recomendacion)
	if err != nil {
		// Si falla y se subió nuevo archivo, eliminarlo
		if oldPDF != "" {
			ctrl.fileUploader.DeleteFile("pdfs", nombrePDF)
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la recomendación nutricional",
			"error":   err.Error(),
		})
		return
	}

	// Si todo salió bien y se subió nuevo archivo, eliminar el archivo viejo
	if oldPDF != "" {
		if err := ctrl.fileUploader.DeleteFile("pdfs", oldPDF); err != nil {
			// Logear error pero no afectar la respuesta al cliente
			fmt.Printf("Advertencia: No se pudo eliminar el archivo anterior %s: %v\n", oldPDF, err)
		}
	}

	c.JSON(http.StatusOK, updatedRecomendacion)
}