package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "OPERATINGSYSTEMVOLUME":
            return OPERATINGSYSTEMVOLUME_VOLUMETYPE, nil
        case "FIXEDDATAVOLUME":
            return FIXEDDATAVOLUME_VOLUMETYPE, nil
        case "REMOVABLEDATAVOLUME":
            return REMOVABLEDATAVOLUME_VOLUMETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_VOLUMETYPE, nil
    }
    return 0, errors.New("Unknown VolumeType value: " + v)
}
func SerializeVolumeType(values []VolumeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
