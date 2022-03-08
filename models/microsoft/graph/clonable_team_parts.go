package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the clone method.
type ClonableTeamParts int

const (
    APPS_CLONABLETEAMPARTS ClonableTeamParts = iota
    TABS_CLONABLETEAMPARTS
    SETTINGS_CLONABLETEAMPARTS
    CHANNELS_CLONABLETEAMPARTS
    MEMBERS_CLONABLETEAMPARTS
)

func (i ClonableTeamParts) String() string {
    return []string{"APPS", "TABS", "SETTINGS", "CHANNELS", "MEMBERS"}[i]
}
func ParseClonableTeamParts(v string) (interface{}, error) {
    result := APPS_CLONABLETEAMPARTS
    switch strings.ToUpper(v) {
        case "APPS":
            result = APPS_CLONABLETEAMPARTS
        case "TABS":
            result = TABS_CLONABLETEAMPARTS
        case "SETTINGS":
            result = SETTINGS_CLONABLETEAMPARTS
        case "CHANNELS":
            result = CHANNELS_CLONABLETEAMPARTS
        case "MEMBERS":
            result = MEMBERS_CLONABLETEAMPARTS
        default:
            return 0, errors.New("Unknown ClonableTeamParts value: " + v)
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
