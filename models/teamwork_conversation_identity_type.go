package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of chat entities.
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
    result := TEAM_TEAMWORKCONVERSATIONIDENTITYTYPE
    switch strings.ToUpper(v) {
        case "TEAM":
            result = TEAM_TEAMWORKCONVERSATIONIDENTITYTYPE
        case "CHANNEL":
            result = CHANNEL_TEAMWORKCONVERSATIONIDENTITYTYPE
        case "CHAT":
            result = CHAT_TEAMWORKCONVERSATIONIDENTITYTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMWORKCONVERSATIONIDENTITYTYPE
        default:
            return 0, errors.New("Unknown TeamworkConversationIdentityType value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkConversationIdentityType(values []TeamworkConversationIdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
