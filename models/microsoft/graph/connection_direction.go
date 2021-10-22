package graph
import (
    "strings"
    "errors"
)
type ConnectionDirection int

const (
    UNKNOWN_CONNECTIONDIRECTION ConnectionDirection = iota
    INBOUND_CONNECTIONDIRECTION
    OUTBOUND_CONNECTIONDIRECTION
    UNKNOWNFUTUREVALUE_CONNECTIONDIRECTION
)

func (i ConnectionDirection) String() string {
    return []string{"UNKNOWN", "INBOUND", "OUTBOUND", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConnectionDirection(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_CONNECTIONDIRECTION, nil
        case "INBOUND":
            return INBOUND_CONNECTIONDIRECTION, nil
        case "OUTBOUND":
            return OUTBOUND_CONNECTIONDIRECTION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONNECTIONDIRECTION, nil
    }
    return 0, errors.New("Unknown ConnectionDirection value: " + v)
}
func SerializeConnectionDirection(values []ConnectionDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
