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
    switch strings.ToUpper(v) {
        case "OTHER":
            return OTHER_PROVISIONINGACTION, nil
        case "CREATE":
            return CREATE_PROVISIONINGACTION, nil
        case "DELETE":
            return DELETE_PROVISIONINGACTION, nil
        case "DISABLE":
            return DISABLE_PROVISIONINGACTION, nil
        case "UPDATE":
            return UPDATE_PROVISIONINGACTION, nil
        case "STAGEDDELETE":
            return STAGEDDELETE_PROVISIONINGACTION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PROVISIONINGACTION, nil
    }
    return 0, errors.New("Unknown ProvisioningAction value: " + v)
}
func SerializeProvisioningAction(values []ProvisioningAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
