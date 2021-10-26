package graph
import (
    "strings"
    "errors"
)
// 
type UserFlowType int

const (
    SIGNUP_USERFLOWTYPE UserFlowType = iota
    SIGNIN_USERFLOWTYPE
    SIGNUPORSIGNIN_USERFLOWTYPE
    PASSWORDRESET_USERFLOWTYPE
    PROFILEUPDATE_USERFLOWTYPE
    RESOURCEOWNER_USERFLOWTYPE
    UNKNOWNFUTUREVALUE_USERFLOWTYPE
)

func (i UserFlowType) String() string {
    return []string{"SIGNUP", "SIGNIN", "SIGNUPORSIGNIN", "PASSWORDRESET", "PROFILEUPDATE", "RESOURCEOWNER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUserFlowType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "SIGNUP":
            return SIGNUP_USERFLOWTYPE, nil
        case "SIGNIN":
            return SIGNIN_USERFLOWTYPE, nil
        case "SIGNUPORSIGNIN":
            return SIGNUPORSIGNIN_USERFLOWTYPE, nil
        case "PASSWORDRESET":
            return PASSWORDRESET_USERFLOWTYPE, nil
        case "PROFILEUPDATE":
            return PROFILEUPDATE_USERFLOWTYPE, nil
        case "RESOURCEOWNER":
            return RESOURCEOWNER_USERFLOWTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_USERFLOWTYPE, nil
    }
    return 0, errors.New("Unknown UserFlowType value: " + v)
}
func SerializeUserFlowType(values []UserFlowType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
