package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type StatisticsRepo struct {
	db *sql.DB
}

func NewStatisticsRepo(db *sql.DB) *StatisticsRepo {
	return &StatisticsRepo{db: db}
}

func (r *StatisticsRepo) Record(ctx context.Context, metric string, value float64) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO statistics (id, metric, value, recorded_at) VALUES (?, ?, ?, ?)`,
		uuidOrRandom(), metric, value, time.Now().UTC().Format(time.RFC3339))
	return err
}

func (r *StatisticsRepo) Query(ctx context.Context, metric string, limit int) ([]StatRecord, error) {
	if limit <= 0 {
		limit = 100
	}
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, metric, value, recorded_at FROM statistics WHERE metric = ? ORDER BY recorded_at DESC LIMIT ?`, metric, limit)
	if err != nil {
		return nil, fmt.Errorf("query statistics: %w", err)
	}
	defer rows.Close()

	var records []StatRecord
	for rows.Next() {
		var rec StatRecord
		var recordedAt string
		if err := rows.Scan(&rec.ID, &rec.Metric, &rec.Value, &recordedAt); err != nil {
			return nil, err
		}
		rec.RecordedAt, _ = time.Parse(time.RFC3339, recordedAt)
		records = append(records, rec)
	}
	return records, rows.Err()
}

type StatRecord struct {
	ID         string    `json:"id"`
	Metric     string    `json:"metric"`
	Value      float64   `json:"value"`
	RecordedAt time.Time `json:"recorded_at"`
}

func uuidOrRandom() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = byte(time.Now().UnixNano() & 0xff)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
