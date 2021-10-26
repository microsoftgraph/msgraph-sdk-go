package graph
import (
    "strings"
    "errors"
)
// 
type SelectionLikelihoodInfo int

const (
    NOTSPECIFIED_SELECTIONLIKELIHOODINFO SelectionLikelihoodInfo = iota
    HIGH_SELECTIONLIKELIHOODINFO
)

func (i SelectionLikelihoodInfo) String() string {
    return []string{"NOTSPECIFIED", "HIGH"}[i]
}
func ParseSelectionLikelihoodInfo(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTSPECIFIED":
            return NOTSPECIFIED_SELECTIONLIKELIHOODINFO, nil
        case "HIGH":
            return HIGH_SELECTIONLIKELIHOODINFO, nil
    }
    return 0, errors.New("Unknown SelectionLikelihoodInfo value: " + v)
}
func SerializeSelectionLikelihoodInfo(values []SelectionLikelihoodInfo) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
