package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityContainer singleton.
type ConditionalAccessGrantControl int

const (
    BLOCK_CONDITIONALACCESSGRANTCONTROL ConditionalAccessGrantControl = iota
    MFA_CONDITIONALACCESSGRANTCONTROL
    COMPLIANTDEVICE_CONDITIONALACCESSGRANTCONTROL
    DOMAINJOINEDDEVICE_CONDITIONALACCESSGRANTCONTROL
    APPROVEDAPPLICATION_CONDITIONALACCESSGRANTCONTROL
    COMPLIANTAPPLICATION_CONDITIONALACCESSGRANTCONTROL
    PASSWORDCHANGE_CONDITIONALACCESSGRANTCONTROL
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSGRANTCONTROL
)

func (i ConditionalAccessGrantControl) String() string {
    return []string{"BLOCK", "MFA", "COMPLIANTDEVICE", "DOMAINJOINEDDEVICE", "APPROVEDAPPLICATION", "COMPLIANTAPPLICATION", "PASSWORDCHANGE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConditionalAccessGrantControl(v string) (interface{}, error) {
    result := BLOCK_CONDITIONALACCESSGRANTCONTROL
    switch strings.ToUpper(v) {
        case "BLOCK":
            result = BLOCK_CONDITIONALACCESSGRANTCONTROL
        case "MFA":
            result = MFA_CONDITIONALACCESSGRANTCONTROL
        case "COMPLIANTDEVICE":
            result = COMPLIANTDEVICE_CONDITIONALACCESSGRANTCONTROL
        case "DOMAINJOINEDDEVICE":
            result = DOMAINJOINEDDEVICE_CONDITIONALACCESSGRANTCONTROL
        case "APPROVEDAPPLICATION":
            result = APPROVEDAPPLICATION_CONDITIONALACCESSGRANTCONTROL
        case "COMPLIANTAPPLICATION":
            result = COMPLIANTAPPLICATION_CONDITIONALACCESSGRANTCONTROL
        case "PASSWORDCHANGE":
            result = PASSWORDCHANGE_CONDITIONALACCESSGRANTCONTROL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONDITIONALACCESSGRANTCONTROL
        default:
            return 0, errors.New("Unknown ConditionalAccessGrantControl value: " + v)
    }
    return &result, nil
}
func SerializeConditionalAccessGrantControl(values []ConditionalAccessGrantControl) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
