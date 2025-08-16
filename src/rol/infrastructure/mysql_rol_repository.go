//mysql_rol_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/APIGOINIFAP/src/core"
	repositories "github.com/vicpoo/APIGOINIFAP/src/rol/domain"
	"github.com/vicpoo/APIGOINIFAP/src/rol/domain/entities"
)

type MySQLRolRepository struct {
	conn *sql.DB
}

func NewMySQLRolRepository() repositories.IRol {
	conn := core.GetBD()
	return &MySQLRolRepository{conn: conn}
}

func (mysql *MySQLRolRepository) Save(rol *entities.Rol) error {
	query := `
		INSERT INTO rol (titulo)
		VALUES (?)
	`
	result, err := mysql.conn.Exec(query, rol.Titulo)
	if err != nil {
		log.Println("Error al guardar el rol:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	rol.ID = int32(id)

	return nil
}

func (mysql *MySQLRolRepository) Update(rol *entities.Rol) error {
	query := `
		UPDATE rol
		SET titulo = ?
		WHERE id_rol = ?
	`
	result, err := mysql.conn.Exec(query, rol.Titulo, rol.ID)
	if err != nil {
		log.Println("Error al actualizar el rol:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rol con ID %d no encontrado", rol.ID)
	}

	return nil
}

func (mysql *MySQLRolRepository) Delete(id int32) error {
	query := "DELETE FROM rol WHERE id_rol = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el rol:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rol con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MySQLRolRepository) GetById(id int32) (*entities.Rol, error) {
	query := `
		SELECT id_rol, titulo
		FROM rol
		WHERE id_rol = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var rol entities.Rol
	err := row.Scan(&rol.ID, &rol.Titulo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("rol con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el rol por ID:", err)
		return nil, err
	}

	return &rol, nil
}

func (mysql *MySQLRolRepository) GetAll() ([]entities.Rol, error) {
	query := `
		SELECT id_rol, titulo
		FROM rol
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los roles:", err)
		return nil, err
	}
	defer rows.Close()

	var roles []entities.Rol
	for rows.Next() {
		var rol entities.Rol
		err := rows.Scan(&rol.ID, &rol.Titulo)
		if err != nil {
			log.Println("Error al escanear el rol:", err)
			return nil, err
		}
		roles = append(roles, rol)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return roles, nil
}
