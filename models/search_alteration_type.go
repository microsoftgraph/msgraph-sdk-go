package models
import (
    "strings"
    "errors"
)
// Provides operations to call the query method.
type SearchAlterationType int

const (
    SUGGESTION_SEARCHALTERATIONTYPE SearchAlterationType = iota
    MODIFICATION_SEARCHALTERATIONTYPE
    UNKNOWNFUTUREVALUE_SEARCHALTERATIONTYPE
)

func (i SearchAlterationType) String() string {
    return []string{"SUGGESTION", "MODIFICATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSearchAlterationType(v string) (interface{}, error) {
    result := SUGGESTION_SEARCHALTERATIONTYPE
    switch strings.ToUpper(v) {
        case "SUGGESTION":
            result = SUGGESTION_SEARCHALTERATIONTYPE
        case "MODIFICATION":
            result = MODIFICATION_SEARCHALTERATIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SEARCHALTERATIONTYPE
        default:
            return 0, errors.New("Unknown SearchAlterationType value: " + v)
    }
    return &result, nil
}
func SerializeSearchAlterationType(values []SearchAlterationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
