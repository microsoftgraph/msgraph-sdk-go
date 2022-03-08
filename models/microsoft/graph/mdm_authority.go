package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of organization entities.
type MdmAuthority int

const (
    UNKNOWN_MDMAUTHORITY MdmAuthority = iota
    INTUNE_MDMAUTHORITY
    SCCM_MDMAUTHORITY
    OFFICE365_MDMAUTHORITY
)

func (i MdmAuthority) String() string {
    return []string{"UNKNOWN", "INTUNE", "SCCM", "OFFICE365"}[i]
}
func ParseMdmAuthority(v string) (interface{}, error) {
    result := UNKNOWN_MDMAUTHORITY
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_MDMAUTHORITY
        case "INTUNE":
            result = INTUNE_MDMAUTHORITY
        case "SCCM":
            result = SCCM_MDMAUTHORITY
        case "OFFICE365":
            result = OFFICE365_MDMAUTHORITY
        default:
            return 0, errors.New("Unknown MdmAuthority value: " + v)
    }
    return &result, nil
}
func SerializeMdmAuthority(values []MdmAuthority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
