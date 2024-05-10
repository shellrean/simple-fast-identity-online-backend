package service

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fido-bio/domain"
	"fido-bio/dto"
	"fido-bio/internal/util"
	"github.com/google/uuid"
	"time"
)

type challengeService struct {
	challengeRepository domain.ChallengeRepository
	userRepository      domain.UserRepository
}

func NewChallenge(challengeRepository domain.ChallengeRepository,
	userRepository domain.UserRepository) domain.ChallengeService {
	return &challengeService{
		challengeRepository: challengeRepository,
		userRepository:      userRepository,
	}
}

func (c challengeService) Generate(ctx context.Context) (dto.ChallengeData, error) {
	challenge := domain.Challenge{
		Id:        uuid.NewString(),
		Key:       util.RandomString(10),
		ExpiredAt: time.Now().Add(5 * time.Minute).Unix(),
	}
	err := c.challengeRepository.Save(ctx, &challenge)
	if err != nil {
		return dto.ChallengeData{}, err
	}

	return dto.ChallengeData{
		Id:  challenge.Id,
		Key: challenge.Key,
	}, nil
}

func (c challengeService) Validate(ctx context.Context, req dto.ChallengeValidate) (dto.UserData, error) {
	challenge, err := c.challengeRepository.FindById(ctx, req.Id)
	if err != nil {
		return dto.UserData{}, err
	}
	if challenge.Id == "" {
		return dto.UserData{}, errors.New("challenge not found")
	}
	if challenge.ExpiredAt < time.Now().Unix() {
		return dto.UserData{}, errors.New("challenge already expired")
	}
	if challenge.ValidatedAt > 0 {
		return dto.UserData{}, errors.New("challenge already validated")
	}

	user, err := c.userRepository.FindByDeviceId(ctx, req.DeviceId)
	if err != nil {
		return dto.UserData{}, err
	}
	if user.Id == "" {
		return dto.UserData{}, errors.New("device not found")
	}

	publicKeyBase64, _ := base64.StdEncoding.DecodeString(user.PublicKey)
	signBase64, _ := base64.StdEncoding.DecodeString(req.Sign)

	publicKey := ed25519.PublicKey(publicKeyBase64)
	if ed25519.Verify(publicKey, []byte(challenge.Key), signBase64) {

		challenge.ValidatedAt = time.Now().Unix()
		_ = c.challengeRepository.Update(ctx, &challenge)

		return dto.UserData{
			Id:   user.Id,
			Name: user.Name,
		}, nil
	}
	return dto.UserData{}, errors.New("challenge validation failed")
}
