package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"cargo-rest-api/pkg/security"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

// UserRepo is a struct to store db connection.
type UserRepo struct {
	db *gorm.DB
}

// NewUserRepository will initialize user repository.
func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

// UserRepo implements the repository.UserRepository interface.
var _ repository.UserRepository = &UserRepo{}

// SaveUser will create a new user.
func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			errDesc["email"] = exception.ErrorTextUserEmailAlreadyTaken.Error()
			return nil, errDesc, exception.ErrorTextUnprocessableEntity
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return user, nil, nil
}

// UpdateUser will create a new user.
func (r *UserRepo) UpdateUser(uuid string, user *entity.User) (*entity.User, map[string]string, error) {
	errDesc := map[string]string{}
	userData := &entity.User{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	if user.Password != "" {
		userData.Password = user.Password
	}

	err := r.db.First(&user, "uuid = ?", uuid).Updates(userData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextUserInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextUserNotFound
		}
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			errDesc["email"] = exception.ErrorTextUserEmailAlreadyTaken.Error()
			return nil, errDesc, exception.ErrorTextUnprocessableEntity
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return user, nil, nil
}

// DeleteUser will return user detail.
func (r *UserRepo) DeleteUser(uuid string) error {
	var user entity.User
	err := r.db.Where("uuid = ?", uuid).Take(&user).Delete(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextUserNotFound
		}
		return err
	}
	return nil
}

// GetUser will return user detail.
func (r *UserRepo) GetUser(uuid string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("uuid = ?", uuid).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// GetUserRoles will return user roles.
func (r *UserRepo) GetUserRoles(uuid string) ([]entity.UserRole, error) {
	var roles []entity.UserRole
	err := r.db.Preload("Role").Where("user_uuid = ?", uuid).Find(&roles).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextUserNotFound
		}
		return nil, err
	}
	return roles, nil
}

// GetUserWithRoles will return user detail with roles.
func (r *UserRepo) GetUserWithRoles(uuid string) (*entity.User, error) {
	var user entity.User
	err := r.db.Preload("UserRoles.Role").Where("uuid = ?", uuid).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// GetUsers will return user list.
func (r *UserRepo) GetUsers(p *repository.Parameters) ([]*entity.User, *repository.Meta, error) {
	var total int64
	var users []*entity.User
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&users).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Order(p.Order).Limit(p.Limit).Offset(p.Offset).Find(&users).Error
	if errTotal != nil {
		return nil, nil, errTotal
	}
	if errList != nil {
		return nil, nil, errList
	}
	meta := repository.NewMeta(p, total)
	return users, meta, nil
}

// GetUserByEmail will find user by email.
func (r *UserRepo) GetUserByEmail(u *entity.User) (*entity.User, map[string]string, error) {
	var user entity.User
	errDesc := map[string]string{}
	err := r.db.Where("email = ?", u.Email).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["email"] = exception.ErrorTextUserEmailNotRegistered.Error()
			return nil, errDesc, exception.ErrorTextUserEmailNotRegistered
		}
		return nil, errDesc, err
	}

	return &user, nil, nil
}

func (r *UserRepo) GetUserByPhone(u *entity.User) (*entity.User, map[string]string, error) {
	var user entity.User
	errDesc := map[string]string{}
	err := r.db.Where("phone = ?", u.Phone).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["phone"] = exception.ErrorTextUserPhoneNotRegistered.Error()
			return nil, errDesc, exception.ErrorTextUserPhoneNotRegistered
		}
		return nil, errDesc, err
	}

	return &user, nil, nil
}

// GetUserByEmailAndPassword will find user by email and password.
func (r *UserRepo) GetUserByEmailAndPassword(u *entity.User) (*entity.User, map[string]string, error) {
	var user entity.User
	errDesc := map[string]string{}
	err := r.db.Where("email = ?", u.Email).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["email"] = exception.ErrorTextUserEmailNotRegistered.Error()
			return nil, errDesc, exception.ErrorTextUserEmailNotRegistered
		}
		return nil, errDesc, err
	}

	err = security.VerifyPassword(user.Password, u.Password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			errDesc["password"] = exception.ErrorTextUserInvalidPassword.Error()
			return nil, errDesc, exception.ErrorTextUserInvalidUsernameAndPassword
		}
	}
	return &user, nil, nil
}

// UpdateUserAvatar will create a new user.
func (r *UserRepo) UpdateUserAvatar(uuid string, user *entity.User) (*entity.User, map[string]string, error) {
	errDesc := map[string]string{}
	userData := &entity.User{
		AvatarUUID: user.AvatarUUID,
	}
	err := r.db.First(&user, "uuid = ?", uuid).Updates(userData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextUserInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextUserNotFound
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return user, nil, nil
}
