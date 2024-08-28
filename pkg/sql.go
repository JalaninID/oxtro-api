package pkg

import (
	"fmt"
	"strings"
)

func SQLXAndMatch(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " AND " + column + " = ?"
	args = append(args, val)
	return where, args
}

func SQLXAndLike(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " AND " + column + " LIKE ?"
	args = append(args, val)
	return where, args
}
func SQLXOrLike(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " OR " + column + " LIKE ?"
	args = append(args, val)
	return where, args
}
func SQLXAndLikeLower(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " AND LOWER(" + column + ") LIKE LOWER(?)"
	args = append(args, val)
	return where, args
}

func SQLXOrLikeLower(where string, args []interface{}, columns []string, val interface{}) (string, []interface{}) {
	inner := " AND ("
	or := []string{}
	for _, column := range columns {
		or = append(or, "LOWER("+column+") LIKE LOWER(?)")
		args = append(args, val)
	}
	inner += strings.Join(or, " OR ")
	inner += ")"
	where += inner
	return where, args
}

func SQLXAndIn(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " AND " + column + " IN (?)"
	args = append(args, val)
	return where, args
}

func SQLXAndNotIn(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " AND " + column + " NOT IN (?)"
	args = append(args, val)
	return where, args
}

func DebugQuery(query string, args ...interface{}) {
	// replce ? to args
	for _, arg := range args {
		query = strings.Replace(query, "?", fmt.Sprintf("%v", arg), 1)
	}
	colorGreen := "\033[32m"
	colorReset := "\033[0m"

	fmt.Printf("%s%s%s\n", colorGreen, query, colorReset)
}
