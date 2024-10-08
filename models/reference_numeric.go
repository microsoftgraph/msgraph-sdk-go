package models
type ReferenceNumeric int

const (
    MINUS_INF_REFERENCENUMERIC ReferenceNumeric = iota
    INF_REFERENCENUMERIC
    NAN_REFERENCENUMERIC
)

func (i ReferenceNumeric) String() string {
    return []string{"-INF", "INF", "NaN"}[i]
}
func ParseReferenceNumeric(v string) (any, error) {
    result := MINUS_INF_REFERENCENUMERIC
    switch v {
        case "-INF":
            result = MINUS_INF_REFERENCENUMERIC
        case "INF":
            result = INF_REFERENCENUMERIC
        case "NaN":
            result = NAN_REFERENCENUMERIC
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeReferenceNumeric(values []ReferenceNumeric) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReferenceNumeric) isMultiValue() bool {
    return false
}
