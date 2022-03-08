package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the changeScreenSharingRole method.
type ScreenSharingRole int

const (
    VIEWER_SCREENSHARINGROLE ScreenSharingRole = iota
    SHARER_SCREENSHARINGROLE
)

func (i ScreenSharingRole) String() string {
    return []string{"VIEWER", "SHARER"}[i]
}
func ParseScreenSharingRole(v string) (interface{}, error) {
    result := VIEWER_SCREENSHARINGROLE
    switch strings.ToUpper(v) {
        case "VIEWER":
            result = VIEWER_SCREENSHARINGROLE
        case "SHARER":
            result = SHARER_SCREENSHARINGROLE
        default:
            return 0, errors.New("Unknown ScreenSharingRole value: " + v)
    }
    return &result, nil
}
func SerializeScreenSharingRole(values []ScreenSharingRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
