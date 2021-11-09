package ediscovery
import (
    "strings"
    "errors"
)
// 
type LegalHoldStatus int

const (
    PENDING_LEGALHOLDSTATUS LegalHoldStatus = iota
    ERROR_LEGALHOLDSTATUS
    SUCCESS_LEGALHOLDSTATUS
    UNKNOWNFUTUREVALUE_LEGALHOLDSTATUS
)

func (i LegalHoldStatus) String() string {
    return []string{"PENDING", "ERROR", "SUCCESS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseLegalHoldStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "PENDING":
            return PENDING_LEGALHOLDSTATUS, nil
        case "ERROR":
            return ERROR_LEGALHOLDSTATUS, nil
        case "SUCCESS":
            return SUCCESS_LEGALHOLDSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_LEGALHOLDSTATUS, nil
    }
    return 0, errors.New("Unknown LegalHoldStatus value: " + v)
}
func SerializeLegalHoldStatus(values []LegalHoldStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
