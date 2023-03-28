package models
import (
    "errors"
)
// 
type AuthenticationMethodsPolicyMigrationState int

const (
    PREMIGRATION_AUTHENTICATIONMETHODSPOLICYMIGRATIONSTATE AuthenticationMethodsPolicyMigrationState = iota
    MIGRATIONINPROGRESS_AUTHENTICATIONMETHODSPOLICYMIGRATIONSTATE
    MIGRATIONCOMPLETE_AUTHENTICATIONMETHODSPOLICYMIGRATIONSTATE
    UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODSPOLICYMIGRATIONSTATE
)

func (i AuthenticationMethodsPolicyMigrationState) String() string {
    return []string{"preMigration", "migrationInProgress", "migrationComplete", "unknownFutureValue"}[i]
}
func ParseAuthenticationMethodsPolicyMigrationState(v string) (any, error) {
    result := PREMIGRATION_AUTHENTICATIONMETHODSPOLICYMIGRATIONSTATE
    switch v {
        case "preMigration":
            result = PREMIGRATION_AUTHENTICATIONMETHODSPOLICYMIGRATIONSTATE
        case "migrationInProgress":
            result = MIGRATIONINPROGRESS_AUTHENTICATIONMETHODSPOLICYMIGRATIONSTATE
        case "migrationComplete":
            result = MIGRATIONCOMPLETE_AUTHENTICATIONMETHODSPOLICYMIGRATIONSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODSPOLICYMIGRATIONSTATE
        default:
            return 0, errors.New("Unknown AuthenticationMethodsPolicyMigrationState value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationMethodsPolicyMigrationState(values []AuthenticationMethodsPolicyMigrationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
