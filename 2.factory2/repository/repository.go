package repository

import (
	"fmt"

	"github.com/gsaldanab/design-pattern/2.factory2/config"
	"github.com/gsaldanab/design-pattern/2.factory2/repository/mysql"
	"github.com/gsaldanab/design-pattern/2.factory2/repository/sqlserver"
)

type Repository interface {
	Find(id int) string
	Save(data string) error
}

func New(config *config.Configuration) (Repository, error) {
	var repo Repository
	var err error

	switch config.Engine {
	case "mysql":
		repo = mysql.New()
	case "sqlserver":
		repo = sqlserver.New()
	default:
		err = fmt.Errorf("invalid engine %s", config.Engine)
	}
	return repo, err
}
