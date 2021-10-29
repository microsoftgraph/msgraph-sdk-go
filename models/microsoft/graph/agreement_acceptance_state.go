package graph
import (
    "strings"
    "errors"
)
// 
type AgreementAcceptanceState int

const (
    ACCEPTED_AGREEMENTACCEPTANCESTATE AgreementAcceptanceState = iota
    DECLINED_AGREEMENTACCEPTANCESTATE
    UNKNOWNFUTUREVALUE_AGREEMENTACCEPTANCESTATE
)

func (i AgreementAcceptanceState) String() string {
    return []string{"ACCEPTED", "DECLINED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAgreementAcceptanceState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ACCEPTED":
            return ACCEPTED_AGREEMENTACCEPTANCESTATE, nil
        case "DECLINED":
            return DECLINED_AGREEMENTACCEPTANCESTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_AGREEMENTACCEPTANCESTATE, nil
    }
    return 0, errors.New("Unknown AgreementAcceptanceState value: " + v)
}
func SerializeAgreementAcceptanceState(values []AgreementAcceptanceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
