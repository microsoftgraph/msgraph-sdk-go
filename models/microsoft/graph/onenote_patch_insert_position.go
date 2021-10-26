package graph
import (
    "strings"
    "errors"
)
// 
type OnenotePatchInsertPosition int

const (
    AFTER_ONENOTEPATCHINSERTPOSITION OnenotePatchInsertPosition = iota
    BEFORE_ONENOTEPATCHINSERTPOSITION
)

func (i OnenotePatchInsertPosition) String() string {
    return []string{"AFTER", "BEFORE"}[i]
}
func ParseOnenotePatchInsertPosition(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "AFTER":
            return AFTER_ONENOTEPATCHINSERTPOSITION, nil
        case "BEFORE":
            return BEFORE_ONENOTEPATCHINSERTPOSITION, nil
    }
    return 0, errors.New("Unknown OnenotePatchInsertPosition value: " + v)
}
func SerializeOnenotePatchInsertPosition(values []OnenotePatchInsertPosition) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
