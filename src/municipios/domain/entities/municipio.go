// municipio.go
package entities

type Municipio struct {
	ID             int32  `json:"id_municipio" gorm:"column:id_municipio;primaryKey;autoIncrement"`
	ClaveEstado    int32  `json:"clave_estado" gorm:"column:clave_estado;not null"`
	ClaveMunicipio int32  `json:"clave_municipio" gorm:"column:clave_municipio;not null"`
	Nombre         string `json:"nombre" gorm:"column:nombre;not null"`
}

// Setters
func (m *Municipio) SetID(id int32) {
	m.ID = id
}

func (m *Municipio) SetClaveEstado(claveEstado int32) {
	m.ClaveEstado = claveEstado
}

func (m *Municipio) SetClaveMunicipio(claveMunicipio int32) {
	m.ClaveMunicipio = claveMunicipio
}

func (m *Municipio) SetNombre(nombre string) {
	m.Nombre = nombre
}

// Getters
func (m *Municipio) GetID() int32 {
	return m.ID
}

func (m *Municipio) GetClaveEstado() int32 {
	return m.ClaveEstado
}

func (m *Municipio) GetClaveMunicipio() int32 {
	return m.ClaveMunicipio
}

func (m *Municipio) GetNombre() string {
	return m.Nombre
}

// Constructor
func NewMunicipio(claveEstado int32, claveMunicipio int32, nombre string) *Municipio {
	return &Municipio{
		ClaveEstado:    claveEstado,
		ClaveMunicipio: claveMunicipio,
		Nombre:         nombre,
	}
}