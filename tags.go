package easy_sql

import "reflect"

func ReadTags(v any) []string {
	var tags []string
	tv := reflect.TypeOf(v)
	for i := range tv.NumField() {
		field := tv.Field(i)
		tags = append(tags, field.Tag.Get("sql"))
	}
	return tags
}
