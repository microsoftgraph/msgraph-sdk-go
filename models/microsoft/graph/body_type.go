package graph
import (
    "strings"
    "errors"
)
// 
type BodyType int

const (
    TEXT_BODYTYPE BodyType = iota
    HTML_BODYTYPE
)

func (i BodyType) String() string {
    return []string{"TEXT", "HTML"}[i]
}
func ParseBodyType(v string) (interface{}, error) {
    result := TEXT_BODYTYPE
    switch strings.ToUpper(v) {
        case "TEXT":
            result = TEXT_BODYTYPE
        case "HTML":
            result = HTML_BODYTYPE
        default:
            return 0, errors.New("Unknown BodyType value: " + v)
    }
    return &result, nil
}
func SerializeBodyType(values []BodyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
