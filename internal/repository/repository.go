package repository

//internal/repository/repository.go
import "context"

type URLRepository interface {
	Create(ctx context.Context, urlToSave, alias string) error
	Get(ctx context.Context, alias string) (string, error)
	Delete(ctx context.Context, alias string) error
}
