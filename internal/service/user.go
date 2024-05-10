package service

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fido-bio/domain"
	"fido-bio/dto"
	"github.com/google/uuid"
	"time"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (d userService) Register(ctx context.Context, req dto.RegisterUser) error {
	dvc, err := d.userRepository.FindByDeviceId(ctx, req.DeviceId)
	if err != nil {
		return err
	}
	if dvc.Id != "" {
		return errors.New("device already registered")
	}
	publicKey, _ := base64.StdEncoding.DecodeString(req.PublicKey)
	if l := len(publicKey); l != ed25519.PublicKeySize {
		return errors.New("invalid public-key")
	}

	device := domain.User{
		Id:        uuid.NewString(),
		DeviceId:  req.DeviceId,
		Name:      req.Name,
		PublicKey: req.PublicKey,
		CreatedAt: time.Now().Unix(),
	}
	err = d.userRepository.Save(ctx, &device)
	return err
}
