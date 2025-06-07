package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"

	"github.com/Jesuloba-world/flowcast/internal/domain/models"
	"github.com/Jesuloba-world/flowcast/internal/domain/repositories"
)

type userRepository struct {
	db *bun.DB
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
)

func NewUserRepository(db *bun.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *userRepository) GetById(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	err := r.db.NewSelect().Model(user).Where("username = ?", username).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	_, err := r.db.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(ctx)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewDelete().Model((*models.User)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return r.db.NewSelect().Model((*models.User)(nil)).Where("email = ?", email).Exists(ctx)
}

func (r *userRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	return r.db.NewSelect().Model((*models.User)(nil)).Where("username = ?", username).Exists(ctx)
}

func (r *userRepository) ExistsByEmailOrUserName(ctx context.Context, email, username string) (bool, error) {
	return r.db.NewSelect().Model((*models.User)(nil)).Where("email = ? OR username == ?", email, username).Exists(ctx)
}

func (r *userRepository) GetActiveUsers(ctx context.Context, limit, offset int) ([]*models.User, error) {
	var users []*models.User
	err := r.db.NewSelect().Model(&users).
		Where("is_active = ?", true).
		Limit(limit).
		Offset(offset).
		Scan(ctx)
	return users, err
}

func (r *userRepository) CountUsers(ctx context.Context) (int, error) {
	return r.db.NewSelect().Model((*models.User)(nil)).Count(ctx)
}

func (r *userRepository) CreateBatch(ctx context.Context, users []*models.User) error {
	_, err := r.db.NewInsert().Model(&users).Exec(ctx)
	return err
}
func (r *userRepository) UpdateBatch(ctx context.Context, users []*models.User) error {
	_, err := r.db.NewUpdate().Model(&users).Bulk().Exec(ctx)
	return err
}
