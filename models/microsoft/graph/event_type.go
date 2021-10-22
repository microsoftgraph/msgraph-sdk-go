package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "SINGLEINSTANCE":
            return SINGLEINSTANCE_EVENTTYPE, nil
        case "OCCURRENCE":
            return OCCURRENCE_EVENTTYPE, nil
        case "EXCEPTION":
            return EXCEPTION_EVENTTYPE, nil
        case "SERIESMASTER":
            return SERIESMASTER_EVENTTYPE, nil
    }
    return 0, errors.New("Unknown EventType value: " + v)
}
func SerializeEventType(values []EventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
