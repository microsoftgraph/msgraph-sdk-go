package graph
import (
    "strings"
    "errors"
)
// 
type TokenIssuerType int

const (
    AZUREAD_TOKENISSUERTYPE TokenIssuerType = iota
    ADFEDERATIONSERVICES_TOKENISSUERTYPE
    UNKNOWNFUTUREVALUE_TOKENISSUERTYPE
    AZUREADBACKUPAUTH_TOKENISSUERTYPE
    ADFEDERATIONSERVICESMFAADAPTER_TOKENISSUERTYPE
    NPSEXTENSION_TOKENISSUERTYPE
)

func (i TokenIssuerType) String() string {
    return []string{"AZUREAD", "ADFEDERATIONSERVICES", "UNKNOWNFUTUREVALUE", "AZUREADBACKUPAUTH", "ADFEDERATIONSERVICESMFAADAPTER", "NPSEXTENSION"}[i]
}
func ParseTokenIssuerType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "AZUREAD":
            return AZUREAD_TOKENISSUERTYPE, nil
        case "ADFEDERATIONSERVICES":
            return ADFEDERATIONSERVICES_TOKENISSUERTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TOKENISSUERTYPE, nil
        case "AZUREADBACKUPAUTH":
            return AZUREADBACKUPAUTH_TOKENISSUERTYPE, nil
        case "ADFEDERATIONSERVICESMFAADAPTER":
            return ADFEDERATIONSERVICESMFAADAPTER_TOKENISSUERTYPE, nil
        case "NPSEXTENSION":
            return NPSEXTENSION_TOKENISSUERTYPE, nil
    }
    return 0, errors.New("Unknown TokenIssuerType value: " + v)
}
func SerializeTokenIssuerType(values []TokenIssuerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
