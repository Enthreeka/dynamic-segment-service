package entity

type User struct {
	ID       string    `json:"id"`
	Segments []Segment `json:"segments"`
}
