package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	avitotest "github.com/ulnesterova/avito_test"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) Create(segment avitotest.Segment) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO segments (slug) VALUES ('%s') RETURNING id", segment.Slug)
	err := r.db.QueryRow(query).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SegmentPostgres) Delete(slug string) error {
	query_userSegments := fmt.Sprintf("DELETE FROM %s WHERE segment_id = (SELECT id FROM %s WHERE slug = $1)", usersSegmantsTable, segmentsTable)
	_, err := r.db.Exec(query_userSegments, slug)
	if err != nil {
		return err
	}

	query_segments := fmt.Sprintf("DELETE FROM %s WHERE slug = $1", segmentsTable)
	_, err = r.db.Exec(query_segments, slug)

	return err
}

func (r *SegmentPostgres) GetAll(userId int) ([]string, error) {
	var usersSegments []string

	query := fmt.Sprintf("SELECT slug FROM %s seg WHERE seg.id in (SELECT segment_id FROM %s WHERE user_id=$1)", segmentsTable, usersSegmantsTable)
	err := r.db.Select(&usersSegments, query, userId)

	return usersSegments, err
}

func (r *SegmentPostgres) AddSegmentToUser(userId int, slug string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var slug_id int
	findSegIdQuery := fmt.Sprintf("SELECT id FROM segments WHERE slug ='%s'", slug)
	fmt.Print(slug_id)
	err = tx.QueryRow(findSegIdQuery).Scan(&slug_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	createUsersSegmentsQuery := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2)", usersSegmantsTable)
	_, err = tx.Exec(createUsersSegmentsQuery, userId, slug_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *SegmentPostgres) DeleteSegmentFromUser(userId int, slug string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var slug_id int
	findSegIdQuery := fmt.Sprintf("SELECT id FROM segments WHERE slug ='%s'", slug)
	fmt.Print(slug_id)
	err = tx.QueryRow(findSegIdQuery).Scan(&slug_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteUsersSegmentsQuery := fmt.Sprintf("DELETE FROM %s WHERE segment_id = $1 AND user_id = $2", usersSegmantsTable)
	_, err = tx.Exec(deleteUsersSegmentsQuery, userId, slug_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
