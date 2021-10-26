package graph
import (
    "strings"
    "errors"
)
// 
type FilterMode int

const (
    INCLUDE_FILTERMODE FilterMode = iota
    EXCLUDE_FILTERMODE
)

func (i FilterMode) String() string {
    return []string{"INCLUDE", "EXCLUDE"}[i]
}
func ParseFilterMode(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "INCLUDE":
            return INCLUDE_FILTERMODE, nil
        case "EXCLUDE":
            return EXCLUDE_FILTERMODE, nil
    }
    return 0, errors.New("Unknown FilterMode value: " + v)
}
func SerializeFilterMode(values []FilterMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
