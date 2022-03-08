package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the print singleton.
type PrintScaling int

const (
    AUTO_PRINTSCALING PrintScaling = iota
    SHRINKTOFIT_PRINTSCALING
    FILL_PRINTSCALING
    FIT_PRINTSCALING
    NONE_PRINTSCALING
    UNKNOWNFUTUREVALUE_PRINTSCALING
)

func (i PrintScaling) String() string {
    return []string{"AUTO", "SHRINKTOFIT", "FILL", "FIT", "NONE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintScaling(v string) (interface{}, error) {
    result := AUTO_PRINTSCALING
    switch strings.ToUpper(v) {
        case "AUTO":
            result = AUTO_PRINTSCALING
        case "SHRINKTOFIT":
            result = SHRINKTOFIT_PRINTSCALING
        case "FILL":
            result = FILL_PRINTSCALING
        case "FIT":
            result = FIT_PRINTSCALING
        case "NONE":
            result = NONE_PRINTSCALING
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTSCALING
        default:
            return 0, errors.New("Unknown PrintScaling value: " + v)
    }
    return &result, nil
}
func SerializePrintScaling(values []PrintScaling) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
