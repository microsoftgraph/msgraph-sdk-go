package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of chat entities.
type TeamworkUserIdentityType int

const (
    AADUSER_TEAMWORKUSERIDENTITYTYPE TeamworkUserIdentityType = iota
    ONPREMISEAADUSER_TEAMWORKUSERIDENTITYTYPE
    ANONYMOUSGUEST_TEAMWORKUSERIDENTITYTYPE
    FEDERATEDUSER_TEAMWORKUSERIDENTITYTYPE
    PERSONALMICROSOFTACCOUNTUSER_TEAMWORKUSERIDENTITYTYPE
    SKYPEUSER_TEAMWORKUSERIDENTITYTYPE
    PHONEUSER_TEAMWORKUSERIDENTITYTYPE
    UNKNOWNFUTUREVALUE_TEAMWORKUSERIDENTITYTYPE
    EMAILUSER_TEAMWORKUSERIDENTITYTYPE
)

func (i TeamworkUserIdentityType) String() string {
    return []string{"AADUSER", "ONPREMISEAADUSER", "ANONYMOUSGUEST", "FEDERATEDUSER", "PERSONALMICROSOFTACCOUNTUSER", "SKYPEUSER", "PHONEUSER", "UNKNOWNFUTUREVALUE", "EMAILUSER"}[i]
}
func ParseTeamworkUserIdentityType(v string) (interface{}, error) {
    result := AADUSER_TEAMWORKUSERIDENTITYTYPE
    switch strings.ToUpper(v) {
        case "AADUSER":
            result = AADUSER_TEAMWORKUSERIDENTITYTYPE
        case "ONPREMISEAADUSER":
            result = ONPREMISEAADUSER_TEAMWORKUSERIDENTITYTYPE
        case "ANONYMOUSGUEST":
            result = ANONYMOUSGUEST_TEAMWORKUSERIDENTITYTYPE
        case "FEDERATEDUSER":
            result = FEDERATEDUSER_TEAMWORKUSERIDENTITYTYPE
        case "PERSONALMICROSOFTACCOUNTUSER":
            result = PERSONALMICROSOFTACCOUNTUSER_TEAMWORKUSERIDENTITYTYPE
        case "SKYPEUSER":
            result = SKYPEUSER_TEAMWORKUSERIDENTITYTYPE
        case "PHONEUSER":
            result = PHONEUSER_TEAMWORKUSERIDENTITYTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMWORKUSERIDENTITYTYPE
        case "EMAILUSER":
            result = EMAILUSER_TEAMWORKUSERIDENTITYTYPE
        default:
            return 0, errors.New("Unknown TeamworkUserIdentityType value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkUserIdentityType(values []TeamworkUserIdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
