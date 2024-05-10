package domain

import (
	"context"
	"fido-bio/dto"
)

type Challenge struct {
	Id          string `db:"id"`
	Key         string `db:"key"`
	ExpiredAt   int64  `db:"expired_at"`
	ValidatedAt int64  `db:"validated_at"`
}

type ChallengeRepository interface {
	Save(ctx context.Context, challenge *Challenge) error
	Update(ctx context.Context, challenge *Challenge) error
	FindById(ctx context.Context, id string) (Challenge, error)
}

type ChallengeService interface {
	Generate(ctx context.Context) (dto.ChallengeData, error)
	Validate(ctx context.Context, req dto.ChallengeValidate) (dto.UserData, error)
}
