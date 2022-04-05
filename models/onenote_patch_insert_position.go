package models
import (
    "strings"
    "errors"
)
// Provides operations to call the onenotePatchContent method.
type OnenotePatchInsertPosition int

const (
    AFTER_ONENOTEPATCHINSERTPOSITION OnenotePatchInsertPosition = iota
    BEFORE_ONENOTEPATCHINSERTPOSITION
)

func (i OnenotePatchInsertPosition) String() string {
    return []string{"AFTER", "BEFORE"}[i]
}
func ParseOnenotePatchInsertPosition(v string) (interface{}, error) {
    result := AFTER_ONENOTEPATCHINSERTPOSITION
    switch strings.ToUpper(v) {
        case "AFTER":
            result = AFTER_ONENOTEPATCHINSERTPOSITION
        case "BEFORE":
            result = BEFORE_ONENOTEPATCHINSERTPOSITION
        default:
            return 0, errors.New("Unknown OnenotePatchInsertPosition value: " + v)
    }
    return &result, nil
}
func SerializeOnenotePatchInsertPosition(values []OnenotePatchInsertPosition) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
