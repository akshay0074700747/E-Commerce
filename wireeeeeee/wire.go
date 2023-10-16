package wireeeeeee

import (
	"ecommerce/internal/adapters"
	routes "ecommerce/web/api/Routes"
	"ecommerce/web/database"

	"github.com/google/wire"
)

func InitializeAPI(envs map[string]string) (*routes.EchoEngine, error) {

	wire.Build(database.Connect_to,adapters.NewUserAdapter,)
	
	return &routes.EchoEngine{},nil

}
