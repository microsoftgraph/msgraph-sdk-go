package graph
import (
    "strings"
    "errors"
)
type BodyType int

const (
    TEXT_BODYTYPE BodyType = iota
    HTML_BODYTYPE
)

func (i BodyType) String() string {
    return []string{"TEXT", "HTML"}[i]
}
func ParseBodyType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "TEXT":
            return TEXT_BODYTYPE, nil
        case "HTML":
            return HTML_BODYTYPE, nil
    }
    return 0, errors.New("Unknown BodyType value: " + v)
}
func SerializeBodyType(values []BodyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
