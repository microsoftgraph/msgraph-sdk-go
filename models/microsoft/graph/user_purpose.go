package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type UserPurpose int

const (
    USER_USERPURPOSE UserPurpose = iota
    LINKED_USERPURPOSE
    SHARED_USERPURPOSE
    ROOM_USERPURPOSE
    EQUIPMENT_USERPURPOSE
    OTHERS_USERPURPOSE
    UNKNOWNFUTUREVALUE_USERPURPOSE
)

func (i UserPurpose) String() string {
    return []string{"USER", "LINKED", "SHARED", "ROOM", "EQUIPMENT", "OTHERS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUserPurpose(v string) (interface{}, error) {
    result := USER_USERPURPOSE
    switch strings.ToUpper(v) {
        case "USER":
            result = USER_USERPURPOSE
        case "LINKED":
            result = LINKED_USERPURPOSE
        case "SHARED":
            result = SHARED_USERPURPOSE
        case "ROOM":
            result = ROOM_USERPURPOSE
        case "EQUIPMENT":
            result = EQUIPMENT_USERPURPOSE
        case "OTHERS":
            result = OTHERS_USERPURPOSE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_USERPURPOSE
        default:
            return 0, errors.New("Unknown UserPurpose value: " + v)
    }
    return &result, nil
}
func SerializeUserPurpose(values []UserPurpose) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
