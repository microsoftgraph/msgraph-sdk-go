package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UPLOADPENDING":
            return UPLOADPENDING_PRINTJOBSTATEDETAIL, nil
        case "TRANSFORMING":
            return TRANSFORMING_PRINTJOBSTATEDETAIL, nil
        case "COMPLETEDSUCCESSFULLY":
            return COMPLETEDSUCCESSFULLY_PRINTJOBSTATEDETAIL, nil
        case "COMPLETEDWITHWARNINGS":
            return COMPLETEDWITHWARNINGS_PRINTJOBSTATEDETAIL, nil
        case "COMPLETEDWITHERRORS":
            return COMPLETEDWITHERRORS_PRINTJOBSTATEDETAIL, nil
        case "RELEASEWAIT":
            return RELEASEWAIT_PRINTJOBSTATEDETAIL, nil
        case "INTERPRETING":
            return INTERPRETING_PRINTJOBSTATEDETAIL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTJOBSTATEDETAIL, nil
    }
    return 0, errors.New("Unknown PrintJobStateDetail value: " + v)
}
func SerializePrintJobStateDetail(values []PrintJobStateDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
