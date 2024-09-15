package article_service

import "errors"

var (
	ErrCreation = errors.New("article creation failed")
	ErrUpdating = errors.New("article updating failed")
	ErrDeletion = errors.New("article deletion failed")
	ErrNotFound = errors.New("article not found")
	ErrSearch   = errors.New("search operation failed")
)
