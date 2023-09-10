package models
import (
    "errors"
    "strings"
)
// 
type ClonableTeamParts int

const (
    APPS_CLONABLETEAMPARTS ClonableTeamParts = iota
    TABS_CLONABLETEAMPARTS
    SETTINGS_CLONABLETEAMPARTS
    CHANNELS_CLONABLETEAMPARTS
    MEMBERS_CLONABLETEAMPARTS
)

func (i ClonableTeamParts) String() string {
    var values []string
    for p := ClonableTeamParts(1); p <= MEMBERS_CLONABLETEAMPARTS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"apps", "tabs", "settings", "channels", "members"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseClonableTeamParts(v string) (any, error) {
    var result ClonableTeamParts
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "apps":
                result |= APPS_CLONABLETEAMPARTS
            case "tabs":
                result |= TABS_CLONABLETEAMPARTS
            case "settings":
                result |= SETTINGS_CLONABLETEAMPARTS
            case "channels":
                result |= CHANNELS_CLONABLETEAMPARTS
            case "members":
                result |= MEMBERS_CLONABLETEAMPARTS
            default:
                return 0, errors.New("Unknown ClonableTeamParts value: " + v)
        }
    }
    return &result, nil
}
func SerializeClonableTeamParts(values []ClonableTeamParts) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ClonableTeamParts) isMultiValue() bool {
    return true
}
