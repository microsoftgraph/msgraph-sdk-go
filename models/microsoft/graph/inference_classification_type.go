package graph
import (
    "strings"
    "errors"
)
type InferenceClassificationType int

const (
    FOCUSED_INFERENCECLASSIFICATIONTYPE InferenceClassificationType = iota
    OTHER_INFERENCECLASSIFICATIONTYPE
)

func (i InferenceClassificationType) String() string {
    return []string{"FOCUSED", "OTHER"}[i]
}
func ParseInferenceClassificationType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "FOCUSED":
            return FOCUSED_INFERENCECLASSIFICATIONTYPE, nil
        case "OTHER":
            return OTHER_INFERENCECLASSIFICATIONTYPE, nil
    }
    return 0, errors.New("Unknown InferenceClassificationType value: " + v)
}
func SerializeInferenceClassificationType(values []InferenceClassificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
