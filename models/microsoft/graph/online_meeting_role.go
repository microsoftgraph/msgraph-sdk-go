package graph
import (
    "strings"
    "errors"
)
// 
type OnlineMeetingRole int

const (
    ATTENDEE_ONLINEMEETINGROLE OnlineMeetingRole = iota
    PRESENTER_ONLINEMEETINGROLE
    UNKNOWNFUTUREVALUE_ONLINEMEETINGROLE
    PRODUCER_ONLINEMEETINGROLE
)

func (i OnlineMeetingRole) String() string {
    return []string{"ATTENDEE", "PRESENTER", "UNKNOWNFUTUREVALUE", "PRODUCER"}[i]
}
func ParseOnlineMeetingRole(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ATTENDEE":
            return ATTENDEE_ONLINEMEETINGROLE, nil
        case "PRESENTER":
            return PRESENTER_ONLINEMEETINGROLE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ONLINEMEETINGROLE, nil
        case "PRODUCER":
            return PRODUCER_ONLINEMEETINGROLE, nil
    }
    return 0, errors.New("Unknown OnlineMeetingRole value: " + v)
}
func SerializeOnlineMeetingRole(values []OnlineMeetingRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
