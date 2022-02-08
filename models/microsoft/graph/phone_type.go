package graph
import (
    "strings"
    "errors"
)
// 
type PhoneType int

const (
    HOME_PHONETYPE PhoneType = iota
    BUSINESS_PHONETYPE
    MOBILE_PHONETYPE
    OTHER_PHONETYPE
    ASSISTANT_PHONETYPE
    HOMEFAX_PHONETYPE
    BUSINESSFAX_PHONETYPE
    OTHERFAX_PHONETYPE
    PAGER_PHONETYPE
    RADIO_PHONETYPE
)

func (i PhoneType) String() string {
    return []string{"HOME", "BUSINESS", "MOBILE", "OTHER", "ASSISTANT", "HOMEFAX", "BUSINESSFAX", "OTHERFAX", "PAGER", "RADIO"}[i]
}
func ParsePhoneType(v string) (interface{}, error) {
    result := HOME_PHONETYPE
    switch strings.ToUpper(v) {
        case "HOME":
            result = HOME_PHONETYPE
        case "BUSINESS":
            result = BUSINESS_PHONETYPE
        case "MOBILE":
            result = MOBILE_PHONETYPE
        case "OTHER":
            result = OTHER_PHONETYPE
        case "ASSISTANT":
            result = ASSISTANT_PHONETYPE
        case "HOMEFAX":
            result = HOMEFAX_PHONETYPE
        case "BUSINESSFAX":
            result = BUSINESSFAX_PHONETYPE
        case "OTHERFAX":
            result = OTHERFAX_PHONETYPE
        case "PAGER":
            result = PAGER_PHONETYPE
        case "RADIO":
            result = RADIO_PHONETYPE
        default:
            return 0, errors.New("Unknown PhoneType value: " + v)
    }
    return &result, nil
}
func SerializePhoneType(values []PhoneType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
