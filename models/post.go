package models

type RealEstateType string

const (
	Land      RealEstateType = "Land"
	House     RealEstateType = "House"
	Apartment RealEstateType = "Apartment"
)

type ServiceType string

const (
	Rent ServiceType = "Rent"
	Sell ServiceType = "Sell"
)

type Post struct {
	ID              string         `json:"id" db:"id"`
	RealEstateType  RealEstateType `json:"real_estate_type" db:"real_estate_type"`
	ServiceType     ServiceType    `json:"service_type" db:"service_type"`
	UserID          string         `json:"user_id" db:"user_id"`
	Title           string         `json:"title" db:"title"`
	Description     string         `json:"description" db:"description"`
	Region          string         `json:"region" db:"region"`
	Address         string         `json:"address" db:"address"`
	ContactDetails  string         `json:"contact_details" db:"contact_details"`
	Area            int            `json:"area" db:"area"`
	NumberOfRooms   int            `json:"number_of_rooms" db:"number_of_rooms"`
	FloorNumber     int            `json:"floor_number" db:"floor_number"`
	Price           float64        `json:"price" db:"price"`
	RentPrice       float64        `json:"rent_price" db:"rent_price"`
	SpecialBenefits []string       `json:"special_benefits" db:"special_benefits"`
	Images          []string       `json:"images" db:"images"`
	Status          string         `json:"status"`
	CreatedAt       string         `json:"created_at" db:"created_at"`
	UpdatedAt       string         `json:"updated_at" db:"updated_at"`
}

type PostPrimaryKey struct {
	ID string `json:"id" db:"id"`
}

type CreatePost struct {
	RealEstateType  RealEstateType `json:"real_estate_type" db:"real_estate_type" validate:"required,oneof=Land House Apartment"`
	ServiceType     ServiceType    `json:"service_type" db:"service_type" validate:"required,oneof=Rent Sell"`
	UserID          string         `json:"user_id" db:"user_id" validate:"required,uuid"`
	Title           string         `json:"title" db:"title" validate:"required"`
	Description     string         `json:"description" db:"description"`
	Region          string         `json:"region" db:"region" validate:"required"`
	Address         string         `json:"address" db:"address" validate:"required"`
	ContactDetails  string         `json:"contact_details" db:"contact_details" validate:"required"`
	Area            int            `json:"area" db:"area" validate:"required,min=0"`
	NumberOfRooms   int            `json:"number_of_rooms" db:"number_of_rooms" validate:"required,min=0"`
	FloorNumber     int            `json:"floor_number" db:"floor_number" validate:"min=0"`
	Price           float64        `json:"price" db:"price" validate:"required,min=0"`
	RentPrice       float64        `json:"rent_price" db:"rent_price" validate:"min=0"`
	SpecialBenefits []string       `json:"special_benefits" db:"special_benefits"`
	Images          []string       `json:"images" db:"images"`
}

type UpdatePost struct {
	ID              string         `json:"id" db:"id" validate:"required,uuid"`
	RealEstateType  RealEstateType `json:"real_estate_type" db:"real_estate_type" validate:"oneof=Land House Apartment"`
	ServiceType     ServiceType    `json:"service_type" db:"service_type" validate:"oneof=Rent Sell"`
	UserID          string         `json:"user_id" db:"user_id" validate:"uuid"`
	Title           string         `json:"title" db:"title"`
	Description     string         `json:"description" db:"description"`
	Region          string         `json:"region" db:"region"`
	Address         string         `json:"address" db:"address"`
	ContactDetails  string         `json:"contact_details" db:"contact_details"`
	Area            int            `json:"area" db:"area" validate:"min=0"`
	NumberOfRooms   int            `json:"number_of_rooms" db:"number_of_rooms" validate:"min=0"`
	FloorNumber     int            `json:"floor_number" db:"floor_number" validate:"min=0"`
	Price           float64        `json:"price" db:"price" validate:"min=0"`
	RentPrice       float64        `json:"rent_price" db:"rent_price" validate:"min=0"`
	SpecialBenefits []string       `json:"special_benefits" db:"special_benefits"`
	Images          []string       `json:"images" db:"images"`
	UpdatedAt       string         `json:"updated_at" db:"updated_at"`
}

type UpdateStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type GetListPostRequest struct {
	Offset         int            `json:"offset" db:"offset"`
	Limit          int            `json:"limit" db:"limit"`
	Search         string         `json:"search" db:"search"`
	RealEstateType RealEstateType `json:"real_estate_type" db:"real_estate_type"`
	Region         string         `json:"region" db:"region"`
	UserId         string         `json:"user_id"`
	Status         string         `json:"status"`
}

type GetListPostResponse struct {
	Count int     `json:"count"`
	Posts []*Post `json:"posts"`
}
