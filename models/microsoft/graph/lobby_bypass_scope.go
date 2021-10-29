package graph
import (
    "strings"
    "errors"
)
// 
type LobbyBypassScope int

const (
    ORGANIZER_LOBBYBYPASSSCOPE LobbyBypassScope = iota
    ORGANIZATION_LOBBYBYPASSSCOPE
    ORGANIZATIONANDFEDERATED_LOBBYBYPASSSCOPE
    EVERYONE_LOBBYBYPASSSCOPE
    UNKNOWNFUTUREVALUE_LOBBYBYPASSSCOPE
)

func (i LobbyBypassScope) String() string {
    return []string{"ORGANIZER", "ORGANIZATION", "ORGANIZATIONANDFEDERATED", "EVERYONE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseLobbyBypassScope(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ORGANIZER":
            return ORGANIZER_LOBBYBYPASSSCOPE, nil
        case "ORGANIZATION":
            return ORGANIZATION_LOBBYBYPASSSCOPE, nil
        case "ORGANIZATIONANDFEDERATED":
            return ORGANIZATIONANDFEDERATED_LOBBYBYPASSSCOPE, nil
        case "EVERYONE":
            return EVERYONE_LOBBYBYPASSSCOPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_LOBBYBYPASSSCOPE, nil
    }
    return 0, errors.New("Unknown LobbyBypassScope value: " + v)
}
func SerializeLobbyBypassScope(values []LobbyBypassScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
