// user.go
package entities

type User struct {
	IDUser           int32  `json:"id_user"`
	Nombre           string `json:"nombre"`
	Apellido         string `json:"apellido"`
	Correo           string `json:"correo"`
	NumeroTelefonico string `json:"numero_telefonico"`
	Password         string `json:"password"`
	RolIDFK          int32  `json:"rol_id_FK"`
}

// Constructor
func NewUser(nombre, apellido, correo, numeroTelefonico, password string, rolID int32) *User {
	return &User{
		Nombre:           nombre,
		Apellido:         apellido,
		Correo:           correo,
		NumeroTelefonico: numeroTelefonico,
		Password:         password,
		RolIDFK:          rolID,
	}
}

// Setters
func (u *User) SetIDUser(id int32) {
	u.IDUser = id
}

func (u *User) SetNombre(nombre string) {
	u.Nombre = nombre
}

func (u *User) SetApellido(apellido string) {
	u.Apellido = apellido
}

func (u *User) SetCorreo(correo string) {
	u.Correo = correo
}

func (u *User) SetNumeroTelefonico(numero string) {
	u.NumeroTelefonico = numero
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetRolIDFK(rolID int32) {
	u.RolIDFK = rolID
}

// Getters
func (u *User) GetIDUser() int32 {
	return u.IDUser
}

func (u *User) GetNombre() string {
	return u.Nombre
}

func (u *User) GetApellido() string {
	return u.Apellido
}

func (u *User) GetCorreo() string {
	return u.Correo
}

func (u *User) GetNumeroTelefonico() string {
	return u.NumeroTelefonico
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetRolIDFK() int32 {
	return u.RolIDFK
}
