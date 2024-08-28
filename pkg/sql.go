package pkg

import (
	"fmt"
	"strings"
)

// SQLXAndMatch is a function to add a match condition to the query,
// where is the current where clause,
// args is the current args,
// column is the column to match,
// val is the value to match
func SQLXAndMatch(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " AND " + column + " = ?"
	args = append(args, val)
	return where, args
}

// SQLXAndLike is a function to add a like condition to the query,
// where is the current where clause,
// args is the current args,
// column is the column to like,
// val is the value to like
func SQLXAndLike(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " AND " + column + " LIKE ?"
	args = append(args, val)
	return where, args
}

// SQLXOrLike is a function to add a like condition to the query,
// where is the current where clause,
// args is the current args,
// column is the column to like,
// val is the value to like
func SQLXOrLike(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " OR " + column + " LIKE ?"
	args = append(args, val)
	return where, args
}

// SQLXAndLikeLower is a function to add a like condition to the query,
// where is the current where clause,
// args is the current args,
// column is the column to like,
// val is the value to like
func SQLXAndLikeLower(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " AND LOWER(" + column + ") LIKE LOWER(?)"
	args = append(args, val)
	return where, args
}

// SQLXOrLikeLower is a function to add a like condition to the query,
// where is the current where clause,
// args is the current args,
// columns is the columns to like,
// val is the value to like
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

// SQLXAndIn is a function to add a in condition to the query,
// where is the current where clause,
// args is the current args,
// column is the column to in,
// val is the value to in
func SQLXAndIn(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " AND " + column + " IN (?)"
	args = append(args, val)
	return where, args
}

// SQLXAndNotIn is a function to add a not in condition to the query,
// where is the current where clause,
// args is the current args,
// column is the column to not in,
// val is the value to not in
func SQLXAndNotIn(where string, args []interface{}, column string, val interface{}) (string, []interface{}) {
	where += " AND " + column + " NOT IN (?)"
	args = append(args, val)
	return where, args
}

// DebugQuery is a function to debug the query,
// query is the query to debug
// args is the args to debug
func DebugQuery(query string, args ...interface{}) {
	// replce ? to args
	for _, arg := range args {
		query = strings.Replace(query, "?", fmt.Sprintf("%v", arg), 1)
	}
	colorGreen := "\033[32m"
	colorReset := "\033[0m"

	fmt.Printf("%s%s%s\n", colorGreen, query, colorReset)
}
