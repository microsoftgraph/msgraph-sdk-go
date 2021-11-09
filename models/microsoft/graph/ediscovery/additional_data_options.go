package ediscovery
import (
    "strings"
    "errors"
)
// 
type AdditionalDataOptions int

const (
    ALLVERSIONS_ADDITIONALDATAOPTIONS AdditionalDataOptions = iota
    LINKEDFILES_ADDITIONALDATAOPTIONS
    UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS
)

func (i AdditionalDataOptions) String() string {
    return []string{"ALLVERSIONS", "LINKEDFILES", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAdditionalDataOptions(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ALLVERSIONS":
            return ALLVERSIONS_ADDITIONALDATAOPTIONS, nil
        case "LINKEDFILES":
            return LINKEDFILES_ADDITIONALDATAOPTIONS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS, nil
    }
    return 0, errors.New("Unknown AdditionalDataOptions value: " + v)
}
func SerializeAdditionalDataOptions(values []AdditionalDataOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
