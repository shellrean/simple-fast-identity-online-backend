package domain

import (
	"context"
	"fido-bio/dto"
)

type User struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	DeviceId  string `db:"device_id"`
	PublicKey string `db:"public_key"`
	CreatedAt int64  `db:"created_at"`
}

type UserRepository interface {
	Save(ctx context.Context, device *User) error
	FindById(ctx context.Context, id string) (User, error)
	FindByDeviceId(ctx context.Context, id string) (User, error)
}

type UserService interface {
	Register(ctx context.Context, req dto.RegisterUser) error
}
