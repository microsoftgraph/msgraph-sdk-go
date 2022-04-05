package models
import (
    "strings"
    "errors"
)
// Provides operations to call the onenotePatchContent method.
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
    result := REPLACE_ONENOTEPATCHACTIONTYPE
    switch strings.ToUpper(v) {
        case "REPLACE":
            result = REPLACE_ONENOTEPATCHACTIONTYPE
        case "APPEND":
            result = APPEND_ONENOTEPATCHACTIONTYPE
        case "DELETE":
            result = DELETE_ONENOTEPATCHACTIONTYPE
        case "INSERT":
            result = INSERT_ONENOTEPATCHACTIONTYPE
        case "PREPEND":
            result = PREPEND_ONENOTEPATCHACTIONTYPE
        default:
            return 0, errors.New("Unknown OnenotePatchActionType value: " + v)
    }
    return &result, nil
}
func SerializeOnenotePatchActionType(values []OnenotePatchActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
