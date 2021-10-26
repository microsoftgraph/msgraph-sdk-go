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
    switch strings.ToUpper(v) {
        case "HOME":
            return HOME_PHONETYPE, nil
        case "BUSINESS":
            return BUSINESS_PHONETYPE, nil
        case "MOBILE":
            return MOBILE_PHONETYPE, nil
        case "OTHER":
            return OTHER_PHONETYPE, nil
        case "ASSISTANT":
            return ASSISTANT_PHONETYPE, nil
        case "HOMEFAX":
            return HOMEFAX_PHONETYPE, nil
        case "BUSINESSFAX":
            return BUSINESSFAX_PHONETYPE, nil
        case "OTHERFAX":
            return OTHERFAX_PHONETYPE, nil
        case "PAGER":
            return PAGER_PHONETYPE, nil
        case "RADIO":
            return RADIO_PHONETYPE, nil
    }
    return 0, errors.New("Unknown PhoneType value: " + v)
}
func SerializePhoneType(values []PhoneType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
