package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type ManagedAppDataStorageLocation int

const (
    ONEDRIVEFORBUSINESS_MANAGEDAPPDATASTORAGELOCATION ManagedAppDataStorageLocation = iota
    SHAREPOINT_MANAGEDAPPDATASTORAGELOCATION
    BOX_MANAGEDAPPDATASTORAGELOCATION
    LOCALSTORAGE_MANAGEDAPPDATASTORAGELOCATION
)

func (i ManagedAppDataStorageLocation) String() string {
    return []string{"ONEDRIVEFORBUSINESS", "SHAREPOINT", "BOX", "LOCALSTORAGE"}[i]
}
func ParseManagedAppDataStorageLocation(v string) (interface{}, error) {
    result := ONEDRIVEFORBUSINESS_MANAGEDAPPDATASTORAGELOCATION
    switch strings.ToUpper(v) {
        case "ONEDRIVEFORBUSINESS":
            result = ONEDRIVEFORBUSINESS_MANAGEDAPPDATASTORAGELOCATION
        case "SHAREPOINT":
            result = SHAREPOINT_MANAGEDAPPDATASTORAGELOCATION
        case "BOX":
            result = BOX_MANAGEDAPPDATASTORAGELOCATION
        case "LOCALSTORAGE":
            result = LOCALSTORAGE_MANAGEDAPPDATASTORAGELOCATION
        default:
            return 0, errors.New("Unknown ManagedAppDataStorageLocation value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppDataStorageLocation(values []ManagedAppDataStorageLocation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
