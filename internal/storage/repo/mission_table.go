package repo

import (
	"context"
	"errors"
	"fmt"
	"spy-cats/internal/models"
	"spy-cats/internal/storage"
	db_connection "spy-cats/pkg/pg_connection"

	"github.com/jackc/pgx/v4"
)

type missionTableImpl struct {
	client db_connection.Client
}

func NewMissionRepository(client db_connection.Client) storage.MissionRepository {
	return &missionTableImpl{client: client}
}

func (r *missionTableImpl) CreateMission(ctx context.Context, mission *models.MissionModel) (int64, error) {
	query := `
        INSERT INTO missions (cat_id, is_completed, created_at, updated_at) 
		VALUES ($1, $2, NOW(), NOW()) 
		RETURNING id`

	err := r.client.QueryRow(ctx, query,
		mission.CatID,
		mission.IsCompleted,
	).Scan(&mission.ID)

	if err != nil {
		return 0, fmt.Errorf("failed to create mission: %w", err)
	}

	return int64(mission.ID), nil
}

func (r *missionTableImpl) DeleteMission(ctx context.Context, id int64) error {
	query := `DELETE FROM missions WHERE id = $1 AND cat_id IS NULL`
	result, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("mission cannot be deleted because it is assigned to a cat")
	}
	return nil
}

func (r *missionTableImpl) GetMission(ctx context.Context, id int64) (*models.MissionModel, error) {
	query := `
		SELECT id, cat_id, is_completed, created_at, updated_at 
		FROM missions 
		WHERE id = $1`

	var mission models.MissionModel
	err := r.client.QueryRow(ctx, query, id).Scan(
		&mission.ID,
		&mission.CatID,
		&mission.IsCompleted,
		&mission.CreatedAt,
		&mission.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get mission: %w", err)
	}

	return &mission, nil
}

func (r *missionTableImpl) ListMissions(ctx context.Context) ([]models.MissionModel, error) {
	query := `SELECT id, cat_id, is_completed, created_at, 
	updated_at FROM missions`

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query missions: %w", err)
	}
	defer rows.Close()

	var missions []models.MissionModel
	for rows.Next() {
		var mission models.MissionModel
		err := rows.Scan(
			&mission.ID,
			&mission.CatID,
			&mission.IsCompleted,
			&mission.CreatedAt,
			&mission.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan mission: %w", err)
		}
		missions = append(missions, mission)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating mission rows: %w", err)
	}

	return missions, nil
}

func (r *missionTableImpl) CompleteMission(ctx context.Context, id int64) error {
	query := `UPDATE missions SET is_completed = TRUE, updated_at = NOW() WHERE id = $1`

	commandTag, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to change mission`s status: %w", err)
	}
	if commandTag.RowsAffected() == 0 {
		return storage.ErrNotFound
	}

	return nil
}
