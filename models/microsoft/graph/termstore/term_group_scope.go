package termstore
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type TermGroupScope int

const (
    GLOBAL_TERMGROUPSCOPE TermGroupScope = iota
    SYSTEM_TERMGROUPSCOPE
    SITECOLLECTION_TERMGROUPSCOPE
    UNKNOWNFUTUREVALUE_TERMGROUPSCOPE
)

func (i TermGroupScope) String() string {
    return []string{"GLOBAL", "SYSTEM", "SITECOLLECTION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTermGroupScope(v string) (interface{}, error) {
    result := GLOBAL_TERMGROUPSCOPE
    switch strings.ToUpper(v) {
        case "GLOBAL":
            result = GLOBAL_TERMGROUPSCOPE
        case "SYSTEM":
            result = SYSTEM_TERMGROUPSCOPE
        case "SITECOLLECTION":
            result = SITECOLLECTION_TERMGROUPSCOPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TERMGROUPSCOPE
        default:
            return 0, errors.New("Unknown TermGroupScope value: " + v)
    }
    return &result, nil
}
func SerializeTermGroupScope(values []TermGroupScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
