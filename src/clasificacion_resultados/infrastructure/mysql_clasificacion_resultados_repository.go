// mysql_clasificacion_resultados_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/vicpoo/APIGOINIFAP/src/core"
	repositories "github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain/entities"
)

type MySQLClasificacionResultadosRepository struct {
	conn *sql.DB
}

func NewMySQLClasificacionResultadosRepository() repositories.IClasificacionResultados {
	conn := core.GetBD()
	return &MySQLClasificacionResultadosRepository{conn: conn}
}

func (mysql *MySQLClasificacionResultadosRepository) Save(clasificacion *entities.ClasificacionResultados) error {
	query := `
		INSERT INTO clasificacion_resultados 
		(municipio_id_FK, analisis_tipo, fecha_analisis, resultado_general, 
		 nutrientes_criticos, comentario, imagen, user_id_FK, fecha_creacion)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(query, 
		clasificacion.MunicipioID, 
		clasificacion.AnalisisTipo, 
		clasificacion.FechaAnalisis,
		clasificacion.ResultadoGeneral,
		clasificacion.NutrientesCriticos,
		clasificacion.Comentario,
		clasificacion.Imagen,
		clasificacion.UserID,
		clasificacion.FechaCreacion,
	)
	if err != nil {
		log.Println("Error al guardar la clasificación de resultados:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	clasificacion.ID = int32(id)

	return nil
}

func (mysql *MySQLClasificacionResultadosRepository) Update(clasificacion *entities.ClasificacionResultados) error {
	query := `
		UPDATE clasificacion_resultados
		SET municipio_id_FK = ?, analisis_tipo = ?, fecha_analisis = ?, 
			resultado_general = ?, nutrientes_criticos = ?, comentario = ?, 
			imagen = ?, user_id_FK = ?
		WHERE id = ?
	`
	result, err := mysql.conn.Exec(query, 
		clasificacion.MunicipioID, 
		clasificacion.AnalisisTipo, 
		clasificacion.FechaAnalisis,
		clasificacion.ResultadoGeneral,
		clasificacion.NutrientesCriticos,
		clasificacion.Comentario,
		clasificacion.Imagen,
		clasificacion.UserID,
		clasificacion.ID,
	)
	if err != nil {
		log.Println("Error al actualizar la clasificación de resultados:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("clasificación de resultados con ID %d no encontrada", clasificacion.ID)
	}

	return nil
}

func (mysql *MySQLClasificacionResultadosRepository) Delete(id int32) error {
	query := "DELETE FROM clasificacion_resultados WHERE id = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la clasificación de resultados:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("clasificación de resultados con ID %d no encontrada", id)
	}

	return nil
}

func (mysql *MySQLClasificacionResultadosRepository) GetById(id int32) (*entities.ClasificacionResultados, error) {
	query := `
		SELECT id, municipio_id_FK, analisis_tipo, fecha_analisis, resultado_general,
			   nutrientes_criticos, comentario, imagen, user_id_FK, fecha_creacion
		FROM clasificacion_resultados
		WHERE id = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var clasificacion entities.ClasificacionResultados
	var fechaAnalisis, fechaCreacion time.Time
	
	err := row.Scan(
		&clasificacion.ID,
		&clasificacion.MunicipioID,
		&clasificacion.AnalisisTipo,
		&fechaAnalisis,
		&clasificacion.ResultadoGeneral,
		&clasificacion.NutrientesCriticos,
		&clasificacion.Comentario,
		&clasificacion.Imagen,
		&clasificacion.UserID,
		&fechaCreacion,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("clasificación de resultados con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la clasificación de resultados por ID:", err)
		return nil, err
	}

	clasificacion.FechaAnalisis = fechaAnalisis
	clasificacion.FechaCreacion = fechaCreacion

	return &clasificacion, nil
}

func (mysql *MySQLClasificacionResultadosRepository) GetAll() ([]entities.ClasificacionResultados, error) {
	query := `
		SELECT id, municipio_id_FK, analisis_tipo, fecha_analisis, resultado_general,
			   nutrientes_criticos, comentario, imagen, user_id_FK, fecha_creacion
		FROM clasificacion_resultados
		ORDER BY fecha_creacion DESC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todas las clasificaciones de resultados:", err)
		return nil, err
	}
	defer rows.Close()

	var clasificaciones []entities.ClasificacionResultados
	for rows.Next() {
		var clasificacion entities.ClasificacionResultados
		var fechaAnalisis, fechaCreacion time.Time
		
		err := rows.Scan(
			&clasificacion.ID,
			&clasificacion.MunicipioID,
			&clasificacion.AnalisisTipo,
			&fechaAnalisis,
			&clasificacion.ResultadoGeneral,
			&clasificacion.NutrientesCriticos,
			&clasificacion.Comentario,
			&clasificacion.Imagen,
			&clasificacion.UserID,
			&fechaCreacion,
		)
		if err != nil {
			log.Println("Error al escanear la clasificación de resultados:", err)
			return nil, err
		}
		
		clasificacion.FechaAnalisis = fechaAnalisis
		clasificacion.FechaCreacion = fechaCreacion
		clasificaciones = append(clasificaciones, clasificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return clasificaciones, nil
}

func (mysql *MySQLClasificacionResultadosRepository) GetByMunicipioID(municipioID int32) ([]entities.ClasificacionResultados, error) {
	query := `
		SELECT id, municipio_id_FK, analisis_tipo, fecha_analisis, resultado_general,
			   nutrientes_criticos, comentario, imagen, user_id_FK, fecha_creacion
		FROM clasificacion_resultados
		WHERE municipio_id_FK = ?
		ORDER BY fecha_creacion DESC
	`
	rows, err := mysql.conn.Query(query, municipioID)
	if err != nil {
		log.Println("Error al obtener clasificaciones por municipio:", err)
		return nil, err
	}
	defer rows.Close()

	var clasificaciones []entities.ClasificacionResultados
	for rows.Next() {
		var clasificacion entities.ClasificacionResultados
		var fechaAnalisis, fechaCreacion time.Time
		
		err := rows.Scan(
			&clasificacion.ID,
			&clasificacion.MunicipioID,
			&clasificacion.AnalisisTipo,
			&fechaAnalisis,
			&clasificacion.ResultadoGeneral,
			&clasificacion.NutrientesCriticos,
			&clasificacion.Comentario,
			&clasificacion.Imagen,
			&clasificacion.UserID,
			&fechaCreacion,
		)
		if err != nil {
			log.Println("Error al escanear la clasificación de resultados:", err)
			return nil, err
		}
		
		clasificacion.FechaAnalisis = fechaAnalisis
		clasificacion.FechaCreacion = fechaCreacion
		clasificaciones = append(clasificaciones, clasificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return clasificaciones, nil
}