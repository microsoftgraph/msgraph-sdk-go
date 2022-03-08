package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the print singleton.
type PrintDuplexMode int

const (
    FLIPONLONGEDGE_PRINTDUPLEXMODE PrintDuplexMode = iota
    FLIPONSHORTEDGE_PRINTDUPLEXMODE
    ONESIDED_PRINTDUPLEXMODE
    UNKNOWNFUTUREVALUE_PRINTDUPLEXMODE
)

func (i PrintDuplexMode) String() string {
    return []string{"FLIPONLONGEDGE", "FLIPONSHORTEDGE", "ONESIDED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintDuplexMode(v string) (interface{}, error) {
    result := FLIPONLONGEDGE_PRINTDUPLEXMODE
    switch strings.ToUpper(v) {
        case "FLIPONLONGEDGE":
            result = FLIPONLONGEDGE_PRINTDUPLEXMODE
        case "FLIPONSHORTEDGE":
            result = FLIPONSHORTEDGE_PRINTDUPLEXMODE
        case "ONESIDED":
            result = ONESIDED_PRINTDUPLEXMODE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTDUPLEXMODE
        default:
            return 0, errors.New("Unknown PrintDuplexMode value: " + v)
    }
    return &result, nil
}
func SerializePrintDuplexMode(values []PrintDuplexMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
