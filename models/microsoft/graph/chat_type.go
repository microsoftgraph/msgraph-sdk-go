package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ONEONONE":
            return ONEONONE_CHATTYPE, nil
        case "GROUP":
            return GROUP_CHATTYPE, nil
        case "MEETING":
            return MEETING_CHATTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CHATTYPE, nil
    }
    return 0, errors.New("Unknown ChatType value: " + v)
}
func SerializeChatType(values []ChatType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
