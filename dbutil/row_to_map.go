package dbutil

import (
	"database/sql"
	"strconv"
	"strings"
)

func RowToMap(rows *sql.Rows) []map[string]interface{} {
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	results := make([]map[string]interface{}, 0)
	index := 0
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		rowResult := make(map[string]interface{})
		for i, value := range values {

			switch value.(type) {
			case []byte:
				valueStr := string(value.([]byte))
				valueNum, err := strconv.ParseFloat(valueStr, 64)
				if err != nil {
					rowResult[columns[i]] = valueStr
				} else {
					if strings.Index(valueStr, ".") >= 0 {
						rowResult[columns[i]] = valueNum
					} else {
						rowResult[columns[i]] = int(valueNum)
					}
				}
			default:
				rowResult[columns[i]] = value
			}
		}
		results = append(results, rowResult)
		index++
	}
	return results
}
