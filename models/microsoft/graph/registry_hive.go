package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_REGISTRYHIVE, nil
        case "CURRENTCONFIG":
            return CURRENTCONFIG_REGISTRYHIVE, nil
        case "CURRENTUSER":
            return CURRENTUSER_REGISTRYHIVE, nil
        case "LOCALMACHINESAM":
            return LOCALMACHINESAM_REGISTRYHIVE, nil
        case "LOCALMACHINESECURITY":
            return LOCALMACHINESECURITY_REGISTRYHIVE, nil
        case "LOCALMACHINESOFTWARE":
            return LOCALMACHINESOFTWARE_REGISTRYHIVE, nil
        case "LOCALMACHINESYSTEM":
            return LOCALMACHINESYSTEM_REGISTRYHIVE, nil
        case "USERSDEFAULT":
            return USERSDEFAULT_REGISTRYHIVE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_REGISTRYHIVE, nil
    }
    return 0, errors.New("Unknown RegistryHive value: " + v)
}
func SerializeRegistryHive(values []RegistryHive) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
