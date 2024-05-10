package repository

import (
	"context"
	"database/sql"
	"fido-bio/domain"
	"github.com/doug-martin/goqu/v9"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", con),
	}
}

func (d userRepository) Save(ctx context.Context, user *domain.User) error {
	dataset := d.db.Insert("users").Rows(user).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

func (d userRepository) FindById(ctx context.Context, id string) (device domain.User, err error) {
	dataset := d.db.From("users").Where(goqu.I("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &device)
	return
}

func (d userRepository) FindByDeviceId(ctx context.Context, id string) (device domain.User, err error) {
	dataset := d.db.From("users").Where(goqu.I("device_id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &device)
	return
}
