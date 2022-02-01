package termstore
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "GLOBAL":
            return GLOBAL_TERMGROUPSCOPE, nil
        case "SYSTEM":
            return SYSTEM_TERMGROUPSCOPE, nil
        case "SITECOLLECTION":
            return SITECOLLECTION_TERMGROUPSCOPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TERMGROUPSCOPE, nil
    }
    return 0, errors.New("Unknown TermGroupScope value: " + v)
}
func SerializeTermGroupScope(values []TermGroupScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
