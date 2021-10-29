package graph
import (
    "strings"
    "errors"
)
// 
type RoutingType int

const (
    FORWARDED_ROUTINGTYPE RoutingType = iota
    LOOKUP_ROUTINGTYPE
    SELFFORK_ROUTINGTYPE
    UNKNOWNFUTUREVALUE_ROUTINGTYPE
)

func (i RoutingType) String() string {
    return []string{"FORWARDED", "LOOKUP", "SELFFORK", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRoutingType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "FORWARDED":
            return FORWARDED_ROUTINGTYPE, nil
        case "LOOKUP":
            return LOOKUP_ROUTINGTYPE, nil
        case "SELFFORK":
            return SELFFORK_ROUTINGTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ROUTINGTYPE, nil
    }
    return 0, errors.New("Unknown RoutingType value: " + v)
}
func SerializeRoutingType(values []RoutingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
