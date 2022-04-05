package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type LobbyBypassScope int

const (
    ORGANIZER_LOBBYBYPASSSCOPE LobbyBypassScope = iota
    ORGANIZATION_LOBBYBYPASSSCOPE
    ORGANIZATIONANDFEDERATED_LOBBYBYPASSSCOPE
    EVERYONE_LOBBYBYPASSSCOPE
    UNKNOWNFUTUREVALUE_LOBBYBYPASSSCOPE
    INVITED_LOBBYBYPASSSCOPE
    ORGANIZATIONEXCLUDINGGUESTS_LOBBYBYPASSSCOPE
)

func (i LobbyBypassScope) String() string {
    return []string{"ORGANIZER", "ORGANIZATION", "ORGANIZATIONANDFEDERATED", "EVERYONE", "UNKNOWNFUTUREVALUE", "INVITED", "ORGANIZATIONEXCLUDINGGUESTS"}[i]
}
func ParseLobbyBypassScope(v string) (interface{}, error) {
    result := ORGANIZER_LOBBYBYPASSSCOPE
    switch strings.ToUpper(v) {
        case "ORGANIZER":
            result = ORGANIZER_LOBBYBYPASSSCOPE
        case "ORGANIZATION":
            result = ORGANIZATION_LOBBYBYPASSSCOPE
        case "ORGANIZATIONANDFEDERATED":
            result = ORGANIZATIONANDFEDERATED_LOBBYBYPASSSCOPE
        case "EVERYONE":
            result = EVERYONE_LOBBYBYPASSSCOPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_LOBBYBYPASSSCOPE
        case "INVITED":
            result = INVITED_LOBBYBYPASSSCOPE
        case "ORGANIZATIONEXCLUDINGGUESTS":
            result = ORGANIZATIONEXCLUDINGGUESTS_LOBBYBYPASSSCOPE
        default:
            return 0, errors.New("Unknown LobbyBypassScope value: " + v)
    }
    return &result, nil
}
func SerializeLobbyBypassScope(values []LobbyBypassScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
