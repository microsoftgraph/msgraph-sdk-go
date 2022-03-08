package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the security singleton.
type RegistryHive int

const (
    UNKNOWN_REGISTRYHIVE RegistryHive = iota
    CURRENTCONFIG_REGISTRYHIVE
    CURRENTUSER_REGISTRYHIVE
    LOCALMACHINESAM_REGISTRYHIVE
    LOCALMACHINESECURITY_REGISTRYHIVE
    LOCALMACHINESOFTWARE_REGISTRYHIVE
    LOCALMACHINESYSTEM_REGISTRYHIVE
    USERSDEFAULT_REGISTRYHIVE
    UNKNOWNFUTUREVALUE_REGISTRYHIVE
)

func (i RegistryHive) String() string {
    return []string{"UNKNOWN", "CURRENTCONFIG", "CURRENTUSER", "LOCALMACHINESAM", "LOCALMACHINESECURITY", "LOCALMACHINESOFTWARE", "LOCALMACHINESYSTEM", "USERSDEFAULT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRegistryHive(v string) (interface{}, error) {
    result := UNKNOWN_REGISTRYHIVE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_REGISTRYHIVE
        case "CURRENTCONFIG":
            result = CURRENTCONFIG_REGISTRYHIVE
        case "CURRENTUSER":
            result = CURRENTUSER_REGISTRYHIVE
        case "LOCALMACHINESAM":
            result = LOCALMACHINESAM_REGISTRYHIVE
        case "LOCALMACHINESECURITY":
            result = LOCALMACHINESECURITY_REGISTRYHIVE
        case "LOCALMACHINESOFTWARE":
            result = LOCALMACHINESOFTWARE_REGISTRYHIVE
        case "LOCALMACHINESYSTEM":
            result = LOCALMACHINESYSTEM_REGISTRYHIVE
        case "USERSDEFAULT":
            result = USERSDEFAULT_REGISTRYHIVE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_REGISTRYHIVE
        default:
            return 0, errors.New("Unknown RegistryHive value: " + v)
    }
    return &result, nil
}
func SerializeRegistryHive(values []RegistryHive) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
