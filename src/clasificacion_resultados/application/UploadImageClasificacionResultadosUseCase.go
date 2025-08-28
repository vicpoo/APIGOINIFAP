// UploadImageClasificacionResultadosUseCase.go
package application

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	repositories "github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain/entities"
)

type UploadImageClasificacionResultadosUseCase struct {
	repo repositories.IClasificacionResultados
}

func NewUploadImageClasificacionResultadosUseCase(repo repositories.IClasificacionResultados) *UploadImageClasificacionResultadosUseCase {
	return &UploadImageClasificacionResultadosUseCase{repo: repo}
}

func (uc *UploadImageClasificacionResultadosUseCase) Run(
	clasificacionID int32,
	file *multipart.FileHeader,
	uploadDir string,
) (*entities.ClasificacionResultados, error) {
	// Obtener la clasificación existente
	clasificacion, err := uc.repo.GetById(clasificacionID)
	if err != nil {
		return nil, err
	}

	// Generar nombre único para la imagen
	extension := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("clasificacion_%d_%d%s", clasificacionID, time.Now().Unix(), extension)
	
	// La variable filePath se usa para guardar el archivo (comentado)
	// filePath := filepath.Join(uploadDir, newFilename)

	// Guardar la imagen en el sistema de archivos (esto debería implementarse)
	// src, err := file.Open()
	// if err != nil {
	//     return nil, err
	// }
	// defer src.Close()
	// 
	// dst, err := os.Create(filePath)
	// if err != nil {
	//     return nil, err
	// }
	// defer dst.Close()
	// 
	// if _, err = io.Copy(dst, src); err != nil {
	//     return nil, err
	// }

	// Actualizar la ruta de la imagen en la entidad
	clasificacion.SetImagen(newFilename)

	// Guardar los cambios en la base de datos
	err = uc.repo.Update(clasificacion)
	if err != nil {
		return nil, err
	}

	return clasificacion, nil
}