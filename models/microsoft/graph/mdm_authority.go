package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_MDMAUTHORITY, nil
        case "INTUNE":
            return INTUNE_MDMAUTHORITY, nil
        case "SCCM":
            return SCCM_MDMAUTHORITY, nil
        case "OFFICE365":
            return OFFICE365_MDMAUTHORITY, nil
    }
    return 0, errors.New("Unknown MdmAuthority value: " + v)
}
func SerializeMdmAuthority(values []MdmAuthority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
