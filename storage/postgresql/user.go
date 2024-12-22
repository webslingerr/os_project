package postgresql

import (
	"context"
	"fmt"

	"app/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {
	id := uuid.New().String()

	query := `
		INSERT INTO users (
			id,
			email,
			fullname,
			password,
			address,
			type,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
	`

	_, err := u.db.Exec(ctx, query,
		id,
		req.Email,
		req.Fullname,
		req.Password,
		req.Address,
		req.Type,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (u *userRepo) GetById(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {
	query := `
        SELECT 
            id,
            email,
            fullname,
            password, 
            address,
            type
        FROM users
        WHERE id = $1
    `

	row := u.db.QueryRow(ctx, query, req.ID)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Fullname,
		&user.Password,
		&user.Address,
		&user.Type,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return &user, nil
}

func (u *userRepo) GetByLoginAndPassword(ctx context.Context, req *models.Login) (*models.User, error) {
	query := `
        SELECT 
            id,
            email,
            fullname,
            password, 
            address,
            type
        FROM users
        WHERE password = $1 AND email = $2
    `

	row := u.db.QueryRow(ctx, query, req.Password, req.Email)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Fullname,
		&user.Password,
		&user.Address,
		&user.Type,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &models.User{}, nil
		}
		return nil, fmt.Errorf("failed to get user by email and password: %w", err)
	}

	return &user, nil
}

func (u *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {
	resp := models.GetListUserResponse{}

	query := `
		SELECT
			id,
			email,
			fullname,
			password, 
			address,
			type
		FROM users
		WHERE fullname ILIKE $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	searchPattern := fmt.Sprintf("%%%s%%", req.Search)
	rows, err := u.db.Query(ctx, query, searchPattern, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get user list: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Fullname,
			&user.Password,
			&user.Address,
			&user.Type,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		resp.Users = append(resp.Users, &user)
	}

	countQuery := `
		SELECT COUNT(*) FROM users WHERE fullname ILIKE $1
	`
	err = u.db.QueryRow(ctx, countQuery, searchPattern).Scan(&resp.Count)
	if err != nil {
		return nil, fmt.Errorf("failed to get user count: %w", err)
	}

	return &resp, nil
}

func (u *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
	query := `
		UPDATE users
		SET 
			email = $1,
			fullname = $2,
			password = $3,
			address = $4,
			type = $5,
			updated_at = NOW()
		WHERE id = $6
	`

	cmdTag, err := u.db.Exec(ctx, query,
		req.Email,
		req.Fullname,
		req.Password,
		req.Address,
		req.Type,
		req.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to update user: %w", err)
	}

	return cmdTag.RowsAffected(), nil
}

func (u *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) (int64, error) {
	query := `DELETE FROM users WHERE id = $1`

	cmdTag, err := u.db.Exec(ctx, query, req.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to delete user: %w", err)
	}

	return cmdTag.RowsAffected(), nil
}
