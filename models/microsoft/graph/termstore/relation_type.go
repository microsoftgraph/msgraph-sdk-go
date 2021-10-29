package termstore
import (
    "strings"
    "errors"
)
// 
type RelationType int

const (
    PIN_RELATIONTYPE RelationType = iota
    REUSE_RELATIONTYPE
    UNKNOWNFUTUREVALUE_RELATIONTYPE
)

func (i RelationType) String() string {
    return []string{"PIN", "REUSE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRelationType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "PIN":
            return PIN_RELATIONTYPE, nil
        case "REUSE":
            return REUSE_RELATIONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_RELATIONTYPE, nil
    }
    return 0, errors.New("Unknown RelationType value: " + v)
}
func SerializeRelationType(values []RelationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
