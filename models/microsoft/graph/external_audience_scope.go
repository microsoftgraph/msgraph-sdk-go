package graph
import (
    "strings"
    "errors"
)
// 
type ExternalAudienceScope int

const (
    NONE_EXTERNALAUDIENCESCOPE ExternalAudienceScope = iota
    CONTACTSONLY_EXTERNALAUDIENCESCOPE
    ALL_EXTERNALAUDIENCESCOPE
)

func (i ExternalAudienceScope) String() string {
    return []string{"NONE", "CONTACTSONLY", "ALL"}[i]
}
func ParseExternalAudienceScope(v string) (interface{}, error) {
    result := NONE_EXTERNALAUDIENCESCOPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_EXTERNALAUDIENCESCOPE
        case "CONTACTSONLY":
            result = CONTACTSONLY_EXTERNALAUDIENCESCOPE
        case "ALL":
            result = ALL_EXTERNALAUDIENCESCOPE
        default:
            return 0, errors.New("Unknown ExternalAudienceScope value: " + v)
    }
    return &result, nil
}
func SerializeExternalAudienceScope(values []ExternalAudienceScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
