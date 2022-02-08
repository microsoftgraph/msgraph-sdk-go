package graph
import (
    "strings"
    "errors"
)
// 
type ChatType int

const (
    ONEONONE_CHATTYPE ChatType = iota
    GROUP_CHATTYPE
    MEETING_CHATTYPE
    UNKNOWNFUTUREVALUE_CHATTYPE
)

func (i ChatType) String() string {
    return []string{"ONEONONE", "GROUP", "MEETING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseChatType(v string) (interface{}, error) {
    result := ONEONONE_CHATTYPE
    switch strings.ToUpper(v) {
        case "ONEONONE":
            result = ONEONONE_CHATTYPE
        case "GROUP":
            result = GROUP_CHATTYPE
        case "MEETING":
            result = MEETING_CHATTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CHATTYPE
        default:
            return 0, errors.New("Unknown ChatType value: " + v)
    }
    return &result, nil
}
func SerializeChatType(values []ChatType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
