package graph
import (
    "strings"
    "errors"
)
// 
type MediaState int

const (
    ACTIVE_MEDIASTATE MediaState = iota
    INACTIVE_MEDIASTATE
    UNKNOWNFUTUREVALUE_MEDIASTATE
)

func (i MediaState) String() string {
    return []string{"ACTIVE", "INACTIVE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMediaState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_MEDIASTATE, nil
        case "INACTIVE":
            return INACTIVE_MEDIASTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MEDIASTATE, nil
    }
    return 0, errors.New("Unknown MediaState value: " + v)
}
func SerializeMediaState(values []MediaState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
