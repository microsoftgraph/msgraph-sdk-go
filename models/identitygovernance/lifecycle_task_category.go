package identitygovernance
import (
    "errors"
)
// 
type LifecycleTaskCategory int

const (
    JOINER_LIFECYCLETASKCATEGORY LifecycleTaskCategory = iota
    LEAVER_LIFECYCLETASKCATEGORY
    UNKNOWNFUTUREVALUE_LIFECYCLETASKCATEGORY
    MOVER_LIFECYCLETASKCATEGORY
)

func (i LifecycleTaskCategory) String() string {
    return []string{"joiner", "leaver", "unknownFutureValue", "mover"}[i]
}
func ParseLifecycleTaskCategory(v string) (any, error) {
    result := JOINER_LIFECYCLETASKCATEGORY
    switch v {
        case "joiner":
            result = JOINER_LIFECYCLETASKCATEGORY
        case "leaver":
            result = LEAVER_LIFECYCLETASKCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LIFECYCLETASKCATEGORY
        case "mover":
            result = MOVER_LIFECYCLETASKCATEGORY
        default:
            return 0, errors.New("Unknown LifecycleTaskCategory value: " + v)
    }
    return &result, nil
}
func SerializeLifecycleTaskCategory(values []LifecycleTaskCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
