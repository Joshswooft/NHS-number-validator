package utils

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

// String returns gender as a string
func (g Gender) String() string {
	return string(g)
}
