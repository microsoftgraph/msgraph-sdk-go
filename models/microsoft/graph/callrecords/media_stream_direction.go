package callrecords
import (
    "strings"
    "errors"
)
// 
type MediaStreamDirection int

const (
    CALLERTOCALLEE_MEDIASTREAMDIRECTION MediaStreamDirection = iota
    CALLEETOCALLER_MEDIASTREAMDIRECTION
)

func (i MediaStreamDirection) String() string {
    return []string{"CALLERTOCALLEE", "CALLEETOCALLER"}[i]
}
func ParseMediaStreamDirection(v string) (interface{}, error) {
    result := CALLERTOCALLEE_MEDIASTREAMDIRECTION
    switch strings.ToUpper(v) {
        case "CALLERTOCALLEE":
            result = CALLERTOCALLEE_MEDIASTREAMDIRECTION
        case "CALLEETOCALLER":
            result = CALLEETOCALLER_MEDIASTREAMDIRECTION
        default:
            return 0, errors.New("Unknown MediaStreamDirection value: " + v)
    }
    return &result, nil
}
func SerializeMediaStreamDirection(values []MediaStreamDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
