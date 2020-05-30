package main

import (
	"github.com/kyohei0423/oidc-sample/backend/app/infra"
	// "github.com/kyohei0423/oidc-sample/backend/app/infra/db"
)

func main() {
	// db := db.NewDBconnection()
	// defer db.Close()
	infra.Router()
}
