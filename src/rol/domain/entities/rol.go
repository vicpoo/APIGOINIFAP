// rol.go
package entities

type Rol struct {
	ID     int32  `json:"id_rol" gorm:"column:id_rol;primaryKey;autoIncrement"`
	Titulo string `json:"titulo" gorm:"column:titulo;not null"`
}

// Setters
func (r *Rol) SetID(id int32) {
	r.ID = id
}

func (r *Rol) SetTitulo(titulo string) {
	r.Titulo = titulo
}

// Getters
func (r *Rol) GetID() int32 {
	return r.ID
}

func (r *Rol) GetTitulo() string {
	return r.Titulo
}

// Constructor
func NewRol(titulo string) *Rol {
	return &Rol{
		Titulo: titulo,
	}
}
