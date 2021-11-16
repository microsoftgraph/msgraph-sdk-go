package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION, nil
        case "BLOCKSIGNIN":
            return BLOCKSIGNIN_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION, nil
        case "BLOCKSIGNINANDDELETE":
            return BLOCKSIGNINANDDELETE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION, nil
    }
    return 0, errors.New("Unknown AccessPackageExternalUserLifecycleAction value: " + v)
}
func SerializeAccessPackageExternalUserLifecycleAction(values []AccessPackageExternalUserLifecycleAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
