package graph
import (
    "strings"
    "errors"
)
// 
type ApplicationType int

const (
    UNIVERSAL_APPLICATIONTYPE ApplicationType = iota
    DESKTOP_APPLICATIONTYPE
)

func (i ApplicationType) String() string {
    return []string{"UNIVERSAL", "DESKTOP"}[i]
}
func ParseApplicationType(v string) (interface{}, error) {
    result := UNIVERSAL_APPLICATIONTYPE
    switch strings.ToUpper(v) {
        case "UNIVERSAL":
            result = UNIVERSAL_APPLICATIONTYPE
        case "DESKTOP":
            result = DESKTOP_APPLICATIONTYPE
        default:
            return 0, errors.New("Unknown ApplicationType value: " + v)
    }
    return &result, nil
}
func SerializeApplicationType(values []ApplicationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
