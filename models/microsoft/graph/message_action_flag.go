package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ANY":
            return ANY_MESSAGEACTIONFLAG, nil
        case "CALL":
            return CALL_MESSAGEACTIONFLAG, nil
        case "DONOTFORWARD":
            return DONOTFORWARD_MESSAGEACTIONFLAG, nil
        case "FOLLOWUP":
            return FOLLOWUP_MESSAGEACTIONFLAG, nil
        case "FYI":
            return FYI_MESSAGEACTIONFLAG, nil
        case "FORWARD":
            return FORWARD_MESSAGEACTIONFLAG, nil
        case "NORESPONSENECESSARY":
            return NORESPONSENECESSARY_MESSAGEACTIONFLAG, nil
        case "READ":
            return READ_MESSAGEACTIONFLAG, nil
        case "REPLY":
            return REPLY_MESSAGEACTIONFLAG, nil
        case "REPLYTOALL":
            return REPLYTOALL_MESSAGEACTIONFLAG, nil
        case "REVIEW":
            return REVIEW_MESSAGEACTIONFLAG, nil
    }
    return 0, errors.New("Unknown MessageActionFlag value: " + v)
}
func SerializeMessageActionFlag(values []MessageActionFlag) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
