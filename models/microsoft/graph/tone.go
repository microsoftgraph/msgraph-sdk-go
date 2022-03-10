package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type Tone int

const (
    TONE0_TONE Tone = iota
    TONE1_TONE
    TONE2_TONE
    TONE3_TONE
    TONE4_TONE
    TONE5_TONE
    TONE6_TONE
    TONE7_TONE
    TONE8_TONE
    TONE9_TONE
    STAR_TONE
    POUND_TONE
    A_TONE
    B_TONE
    C_TONE
    D_TONE
    FLASH_TONE
)

func (i Tone) String() string {
    return []string{"TONE0", "TONE1", "TONE2", "TONE3", "TONE4", "TONE5", "TONE6", "TONE7", "TONE8", "TONE9", "STAR", "POUND", "A", "B", "C", "D", "FLASH"}[i]
}
func ParseTone(v string) (interface{}, error) {
    result := TONE0_TONE
    switch strings.ToUpper(v) {
        case "TONE0":
            result = TONE0_TONE
        case "TONE1":
            result = TONE1_TONE
        case "TONE2":
            result = TONE2_TONE
        case "TONE3":
            result = TONE3_TONE
        case "TONE4":
            result = TONE4_TONE
        case "TONE5":
            result = TONE5_TONE
        case "TONE6":
            result = TONE6_TONE
        case "TONE7":
            result = TONE7_TONE
        case "TONE8":
            result = TONE8_TONE
        case "TONE9":
            result = TONE9_TONE
        case "STAR":
            result = STAR_TONE
        case "POUND":
            result = POUND_TONE
        case "A":
            result = A_TONE
        case "B":
            result = B_TONE
        case "C":
            result = C_TONE
        case "D":
            result = D_TONE
        case "FLASH":
            result = FLASH_TONE
        default:
            return 0, errors.New("Unknown Tone value: " + v)
    }
    return &result, nil
}
func SerializeTone(values []Tone) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
