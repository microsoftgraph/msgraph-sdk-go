package models
import (
    "errors"
    "strings"
)
// 
type ConditionalAccessGuestOrExternalUserTypes int

const (
    NONE_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES ConditionalAccessGuestOrExternalUserTypes = iota
    INTERNALGUEST_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
    B2BCOLLABORATIONGUEST_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
    B2BCOLLABORATIONMEMBER_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
    B2BDIRECTCONNECTUSER_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
    OTHEREXTERNALUSER_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
    SERVICEPROVIDER_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
)

func (i ConditionalAccessGuestOrExternalUserTypes) String() string {
    var values []string
    for p := ConditionalAccessGuestOrExternalUserTypes(1); p <= UNKNOWNFUTUREVALUE_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "internalGuest", "b2bCollaborationGuest", "b2bCollaborationMember", "b2bDirectConnectUser", "otherExternalUser", "serviceProvider", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseConditionalAccessGuestOrExternalUserTypes(v string) (any, error) {
    var result ConditionalAccessGuestOrExternalUserTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
            case "internalGuest":
                result |= INTERNALGUEST_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
            case "b2bCollaborationGuest":
                result |= B2BCOLLABORATIONGUEST_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
            case "b2bCollaborationMember":
                result |= B2BCOLLABORATIONMEMBER_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
            case "b2bDirectConnectUser":
                result |= B2BDIRECTCONNECTUSER_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
            case "otherExternalUser":
                result |= OTHEREXTERNALUSER_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
            case "serviceProvider":
                result |= SERVICEPROVIDER_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CONDITIONALACCESSGUESTOREXTERNALUSERTYPES
            default:
                return 0, errors.New("Unknown ConditionalAccessGuestOrExternalUserTypes value: " + v)
        }
    }
    return &result, nil
}
func SerializeConditionalAccessGuestOrExternalUserTypes(values []ConditionalAccessGuestOrExternalUserTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConditionalAccessGuestOrExternalUserTypes) isMultiValue() bool {
    return true
}
