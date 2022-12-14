package handlers

import (
	"github.com/matheusrbarbosa/squilo/application/services"
	v "github.com/matheusrbarbosa/squilo/application/validators"
	"github.com/matheusrbarbosa/squilo/domain/dtos"
	"github.com/matheusrbarbosa/squilo/domain/exceptions"
	i "github.com/matheusrbarbosa/squilo/domain/interfaces"
	"github.com/matheusrbarbosa/squilo/infra/database/repositories"
	"gorm.io/gorm"
)

type userHanlder struct {
	userService    i.UserService
	userRepository i.UserRepository
}

func UserHandler() *userHanlder {
	return &userHanlder{
		userService:    services.UserService(),
		userRepository: repositories.UserRepository(),
	}
}

func (h *userHanlder) HandleSignup(request v.SignupRequest) (dtos.UserDto, error) {
	user := request.ParseToUser()

	_, err := h.userRepository.GetByEmail(user.Email)
	if err != nil && err == gorm.ErrRecordNotFound {
		h.userService.PrepareToCreate(&user)
		user = h.userRepository.Create(user)
		return user.ParseDto(), nil
	} else {
		return user.ParseDto(), exceptions.EMAIL_ALREADY_EXIST
	}
}
