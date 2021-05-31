package internal

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	UserWriter
	UserReader
}

type UserWriter interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
}

type UserReader interface {
	ReadUser(ctx context.Context, id int) (*User, error)
	ReadUserByEmail(ctx context.Context, email string) (*User, error)
}

type User struct {
	ID       int    `gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email" gorm:"unique;not null;size:255"`
	Password string `json:"password" validate:"required" gorm:"not null;size:255"`

	Picture string `json:"picture" gorm:"size:255"`
	Locale  string `json:"locale" gorm:"size:31"`

	Key string `json:"-" gorm:"size:31"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewUser(email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost+1)
	if err != nil {
		return nil, err
	}

	key, err := generateRandomKey()
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:    email,
		Password: string(hashedPassword),
		Key:      key,
	}

	v := validator.New()
	return user, v.Struct(user)
}

func (u *User) HasPassword(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		log.Println(err)
	}
	return err == nil
}

func generateRandomKey() (string, error) {
	const length = 31

	buff := make([]byte, length)
	if _, err := rand.Read(buff); err != nil {
		return "", err
	}
	b64encoded := base64.StdEncoding.EncodeToString(buff)
	return b64encoded[:length], nil
}
