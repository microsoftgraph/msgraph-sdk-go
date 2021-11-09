package ediscovery
import (
    "strings"
    "errors"
)
// 
type CaseAction int

const (
    CONTENTEXPORT_CASEACTION CaseAction = iota
    APPLYTAGS_CASEACTION
    CONVERTTOPDF_CASEACTION
    INDEX_CASEACTION
    ESTIMATESTATISTICS_CASEACTION
    ADDTOREVIEWSET_CASEACTION
    UNKNOWNFUTUREVALUE_CASEACTION
)

func (i CaseAction) String() string {
    return []string{"CONTENTEXPORT", "APPLYTAGS", "CONVERTTOPDF", "INDEX", "ESTIMATESTATISTICS", "ADDTOREVIEWSET", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCaseAction(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CONTENTEXPORT":
            return CONTENTEXPORT_CASEACTION, nil
        case "APPLYTAGS":
            return APPLYTAGS_CASEACTION, nil
        case "CONVERTTOPDF":
            return CONVERTTOPDF_CASEACTION, nil
        case "INDEX":
            return INDEX_CASEACTION, nil
        case "ESTIMATESTATISTICS":
            return ESTIMATESTATISTICS_CASEACTION, nil
        case "ADDTOREVIEWSET":
            return ADDTOREVIEWSET_CASEACTION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CASEACTION, nil
    }
    return 0, errors.New("Unknown CaseAction value: " + v)
}
func SerializeCaseAction(values []CaseAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
