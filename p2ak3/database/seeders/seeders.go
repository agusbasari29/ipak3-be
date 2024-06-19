package seeders

import (
	"github.com/agusbasari29/simru-be/p2ak3/database"
	"github.com/agusbasari29/simru-be/p2ak3/repository"

	"gorm.io/gorm"
)

var (
	db           *gorm.DB = database.SetupDBConnection()
	roleRepo              = repository.NewRoleRepository(db)
	companyRepo           = repository.NewCompanyRepository(db)
	// businessRepo          = repository.NewBusinessRepository(db)
	// contactRepo           = repository.NewContactRepository(db)
	userRepo              = repository.NewUserRepository(db)
)

func Seeders() {
	seedRoles()
	insertCompany()
	insertUsers()
	// seedCompanies()
}
