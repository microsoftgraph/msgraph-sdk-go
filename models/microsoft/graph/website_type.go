package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "OTHER":
            return OTHER_WEBSITETYPE, nil
        case "HOME":
            return HOME_WEBSITETYPE, nil
        case "WORK":
            return WORK_WEBSITETYPE, nil
        case "BLOG":
            return BLOG_WEBSITETYPE, nil
        case "PROFILE":
            return PROFILE_WEBSITETYPE, nil
    }
    return 0, errors.New("Unknown WebsiteType value: " + v)
}
func SerializeWebsiteType(values []WebsiteType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
