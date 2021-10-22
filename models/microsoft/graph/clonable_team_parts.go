package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "APPS":
            return APPS_CLONABLETEAMPARTS, nil
        case "TABS":
            return TABS_CLONABLETEAMPARTS, nil
        case "SETTINGS":
            return SETTINGS_CLONABLETEAMPARTS, nil
        case "CHANNELS":
            return CHANNELS_CLONABLETEAMPARTS, nil
        case "MEMBERS":
            return MEMBERS_CLONABLETEAMPARTS, nil
    }
    return 0, errors.New("Unknown ClonableTeamParts value: " + v)
}
func SerializeClonableTeamParts(values []ClonableTeamParts) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
