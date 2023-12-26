package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/haziqam/course-scheduler-backend/packages/database"
	"github.com/haziqam/course-scheduler-backend/packages/models"
	"github.com/haziqam/course-scheduler-backend/packages/utils"
)

type FakultasRepo struct {
	ctx context.Context
}

var fakultasRepoInstance *FakultasRepo = nil

func GetFakultasRepoInstance() *FakultasRepo {
	if fakultasRepoInstance == nil {
		fakultasRepoInstance = new(FakultasRepo)
		fakultasRepoInstance.ctx = context.Background()
	}
	return fakultasRepoInstance
}

func (repo *FakultasRepo) GetAllFakultas() ([]models.Fakultas, error) {
	query := `
		SELECT * 
		FROM fakultas
	`
	rows, err := database.DbInstance.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var fakultasArr []models.Fakultas

	for rows.Next() {
		fakultas, err := scanIntoFakultas(rows)
		if err != nil {
			return nil, err
		}
		fakultasArr = append(fakultasArr, *fakultas)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return fakultasArr, nil
}

func (repo *FakultasRepo) GetFakultasByName(name string) (*models.Fakultas, error) {
	query := `
		SELECT *
		FROM fakultas
		WHERE nama_fakultas = ($1)
	`
	row := database.DbInstance.QueryRow(query, name)
	return scanIntoFakultas(row)
}

func (repo *FakultasRepo) AddFakultas(newFakultas []models.Fakultas) error {
	tx, err := database.DbInstance.BeginTx(repo.ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	if err != nil {
		return errors.New("Failed to begin transaction: " + err.Error())
	}

	query := `
		INSERT INTO fakultas(nama_fakultas)
		VALUES ($1)
	`

	for _, fakultas := range newFakultas {
		_, err := tx.Exec(query, fakultas.NamaFakultas)
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

func (repo *FakultasRepo) RemoveFakultasByName(fakultasName string) error {
	query := `
		DELETE FROM fakultas
		WHERE nama_fakultas = ($1)
	`

	_, err := database.DbInstance.Exec(query, fakultasName)

	if err != nil {
		return errors.New("Failed to delete fakultas: " + err.Error())
	}

	return nil
}

func (repo *FakultasRepo) UpdateFakultas(oldFakultasName string, newFakultasName string) error {
	query := `
		UPDATE fakultas
		SET nama_fakultas = ($1)
		WHERE nama_fakultas = ($2)
	`

	_, err := database.DbInstance.Exec(query, newFakultasName, oldFakultasName)

	if err != nil {
		return errors.New("Failed to update fakultas: " + err.Error())
	}

	return nil
}

func scanIntoFakultas(scanner utils.IScanner) (*models.Fakultas, error) {
	fakultas := new(models.Fakultas)
	err := scanner.Scan(&fakultas.NamaFakultas)
	if err != nil {
		return nil, err
	}
	return fakultas, nil
}
