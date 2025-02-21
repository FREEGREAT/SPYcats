package postgre

import (
	"context"
	"database/sql"
	"errors"
	"spy-cats/internal/models"
	"spy-cats/internal/storage"
	db_connection "spy-cats/pkg/pg_connection"
)

type targetRepository struct {
	client db_connection.Client
}

func NewTargetRepository(client db_connection.Client) storage.TargetRepository {
	return &targetRepository{client: client}
}

func (r *targetRepository) CreateTarget(ctx context.Context, target *models.TargetModel) (int, error) {
	var isCompleted bool

	err := r.client.QueryRow(ctx, `SELECT is_completed FROM missions WHERE id = $1`, target.MissionID).Scan(&isCompleted)
	if err != nil {
		return 0, err
	}
	if isCompleted {
		return 0, errors.New("cannot add targets to a completed mission")
	}

	query := `
		INSERT INTO targets (mission_id, name, country, is_completed, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, NOW(), NOW()) 
		RETURNING id`

	var id int
	err = r.client.QueryRow(ctx, query, target.MissionID, target.Name, target.Country, target.IsCompleted).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *targetRepository) DeleteTarget(ctx context.Context, id int64) error {
	query := `DELETE FROM targets WHERE id = $1 AND is_completed = FALSE`

	result, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("target cannot be deleted because it is already completed")
	}
	return nil
}

func (r *targetRepository) GetTarget(ctx context.Context, id int64) (*models.TargetModel, error) {
	query := `
	SELECT id, mission_id, name, country, is_completed, created_at, updated_at 
	FROM targets WHERE id = $1`
	target := &models.TargetModel{}

	err := r.client.QueryRow(ctx, query, id).Scan(
		&target.ID, &target.MissionID, &target.Name, &target.Country, &target.IsCompleted, &target.CreatedAt, &target.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return target, nil
}

func (r *targetRepository) ListTargetsByMission(ctx context.Context, missionID int64) ([]models.TargetModel, error) {
	query := `SELECT id, mission_id, name, country, is_completed, created_at, updated_at FROM targets WHERE mission_id = $1`

	rows, err := r.client.Query(ctx, query, missionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var targets []models.TargetModel
	for rows.Next() {
		var target models.TargetModel
		if err := rows.Scan(&target.ID, &target.MissionID, &target.Name, &target.Country, &target.IsCompleted, &target.CreatedAt, &target.UpdatedAt); err != nil {
			return nil, err
		}
		targets = append(targets, target)
	}

	return targets, nil
}

func (r *targetRepository) CompleteTarget(ctx context.Context, id int64) error {
	query := `UPDATE targets SET is_completed = TRUE, updated_at = NOW() WHERE id = $1`
	_, err := r.client.Exec(ctx, query, id)
	return err
}
