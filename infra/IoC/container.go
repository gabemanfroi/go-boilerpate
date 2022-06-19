package IoC

import (
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
	"showcaseme/infra/db"
	"showcaseme/internal/utils"
)

func InitContainer() {

}

func bindCore() {
	err := container.Singleton(func() *gorm.DB { return db.CreateSqlInstance() })
	utils.Check(err, "error while creating container bindings [database]")
}

func bindRepositories() {
	// register your repositories interfaces binders here
}

func bindServices() {
	// register your services interfaces binders here
}

func bindControllers() {
	// register your controllers interfaces binder here
}
