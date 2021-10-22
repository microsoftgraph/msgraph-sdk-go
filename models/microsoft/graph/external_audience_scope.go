package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_EXTERNALAUDIENCESCOPE, nil
        case "CONTACTSONLY":
            return CONTACTSONLY_EXTERNALAUDIENCESCOPE, nil
        case "ALL":
            return ALL_EXTERNALAUDIENCESCOPE, nil
    }
    return 0, errors.New("Unknown ExternalAudienceScope value: " + v)
}
func SerializeExternalAudienceScope(values []ExternalAudienceScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
