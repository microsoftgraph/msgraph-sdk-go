package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type FreeBusyStatus int

const (
    UNKNOWN_FREEBUSYSTATUS FreeBusyStatus = iota
    FREE_FREEBUSYSTATUS
    TENTATIVE_FREEBUSYSTATUS
    BUSY_FREEBUSYSTATUS
    OOF_FREEBUSYSTATUS
    WORKINGELSEWHERE_FREEBUSYSTATUS
)

func (i FreeBusyStatus) String() string {
    return []string{"UNKNOWN", "FREE", "TENTATIVE", "BUSY", "OOF", "WORKINGELSEWHERE"}[i]
}
func ParseFreeBusyStatus(v string) (interface{}, error) {
    result := UNKNOWN_FREEBUSYSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_FREEBUSYSTATUS
        case "FREE":
            result = FREE_FREEBUSYSTATUS
        case "TENTATIVE":
            result = TENTATIVE_FREEBUSYSTATUS
        case "BUSY":
            result = BUSY_FREEBUSYSTATUS
        case "OOF":
            result = OOF_FREEBUSYSTATUS
        case "WORKINGELSEWHERE":
            result = WORKINGELSEWHERE_FREEBUSYSTATUS
        default:
            return 0, errors.New("Unknown FreeBusyStatus value: " + v)
    }
    return &result, nil
}
func SerializeFreeBusyStatus(values []FreeBusyStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
