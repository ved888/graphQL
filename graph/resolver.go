package graph

import (
	"github.com/99designs/gqlgen/plugin/federation/testdata/entityresolver/generated"
	"grapgQL/dbhelper"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DAO dbhelper.DAO
}

func (r *Resolver) Entity() generated.EntityResolver {
	//TODO implement me
	panic("implement me")
}
