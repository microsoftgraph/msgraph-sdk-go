package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "AUTO":
            return AUTO_PRINTSCALING, nil
        case "SHRINKTOFIT":
            return SHRINKTOFIT_PRINTSCALING, nil
        case "FILL":
            return FILL_PRINTSCALING, nil
        case "FIT":
            return FIT_PRINTSCALING, nil
        case "NONE":
            return NONE_PRINTSCALING, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTSCALING, nil
    }
    return 0, errors.New("Unknown PrintScaling value: " + v)
}
func SerializePrintScaling(values []PrintScaling) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
