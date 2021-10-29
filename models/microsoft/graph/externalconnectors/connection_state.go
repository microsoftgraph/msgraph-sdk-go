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
    switch strings.ToUpper(v) {
        case "DRAFT":
            return DRAFT_CONNECTIONSTATE, nil
        case "READY":
            return READY_CONNECTIONSTATE, nil
        case "OBSOLETE":
            return OBSOLETE_CONNECTIONSTATE, nil
        case "LIMITEXCEEDED":
            return LIMITEXCEEDED_CONNECTIONSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONNECTIONSTATE, nil
    }
    return 0, errors.New("Unknown ConnectionState value: " + v)
}
func SerializeConnectionState(values []ConnectionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
