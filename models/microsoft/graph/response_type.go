package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
type ResponseType int

const (
    NONE_RESPONSETYPE ResponseType = iota
    ORGANIZER_RESPONSETYPE
    TENTATIVELYACCEPTED_RESPONSETYPE
    ACCEPTED_RESPONSETYPE
    DECLINED_RESPONSETYPE
    NOTRESPONDED_RESPONSETYPE
)

func (i ResponseType) String() string {
    return []string{"NONE", "ORGANIZER", "TENTATIVELYACCEPTED", "ACCEPTED", "DECLINED", "NOTRESPONDED"}[i]
}
func ParseResponseType(v string) (interface{}, error) {
    result := NONE_RESPONSETYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_RESPONSETYPE
        case "ORGANIZER":
            result = ORGANIZER_RESPONSETYPE
        case "TENTATIVELYACCEPTED":
            result = TENTATIVELYACCEPTED_RESPONSETYPE
        case "ACCEPTED":
            result = ACCEPTED_RESPONSETYPE
        case "DECLINED":
            result = DECLINED_RESPONSETYPE
        case "NOTRESPONDED":
            result = NOTRESPONDED_RESPONSETYPE
        default:
            return 0, errors.New("Unknown ResponseType value: " + v)
    }
    return &result, nil
}
func SerializeResponseType(values []ResponseType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
