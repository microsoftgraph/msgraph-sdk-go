package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "BLACKANDWHITE":
            return BLACKANDWHITE_PRINTCOLORMODE, nil
        case "GRAYSCALE":
            return GRAYSCALE_PRINTCOLORMODE, nil
        case "COLOR":
            return COLOR_PRINTCOLORMODE, nil
        case "AUTO":
            return AUTO_PRINTCOLORMODE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTCOLORMODE, nil
    }
    return 0, errors.New("Unknown PrintColorMode value: " + v)
}
func SerializePrintColorMode(values []PrintColorMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
