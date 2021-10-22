package graph
import (
    "strings"
    "errors"
)
type ScreenSharingRole int

const (
    VIEWER_SCREENSHARINGROLE ScreenSharingRole = iota
    SHARER_SCREENSHARINGROLE
)

func (i ScreenSharingRole) String() string {
    return []string{"VIEWER", "SHARER"}[i]
}
func ParseScreenSharingRole(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "VIEWER":
            return VIEWER_SCREENSHARINGROLE, nil
        case "SHARER":
            return SHARER_SCREENSHARINGROLE, nil
    }
    return 0, errors.New("Unknown ScreenSharingRole value: " + v)
}
func SerializeScreenSharingRole(values []ScreenSharingRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
