package repository

import (
	"context"
	"database/sql"
	"fido-bio/domain"
	"github.com/doug-martin/goqu/v9"
)

type challengeRepository struct {
	db *goqu.Database
}

func NewChallenge(con *sql.DB) domain.ChallengeRepository {
	return &challengeRepository{
		db: goqu.New("default", con),
	}
}

func (c challengeRepository) Save(ctx context.Context, challenge *domain.Challenge) error {
	executor := c.db.Insert("challenges").Rows(challenge).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (c challengeRepository) Update(ctx context.Context, challenge *domain.Challenge) error {
	executor := c.db.Update("challenges").Set(challenge).Where(goqu.C("id").Eq(challenge.Id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (c challengeRepository) FindById(ctx context.Context, id string) (challenge domain.Challenge, err error) {
	dataset := c.db.From("challenges").Where(goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &challenge)
	return
}
