package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/config"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func (useCase *UseCase) GetAuthUser(ctx context.Context, ID int64) (entity.User, error) {

	user, err := useCase.user.GetUserByID(ctx, ID)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "useCase.user.GetUserByID() error - GetUserByID")
		return user, err
	}

	return user, nil
}

func (useCase *UseCase) Login(ctx context.Context, input LoginRequest) (entity.User, error) {

	user, err := useCase.user.GetUserByEmail(ctx, input.Email)
	if err != nil {
		logger.Trace(ctx, user, err, "useCase.user.getUserByEmail() error - Login")
		return user, err
	}

	if !CheckPasswordHash(input.Password, user.Password) {
		return entity.User{}, err
	}

	return user, nil
}

func (useCase *UseCase) Register(ctx context.Context, input RegisterRequest) (entity.User, error) {

	hashedPass, err := HashPassword(input.Password)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Name:     input.Name,
		Email:    input.Password,
		Password: hashedPass,
	}
	user, err = useCase.user.AddUser(ctx, user)
	if err != nil {
		logger.Trace(ctx, user, err, "useCase.cronjob.AddCronjob() error - AddCronjob")
		return user, err
	}

	return user, nil
}

func (useCase *UseCase) GenerateToken(ctx context.Context, user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": user.Email,
			"name":  user.Name,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(config.Get("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
