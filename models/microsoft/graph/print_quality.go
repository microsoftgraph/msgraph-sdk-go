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
    switch strings.ToUpper(v) {
        case "LOW":
            return LOW_PRINTQUALITY, nil
        case "MEDIUM":
            return MEDIUM_PRINTQUALITY, nil
        case "HIGH":
            return HIGH_PRINTQUALITY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTQUALITY, nil
    }
    return 0, errors.New("Unknown PrintQuality value: " + v)
}
func SerializePrintQuality(values []PrintQuality) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
