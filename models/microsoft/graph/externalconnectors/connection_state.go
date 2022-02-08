package externalconnectors
import (
    "strings"
    "errors"
)
// 
type ConnectionState int

const (
    DRAFT_CONNECTIONSTATE ConnectionState = iota
    READY_CONNECTIONSTATE
    OBSOLETE_CONNECTIONSTATE
    LIMITEXCEEDED_CONNECTIONSTATE
    UNKNOWNFUTUREVALUE_CONNECTIONSTATE
)

func (i ConnectionState) String() string {
    return []string{"DRAFT", "READY", "OBSOLETE", "LIMITEXCEEDED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConnectionState(v string) (interface{}, error) {
    result := DRAFT_CONNECTIONSTATE
    switch strings.ToUpper(v) {
        case "DRAFT":
            result = DRAFT_CONNECTIONSTATE
        case "READY":
            result = READY_CONNECTIONSTATE
        case "OBSOLETE":
            result = OBSOLETE_CONNECTIONSTATE
        case "LIMITEXCEEDED":
            result = LIMITEXCEEDED_CONNECTIONSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONNECTIONSTATE
        default:
            return 0, errors.New("Unknown ConnectionState value: " + v)
    }
    return &result, nil
}
func SerializeConnectionState(values []ConnectionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
