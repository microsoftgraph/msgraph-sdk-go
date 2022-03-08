package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the print singleton.
type PrintColorMode int

const (
    BLACKANDWHITE_PRINTCOLORMODE PrintColorMode = iota
    GRAYSCALE_PRINTCOLORMODE
    COLOR_PRINTCOLORMODE
    AUTO_PRINTCOLORMODE
    UNKNOWNFUTUREVALUE_PRINTCOLORMODE
)

func (i PrintColorMode) String() string {
    return []string{"BLACKANDWHITE", "GRAYSCALE", "COLOR", "AUTO", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintColorMode(v string) (interface{}, error) {
    result := BLACKANDWHITE_PRINTCOLORMODE
    switch strings.ToUpper(v) {
        case "BLACKANDWHITE":
            result = BLACKANDWHITE_PRINTCOLORMODE
        case "GRAYSCALE":
            result = GRAYSCALE_PRINTCOLORMODE
        case "COLOR":
            result = COLOR_PRINTCOLORMODE
        case "AUTO":
            result = AUTO_PRINTCOLORMODE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTCOLORMODE
        default:
            return 0, errors.New("Unknown PrintColorMode value: " + v)
    }
    return &result, nil
}
func SerializePrintColorMode(values []PrintColorMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
