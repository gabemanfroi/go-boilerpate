package IoC

import (
	"github.com/gabemanfroi/go-boilerplate/infra/db"
	"github.com/gabemanfroi/go-boilerplate/internal/utils"
	"github.com/golobby/container/v3"

	"gorm.io/gorm"
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
