package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type AccessPackageExternalUserLifecycleAction int

const (
    NONE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION AccessPackageExternalUserLifecycleAction = iota
    BLOCKSIGNIN_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
    BLOCKSIGNINANDDELETE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
    UNKNOWNFUTUREVALUE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
)

func (i AccessPackageExternalUserLifecycleAction) String() string {
    return []string{"NONE", "BLOCKSIGNIN", "BLOCKSIGNINANDDELETE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessPackageExternalUserLifecycleAction(v string) (interface{}, error) {
    result := NONE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
        case "BLOCKSIGNIN":
            result = BLOCKSIGNIN_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
        case "BLOCKSIGNINANDDELETE":
            result = BLOCKSIGNINANDDELETE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
        default:
            return 0, errors.New("Unknown AccessPackageExternalUserLifecycleAction value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageExternalUserLifecycleAction(values []AccessPackageExternalUserLifecycleAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
