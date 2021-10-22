package graph
import (
    "strings"
    "errors"
)
type PrintEvent int

const (
    JOBSTARTED_PRINTEVENT PrintEvent = iota
    UNKNOWNFUTUREVALUE_PRINTEVENT
)

func (i PrintEvent) String() string {
    return []string{"JOBSTARTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintEvent(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "JOBSTARTED":
            return JOBSTARTED_PRINTEVENT, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTEVENT, nil
    }
    return 0, errors.New("Unknown PrintEvent value: " + v)
}
func SerializePrintEvent(values []PrintEvent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
