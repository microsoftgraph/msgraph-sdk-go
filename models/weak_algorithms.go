package models
import (
    "errors"
    "strings"
)
// 
type WeakAlgorithms int

const (
    RSASHA1_WEAKALGORITHMS WeakAlgorithms = iota
    UNKNOWNFUTUREVALUE_WEAKALGORITHMS
)

func (i WeakAlgorithms) String() string {
    var values []string
    for p := WeakAlgorithms(1); p <= UNKNOWNFUTUREVALUE_WEAKALGORITHMS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"rsaSha1", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseWeakAlgorithms(v string) (any, error) {
    var result WeakAlgorithms
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "rsaSha1":
                result |= RSASHA1_WEAKALGORITHMS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_WEAKALGORITHMS
            default:
                return 0, errors.New("Unknown WeakAlgorithms value: " + v)
        }
    }
    return &result, nil
}
func SerializeWeakAlgorithms(values []WeakAlgorithms) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WeakAlgorithms) isMultiValue() bool {
    return true
}
