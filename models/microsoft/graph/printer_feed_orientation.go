package graph
import (
    "strings"
    "errors"
)
type PrinterFeedOrientation int

const (
    LONGEDGEFIRST_PRINTERFEEDORIENTATION PrinterFeedOrientation = iota
    SHORTEDGEFIRST_PRINTERFEEDORIENTATION
    UNKNOWNFUTUREVALUE_PRINTERFEEDORIENTATION
)

func (i PrinterFeedOrientation) String() string {
    return []string{"LONGEDGEFIRST", "SHORTEDGEFIRST", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrinterFeedOrientation(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "LONGEDGEFIRST":
            return LONGEDGEFIRST_PRINTERFEEDORIENTATION, nil
        case "SHORTEDGEFIRST":
            return SHORTEDGEFIRST_PRINTERFEEDORIENTATION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTERFEEDORIENTATION, nil
    }
    return 0, errors.New("Unknown PrinterFeedOrientation value: " + v)
}
func SerializePrinterFeedOrientation(values []PrinterFeedOrientation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
