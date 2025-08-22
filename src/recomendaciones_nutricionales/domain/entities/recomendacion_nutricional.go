// recomendacion_nutricional.go
package entities

import (
	"time"
)

type RecomendacionNutricional struct {
	ID           int32     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	MunicipioID  int32     `json:"municipio_id_FK" gorm:"column:municipio_id_FK;not null"`
	NombrePDF    string    `json:"nombre_pdf" gorm:"column:nombre_pdf;not null"`
	RutaPDF      string    `json:"ruta_pdf" gorm:"column:ruta_pdf;not null"`
	FechaSubida  time.Time `json:"fecha_subida" gorm:"column:fecha_subida;not null"`
	UserID       int32     `json:"user_id_FK" gorm:"column:user_id_FK;not null"`
}

// Setters
func (r *RecomendacionNutricional) SetID(id int32) {
	r.ID = id
}

func (r *RecomendacionNutricional) SetMunicipioID(municipioID int32) {
	r.MunicipioID = municipioID
}

func (r *RecomendacionNutricional) SetNombrePDF(nombrePDF string) {
	r.NombrePDF = nombrePDF
}

func (r *RecomendacionNutricional) SetRutaPDF(rutaPDF string) {
	r.RutaPDF = rutaPDF
}

func (r *RecomendacionNutricional) SetFechaSubida(fechaSubida time.Time) {
	r.FechaSubida = fechaSubida
}

func (r *RecomendacionNutricional) SetUserID(userID int32) {
	r.UserID = userID
}

// Getters
func (r *RecomendacionNutricional) GetID() int32 {
	return r.ID
}

func (r *RecomendacionNutricional) GetMunicipioID() int32 {
	return r.MunicipioID
}

func (r *RecomendacionNutricional) GetNombrePDF() string {
	return r.NombrePDF
}

func (r *RecomendacionNutricional) GetRutaPDF() string {
	return r.RutaPDF
}

func (r *RecomendacionNutricional) GetFechaSubida() time.Time {
	return r.FechaSubida
}

func (r *RecomendacionNutricional) GetUserID() int32 {
	return r.UserID
}

// Constructor
func NewRecomendacionNutricional(municipioID int32, nombrePDF string, rutaPDF string, userID int32) *RecomendacionNutricional {
	return &RecomendacionNutricional{
		MunicipioID: municipioID,
		NombrePDF:   nombrePDF,
		RutaPDF:     rutaPDF,
		FechaSubida: time.Now(),
		UserID:      userID,
	}
}