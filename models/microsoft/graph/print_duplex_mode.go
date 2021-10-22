package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "FLIPONLONGEDGE":
            return FLIPONLONGEDGE_PRINTDUPLEXMODE, nil
        case "FLIPONSHORTEDGE":
            return FLIPONSHORTEDGE_PRINTDUPLEXMODE, nil
        case "ONESIDED":
            return ONESIDED_PRINTDUPLEXMODE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTDUPLEXMODE, nil
    }
    return 0, errors.New("Unknown PrintDuplexMode value: " + v)
}
func SerializePrintDuplexMode(values []PrintDuplexMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
