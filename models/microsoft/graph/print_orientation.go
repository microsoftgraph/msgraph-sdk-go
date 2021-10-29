package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "PORTRAIT":
            return PORTRAIT_PRINTORIENTATION, nil
        case "LANDSCAPE":
            return LANDSCAPE_PRINTORIENTATION, nil
        case "REVERSELANDSCAPE":
            return REVERSELANDSCAPE_PRINTORIENTATION, nil
        case "REVERSEPORTRAIT":
            return REVERSEPORTRAIT_PRINTORIENTATION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTORIENTATION, nil
    }
    return 0, errors.New("Unknown PrintOrientation value: " + v)
}
func SerializePrintOrientation(values []PrintOrientation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
