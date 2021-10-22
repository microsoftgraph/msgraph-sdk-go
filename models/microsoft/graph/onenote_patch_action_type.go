package graph
import (
    "strings"
    "errors"
)
type OnenotePatchActionType int

const (
    REPLACE_ONENOTEPATCHACTIONTYPE OnenotePatchActionType = iota
    APPEND_ONENOTEPATCHACTIONTYPE
    DELETE_ONENOTEPATCHACTIONTYPE
    INSERT_ONENOTEPATCHACTIONTYPE
    PREPEND_ONENOTEPATCHACTIONTYPE
)

func (i OnenotePatchActionType) String() string {
    return []string{"REPLACE", "APPEND", "DELETE", "INSERT", "PREPEND"}[i]
}
func ParseOnenotePatchActionType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "REPLACE":
            return REPLACE_ONENOTEPATCHACTIONTYPE, nil
        case "APPEND":
            return APPEND_ONENOTEPATCHACTIONTYPE, nil
        case "DELETE":
            return DELETE_ONENOTEPATCHACTIONTYPE, nil
        case "INSERT":
            return INSERT_ONENOTEPATCHACTIONTYPE, nil
        case "PREPEND":
            return PREPEND_ONENOTEPATCHACTIONTYPE, nil
    }
    return 0, errors.New("Unknown OnenotePatchActionType value: " + v)
}
func SerializeOnenotePatchActionType(values []OnenotePatchActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
