package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "ROLEISATTENDEE":
            return ROLEISATTENDEE_BROADCASTMEETINGAUDIENCE, nil
        case "ORGANIZATION":
            return ORGANIZATION_BROADCASTMEETINGAUDIENCE, nil
        case "EVERYONE":
            return EVERYONE_BROADCASTMEETINGAUDIENCE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_BROADCASTMEETINGAUDIENCE, nil
    }
    return 0, errors.New("Unknown BroadcastMeetingAudience value: " + v)
}
func SerializeBroadcastMeetingAudience(values []BroadcastMeetingAudience) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
