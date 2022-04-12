package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type CategoryColor int

const (
    NONE_CATEGORYCOLOR CategoryColor = iota
    PRESET0_CATEGORYCOLOR
    PRESET1_CATEGORYCOLOR
    PRESET2_CATEGORYCOLOR
    PRESET3_CATEGORYCOLOR
    PRESET4_CATEGORYCOLOR
    PRESET5_CATEGORYCOLOR
    PRESET6_CATEGORYCOLOR
    PRESET7_CATEGORYCOLOR
    PRESET8_CATEGORYCOLOR
    PRESET9_CATEGORYCOLOR
    PRESET10_CATEGORYCOLOR
    PRESET11_CATEGORYCOLOR
    PRESET12_CATEGORYCOLOR
    PRESET13_CATEGORYCOLOR
    PRESET14_CATEGORYCOLOR
    PRESET15_CATEGORYCOLOR
    PRESET16_CATEGORYCOLOR
    PRESET17_CATEGORYCOLOR
    PRESET18_CATEGORYCOLOR
    PRESET19_CATEGORYCOLOR
    PRESET20_CATEGORYCOLOR
    PRESET21_CATEGORYCOLOR
    PRESET22_CATEGORYCOLOR
    PRESET23_CATEGORYCOLOR
    PRESET24_CATEGORYCOLOR
)

func (i CategoryColor) String() string {
    return []string{"NONE", "PRESET0", "PRESET1", "PRESET2", "PRESET3", "PRESET4", "PRESET5", "PRESET6", "PRESET7", "PRESET8", "PRESET9", "PRESET10", "PRESET11", "PRESET12", "PRESET13", "PRESET14", "PRESET15", "PRESET16", "PRESET17", "PRESET18", "PRESET19", "PRESET20", "PRESET21", "PRESET22", "PRESET23", "PRESET24"}[i]
}
func ParseCategoryColor(v string) (interface{}, error) {
    result := NONE_CATEGORYCOLOR
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_CATEGORYCOLOR
        case "PRESET0":
            result = PRESET0_CATEGORYCOLOR
        case "PRESET1":
            result = PRESET1_CATEGORYCOLOR
        case "PRESET2":
            result = PRESET2_CATEGORYCOLOR
        case "PRESET3":
            result = PRESET3_CATEGORYCOLOR
        case "PRESET4":
            result = PRESET4_CATEGORYCOLOR
        case "PRESET5":
            result = PRESET5_CATEGORYCOLOR
        case "PRESET6":
            result = PRESET6_CATEGORYCOLOR
        case "PRESET7":
            result = PRESET7_CATEGORYCOLOR
        case "PRESET8":
            result = PRESET8_CATEGORYCOLOR
        case "PRESET9":
            result = PRESET9_CATEGORYCOLOR
        case "PRESET10":
            result = PRESET10_CATEGORYCOLOR
        case "PRESET11":
            result = PRESET11_CATEGORYCOLOR
        case "PRESET12":
            result = PRESET12_CATEGORYCOLOR
        case "PRESET13":
            result = PRESET13_CATEGORYCOLOR
        case "PRESET14":
            result = PRESET14_CATEGORYCOLOR
        case "PRESET15":
            result = PRESET15_CATEGORYCOLOR
        case "PRESET16":
            result = PRESET16_CATEGORYCOLOR
        case "PRESET17":
            result = PRESET17_CATEGORYCOLOR
        case "PRESET18":
            result = PRESET18_CATEGORYCOLOR
        case "PRESET19":
            result = PRESET19_CATEGORYCOLOR
        case "PRESET20":
            result = PRESET20_CATEGORYCOLOR
        case "PRESET21":
            result = PRESET21_CATEGORYCOLOR
        case "PRESET22":
            result = PRESET22_CATEGORYCOLOR
        case "PRESET23":
            result = PRESET23_CATEGORYCOLOR
        case "PRESET24":
            result = PRESET24_CATEGORYCOLOR
        default:
            return 0, errors.New("Unknown CategoryColor value: " + v)
    }
    return &result, nil
}
func SerializeCategoryColor(values []CategoryColor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
