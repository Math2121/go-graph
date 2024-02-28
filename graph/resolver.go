package graph

import "github.com/Math2121/go-graph/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CategoryDb *database.Category
	CoursesDb *database.Courses
}
