package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the print singleton.
type PrintJobStateDetail int

const (
    UPLOADPENDING_PRINTJOBSTATEDETAIL PrintJobStateDetail = iota
    TRANSFORMING_PRINTJOBSTATEDETAIL
    COMPLETEDSUCCESSFULLY_PRINTJOBSTATEDETAIL
    COMPLETEDWITHWARNINGS_PRINTJOBSTATEDETAIL
    COMPLETEDWITHERRORS_PRINTJOBSTATEDETAIL
    RELEASEWAIT_PRINTJOBSTATEDETAIL
    INTERPRETING_PRINTJOBSTATEDETAIL
    UNKNOWNFUTUREVALUE_PRINTJOBSTATEDETAIL
)

func (i PrintJobStateDetail) String() string {
    return []string{"UPLOADPENDING", "TRANSFORMING", "COMPLETEDSUCCESSFULLY", "COMPLETEDWITHWARNINGS", "COMPLETEDWITHERRORS", "RELEASEWAIT", "INTERPRETING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintJobStateDetail(v string) (interface{}, error) {
    result := UPLOADPENDING_PRINTJOBSTATEDETAIL
    switch strings.ToUpper(v) {
        case "UPLOADPENDING":
            result = UPLOADPENDING_PRINTJOBSTATEDETAIL
        case "TRANSFORMING":
            result = TRANSFORMING_PRINTJOBSTATEDETAIL
        case "COMPLETEDSUCCESSFULLY":
            result = COMPLETEDSUCCESSFULLY_PRINTJOBSTATEDETAIL
        case "COMPLETEDWITHWARNINGS":
            result = COMPLETEDWITHWARNINGS_PRINTJOBSTATEDETAIL
        case "COMPLETEDWITHERRORS":
            result = COMPLETEDWITHERRORS_PRINTJOBSTATEDETAIL
        case "RELEASEWAIT":
            result = RELEASEWAIT_PRINTJOBSTATEDETAIL
        case "INTERPRETING":
            result = INTERPRETING_PRINTJOBSTATEDETAIL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTJOBSTATEDETAIL
        default:
            return 0, errors.New("Unknown PrintJobStateDetail value: " + v)
    }
    return &result, nil
}
func SerializePrintJobStateDetail(values []PrintJobStateDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
