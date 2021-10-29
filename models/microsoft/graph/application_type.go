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
    switch strings.ToUpper(v) {
        case "UNIVERSAL":
            return UNIVERSAL_APPLICATIONTYPE, nil
        case "DESKTOP":
            return DESKTOP_APPLICATIONTYPE, nil
    }
    return 0, errors.New("Unknown ApplicationType value: " + v)
}
func SerializeApplicationType(values []ApplicationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
