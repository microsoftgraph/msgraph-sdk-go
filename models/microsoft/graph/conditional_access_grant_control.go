package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "BLOCK":
            return BLOCK_CONDITIONALACCESSGRANTCONTROL, nil
        case "MFA":
            return MFA_CONDITIONALACCESSGRANTCONTROL, nil
        case "COMPLIANTDEVICE":
            return COMPLIANTDEVICE_CONDITIONALACCESSGRANTCONTROL, nil
        case "DOMAINJOINEDDEVICE":
            return DOMAINJOINEDDEVICE_CONDITIONALACCESSGRANTCONTROL, nil
        case "APPROVEDAPPLICATION":
            return APPROVEDAPPLICATION_CONDITIONALACCESSGRANTCONTROL, nil
        case "COMPLIANTAPPLICATION":
            return COMPLIANTAPPLICATION_CONDITIONALACCESSGRANTCONTROL, nil
        case "PASSWORDCHANGE":
            return PASSWORDCHANGE_CONDITIONALACCESSGRANTCONTROL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONDITIONALACCESSGRANTCONTROL, nil
    }
    return 0, errors.New("Unknown ConditionalAccessGrantControl value: " + v)
}
func SerializeConditionalAccessGrantControl(values []ConditionalAccessGrantControl) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
