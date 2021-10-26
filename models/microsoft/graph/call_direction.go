package graph
import (
    "strings"
    "errors"
)
// 
type CallDirection int

const (
    INCOMING_CALLDIRECTION CallDirection = iota
    OUTGOING_CALLDIRECTION
)

func (i CallDirection) String() string {
    return []string{"INCOMING", "OUTGOING"}[i]
}
func ParseCallDirection(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "INCOMING":
            return INCOMING_CALLDIRECTION, nil
        case "OUTGOING":
            return OUTGOING_CALLDIRECTION, nil
    }
    return 0, errors.New("Unknown CallDirection value: " + v)
}
func SerializeCallDirection(values []CallDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
