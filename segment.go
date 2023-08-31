package avitotest

type Segment struct {
	Id   int    `json:"-" db:"id"`
	Slug string `json:"slug" binding:"required"`
}
