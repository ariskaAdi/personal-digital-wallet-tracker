package repositories

import (
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.Users) (entity.Users, error)
    Update(ctx context.Context, user entity.Users) (entity.Users, error)
    Delete(ctx context.Context, id int) error
    FindAll(ctx context.Context) ([]entity.Users, error)
    FindById(ctx context.Context, id int) (entity.Users, error)
	FindByEmail(ctx context.Context, email string) (entity.Users, error)
}

type userRepository struct {
	db *sqlx.DB
}

// Composition Root
func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user entity.Users) (entity.Users, error) {
	SQL := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRowContext(
		ctx,
		SQL,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.ID)

	return user, err
}

func (r *userRepository) Update(ctx context.Context, user entity.Users) (entity.Users, error) {
	SQL := `
		UPDATE users
		SET name = $1, email = $2, password = $3
		WHERE id = $4
	` 

	res, err := r.db.ExecContext(
		ctx,
		SQL,
		user.Name,
		user.Email,
		user.Password,
		user.ID,
	)

	if err != nil {
		return entity.Users{}, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return entity.Users{}, err
	}

	if rows == 0 {
		return entity.Users{}, err
	}

	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	SQL := `
		DELETE FROM users
		WHERE id = $1
	`
	res, err := r.db.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.Users, error) {
	var users []entity.Users
	SQL := `
		SELECT id, name, email FROM users
	`
	err := r.db.SelectContext(ctx, &users, SQL)

	return users, err
}

func (r *userRepository) FindById(ctx context.Context, id int) (entity.Users, error) {
	var user entity.Users
	SQL := `
		SELECT id, name, email FROM users
		WHERE id = $1
	`

	err := r.db.GetContext(ctx, &user, SQL, id)

	return user, err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (entity.Users, error) {
	var user entity.Users
	SQL := `
		SELECT email, password FROM users
		WHERE email = $1
	`

	err := r.db.GetContext(ctx, &user, SQL, email)

	return  user, err
}