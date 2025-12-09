package service

import (
	//"glebosyatina/test_project/internal/service/sub"
	"glebosyatina/test_project/internal/service/user"
)

// сервисный слой содержит бизнес логику, а уже каждый из сервисов зависит внутри от репозиторного слоя
type Services struct {
	UserService *user.UserService
	//SubService  *sub.SubService
}
