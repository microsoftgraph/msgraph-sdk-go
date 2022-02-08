package graph
import (
    "strings"
    "errors"
)
// 
type CloudAppSecuritySessionControlType int

const (
    MCASCONFIGURED_CLOUDAPPSECURITYSESSIONCONTROLTYPE CloudAppSecuritySessionControlType = iota
    MONITORONLY_CLOUDAPPSECURITYSESSIONCONTROLTYPE
    BLOCKDOWNLOADS_CLOUDAPPSECURITYSESSIONCONTROLTYPE
    UNKNOWNFUTUREVALUE_CLOUDAPPSECURITYSESSIONCONTROLTYPE
)

func (i CloudAppSecuritySessionControlType) String() string {
    return []string{"MCASCONFIGURED", "MONITORONLY", "BLOCKDOWNLOADS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudAppSecuritySessionControlType(v string) (interface{}, error) {
    result := MCASCONFIGURED_CLOUDAPPSECURITYSESSIONCONTROLTYPE
    switch strings.ToUpper(v) {
        case "MCASCONFIGURED":
            result = MCASCONFIGURED_CLOUDAPPSECURITYSESSIONCONTROLTYPE
        case "MONITORONLY":
            result = MONITORONLY_CLOUDAPPSECURITYSESSIONCONTROLTYPE
        case "BLOCKDOWNLOADS":
            result = BLOCKDOWNLOADS_CLOUDAPPSECURITYSESSIONCONTROLTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDAPPSECURITYSESSIONCONTROLTYPE
        default:
            return 0, errors.New("Unknown CloudAppSecuritySessionControlType value: " + v)
    }
    return &result, nil
}
func SerializeCloudAppSecuritySessionControlType(values []CloudAppSecuritySessionControlType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
