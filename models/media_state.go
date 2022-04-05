package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
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
    result := ACTIVE_MEDIASTATE
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_MEDIASTATE
        case "INACTIVE":
            result = INACTIVE_MEDIASTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MEDIASTATE
        default:
            return 0, errors.New("Unknown MediaState value: " + v)
    }
    return &result, nil
}
func SerializeMediaState(values []MediaState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
