package graph
import (
    "strings"
    "errors"
)
// 
type ProvisioningAction int

const (
    OTHER_PROVISIONINGACTION ProvisioningAction = iota
    CREATE_PROVISIONINGACTION
    DELETE_PROVISIONINGACTION
    DISABLE_PROVISIONINGACTION
    UPDATE_PROVISIONINGACTION
    STAGEDDELETE_PROVISIONINGACTION
    UNKNOWNFUTUREVALUE_PROVISIONINGACTION
)

func (i ProvisioningAction) String() string {
    return []string{"OTHER", "CREATE", "DELETE", "DISABLE", "UPDATE", "STAGEDDELETE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseProvisioningAction(v string) (interface{}, error) {
    result := OTHER_PROVISIONINGACTION
    switch strings.ToUpper(v) {
        case "OTHER":
            result = OTHER_PROVISIONINGACTION
        case "CREATE":
            result = CREATE_PROVISIONINGACTION
        case "DELETE":
            result = DELETE_PROVISIONINGACTION
        case "DISABLE":
            result = DISABLE_PROVISIONINGACTION
        case "UPDATE":
            result = UPDATE_PROVISIONINGACTION
        case "STAGEDDELETE":
            result = STAGEDDELETE_PROVISIONINGACTION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PROVISIONINGACTION
        default:
            return 0, errors.New("Unknown ProvisioningAction value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningAction(values []ProvisioningAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
