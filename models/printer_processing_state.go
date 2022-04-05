package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the print singleton.
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
    result := UNKNOWN_PRINTERPROCESSINGSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_PRINTERPROCESSINGSTATE
        case "IDLE":
            result = IDLE_PRINTERPROCESSINGSTATE
        case "PROCESSING":
            result = PROCESSING_PRINTERPROCESSINGSTATE
        case "STOPPED":
            result = STOPPED_PRINTERPROCESSINGSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTERPROCESSINGSTATE
        default:
            return 0, errors.New("Unknown PrinterProcessingState value: " + v)
    }
    return &result, nil
}
func SerializePrinterProcessingState(values []PrinterProcessingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
