package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_ALLOWINVITESFROM, nil
        case "ADMINSANDGUESTINVITERS":
            return ADMINSANDGUESTINVITERS_ALLOWINVITESFROM, nil
        case "ADMINSGUESTINVITERSANDALLMEMBERS":
            return ADMINSGUESTINVITERSANDALLMEMBERS_ALLOWINVITESFROM, nil
        case "EVERYONE":
            return EVERYONE_ALLOWINVITESFROM, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ALLOWINVITESFROM, nil
    }
    return 0, errors.New("Unknown AllowInvitesFrom value: " + v)
}
func SerializeAllowInvitesFrom(values []AllowInvitesFrom) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
