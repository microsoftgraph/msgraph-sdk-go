package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
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
    result := USEDEVICESETTINGS_MANAGEDAPPDATAENCRYPTIONTYPE
    switch strings.ToUpper(v) {
        case "USEDEVICESETTINGS":
            result = USEDEVICESETTINGS_MANAGEDAPPDATAENCRYPTIONTYPE
        case "AFTERDEVICERESTART":
            result = AFTERDEVICERESTART_MANAGEDAPPDATAENCRYPTIONTYPE
        case "WHENDEVICELOCKEDEXCEPTOPENFILES":
            result = WHENDEVICELOCKEDEXCEPTOPENFILES_MANAGEDAPPDATAENCRYPTIONTYPE
        case "WHENDEVICELOCKED":
            result = WHENDEVICELOCKED_MANAGEDAPPDATAENCRYPTIONTYPE
        default:
            return 0, errors.New("Unknown ManagedAppDataEncryptionType value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppDataEncryptionType(values []ManagedAppDataEncryptionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
