package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the informationProtection singleton.
type VolumeType int

const (
    OPERATINGSYSTEMVOLUME_VOLUMETYPE VolumeType = iota
    FIXEDDATAVOLUME_VOLUMETYPE
    REMOVABLEDATAVOLUME_VOLUMETYPE
    UNKNOWNFUTUREVALUE_VOLUMETYPE
)

func (i VolumeType) String() string {
    return []string{"OPERATINGSYSTEMVOLUME", "FIXEDDATAVOLUME", "REMOVABLEDATAVOLUME", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseVolumeType(v string) (interface{}, error) {
    result := OPERATINGSYSTEMVOLUME_VOLUMETYPE
    switch strings.ToUpper(v) {
        case "OPERATINGSYSTEMVOLUME":
            result = OPERATINGSYSTEMVOLUME_VOLUMETYPE
        case "FIXEDDATAVOLUME":
            result = FIXEDDATAVOLUME_VOLUMETYPE
        case "REMOVABLEDATAVOLUME":
            result = REMOVABLEDATAVOLUME_VOLUMETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_VOLUMETYPE
        default:
            return 0, errors.New("Unknown VolumeType value: " + v)
    }
    return &result, nil
}
func SerializeVolumeType(values []VolumeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
