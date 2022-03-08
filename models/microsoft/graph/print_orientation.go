package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the print singleton.
type PrintOrientation int

const (
    PORTRAIT_PRINTORIENTATION PrintOrientation = iota
    LANDSCAPE_PRINTORIENTATION
    REVERSELANDSCAPE_PRINTORIENTATION
    REVERSEPORTRAIT_PRINTORIENTATION
    UNKNOWNFUTUREVALUE_PRINTORIENTATION
)

func (i PrintOrientation) String() string {
    return []string{"PORTRAIT", "LANDSCAPE", "REVERSELANDSCAPE", "REVERSEPORTRAIT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintOrientation(v string) (interface{}, error) {
    result := PORTRAIT_PRINTORIENTATION
    switch strings.ToUpper(v) {
        case "PORTRAIT":
            result = PORTRAIT_PRINTORIENTATION
        case "LANDSCAPE":
            result = LANDSCAPE_PRINTORIENTATION
        case "REVERSELANDSCAPE":
            result = REVERSELANDSCAPE_PRINTORIENTATION
        case "REVERSEPORTRAIT":
            result = REVERSEPORTRAIT_PRINTORIENTATION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTORIENTATION
        default:
            return 0, errors.New("Unknown PrintOrientation value: " + v)
    }
    return &result, nil
}
func SerializePrintOrientation(values []PrintOrientation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
