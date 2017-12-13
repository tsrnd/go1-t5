package data
import (
	"time"
)

type Model struct {
	ID				uint			`schema:"id"`
	CreatedAt		*time.Time 		`schema:"created_at"`
	UpdatedAt 		*time.Time 		`schema:"updated_at"`
	DeletedAt 		*time.Time 		`schema:"deleted_at"`
}