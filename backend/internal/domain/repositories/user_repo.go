package repositories

import (
	"context"

	"github.com/Jesuloba-world/flowcast/internal/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetById(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error

	// Business specific
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	ExistsByEmailOrUserName(ctx context.Context, username, email string) (bool, error)
	GetActiveUsers(ctx context.Context, limit, offset int) ([]*models.User, error)
	CountUsers(ctx context.Context) (int, error)

	// Batch operations
	CreateBatch(ctx context.Context, users []*models.User) error
	UpdateBatch(ctx context.Context, users []*models.User) error
}
