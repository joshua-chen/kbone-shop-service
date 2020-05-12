package controllers


type CategoryController struct {
	// Our MovieService, it's an interface which
	// is binded from the main application.
	Service services.MovieService
}
