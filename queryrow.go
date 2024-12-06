package easy_sql

func (d DB) QueryRow(q string, s any) {
	keys := ParseQuerys(q)
}
