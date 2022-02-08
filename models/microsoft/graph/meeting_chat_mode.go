package graph
import (
    "strings"
    "errors"
)
// 
type MeetingChatMode int

const (
    ENABLED_MEETINGCHATMODE MeetingChatMode = iota
    DISABLED_MEETINGCHATMODE
    LIMITED_MEETINGCHATMODE
    UNKNOWNFUTUREVALUE_MEETINGCHATMODE
)

func (i MeetingChatMode) String() string {
    return []string{"ENABLED", "DISABLED", "LIMITED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMeetingChatMode(v string) (interface{}, error) {
    result := ENABLED_MEETINGCHATMODE
    switch strings.ToUpper(v) {
        case "ENABLED":
            result = ENABLED_MEETINGCHATMODE
        case "DISABLED":
            result = DISABLED_MEETINGCHATMODE
        case "LIMITED":
            result = LIMITED_MEETINGCHATMODE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MEETINGCHATMODE
        default:
            return 0, errors.New("Unknown MeetingChatMode value: " + v)
    }
    return &result, nil
}
func SerializeMeetingChatMode(values []MeetingChatMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
