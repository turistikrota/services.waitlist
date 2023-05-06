package entity

type fields struct {
	UUID       string
	Email      string
	LeaveToken string
	IsActive   string
	CreatedAt  string
	UpdatedAt  string
}

var Fields = fields{
	UUID:       "_id",
	Email:      "email",
	LeaveToken: "leave_token",
	IsActive:   "is_active",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}
