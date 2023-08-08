package models
import (
    "errors"
)
// 
type MeetingChatHistoryDefaultMode int

const (
    NONE_MEETINGCHATHISTORYDEFAULTMODE MeetingChatHistoryDefaultMode = iota
    ALL_MEETINGCHATHISTORYDEFAULTMODE
    UNKNOWNFUTUREVALUE_MEETINGCHATHISTORYDEFAULTMODE
)

func (i MeetingChatHistoryDefaultMode) String() string {
    return []string{"none", "all", "unknownFutureValue"}[i]
}
func ParseMeetingChatHistoryDefaultMode(v string) (any, error) {
    result := NONE_MEETINGCHATHISTORYDEFAULTMODE
    switch v {
        case "none":
            result = NONE_MEETINGCHATHISTORYDEFAULTMODE
        case "all":
            result = ALL_MEETINGCHATHISTORYDEFAULTMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MEETINGCHATHISTORYDEFAULTMODE
        default:
            return 0, errors.New("Unknown MeetingChatHistoryDefaultMode value: " + v)
    }
    return &result, nil
}
func SerializeMeetingChatHistoryDefaultMode(values []MeetingChatHistoryDefaultMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
