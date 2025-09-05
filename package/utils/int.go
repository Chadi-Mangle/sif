package utils

import "database/sql"

func ConvertInt32ToNullInt32(number int32) sql.NullInt32 {
	return sql.NullInt32{Int32: number, Valid: true}
}
