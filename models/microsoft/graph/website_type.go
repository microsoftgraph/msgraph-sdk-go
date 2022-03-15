package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type WebsiteType int

const (
    OTHER_WEBSITETYPE WebsiteType = iota
    HOME_WEBSITETYPE
    WORK_WEBSITETYPE
    BLOG_WEBSITETYPE
    PROFILE_WEBSITETYPE
)

func (i WebsiteType) String() string {
    return []string{"OTHER", "HOME", "WORK", "BLOG", "PROFILE"}[i]
}
func ParseWebsiteType(v string) (interface{}, error) {
    result := OTHER_WEBSITETYPE
    switch strings.ToUpper(v) {
        case "OTHER":
            result = OTHER_WEBSITETYPE
        case "HOME":
            result = HOME_WEBSITETYPE
        case "WORK":
            result = WORK_WEBSITETYPE
        case "BLOG":
            result = BLOG_WEBSITETYPE
        case "PROFILE":
            result = PROFILE_WEBSITETYPE
        default:
            return 0, errors.New("Unknown WebsiteType value: " + v)
    }
    return &result, nil
}
func SerializeWebsiteType(values []WebsiteType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
