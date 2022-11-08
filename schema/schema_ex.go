package schema

import (
	"reflect"
	"strings"
)

const (
	GormEx_TagKey            = "gormEx"
	GormEx_TagValueTableName = "tableName"
	GormEx_FieldTableName    = "GormEx_TableName"
)

// support special table name which specified by reflect struct tag dynamically
// by ludanfeng@zj.tech
func GetTableNameEx(modelType reflect.Type) string {
	if modelType == nil {
		return ""
	}

	if field, ok := modelType.FieldByName(GormEx_FieldTableName); ok {
		tag := field.Tag.Get(GormEx_TagKey)
		values := strings.Split(tag, ";")
		for _, v := range values {
			parts := strings.SplitN(v, ":", 2)
			if len(parts) == 2 && parts[0] == GormEx_TagValueTableName {
				return parts[1]
			}
		}
	}
	return ""
}