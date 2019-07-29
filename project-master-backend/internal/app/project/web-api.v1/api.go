package api

import (
	"gitlab.com/standard-go/project/internal/app/project"
)



var (
	// Default Variable
	srv 				project.Service

	// Static Variable
	StoragePath 	= "./internal/storage/"
)

// SetService sets the project's data management service
func SetService(s project.Service) {
	srv = s
}
