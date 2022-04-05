package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityProtectionRoot singleton.
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
    result := AZUREAD_TOKENISSUERTYPE
    switch strings.ToUpper(v) {
        case "AZUREAD":
            result = AZUREAD_TOKENISSUERTYPE
        case "ADFEDERATIONSERVICES":
            result = ADFEDERATIONSERVICES_TOKENISSUERTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TOKENISSUERTYPE
        case "AZUREADBACKUPAUTH":
            result = AZUREADBACKUPAUTH_TOKENISSUERTYPE
        case "ADFEDERATIONSERVICESMFAADAPTER":
            result = ADFEDERATIONSERVICESMFAADAPTER_TOKENISSUERTYPE
        case "NPSEXTENSION":
            result = NPSEXTENSION_TOKENISSUERTYPE
        default:
            return 0, errors.New("Unknown TokenIssuerType value: " + v)
    }
    return &result, nil
}
func SerializeTokenIssuerType(values []TokenIssuerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
