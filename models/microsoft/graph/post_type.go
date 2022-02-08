package graph
import (
    "strings"
    "errors"
)
// 
type PostType int

const (
    REGULAR_POSTTYPE PostType = iota
    QUICK_POSTTYPE
    STRATEGIC_POSTTYPE
    UNKNOWNFUTUREVALUE_POSTTYPE
)

func (i PostType) String() string {
    return []string{"REGULAR", "QUICK", "STRATEGIC", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePostType(v string) (interface{}, error) {
    result := REGULAR_POSTTYPE
    switch strings.ToUpper(v) {
        case "REGULAR":
            result = REGULAR_POSTTYPE
        case "QUICK":
            result = QUICK_POSTTYPE
        case "STRATEGIC":
            result = STRATEGIC_POSTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_POSTTYPE
        default:
            return 0, errors.New("Unknown PostType value: " + v)
    }
    return &result, nil
}
func SerializePostType(values []PostType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
