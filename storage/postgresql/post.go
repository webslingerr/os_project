package postgresql

import (
	"app/models"
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postRepo struct {
	db *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) *postRepo {
	return &postRepo{
		db: db,
	}
}

func (p *postRepo) Create(ctx context.Context, req *models.CreatePost) (string, error) {
	id := uuid.New().String()

	query := `
		INSERT INTO posts (
			id,
			real_estate_type,
			service_type,
			user_id,
			title,
			description,
			region,
			address,
			contact_details,
			area,
			number_of_rooms,
			floor_number,
			price,
			rent_price,
			special_benefits,
			images,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, 
			$10, $11, $12, $13, $14, $15, $16, NOW(), NOW()
		)
	`

	_, err := p.db.Exec(ctx, query,
		id,
		req.RealEstateType,
		req.ServiceType,
		req.UserID,
		req.Title,
		req.Description,
		req.Region,
		req.Address,
		req.ContactDetails,
		req.Area,
		req.NumberOfRooms,
		req.FloorNumber,
		req.Price,
		req.RentPrice,
		req.SpecialBenefits,
		req.Images,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create post: %w", err)
	}

	return id, nil
}

func (p *postRepo) GetById(ctx context.Context, req *models.PostPrimaryKey) (*models.Post, error) {
	query := `
		SELECT
			id,
			real_estate_type,
			service_type,
			user_id,
			title,
			description,
			region,
			address,
			contact_details,
			area,
			number_of_rooms,
			floor_number,
			price,
			rent_price,
			status,
			special_benefits,
			images
		FROM posts
		WHERE id = $1
	`

	row := p.db.QueryRow(ctx, query, req.ID)

	var post models.Post
	var specialBenefits []string
	var images []string

	err := row.Scan(
		&post.ID,
		&post.RealEstateType,
		&post.ServiceType,
		&post.UserID,
		&post.Title,
		&post.Description,
		&post.Region,
		&post.Address,
		&post.ContactDetails,
		&post.Area,
		&post.NumberOfRooms,
		&post.FloorNumber,
		&post.Price,
		&post.RentPrice,
		&post.Status,
		&specialBenefits,
		&images,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get post by ID: %w", err)
	}

	post.SpecialBenefits = specialBenefits
	post.Images = images

	return &post, nil
}

func (p *postRepo) GetList(ctx context.Context, req *models.GetListPostRequest) (*models.GetListPostResponse, error) {
	resp := models.GetListPostResponse{}

	fmt.Println("Req->", req)

	// Base query without WHERE, ORDER BY, LIMIT, OFFSET
	baseQuery := `
        SELECT
            id,
            real_estate_type,
            service_type,
            user_id,
            title,
            description,
            region,
            address,
            contact_details,
            area,
            number_of_rooms,
            floor_number,
            price,
            rent_price,
            status,
            special_benefits,
            images
        FROM posts
    `

	var whereClauses []string
	var args []interface{}
	argPos := 1

	// Apply search filter if provided
	if req.Search != "" {
		whereClauses = append(whereClauses, `(title ILIKE $`+fmt.Sprintf("%d", argPos)+` OR description ILIKE $`+fmt.Sprintf("%d", argPos)+`)`)
		args = append(args, fmt.Sprintf("%%%s%%", req.Search))
		argPos++
	}

	// Apply RealEstateType filter if provided
	if req.RealEstateType != "" {
		whereClauses = append(whereClauses, `real_estate_type = $`+fmt.Sprintf("%d", argPos))
		args = append(args, req.RealEstateType)
		argPos++
	}

	// Apply Region filter if provided
	if req.Region != "" {
		whereClauses = append(whereClauses, `region = $`+fmt.Sprintf("%d", argPos))
		args = append(args, req.Region)
		argPos++
	}

	// Apply UserId filter if provided
	if req.UserId != "" {
		whereClauses = append(whereClauses, `user_id = $`+fmt.Sprintf("%d", argPos))
		args = append(args, req.UserId)
		argPos++
	}

	// Apply Status filter if provided
	if req.Status != "" {
		whereClauses = append(whereClauses, `status = $`+fmt.Sprintf("%d", argPos))
		args = append(args, req.Status)
		argPos++
	}

	// Combine WHERE clauses
	whereSQL := ""
	if len(whereClauses) > 0 {
		whereSQL = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Append ORDER BY, LIMIT, OFFSET
	orderLimitOffset := fmt.Sprintf("ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, req.Limit, req.Offset)
	argPos += 2

	// Final query combining base, where, and order/limit/offset clauses
	finalQuery := baseQuery + " " + whereSQL + " " + orderLimitOffset

	fmt.Println("QUERY->", finalQuery)

	// Execute the main query
	rows, err := p.db.Query(ctx, finalQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get post list: %w", err)
	}
	defer rows.Close()

	// Iterate over the result set and populate the response
	for rows.Next() {
		var post models.Post
		var specialBenefits []string
		var images []string

		err := rows.Scan(
			&post.ID,
			&post.RealEstateType,
			&post.ServiceType,
			&post.UserID,
			&post.Title,
			&post.Description,
			&post.Region,
			&post.Address,
			&post.ContactDetails,
			&post.Area,
			&post.NumberOfRooms,
			&post.FloorNumber,
			&post.Price,
			&post.RentPrice,
			&post.Status,
			&specialBenefits,
			&images,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}

		post.SpecialBenefits = specialBenefits
		post.Images = images

		resp.Posts = append(resp.Posts, &post)
	}

	countQuery := `
        SELECT COUNT(*)
        FROM posts
    `

	var count int
	err = p.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("failed to scan count: %w", err)
	}

	resp.Count = count

	return &resp, nil
}

func (p *postRepo) Update(ctx context.Context, req *models.UpdatePost) (int64, error) {
	query := `
		UPDATE posts
		SET 
			real_estate_type = $1,
			service_type = $2,
			user_id = $3,
			title = $4,
			description = $5,
			region = $6,
			address = $7,
			contact_details = $8,
			area = $9,
			number_of_rooms = $10,
			floor_number = $11,
			price = $12,
			rent_price = $13,
			special_benefits = $14,
			images = $15,
			updated_at = NOW()
		WHERE id = $16
	`

	cmdTag, err := p.db.Exec(ctx, query,
		req.RealEstateType,
		req.ServiceType,
		req.UserID,
		req.Title,
		req.Description,
		req.Region,
		req.Address,
		req.ContactDetails,
		req.Area,
		req.NumberOfRooms,
		req.FloorNumber,
		req.Price,
		req.RentPrice,
		req.SpecialBenefits,
		req.Images,
		req.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to update post: %w", err)
	}

	return cmdTag.RowsAffected(), nil
}

func (p *postRepo) UpdateStatus(ctx context.Context, req *models.UpdateStatus) error {
	query := `
    UPDATE posts
    SET 
        status = $2,
        updated_at = NOW()
    WHERE id = $1
`
	_, err := p.db.Exec(ctx, query, req.ID, req.Status)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}
	return nil
}

func (p *postRepo) Delete(ctx context.Context, req *models.PostPrimaryKey) (int64, error) {
	query := `DELETE FROM posts WHERE id = $1`

	cmdTag, err := p.db.Exec(ctx, query, req.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to delete post: %w", err)
	}

	return cmdTag.RowsAffected(), nil
}
