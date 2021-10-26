package graph
import (
    "strings"
    "errors"
)
// 
type TargetedManagedAppGroupType int

const (
    SELECTEDPUBLICAPPS_TARGETEDMANAGEDAPPGROUPTYPE TargetedManagedAppGroupType = iota
    ALLCOREMICROSOFTAPPS_TARGETEDMANAGEDAPPGROUPTYPE
    ALLMICROSOFTAPPS_TARGETEDMANAGEDAPPGROUPTYPE
    ALLAPPS_TARGETEDMANAGEDAPPGROUPTYPE
)

func (i TargetedManagedAppGroupType) String() string {
    return []string{"SELECTEDPUBLICAPPS", "ALLCOREMICROSOFTAPPS", "ALLMICROSOFTAPPS", "ALLAPPS"}[i]
}
func ParseTargetedManagedAppGroupType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "SELECTEDPUBLICAPPS":
            return SELECTEDPUBLICAPPS_TARGETEDMANAGEDAPPGROUPTYPE, nil
        case "ALLCOREMICROSOFTAPPS":
            return ALLCOREMICROSOFTAPPS_TARGETEDMANAGEDAPPGROUPTYPE, nil
        case "ALLMICROSOFTAPPS":
            return ALLMICROSOFTAPPS_TARGETEDMANAGEDAPPGROUPTYPE, nil
        case "ALLAPPS":
            return ALLAPPS_TARGETEDMANAGEDAPPGROUPTYPE, nil
    }
    return 0, errors.New("Unknown TargetedManagedAppGroupType value: " + v)
}
func SerializeTargetedManagedAppGroupType(values []TargetedManagedAppGroupType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
