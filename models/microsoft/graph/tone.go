package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "TONE0":
            return TONE0_TONE, nil
        case "TONE1":
            return TONE1_TONE, nil
        case "TONE2":
            return TONE2_TONE, nil
        case "TONE3":
            return TONE3_TONE, nil
        case "TONE4":
            return TONE4_TONE, nil
        case "TONE5":
            return TONE5_TONE, nil
        case "TONE6":
            return TONE6_TONE, nil
        case "TONE7":
            return TONE7_TONE, nil
        case "TONE8":
            return TONE8_TONE, nil
        case "TONE9":
            return TONE9_TONE, nil
        case "STAR":
            return STAR_TONE, nil
        case "POUND":
            return POUND_TONE, nil
        case "A":
            return A_TONE, nil
        case "B":
            return B_TONE, nil
        case "C":
            return C_TONE, nil
        case "D":
            return D_TONE, nil
        case "FLASH":
            return FLASH_TONE, nil
    }
    return 0, errors.New("Unknown Tone value: " + v)
}
func SerializeTone(values []Tone) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
