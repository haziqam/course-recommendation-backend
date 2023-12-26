package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/haziqam/course-scheduler-backend/packages/database"
	"github.com/haziqam/course-scheduler-backend/packages/models"
	"github.com/haziqam/course-scheduler-backend/packages/utils"
)

type JurusanRepo struct {
	ctx context.Context
}

var jurusanRepoInstance *JurusanRepo = nil

func GetJurusanRepoInstance() *JurusanRepo {
	if jurusanRepoInstance == nil {
		jurusanRepoInstance = new(JurusanRepo)
		jurusanRepoInstance.ctx = context.Background()
	}
	return jurusanRepoInstance
}

func (repo *JurusanRepo) GetAllJurusan() ([]models.Jurusan, error) {
	query := `
		SELECT * 
		FROM jurusan
	`
	rows, err := database.DbInstance.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jurusanArr []models.Jurusan

	for rows.Next() {
		jurusan, err := scanIntoJurusan(rows)
		if err != nil {
			return nil, err
		}
		jurusanArr = append(jurusanArr, *jurusan)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return jurusanArr, nil
}

func (repo *JurusanRepo) GetJurusanByName(name string) (*models.Jurusan, error) {
	query := `
		SELECT *
		FROM jurusan
		WHERE nama_jurusan = ($1)
	`
	row := database.DbInstance.QueryRow(query, name)
	return scanIntoJurusan(row)
}

func (repo *JurusanRepo) AddJurusan(newJurusan []models.Jurusan) error {
	tx, err := database.DbInstance.BeginTx(repo.ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	if err != nil {
		return errors.New("Failed to begin transaction: " + err.Error())
	}

	query := `
		INSERT INTO jurusan(nama_jurusan, nama_fakultas)
		VALUES ($1, $2)
	`

	for _, jurusan := range newJurusan {
		_, err := tx.Exec(query, jurusan.NamaJurusan, jurusan.NamaFakultas)
		if err != nil {
			tx.Rollback()
			return errors.New("Failed to complete transaction: " + err.Error())
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return errors.New("Failed to commit transaction: " + err.Error())
	}

	return nil
}

func (repo *JurusanRepo) RemoveJurusanByName(jurusanName string) error {
	query := `
		DELETE FROM jurusan
		WHERE nama_jurusan = ($1)
	`

	_, err := database.DbInstance.Exec(query, jurusanName)

	if err != nil {
		return errors.New("Failed to delete jurusan: " + err.Error())
	}

	return nil
}

func scanIntoJurusan(scanner utils.IScanner) (*models.Jurusan, error) {
	jurusan := new(models.Jurusan)
	err := scanner.Scan(&jurusan.NamaJurusan, &jurusan.NamaFakultas)
	if err != nil {
		return nil, err
	}
	return jurusan, nil
}
