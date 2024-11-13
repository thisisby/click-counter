package postgre

import (
	"click-counter/internal/models"
	"click-counter/internal/repositories"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type postgreClickRepository struct {
	conn *sqlx.DB
}

func NewPostgreClickRepository(conn *sqlx.DB) repositories.ClickRepository {
	return &postgreClickRepository{
		conn: conn,
	}
}

func (p *postgreClickRepository) Save(bannerID string) error {
	query := `
		INSERT INTO clicks(created_at, banner_id)
		VALUES (NOW(), $1)
	`

	_, err := p.conn.Exec(query, bannerID)
	if err != nil {
		return fmt.Errorf("postgreClickRepository.Save - r.conn.Exec: %w", err)
	}

	return nil
}

func (p *postgreClickRepository) FindByIdInRange(id string, startTimestamp, endTimestamp time.Time) ([]models.Click, error) {
	query := `
		SELECT created_at, banner_id 
		FROM clicks 
		WHERE banner_id = $1 AND created_at > $2 AND created_at < $3
	`

	rows, err := p.conn.Query(query, id, startTimestamp, endTimestamp)
	if err != nil {
		return nil, fmt.Errorf("postgreClickRepository.FindByIdInRange: %w", err)
	}
	defer rows.Close()

	var clicks []models.Click
	for rows.Next() {
		var click models.Click
		if err := rows.Scan(&click.CreatedAt, &click.BannerID); err != nil {
			return nil, fmt.Errorf("postgreClickRepository.FindByIdInRange: row scan error: %w", err)
		}
		clicks = append(clicks, click)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("postgreClickRepository.FindByIdInRange: rows error: %w", err)
	}

	return clicks, nil

}
