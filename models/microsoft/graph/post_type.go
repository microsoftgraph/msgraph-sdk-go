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
    switch strings.ToUpper(v) {
        case "REGULAR":
            return REGULAR_POSTTYPE, nil
        case "QUICK":
            return QUICK_POSTTYPE, nil
        case "STRATEGIC":
            return STRATEGIC_POSTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_POSTTYPE, nil
    }
    return 0, errors.New("Unknown PostType value: " + v)
}
func SerializePostType(values []PostType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
