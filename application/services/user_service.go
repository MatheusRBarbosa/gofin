package services

import (
	vtypes "github.com/matheusrbarbosa/squilo/domain/enums/vault_types"
	"github.com/matheusrbarbosa/squilo/domain/interfaces"
	"github.com/matheusrbarbosa/squilo/domain/models"
)

type userService struct{}

func UserService() interfaces.UserService {
	return &userService{}
}

func (s *userService) PrepareToCreate(user *models.User) {
	defaultVault := models.Vault{
		Name:        "Geral",
		Description: "Cofrinho inicial",
		Configs:     "{}",
		TypeId:      vtypes.General,
	}

	user.Vaults = append(user.Vaults, defaultVault)
}
