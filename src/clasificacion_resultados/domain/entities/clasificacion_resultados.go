// clasificacion_resultados.go
package entities

import (
	"time"
)

type ClasificacionResultados struct {
	ID               int32     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	MunicipioID      int32     `json:"municipio_id_FK" gorm:"column:municipio_id_FK;not null"`
	AnalisisTipo     string    `json:"analisis_tipo" gorm:"column:analisis_tipo;not null;type:varchar(50)"`
	FechaAnalisis    time.Time `json:"fecha_analisis" gorm:"column:fecha_analisis;not null;type:date"`
	ResultadoGeneral string    `json:"resultado_general" gorm:"column:resultado_general;not null;type:varchar(100)"`
	NutrientesCriticos string  `json:"nutrientes_criticos" gorm:"column:nutrientes_criticos;type:text"`
	Comentario       string    `json:"comentario" gorm:"column:comentario;type:text"`
	Imagen           string    `json:"imagen" gorm:"column:imagen;type:varchar(500)"`
	UserID           int32     `json:"user_id_FK" gorm:"column:user_id_FK;not null"`
	FechaCreacion    time.Time `json:"fecha_creacion" gorm:"column:fecha_creacion;not null;default:CURRENT_TIMESTAMP"`
}

// Setters
func (c *ClasificacionResultados) SetID(id int32) {
	c.ID = id
}

func (c *ClasificacionResultados) SetMunicipioID(municipioID int32) {
	c.MunicipioID = municipioID
}

func (c *ClasificacionResultados) SetAnalisisTipo(analisisTipo string) {
	c.AnalisisTipo = analisisTipo
}

func (c *ClasificacionResultados) SetFechaAnalisis(fechaAnalisis time.Time) {
	c.FechaAnalisis = fechaAnalisis
}

func (c *ClasificacionResultados) SetResultadoGeneral(resultadoGeneral string) {
	c.ResultadoGeneral = resultadoGeneral
}

func (c *ClasificacionResultados) SetNutrientesCriticos(nutrientesCriticos string) {
	c.NutrientesCriticos = nutrientesCriticos
}

func (c *ClasificacionResultados) SetComentario(comentario string) {
	c.Comentario = comentario
}

func (c *ClasificacionResultados) SetImagen(imagen string) {
	c.Imagen = imagen
}

func (c *ClasificacionResultados) SetUserID(userID int32) {
	c.UserID = userID
}

func (c *ClasificacionResultados) SetFechaCreacion(fechaCreacion time.Time) {
	c.FechaCreacion = fechaCreacion
}

// Getters
func (c *ClasificacionResultados) GetID() int32 {
	return c.ID
}

func (c *ClasificacionResultados) GetMunicipioID() int32 {
	return c.MunicipioID
}

func (c *ClasificacionResultados) GetAnalisisTipo() string {
	return c.AnalisisTipo
}

func (c *ClasificacionResultados) GetFechaAnalisis() time.Time {
	return c.FechaAnalisis
}

func (c *ClasificacionResultados) GetResultadoGeneral() string {
	return c.ResultadoGeneral
}

func (c *ClasificacionResultados) GetNutrientesCriticos() string {
	return c.NutrientesCriticos
}

func (c *ClasificacionResultados) GetComentario() string {
	return c.Comentario
}

func (c *ClasificacionResultados) GetImagen() string {
	return c.Imagen
}

func (c *ClasificacionResultados) GetUserID() int32 {
	return c.UserID
}

func (c *ClasificacionResultados) GetFechaCreacion() time.Time {
	return c.FechaCreacion
}

// Constructor
func NewClasificacionResultados(
	municipioID int32,
	analisisTipo string,
	fechaAnalisis time.Time,
	resultadoGeneral string,
	nutrientesCriticos string,
	comentario string,
	imagen string,
	userID int32,
) *ClasificacionResultados {
	return &ClasificacionResultados{
		MunicipioID:      municipioID,
		AnalisisTipo:     analisisTipo,
		FechaAnalisis:    fechaAnalisis,
		ResultadoGeneral: resultadoGeneral,
		NutrientesCriticos: nutrientesCriticos,
		Comentario:       comentario,
		Imagen:           imagen,
		UserID:           userID,
		FechaCreacion:    time.Now(),
	}
}