package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
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
    result := FORWARDED_ROUTINGTYPE
    switch strings.ToUpper(v) {
        case "FORWARDED":
            result = FORWARDED_ROUTINGTYPE
        case "LOOKUP":
            result = LOOKUP_ROUTINGTYPE
        case "SELFFORK":
            result = SELFFORK_ROUTINGTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ROUTINGTYPE
        default:
            return 0, errors.New("Unknown RoutingType value: " + v)
    }
    return &result, nil
}
func SerializeRoutingType(values []RoutingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
