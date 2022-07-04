package schema

import (
	"reflect"
	"testing"
)

func TestGetTableNameEx(t *testing.T) {
	type args struct {
		modelType reflect.Type
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "invalid-nil",
		},
		{
			name: "invalid-struct",
			args: args{
				modelType: reflect.TypeOf(struct{}{}),
			},
		},
		{
			name: "invalid-tag",
			args: args{
				modelType: reflect.StructOf([]reflect.StructField{
					{
						Name: "GormEx_TableName",
						Type: reflect.TypeOf(""),
						Tag:  reflect.StructTag(`gormEx:"" gorm:"-" json:"-"`),
					},
				}),
			},
		},
		{
			name: "invalid-tag",
			args: args{
				modelType: reflect.StructOf([]reflect.StructField{
					{
						Name: "GormEx_TableName",
						Type: reflect.TypeOf(""),
						Tag:  reflect.StructTag(`gormEx:"a:c;1" gorm:"-" json:"-"`),
					},
				}),
			},
		},
		{
			name: "valid-1",
			args: args{
				modelType: reflect.StructOf([]reflect.StructField{
					{
						Name: "GormEx_TableName",
						Type: reflect.TypeOf(""),
						Tag:  reflect.StructTag(`gormEx:"a:c;1;tableName:user2" gorm:"-" json:"-"`),
					},
				}),
			},
			want: "user2",
		},
		{
			name: "valid-2",
			args: args{
				modelType: reflect.StructOf([]reflect.StructField{
					{
						Name: "GormEx_TableName",
						Type: reflect.TypeOf(""),
						Tag:  reflect.StructTag(`gormEx:"tableName:user" gorm:"-" json:"-"`),
					},
					{
						Name: "ID",
						Type: reflect.TypeOf(0),
						Tag:  reflect.StructTag(`gorm:"primarykey" json:"id"`),
					},
					{
						Name: "Name",
						Type: reflect.TypeOf(""),
						Tag:  reflect.StructTag(`json:"name"`),
					},
				}),
			},
			want: "user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTableNameEx(tt.args.modelType); got != tt.want {
				t.Errorf("GetTableNameEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
