package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
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
    result := ACCEPTED_AGREEMENTACCEPTANCESTATE
    switch strings.ToUpper(v) {
        case "ACCEPTED":
            result = ACCEPTED_AGREEMENTACCEPTANCESTATE
        case "DECLINED":
            result = DECLINED_AGREEMENTACCEPTANCESTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_AGREEMENTACCEPTANCESTATE
        default:
            return 0, errors.New("Unknown AgreementAcceptanceState value: " + v)
    }
    return &result, nil
}
func SerializeAgreementAcceptanceState(values []AgreementAcceptanceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
