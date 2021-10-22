package graph
import (
    "strings"
    "errors"
)
type ManagedAppDataStorageLocation int

const (
    ONEDRIVEFORBUSINESS_MANAGEDAPPDATASTORAGELOCATION ManagedAppDataStorageLocation = iota
    SHAREPOINT_MANAGEDAPPDATASTORAGELOCATION
    LOCALSTORAGE_MANAGEDAPPDATASTORAGELOCATION
)

func (i ManagedAppDataStorageLocation) String() string {
    return []string{"ONEDRIVEFORBUSINESS", "SHAREPOINT", "LOCALSTORAGE"}[i]
}
func ParseManagedAppDataStorageLocation(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ONEDRIVEFORBUSINESS":
            return ONEDRIVEFORBUSINESS_MANAGEDAPPDATASTORAGELOCATION, nil
        case "SHAREPOINT":
            return SHAREPOINT_MANAGEDAPPDATASTORAGELOCATION, nil
        case "LOCALSTORAGE":
            return LOCALSTORAGE_MANAGEDAPPDATASTORAGELOCATION, nil
    }
    return 0, errors.New("Unknown ManagedAppDataStorageLocation value: " + v)
}
func SerializeManagedAppDataStorageLocation(values []ManagedAppDataStorageLocation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
