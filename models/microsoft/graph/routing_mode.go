package graph
import (
    "strings"
    "errors"
)
// 
type RoutingMode int

const (
    ONETOONE_ROUTINGMODE RoutingMode = iota
    MULTICAST_ROUTINGMODE
    UNKNOWNFUTUREVALUE_ROUTINGMODE
)

func (i RoutingMode) String() string {
    return []string{"ONETOONE", "MULTICAST", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRoutingMode(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ONETOONE":
            return ONETOONE_ROUTINGMODE, nil
        case "MULTICAST":
            return MULTICAST_ROUTINGMODE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ROUTINGMODE, nil
    }
    return 0, errors.New("Unknown RoutingMode value: " + v)
}
func SerializeRoutingMode(values []RoutingMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
