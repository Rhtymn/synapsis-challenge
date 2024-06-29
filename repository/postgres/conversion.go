package repository

import (
	"database/sql"
	"time"
)

func toTimePtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
