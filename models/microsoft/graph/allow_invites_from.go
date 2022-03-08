package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the policyRoot singleton.
type AllowInvitesFrom int

const (
    NONE_ALLOWINVITESFROM AllowInvitesFrom = iota
    ADMINSANDGUESTINVITERS_ALLOWINVITESFROM
    ADMINSGUESTINVITERSANDALLMEMBERS_ALLOWINVITESFROM
    EVERYONE_ALLOWINVITESFROM
    UNKNOWNFUTUREVALUE_ALLOWINVITESFROM
)

func (i AllowInvitesFrom) String() string {
    return []string{"NONE", "ADMINSANDGUESTINVITERS", "ADMINSGUESTINVITERSANDALLMEMBERS", "EVERYONE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAllowInvitesFrom(v string) (interface{}, error) {
    result := NONE_ALLOWINVITESFROM
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ALLOWINVITESFROM
        case "ADMINSANDGUESTINVITERS":
            result = ADMINSANDGUESTINVITERS_ALLOWINVITESFROM
        case "ADMINSGUESTINVITERSANDALLMEMBERS":
            result = ADMINSGUESTINVITERSANDALLMEMBERS_ALLOWINVITESFROM
        case "EVERYONE":
            result = EVERYONE_ALLOWINVITESFROM
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ALLOWINVITESFROM
        default:
            return 0, errors.New("Unknown AllowInvitesFrom value: " + v)
    }
    return &result, nil
}
func SerializeAllowInvitesFrom(values []AllowInvitesFrom) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
