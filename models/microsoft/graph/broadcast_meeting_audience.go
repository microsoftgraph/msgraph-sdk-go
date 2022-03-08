package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type BroadcastMeetingAudience int

const (
    ROLEISATTENDEE_BROADCASTMEETINGAUDIENCE BroadcastMeetingAudience = iota
    ORGANIZATION_BROADCASTMEETINGAUDIENCE
    EVERYONE_BROADCASTMEETINGAUDIENCE
    UNKNOWNFUTUREVALUE_BROADCASTMEETINGAUDIENCE
)

func (i BroadcastMeetingAudience) String() string {
    return []string{"ROLEISATTENDEE", "ORGANIZATION", "EVERYONE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseBroadcastMeetingAudience(v string) (interface{}, error) {
    result := ROLEISATTENDEE_BROADCASTMEETINGAUDIENCE
    switch strings.ToUpper(v) {
        case "ROLEISATTENDEE":
            result = ROLEISATTENDEE_BROADCASTMEETINGAUDIENCE
        case "ORGANIZATION":
            result = ORGANIZATION_BROADCASTMEETINGAUDIENCE
        case "EVERYONE":
            result = EVERYONE_BROADCASTMEETINGAUDIENCE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_BROADCASTMEETINGAUDIENCE
        default:
            return 0, errors.New("Unknown BroadcastMeetingAudience value: " + v)
    }
    return &result, nil
}
func SerializeBroadcastMeetingAudience(values []BroadcastMeetingAudience) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
