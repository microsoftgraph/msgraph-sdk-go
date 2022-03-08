package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the print singleton.
type PrintEvent int

const (
    JOBSTARTED_PRINTEVENT PrintEvent = iota
    UNKNOWNFUTUREVALUE_PRINTEVENT
)

func (i PrintEvent) String() string {
    return []string{"JOBSTARTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintEvent(v string) (interface{}, error) {
    result := JOBSTARTED_PRINTEVENT
    switch strings.ToUpper(v) {
        case "JOBSTARTED":
            result = JOBSTARTED_PRINTEVENT
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTEVENT
        default:
            return 0, errors.New("Unknown PrintEvent value: " + v)
    }
    return &result, nil
}
func SerializePrintEvent(values []PrintEvent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
