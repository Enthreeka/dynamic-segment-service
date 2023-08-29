package entity

// swagger:parameters entity.User
type User struct {
	ID       string    `json:"id"`
	Segments []Segment `json:"segments"`
}
