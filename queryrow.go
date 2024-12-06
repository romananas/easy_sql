package easy_sql

import "github.com/romananas/easy_sql/utils"

func (d DB) QueryRow(q string, s any) {
	keys := utils.ParseQuerys(q)
}
