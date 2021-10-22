package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_CATEGORYCOLOR, nil
        case "PRESET0":
            return PRESET0_CATEGORYCOLOR, nil
        case "PRESET1":
            return PRESET1_CATEGORYCOLOR, nil
        case "PRESET2":
            return PRESET2_CATEGORYCOLOR, nil
        case "PRESET3":
            return PRESET3_CATEGORYCOLOR, nil
        case "PRESET4":
            return PRESET4_CATEGORYCOLOR, nil
        case "PRESET5":
            return PRESET5_CATEGORYCOLOR, nil
        case "PRESET6":
            return PRESET6_CATEGORYCOLOR, nil
        case "PRESET7":
            return PRESET7_CATEGORYCOLOR, nil
        case "PRESET8":
            return PRESET8_CATEGORYCOLOR, nil
        case "PRESET9":
            return PRESET9_CATEGORYCOLOR, nil
        case "PRESET10":
            return PRESET10_CATEGORYCOLOR, nil
        case "PRESET11":
            return PRESET11_CATEGORYCOLOR, nil
        case "PRESET12":
            return PRESET12_CATEGORYCOLOR, nil
        case "PRESET13":
            return PRESET13_CATEGORYCOLOR, nil
        case "PRESET14":
            return PRESET14_CATEGORYCOLOR, nil
        case "PRESET15":
            return PRESET15_CATEGORYCOLOR, nil
        case "PRESET16":
            return PRESET16_CATEGORYCOLOR, nil
        case "PRESET17":
            return PRESET17_CATEGORYCOLOR, nil
        case "PRESET18":
            return PRESET18_CATEGORYCOLOR, nil
        case "PRESET19":
            return PRESET19_CATEGORYCOLOR, nil
        case "PRESET20":
            return PRESET20_CATEGORYCOLOR, nil
        case "PRESET21":
            return PRESET21_CATEGORYCOLOR, nil
        case "PRESET22":
            return PRESET22_CATEGORYCOLOR, nil
        case "PRESET23":
            return PRESET23_CATEGORYCOLOR, nil
        case "PRESET24":
            return PRESET24_CATEGORYCOLOR, nil
    }
    return 0, errors.New("Unknown CategoryColor value: " + v)
}
func SerializeCategoryColor(values []CategoryColor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
