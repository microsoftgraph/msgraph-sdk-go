package graph
import (
    "strings"
    "errors"
)
type ManagedAppDataEncryptionType int

const (
    USEDEVICESETTINGS_MANAGEDAPPDATAENCRYPTIONTYPE ManagedAppDataEncryptionType = iota
    AFTERDEVICERESTART_MANAGEDAPPDATAENCRYPTIONTYPE
    WHENDEVICELOCKEDEXCEPTOPENFILES_MANAGEDAPPDATAENCRYPTIONTYPE
    WHENDEVICELOCKED_MANAGEDAPPDATAENCRYPTIONTYPE
)

func (i ManagedAppDataEncryptionType) String() string {
    return []string{"USEDEVICESETTINGS", "AFTERDEVICERESTART", "WHENDEVICELOCKEDEXCEPTOPENFILES", "WHENDEVICELOCKED"}[i]
}
func ParseManagedAppDataEncryptionType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "USEDEVICESETTINGS":
            return USEDEVICESETTINGS_MANAGEDAPPDATAENCRYPTIONTYPE, nil
        case "AFTERDEVICERESTART":
            return AFTERDEVICERESTART_MANAGEDAPPDATAENCRYPTIONTYPE, nil
        case "WHENDEVICELOCKEDEXCEPTOPENFILES":
            return WHENDEVICELOCKEDEXCEPTOPENFILES_MANAGEDAPPDATAENCRYPTIONTYPE, nil
        case "WHENDEVICELOCKED":
            return WHENDEVICELOCKED_MANAGEDAPPDATAENCRYPTIONTYPE, nil
    }
    return 0, errors.New("Unknown ManagedAppDataEncryptionType value: " + v)
}
func SerializeManagedAppDataEncryptionType(values []ManagedAppDataEncryptionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
