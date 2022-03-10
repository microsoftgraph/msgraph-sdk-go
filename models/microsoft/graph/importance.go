package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type Importance int

const (
    LOW_IMPORTANCE Importance = iota
    NORMAL_IMPORTANCE
    HIGH_IMPORTANCE
)

func (i Importance) String() string {
    return []string{"LOW", "NORMAL", "HIGH"}[i]
}
func ParseImportance(v string) (interface{}, error) {
    result := LOW_IMPORTANCE
    switch strings.ToUpper(v) {
        case "LOW":
            result = LOW_IMPORTANCE
        case "NORMAL":
            result = NORMAL_IMPORTANCE
        case "HIGH":
            result = HIGH_IMPORTANCE
        default:
            return 0, errors.New("Unknown Importance value: " + v)
    }
    return &result, nil
}
func SerializeImportance(values []Importance) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
