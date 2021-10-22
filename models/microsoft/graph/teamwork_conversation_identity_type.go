package graph
import (
    "strings"
    "errors"
)
type TeamworkConversationIdentityType int

const (
    TEAM_TEAMWORKCONVERSATIONIDENTITYTYPE TeamworkConversationIdentityType = iota
    CHANNEL_TEAMWORKCONVERSATIONIDENTITYTYPE
    CHAT_TEAMWORKCONVERSATIONIDENTITYTYPE
    UNKNOWNFUTUREVALUE_TEAMWORKCONVERSATIONIDENTITYTYPE
)

func (i TeamworkConversationIdentityType) String() string {
    return []string{"TEAM", "CHANNEL", "CHAT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamworkConversationIdentityType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "TEAM":
            return TEAM_TEAMWORKCONVERSATIONIDENTITYTYPE, nil
        case "CHANNEL":
            return CHANNEL_TEAMWORKCONVERSATIONIDENTITYTYPE, nil
        case "CHAT":
            return CHAT_TEAMWORKCONVERSATIONIDENTITYTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMWORKCONVERSATIONIDENTITYTYPE, nil
    }
    return 0, errors.New("Unknown TeamworkConversationIdentityType value: " + v)
}
func SerializeTeamworkConversationIdentityType(values []TeamworkConversationIdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
