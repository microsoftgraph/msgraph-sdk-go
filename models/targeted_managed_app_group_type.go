package models
import (
    "strings"
    "errors"
)
// Provides operations to call the targetApps method.
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
    result := SELECTEDPUBLICAPPS_TARGETEDMANAGEDAPPGROUPTYPE
    switch strings.ToUpper(v) {
        case "SELECTEDPUBLICAPPS":
            result = SELECTEDPUBLICAPPS_TARGETEDMANAGEDAPPGROUPTYPE
        case "ALLCOREMICROSOFTAPPS":
            result = ALLCOREMICROSOFTAPPS_TARGETEDMANAGEDAPPGROUPTYPE
        case "ALLMICROSOFTAPPS":
            result = ALLMICROSOFTAPPS_TARGETEDMANAGEDAPPGROUPTYPE
        case "ALLAPPS":
            result = ALLAPPS_TARGETEDMANAGEDAPPGROUPTYPE
        default:
            return 0, errors.New("Unknown TargetedManagedAppGroupType value: " + v)
    }
    return &result, nil
}
func SerializeTargetedManagedAppGroupType(values []TargetedManagedAppGroupType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
