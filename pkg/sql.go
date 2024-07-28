package pkg

import "strings"

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
