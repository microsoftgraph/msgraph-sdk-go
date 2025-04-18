// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
import (
    "math"
    "strings"
)
type ConfirmedBy int

const (
    NONE_CONFIRMEDBY = 1
    USER_CONFIRMEDBY = 2
    MANAGER_CONFIRMEDBY = 4
    UNKNOWNFUTUREVALUE_CONFIRMEDBY = 8
)

func (i ConfirmedBy) String() string {
    var values []string
    options := []string{"none", "user", "manager", "unknownFutureValue"}
    for p := 0; p < 4; p++ {
        mantis := ConfirmedBy(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseConfirmedBy(v string) (any, error) {
    var result ConfirmedBy
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_CONFIRMEDBY
            case "user":
                result |= USER_CONFIRMEDBY
            case "manager":
                result |= MANAGER_CONFIRMEDBY
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CONFIRMEDBY
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeConfirmedBy(values []ConfirmedBy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfirmedBy) isMultiValue() bool {
    return true
}
