package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/vicpoo/APIGOINIFAP/src/core"
	"github.com/vicpoo/APIGOINIFAP/src/users/domain"
	"github.com/vicpoo/APIGOINIFAP/src/users/domain/entities"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

type MySQLUserRepository struct {
	conn *sql.DB
}

func NewMySQLUserRepository() domain.IUser {
	conn := core.GetBD()
	return &MySQLUserRepository{conn: conn}
}

// Save guarda un usuario con la contraseña encriptada
func (repo *MySQLUserRepository) Save(user *entities.User) error {
	// Encriptar contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error al encriptar contraseña:", err)
		return err
	}

	query := `
		INSERT INTO users (nombre, apellido, correo, numero_telefonico, password, rol_id_FK)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := repo.conn.Exec(query, user.Nombre, user.Apellido, user.Correo, user.NumeroTelefonico, string(hash), user.RolIDFK)
	if err != nil {
		log.Println("Error al guardar el usuario:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	user.IDUser = int32(id)
	return nil
}

// Update actualiza un usuario (si se actualiza password se encripta)
func (repo *MySQLUserRepository) Update(user *entities.User) error {
	var query string
	var args []interface{}

	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error al encriptar contraseña:", err)
			return err
		}
		query = `
			UPDATE users
			SET nombre=?, apellido=?, correo=?, numero_telefonico=?, password=?, rol_id_FK=?
			WHERE ID_user=?
		`
		args = []interface{}{user.Nombre, user.Apellido, user.Correo, user.NumeroTelefonico, string(hash), user.RolIDFK, user.IDUser}
	} else {
		query = `
			UPDATE users
			SET nombre=?, apellido=?, correo=?, numero_telefonico=?, rol_id_FK=?
			WHERE ID_user=?
		`
		args = []interface{}{user.Nombre, user.Apellido, user.Correo, user.NumeroTelefonico, user.RolIDFK, user.IDUser}
	}

	result, err := repo.conn.Exec(query, args...)
	if err != nil {
		log.Println("Error al actualizar el usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", user.IDUser)
	}
	return nil
}

// Delete elimina un usuario
func (repo *MySQLUserRepository) Delete(id int32) error {
	query := "DELETE FROM users WHERE ID_user=?"
	result, err := repo.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", id)
	}
	return nil
}

// GetById obtiene un usuario por ID
func (repo *MySQLUserRepository) GetById(id int32) (*entities.User, error) {
	query := `
		SELECT ID_user, nombre, apellido, correo, numero_telefonico, password, rol_id_FK
		FROM users
		WHERE ID_user=?
	`
	row := repo.conn.QueryRow(query, id)
	var user entities.User
	err := row.Scan(&user.IDUser, &user.Nombre, &user.Apellido, &user.Correo, &user.NumeroTelefonico, &user.Password, &user.RolIDFK)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con ID %d no encontrado", id)
		}
		log.Println("Error al obtener usuario por ID:", err)
		return nil, err
	}
	return &user, nil
}

// GetAll obtiene todos los usuarios
func (repo *MySQLUserRepository) GetAll() ([]entities.User, error) {
	query := "SELECT ID_user, nombre, apellido, correo, numero_telefonico, password, rol_id_FK FROM users"
	rows, err := repo.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener usuarios:", err)
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.IDUser, &user.Nombre, &user.Apellido, &user.Correo, &user.NumeroTelefonico, &user.Password, &user.RolIDFK); err != nil {
			log.Println("Error al escanear usuario:", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Login verifica correo y password y devuelve token JWT
func (repo *MySQLUserRepository) Login(correo, password string) (*entities.User, error) {
	query := "SELECT ID_user, nombre, apellido, correo, numero_telefonico, password, rol_id_FK FROM users WHERE correo=?"
	row := repo.conn.QueryRow(query, correo)

	var user entities.User
	err := row.Scan(&user.IDUser, &user.Nombre, &user.Apellido, &user.Correo, &user.NumeroTelefonico, &user.Password, &user.RolIDFK)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, err
	}

	// Validar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("credenciales inválidas")
	}

	// Generar token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_user": user.IDUser,
		"correo":  user.Correo,
		"rol_id":  user.RolIDFK,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expira en 24h
	})

	secret := []byte("MI_SECRET_KEY") // cambiar por variable de entorno
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return nil, fmt.Errorf("error al generar token: %v", err)
	}

	// Retornar usuario con token en Password solo como placeholder
	user.Password = tokenString
	return &user, nil
}
