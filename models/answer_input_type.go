package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the solutionsRoot singleton.
type AnswerInputType int

const (
    TEXT_ANSWERINPUTTYPE AnswerInputType = iota
    RADIOBUTTON_ANSWERINPUTTYPE
    UNKNOWNFUTUREVALUE_ANSWERINPUTTYPE
)

func (i AnswerInputType) String() string {
    return []string{"TEXT", "RADIOBUTTON", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAnswerInputType(v string) (interface{}, error) {
    result := TEXT_ANSWERINPUTTYPE
    switch strings.ToUpper(v) {
        case "TEXT":
            result = TEXT_ANSWERINPUTTYPE
        case "RADIOBUTTON":
            result = RADIOBUTTON_ANSWERINPUTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ANSWERINPUTTYPE
        default:
            return 0, errors.New("Unknown AnswerInputType value: " + v)
    }
    return &result, nil
}
func SerializeAnswerInputType(values []AnswerInputType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
