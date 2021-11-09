package ediscovery
import (
    "strings"
    "errors"
)
// 
type ChildSelectability int

const (
    ONE_CHILDSELECTABILITY ChildSelectability = iota
    MANY_CHILDSELECTABILITY
)

func (i ChildSelectability) String() string {
    return []string{"ONE", "MANY"}[i]
}
func ParseChildSelectability(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ONE":
            return ONE_CHILDSELECTABILITY, nil
        case "MANY":
            return MANY_CHILDSELECTABILITY, nil
    }
    return 0, errors.New("Unknown ChildSelectability value: " + v)
}
func SerializeChildSelectability(values []ChildSelectability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
