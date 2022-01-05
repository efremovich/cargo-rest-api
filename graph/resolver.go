package graph

import (
	"cargo-rest-api/infrastructure/persistence"
)

type Resolver struct {
	DBServices *persistence.Repositories
}
