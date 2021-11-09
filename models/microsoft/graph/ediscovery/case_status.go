package ediscovery
import (
    "strings"
    "errors"
)
// 
type CaseStatus int

const (
    UNKNOWN_CASESTATUS CaseStatus = iota
    ACTIVE_CASESTATUS
    PENDINGDELETE_CASESTATUS
    CLOSING_CASESTATUS
    CLOSED_CASESTATUS
    CLOSEDWITHERROR_CASESTATUS
)

func (i CaseStatus) String() string {
    return []string{"UNKNOWN", "ACTIVE", "PENDINGDELETE", "CLOSING", "CLOSED", "CLOSEDWITHERROR"}[i]
}
func ParseCaseStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_CASESTATUS, nil
        case "ACTIVE":
            return ACTIVE_CASESTATUS, nil
        case "PENDINGDELETE":
            return PENDINGDELETE_CASESTATUS, nil
        case "CLOSING":
            return CLOSING_CASESTATUS, nil
        case "CLOSED":
            return CLOSED_CASESTATUS, nil
        case "CLOSEDWITHERROR":
            return CLOSEDWITHERROR_CASESTATUS, nil
    }
    return 0, errors.New("Unknown CaseStatus value: " + v)
}
func SerializeCaseStatus(values []CaseStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
