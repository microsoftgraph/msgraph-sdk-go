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
    switch strings.ToUpper(v) {
        case "ENABLED":
            return ENABLED_MEETINGCHATMODE, nil
        case "DISABLED":
            return DISABLED_MEETINGCHATMODE, nil
        case "LIMITED":
            return LIMITED_MEETINGCHATMODE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MEETINGCHATMODE, nil
    }
    return 0, errors.New("Unknown MeetingChatMode value: " + v)
}
func SerializeMeetingChatMode(values []MeetingChatMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
