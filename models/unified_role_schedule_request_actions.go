package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the roleManagement singleton.
type UnifiedRoleScheduleRequestActions int

const (
    ADMINASSIGN_UNIFIEDROLESCHEDULEREQUESTACTIONS UnifiedRoleScheduleRequestActions = iota
    ADMINUPDATE_UNIFIEDROLESCHEDULEREQUESTACTIONS
    ADMINREMOVE_UNIFIEDROLESCHEDULEREQUESTACTIONS
    SELFACTIVATE_UNIFIEDROLESCHEDULEREQUESTACTIONS
    SELFDEACTIVATE_UNIFIEDROLESCHEDULEREQUESTACTIONS
    ADMINEXTEND_UNIFIEDROLESCHEDULEREQUESTACTIONS
    ADMINRENEW_UNIFIEDROLESCHEDULEREQUESTACTIONS
    SELFEXTEND_UNIFIEDROLESCHEDULEREQUESTACTIONS
    SELFRENEW_UNIFIEDROLESCHEDULEREQUESTACTIONS
    UNKNOWNFUTUREVALUE_UNIFIEDROLESCHEDULEREQUESTACTIONS
)

func (i UnifiedRoleScheduleRequestActions) String() string {
    return []string{"ADMINASSIGN", "ADMINUPDATE", "ADMINREMOVE", "SELFACTIVATE", "SELFDEACTIVATE", "ADMINEXTEND", "ADMINRENEW", "SELFEXTEND", "SELFRENEW", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUnifiedRoleScheduleRequestActions(v string) (interface{}, error) {
    result := ADMINASSIGN_UNIFIEDROLESCHEDULEREQUESTACTIONS
    switch strings.ToUpper(v) {
        case "ADMINASSIGN":
            result = ADMINASSIGN_UNIFIEDROLESCHEDULEREQUESTACTIONS
        case "ADMINUPDATE":
            result = ADMINUPDATE_UNIFIEDROLESCHEDULEREQUESTACTIONS
        case "ADMINREMOVE":
            result = ADMINREMOVE_UNIFIEDROLESCHEDULEREQUESTACTIONS
        case "SELFACTIVATE":
            result = SELFACTIVATE_UNIFIEDROLESCHEDULEREQUESTACTIONS
        case "SELFDEACTIVATE":
            result = SELFDEACTIVATE_UNIFIEDROLESCHEDULEREQUESTACTIONS
        case "ADMINEXTEND":
            result = ADMINEXTEND_UNIFIEDROLESCHEDULEREQUESTACTIONS
        case "ADMINRENEW":
            result = ADMINRENEW_UNIFIEDROLESCHEDULEREQUESTACTIONS
        case "SELFEXTEND":
            result = SELFEXTEND_UNIFIEDROLESCHEDULEREQUESTACTIONS
        case "SELFRENEW":
            result = SELFRENEW_UNIFIEDROLESCHEDULEREQUESTACTIONS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_UNIFIEDROLESCHEDULEREQUESTACTIONS
        default:
            return 0, errors.New("Unknown UnifiedRoleScheduleRequestActions value: " + v)
    }
    return &result, nil
}
func SerializeUnifiedRoleScheduleRequestActions(values []UnifiedRoleScheduleRequestActions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
