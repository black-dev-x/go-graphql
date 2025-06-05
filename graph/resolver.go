package graph

import "github.com/black-dev-x/go-graphql/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB database.CategoryDB
	CourseDB   database.CourseDB
}
