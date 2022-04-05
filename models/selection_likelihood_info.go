package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type SelectionLikelihoodInfo int

const (
    NOTSPECIFIED_SELECTIONLIKELIHOODINFO SelectionLikelihoodInfo = iota
    HIGH_SELECTIONLIKELIHOODINFO
)

func (i SelectionLikelihoodInfo) String() string {
    return []string{"NOTSPECIFIED", "HIGH"}[i]
}
func ParseSelectionLikelihoodInfo(v string) (interface{}, error) {
    result := NOTSPECIFIED_SELECTIONLIKELIHOODINFO
    switch strings.ToUpper(v) {
        case "NOTSPECIFIED":
            result = NOTSPECIFIED_SELECTIONLIKELIHOODINFO
        case "HIGH":
            result = HIGH_SELECTIONLIKELIHOODINFO
        default:
            return 0, errors.New("Unknown SelectionLikelihoodInfo value: " + v)
    }
    return &result, nil
}
func SerializeSelectionLikelihoodInfo(values []SelectionLikelihoodInfo) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
