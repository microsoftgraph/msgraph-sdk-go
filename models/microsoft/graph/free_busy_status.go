package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_FREEBUSYSTATUS, nil
        case "FREE":
            return FREE_FREEBUSYSTATUS, nil
        case "TENTATIVE":
            return TENTATIVE_FREEBUSYSTATUS, nil
        case "BUSY":
            return BUSY_FREEBUSYSTATUS, nil
        case "OOF":
            return OOF_FREEBUSYSTATUS, nil
        case "WORKINGELSEWHERE":
            return WORKINGELSEWHERE_FREEBUSYSTATUS, nil
    }
    return 0, errors.New("Unknown FreeBusyStatus value: " + v)
}
func SerializeFreeBusyStatus(values []FreeBusyStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
