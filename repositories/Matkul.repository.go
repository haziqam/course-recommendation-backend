package repositories

import (
	"context"

	"database/sql"
	"errors"

	"github.com/haziqam/course-scheduler-backend/packages/database"
	"github.com/haziqam/course-scheduler-backend/packages/models"
	"github.com/haziqam/course-scheduler-backend/packages/utils"
)

type MatkulRepo struct {
	ctx context.Context
}

var matkulRepoInstance *MatkulRepo = nil

func GetMatkulRepoInstance() *MatkulRepo {
	if matkulRepoInstance == nil {
		matkulRepoInstance = new(MatkulRepo)
		matkulRepoInstance.ctx = context.Background()
	}
	return matkulRepoInstance
}

func (repo *MatkulRepo) GetAllMatkul() ([]models.Matkul, error) {
	query := `
		SELECT * 
		FROM matkul
	`
	rows, err := database.DbInstance.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var matkulArr []models.Matkul

	for rows.Next() {
		matkul := new(models.Matkul)
		err = matkul.ScanRow(rows)
		if err != nil {
			return nil, err
		}
		matkulArr = append(matkulArr, *matkul)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return matkulArr, nil
}

func (repo *MatkulRepo) AddMatkul(newMatkul []models.Matkul) error {
	tx, err := database.DbInstance.BeginTx(repo.ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	if err != nil {
		return errors.New("Failed to begin transaction: " + err.Error())
	}

	query := `
		INSERT INTO matkul(nama_matkul, sks, nama_jurusan, min_semester, prediksi) 
		VALUES ($1, $2, $3, $4, $5)
	`

	for _, matkul := range newMatkul {
		_, err := tx.Exec(
			query,
			matkul.NamaMatkul,
			matkul.SKS,
			matkul.NamaJurusan,
			matkul.MinSemester,
			matkul.PrediksiIndeks,
		)
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

func (repo *MatkulRepo) RemoveMatkulByNameAndJurusan(matkulName string, jurusanName string) error {
	query := `
		DELETE FROM matkul
		WHERE nama_matkul = ($1)
		AND nama_jurusan = ($2)
	`

	_, err := database.DbInstance.Exec(query, matkulName)

	if err != nil {
		return errors.New("Failed to delete matkul: " + err.Error())
	}

	return nil
}

func (repo *MatkulRepo) FilterMatkul(namaFakultas string, currentSemester int) ([]models.Matkul, error) {
	query := `
		SELECT nama_matkul, sks, nama_jurusan, min_semester, prediksi
		FROM matkul NATURAL JOIN jurusan
		WHERE nama_fakultas = $1
		AND min_semester <= $2;
	`

	rows, err := database.DbInstance.Query(query, namaFakultas, currentSemester)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var filteredMatkul []models.Matkul

	for rows.Next() {
		matkul := new(models.Matkul)
		err = matkul.ScanRow(rows)
		if err != nil {
			return nil, err
		}
		filteredMatkul = append(filteredMatkul, *matkul)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return filteredMatkul, nil
}

func scanIntoMatkul(scanner utils.IScanner) (*models.Matkul, error) {
	matkul := new(models.Matkul)
	err := scanner.Scan(
		&matkul.NamaMatkul,
		&matkul.SKS,
		&matkul.NamaJurusan,
		&matkul.MinSemester,
		&matkul.PrediksiIndeks,
	)
	if err != nil {
		return nil, err
	}
	return matkul, nil
}
