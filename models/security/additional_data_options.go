package security
import (
    "errors"
    "strings"
)
// 
type AdditionalDataOptions int

const (
    ALLVERSIONS_ADDITIONALDATAOPTIONS AdditionalDataOptions = iota
    LINKEDFILES_ADDITIONALDATAOPTIONS
    UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS
)

func (i AdditionalDataOptions) String() string {
    var values []string
    for p := AdditionalDataOptions(1); p <= UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"allVersions", "linkedFiles", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAdditionalDataOptions(v string) (any, error) {
    var result AdditionalDataOptions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "allVersions":
                result |= ALLVERSIONS_ADDITIONALDATAOPTIONS
            case "linkedFiles":
                result |= LINKEDFILES_ADDITIONALDATAOPTIONS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS
            default:
                return 0, errors.New("Unknown AdditionalDataOptions value: " + v)
        }
    }
    return &result, nil
}
func SerializeAdditionalDataOptions(values []AdditionalDataOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AdditionalDataOptions) isMultiValue() bool {
    return true
}
