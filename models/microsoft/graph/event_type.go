package graph
import (
    "strings"
    "errors"
)
// 
type EventType int

const (
    SINGLEINSTANCE_EVENTTYPE EventType = iota
    OCCURRENCE_EVENTTYPE
    EXCEPTION_EVENTTYPE
    SERIESMASTER_EVENTTYPE
)

func (i EventType) String() string {
    return []string{"SINGLEINSTANCE", "OCCURRENCE", "EXCEPTION", "SERIESMASTER"}[i]
}
func ParseEventType(v string) (interface{}, error) {
    result := SINGLEINSTANCE_EVENTTYPE
    switch strings.ToUpper(v) {
        case "SINGLEINSTANCE":
            result = SINGLEINSTANCE_EVENTTYPE
        case "OCCURRENCE":
            result = OCCURRENCE_EVENTTYPE
        case "EXCEPTION":
            result = EXCEPTION_EVENTTYPE
        case "SERIESMASTER":
            result = SERIESMASTER_EVENTTYPE
        default:
            return 0, errors.New("Unknown EventType value: " + v)
    }
    return &result, nil
}
func SerializeEventType(values []EventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
