package identitygovernance
import (
    "errors"
    "strings"
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
    var values []string
    for p := LifecycleTaskCategory(1); p <= MOVER_LIFECYCLETASKCATEGORY; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"joiner", "leaver", "unknownFutureValue", "mover"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseLifecycleTaskCategory(v string) (any, error) {
    var result LifecycleTaskCategory
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "joiner":
                result |= JOINER_LIFECYCLETASKCATEGORY
            case "leaver":
                result |= LEAVER_LIFECYCLETASKCATEGORY
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_LIFECYCLETASKCATEGORY
            case "mover":
                result |= MOVER_LIFECYCLETASKCATEGORY
            default:
                return 0, errors.New("Unknown LifecycleTaskCategory value: " + v)
        }
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
func (i LifecycleTaskCategory) isMultiValue() bool {
    return true
}
