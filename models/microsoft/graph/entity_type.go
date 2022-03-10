package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the query method.
type EntityType int

const (
    EVENT_ENTITYTYPE EntityType = iota
    MESSAGE_ENTITYTYPE
    DRIVEITEM_ENTITYTYPE
    EXTERNALITEM_ENTITYTYPE
    SITE_ENTITYTYPE
    LIST_ENTITYTYPE
    LISTITEM_ENTITYTYPE
    DRIVE_ENTITYTYPE
    UNKNOWNFUTUREVALUE_ENTITYTYPE
)

func (i EntityType) String() string {
    return []string{"EVENT", "MESSAGE", "DRIVEITEM", "EXTERNALITEM", "SITE", "LIST", "LISTITEM", "DRIVE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEntityType(v string) (interface{}, error) {
    result := EVENT_ENTITYTYPE
    switch strings.ToUpper(v) {
        case "EVENT":
            result = EVENT_ENTITYTYPE
        case "MESSAGE":
            result = MESSAGE_ENTITYTYPE
        case "DRIVEITEM":
            result = DRIVEITEM_ENTITYTYPE
        case "EXTERNALITEM":
            result = EXTERNALITEM_ENTITYTYPE
        case "SITE":
            result = SITE_ENTITYTYPE
        case "LIST":
            result = LIST_ENTITYTYPE
        case "LISTITEM":
            result = LISTITEM_ENTITYTYPE
        case "DRIVE":
            result = DRIVE_ENTITYTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ENTITYTYPE
        default:
            return 0, errors.New("Unknown EntityType value: " + v)
    }
    return &result, nil
}
func SerializeEntityType(values []EntityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
