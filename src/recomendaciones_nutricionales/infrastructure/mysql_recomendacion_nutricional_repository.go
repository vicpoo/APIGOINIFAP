// mysql_recomendacion_nutricional_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/APIGOINIFAP/src/core"
	repositories "github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type MySQLRecomendacionNutricionalRepository struct {
	conn *sql.DB
}

func NewMySQLRecomendacionNutricionalRepository() repositories.IRecomendacionNutricional {
	conn := core.GetBD()
	return &MySQLRecomendacionNutricionalRepository{conn: conn}
}

func (mysql *MySQLRecomendacionNutricionalRepository) Save(recomendacion *entities.RecomendacionNutricional) error {
	query := `
		INSERT INTO recomendaciones_nutricionales (municipio_id_FK, nombre_pdf, ruta_pdf, fecha_subida, user_id_FK)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(query, 
		recomendacion.MunicipioID, 
		recomendacion.NombrePDF, 
		recomendacion.RutaPDF, 
		recomendacion.FechaSubida, 
		recomendacion.UserID,
	)
	if err != nil {
		log.Println("Error al guardar la recomendación nutricional:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	recomendacion.ID = int32(id)

	return nil
}

func (mysql *MySQLRecomendacionNutricionalRepository) Update(recomendacion *entities.RecomendacionNutricional) error {
	query := `
		UPDATE recomendaciones_nutricionales
		SET municipio_id_FK = ?, nombre_pdf = ?, ruta_pdf = ?, fecha_subida = ?, user_id_FK = ?
		WHERE id = ?
	`
	result, err := mysql.conn.Exec(query, 
		recomendacion.MunicipioID, 
		recomendacion.NombrePDF, 
		recomendacion.RutaPDF, 
		recomendacion.FechaSubida, 
		recomendacion.UserID,
		recomendacion.ID,
	)
	if err != nil {
		log.Println("Error al actualizar la recomendación nutricional:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("recomendación nutricional con ID %d no encontrada", recomendacion.ID)
	}

	return nil
}

func (mysql *MySQLRecomendacionNutricionalRepository) Delete(id int32) error {
	query := "DELETE FROM recomendaciones_nutricionales WHERE id = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la recomendación nutricional:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("recomendación nutricional con ID %d no encontrada", id)
	}

	return nil
}

func (mysql *MySQLRecomendacionNutricionalRepository) GetById(id int32) (*entities.RecomendacionNutricional, error) {
	query := `
		SELECT id, municipio_id_FK, nombre_pdf, ruta_pdf, fecha_subida, user_id_FK
		FROM recomendaciones_nutricionales
		WHERE id = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var recomendacion entities.RecomendacionNutricional
	err := row.Scan(
		&recomendacion.ID,
		&recomendacion.MunicipioID,
		&recomendacion.NombrePDF,
		&recomendacion.RutaPDF,
		&recomendacion.FechaSubida,
		&recomendacion.UserID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("recomendación nutricional con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la recomendación nutricional por ID:", err)
		return nil, err
	}

	return &recomendacion, nil
}

func (mysql *MySQLRecomendacionNutricionalRepository) GetAll() ([]entities.RecomendacionNutricional, error) {
	query := `
		SELECT id, municipio_id_FK, nombre_pdf, ruta_pdf, fecha_subida, user_id_FK
		FROM recomendaciones_nutricionales
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todas las recomendaciones nutricionales:", err)
		return nil, err
	}
	defer rows.Close()

	var recomendaciones []entities.RecomendacionNutricional
	for rows.Next() {
		var recomendacion entities.RecomendacionNutricional
		err := rows.Scan(
			&recomendacion.ID,
			&recomendacion.MunicipioID,
			&recomendacion.NombrePDF,
			&recomendacion.RutaPDF,
			&recomendacion.FechaSubida,
			&recomendacion.UserID,
		)
		if err != nil {
			log.Println("Error al escanear la recomendación nutricional:", err)
			return nil, err
		}
		recomendaciones = append(recomendaciones, recomendacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return recomendaciones, nil
}

func (mysql *MySQLRecomendacionNutricionalRepository) GetByMunicipioID(municipioID int32) ([]entities.RecomendacionNutricional, error) {
	query := `
		SELECT id, municipio_id_FK, nombre_pdf, ruta_pdf, fecha_subida, user_id_FK
		FROM recomendaciones_nutricionales
		WHERE municipio_id_FK = ?
		ORDER BY fecha_subida DESC
	`
	rows, err := mysql.conn.Query(query, municipioID)
	if err != nil {
		log.Println("Error al obtener recomendaciones por municipio:", err)
		return nil, err
	}
	defer rows.Close()

	var recomendaciones []entities.RecomendacionNutricional
	for rows.Next() {
		var recomendacion entities.RecomendacionNutricional
		err := rows.Scan(
			&recomendacion.ID,
			&recomendacion.MunicipioID,
			&recomendacion.NombrePDF,
			&recomendacion.RutaPDF,
			&recomendacion.FechaSubida,
			&recomendacion.UserID,
		)
		if err != nil {
			log.Println("Error al escanear la recomendación nutricional:", err)
			return nil, err
		}
		recomendaciones = append(recomendaciones, recomendacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return recomendaciones, nil
}