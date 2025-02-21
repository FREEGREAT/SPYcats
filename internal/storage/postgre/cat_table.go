package postgre

import (
	"context"
	"errors"
	"fmt"
	"spy-cats/internal/models"
	"spy-cats/internal/storage"
	db_connection "spy-cats/pkg/pg_connection"

	"github.com/jackc/pgx/v4"
)

type catTableImpl struct {
	client db_connection.Client
}

func NewCatRepository(client db_connection.Client) storage.CatRepository {
	return &catTableImpl{client: client}
}

func (r *catTableImpl) CreateSpyCat(ctx context.Context, cat *models.CatModel) (int64, error) {
	query := `
        INSERT INTO spy_cats (name, experience_years, breed, salary)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	err := r.client.QueryRow(ctx, query,
		cat.Name,
		cat.ExperienceYears,
		cat.Breed,
		cat.Salary,
	).Scan(&cat.ID)

	if err != nil {
		return 0, fmt.Errorf("failed to create spy cat: %w", err)
	}

	return *cat.ID, err
}

func (r *catTableImpl) DeleteSpyCat(ctx context.Context, id *int64) error {
	query := `DELETE FROM spy_cats WHERE id = $1`

	commandTag, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete spy cat: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return storage.ErrNotFound
	}

	return nil
}

func (r *catTableImpl) GetSpyCat(ctx context.Context, id int64) (*models.CatModel, error) {
	query := `
        SELECT id, name, experience_years, breed, salary, created_at, updated_at
        FROM spy_cats
        WHERE id = $1`

	var cat models.CatModel
	err := r.client.QueryRow(ctx, query, id).Scan(
		&cat.ID,
		&cat.Name,
		&cat.ExperienceYears,
		&cat.Breed,
		&cat.Salary,
		&cat.CreatedAt,
		&cat.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get spy cat: %w", err)
	}

	return &cat, nil
}

func (r *catTableImpl) ListSpyCats(ctx context.Context) ([]models.CatModel, error) {
	query := `
        SELECT id, name, experience_years, breed, salary, created_at, updated_at
        FROM spy_cats`

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query spy cats: %w", err)
	}
	defer rows.Close()

	var cats []models.CatModel
	for rows.Next() {
		var cat models.CatModel
		err := rows.Scan(
			&cat.ID,
			&cat.Name,
			&cat.ExperienceYears,
			&cat.Breed,
			&cat.Salary,
			&cat.CreatedAt,
			&cat.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan spy cat: %w", err)
		}
		cats = append(cats, cat)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating spy cats rows: %w", err)
	}

	return cats, nil
}

func (r *catTableImpl) UpdateSpyCatSalary(ctx context.Context, id int64, newSalary float64) (float64, error) {
	var response float64
	query := `
        UPDATE spy_cats
        SET salary = $1, updated_at = CURRENT_TIMESTAMP
        WHERE id = $2
		RETURNING salary`

	commandTag, err := r.client.Exec(ctx, query, newSalary, id)
	if err != nil {
		return 0, fmt.Errorf("failed to update spy cat salary: %w", err)
	}

	err = r.client.QueryRow(ctx, query,
		newSalary,
		id,
	).Scan(response)

	if commandTag.RowsAffected() == 0 {
		return 0, storage.ErrNotFound
	}

	return response, nil
}
