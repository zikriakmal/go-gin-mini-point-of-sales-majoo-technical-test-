package repositories

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"majoo/models"
)

type UserRepository interface {
	FindById(userId uint64) (models.User, error)
	Login(userName string, password string) (models.User, error)
}

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{conn: db}
}

func (db *userRepository) FindById(userId uint64) (models.User, error) {
	var User models.User
	err := db.conn.First(&User, userId).Error
	return User, err
}

func (db *userRepository) Login(userName string, password string) (models.User, error) {
	var user models.User
	err := db.conn.Where("user_name = ?", userName).First(&user).Error
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err
}
