package graph
import (
    "strings"
    "errors"
)
// 
type PrinterProcessingState int

const (
    UNKNOWN_PRINTERPROCESSINGSTATE PrinterProcessingState = iota
    IDLE_PRINTERPROCESSINGSTATE
    PROCESSING_PRINTERPROCESSINGSTATE
    STOPPED_PRINTERPROCESSINGSTATE
    UNKNOWNFUTUREVALUE_PRINTERPROCESSINGSTATE
)

func (i PrinterProcessingState) String() string {
    return []string{"UNKNOWN", "IDLE", "PROCESSING", "STOPPED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrinterProcessingState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_PRINTERPROCESSINGSTATE, nil
        case "IDLE":
            return IDLE_PRINTERPROCESSINGSTATE, nil
        case "PROCESSING":
            return PROCESSING_PRINTERPROCESSINGSTATE, nil
        case "STOPPED":
            return STOPPED_PRINTERPROCESSINGSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTERPROCESSINGSTATE, nil
    }
    return 0, errors.New("Unknown PrinterProcessingState value: " + v)
}
func SerializePrinterProcessingState(values []PrinterProcessingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
