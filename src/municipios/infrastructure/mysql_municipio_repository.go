// mysql_municipio_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/APIGOINIFAP/src/core"
	repositories "github.com/vicpoo/APIGOINIFAP/src/municipios/domain"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/domain/entities"
)

type MySQLMunicipioRepository struct {
	conn *sql.DB
}

func NewMySQLMunicipioRepository() repositories.IMunicipio {
	conn := core.GetBD()
	return &MySQLMunicipioRepository{conn: conn}
}

func (mysql *MySQLMunicipioRepository) Save(municipio *entities.Municipio) error {
	query := `
		INSERT INTO municipios (clave_estado, clave_municipio, nombre)
		VALUES (?, ?, ?)
	`
	result, err := mysql.conn.Exec(query, municipio.ClaveEstado, municipio.ClaveMunicipio, municipio.Nombre)
	if err != nil {
		log.Println("Error al guardar el municipio:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	municipio.ID = int32(id)

	return nil
}

func (mysql *MySQLMunicipioRepository) Update(municipio *entities.Municipio) error {
	query := `
		UPDATE municipios
		SET clave_estado = ?, clave_municipio = ?, nombre = ?
		WHERE id_municipio = ?
	`
	result, err := mysql.conn.Exec(query, municipio.ClaveEstado, municipio.ClaveMunicipio, municipio.Nombre, municipio.ID)
	if err != nil {
		log.Println("Error al actualizar el municipio:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("municipio con ID %d no encontrado", municipio.ID)
	}

	return nil
}

func (mysql *MySQLMunicipioRepository) Delete(id int32) error {
	query := "DELETE FROM municipios WHERE id_municipio = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el municipio:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("municipio con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MySQLMunicipioRepository) GetById(id int32) (*entities.Municipio, error) {
	query := `
		SELECT id_municipio, clave_estado, clave_municipio, nombre
		FROM municipios
		WHERE id_municipio = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var municipio entities.Municipio
	err := row.Scan(&municipio.ID, &municipio.ClaveEstado, &municipio.ClaveMunicipio, &municipio.Nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("municipio con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el municipio por ID:", err)
		return nil, err
	}

	return &municipio, nil
}

func (mysql *MySQLMunicipioRepository) GetAll() ([]entities.Municipio, error) {
	query := `
		SELECT id_municipio, clave_estado, clave_municipio, nombre
		FROM municipios
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los municipios:", err)
		return nil, err
	}
	defer rows.Close()

	var municipios []entities.Municipio
	for rows.Next() {
		var municipio entities.Municipio
		err := rows.Scan(&municipio.ID, &municipio.ClaveEstado, &municipio.ClaveMunicipio, &municipio.Nombre)
		if err != nil {
			log.Println("Error al escanear el municipio:", err)
			return nil, err
		}
		municipios = append(municipios, municipio)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return municipios, nil
}