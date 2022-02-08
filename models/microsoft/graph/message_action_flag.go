package graph
import (
    "strings"
    "errors"
)
// 
type MessageActionFlag int

const (
    ANY_MESSAGEACTIONFLAG MessageActionFlag = iota
    CALL_MESSAGEACTIONFLAG
    DONOTFORWARD_MESSAGEACTIONFLAG
    FOLLOWUP_MESSAGEACTIONFLAG
    FYI_MESSAGEACTIONFLAG
    FORWARD_MESSAGEACTIONFLAG
    NORESPONSENECESSARY_MESSAGEACTIONFLAG
    READ_MESSAGEACTIONFLAG
    REPLY_MESSAGEACTIONFLAG
    REPLYTOALL_MESSAGEACTIONFLAG
    REVIEW_MESSAGEACTIONFLAG
)

func (i MessageActionFlag) String() string {
    return []string{"ANY", "CALL", "DONOTFORWARD", "FOLLOWUP", "FYI", "FORWARD", "NORESPONSENECESSARY", "READ", "REPLY", "REPLYTOALL", "REVIEW"}[i]
}
func ParseMessageActionFlag(v string) (interface{}, error) {
    result := ANY_MESSAGEACTIONFLAG
    switch strings.ToUpper(v) {
        case "ANY":
            result = ANY_MESSAGEACTIONFLAG
        case "CALL":
            result = CALL_MESSAGEACTIONFLAG
        case "DONOTFORWARD":
            result = DONOTFORWARD_MESSAGEACTIONFLAG
        case "FOLLOWUP":
            result = FOLLOWUP_MESSAGEACTIONFLAG
        case "FYI":
            result = FYI_MESSAGEACTIONFLAG
        case "FORWARD":
            result = FORWARD_MESSAGEACTIONFLAG
        case "NORESPONSENECESSARY":
            result = NORESPONSENECESSARY_MESSAGEACTIONFLAG
        case "READ":
            result = READ_MESSAGEACTIONFLAG
        case "REPLY":
            result = REPLY_MESSAGEACTIONFLAG
        case "REPLYTOALL":
            result = REPLYTOALL_MESSAGEACTIONFLAG
        case "REVIEW":
            result = REVIEW_MESSAGEACTIONFLAG
        default:
            return 0, errors.New("Unknown MessageActionFlag value: " + v)
    }
    return &result, nil
}
func SerializeMessageActionFlag(values []MessageActionFlag) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
