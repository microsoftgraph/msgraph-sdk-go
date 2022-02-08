package graph
import (
    "strings"
    "errors"
)
// 
type PrintQuality int

const (
    LOW_PRINTQUALITY PrintQuality = iota
    MEDIUM_PRINTQUALITY
    HIGH_PRINTQUALITY
    UNKNOWNFUTUREVALUE_PRINTQUALITY
)

func (i PrintQuality) String() string {
    return []string{"LOW", "MEDIUM", "HIGH", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintQuality(v string) (interface{}, error) {
    result := LOW_PRINTQUALITY
    switch strings.ToUpper(v) {
        case "LOW":
            result = LOW_PRINTQUALITY
        case "MEDIUM":
            result = MEDIUM_PRINTQUALITY
        case "HIGH":
            result = HIGH_PRINTQUALITY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTQUALITY
        default:
            return 0, errors.New("Unknown PrintQuality value: " + v)
    }
    return &result, nil
}
func SerializePrintQuality(values []PrintQuality) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
