package storage

import (
	"context"
	"errors"

	"github.com/wormi4ok/menuplanner/internal"
	"gorm.io/gorm"
)

func (s *DB) CreateUser(ctx context.Context, user *internal.User) error {
	return s.db.WithContext(ctx).Create(user).Error
}

func (s *DB) UpdateUser(ctx context.Context, user *internal.User) error {
	if err := s.db.WithContext(ctx).First(user).Error; err != nil {
		return err
	}

	return s.db.WithContext(ctx).Save(user).Error
}

func (s *DB) ReadUser(ctx context.Context, id int) (*internal.User, error) {
	var user internal.User
	return &user, s.db.WithContext(ctx).First(&user, id).Error
}

func (s *DB) ReadUserByEmail(ctx context.Context, email string) (*internal.User, error) {
	var user internal.User
	err := s.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, internal.NewError(err, internal.ErrorNotFound)
	}
	return &user, err
}
