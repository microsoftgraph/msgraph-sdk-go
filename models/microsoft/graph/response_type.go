package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_RESPONSETYPE, nil
        case "ORGANIZER":
            return ORGANIZER_RESPONSETYPE, nil
        case "TENTATIVELYACCEPTED":
            return TENTATIVELYACCEPTED_RESPONSETYPE, nil
        case "ACCEPTED":
            return ACCEPTED_RESPONSETYPE, nil
        case "DECLINED":
            return DECLINED_RESPONSETYPE, nil
        case "NOTRESPONDED":
            return NOTRESPONDED_RESPONSETYPE, nil
    }
    return 0, errors.New("Unknown ResponseType value: " + v)
}
func SerializeResponseType(values []ResponseType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
