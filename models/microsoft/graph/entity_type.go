package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "EVENT":
            return EVENT_ENTITYTYPE, nil
        case "MESSAGE":
            return MESSAGE_ENTITYTYPE, nil
        case "DRIVEITEM":
            return DRIVEITEM_ENTITYTYPE, nil
        case "EXTERNALITEM":
            return EXTERNALITEM_ENTITYTYPE, nil
        case "SITE":
            return SITE_ENTITYTYPE, nil
        case "LIST":
            return LIST_ENTITYTYPE, nil
        case "LISTITEM":
            return LISTITEM_ENTITYTYPE, nil
        case "DRIVE":
            return DRIVE_ENTITYTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ENTITYTYPE, nil
    }
    return 0, errors.New("Unknown EntityType value: " + v)
}
func SerializeEntityType(values []EntityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
