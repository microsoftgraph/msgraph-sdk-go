package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "TEXT":
            return TEXT_ANSWERINPUTTYPE, nil
        case "RADIOBUTTON":
            return RADIOBUTTON_ANSWERINPUTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ANSWERINPUTTYPE, nil
    }
    return 0, errors.New("Unknown AnswerInputType value: " + v)
}
func SerializeAnswerInputType(values []AnswerInputType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
