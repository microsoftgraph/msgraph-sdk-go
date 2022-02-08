package graph
import (
    "strings"
    "errors"
)
// 
type PrintFinishing int

const (
    NONE_PRINTFINISHING PrintFinishing = iota
    STAPLE_PRINTFINISHING
    PUNCH_PRINTFINISHING
    COVER_PRINTFINISHING
    BIND_PRINTFINISHING
    SADDLESTITCH_PRINTFINISHING
    STITCHEDGE_PRINTFINISHING
    STAPLETOPLEFT_PRINTFINISHING
    STAPLEBOTTOMLEFT_PRINTFINISHING
    STAPLETOPRIGHT_PRINTFINISHING
    STAPLEBOTTOMRIGHT_PRINTFINISHING
    STITCHLEFTEDGE_PRINTFINISHING
    STITCHTOPEDGE_PRINTFINISHING
    STITCHRIGHTEDGE_PRINTFINISHING
    STITCHBOTTOMEDGE_PRINTFINISHING
    STAPLEDUALLEFT_PRINTFINISHING
    STAPLEDUALTOP_PRINTFINISHING
    STAPLEDUALRIGHT_PRINTFINISHING
    STAPLEDUALBOTTOM_PRINTFINISHING
    UNKNOWNFUTUREVALUE_PRINTFINISHING
)

func (i PrintFinishing) String() string {
    return []string{"NONE", "STAPLE", "PUNCH", "COVER", "BIND", "SADDLESTITCH", "STITCHEDGE", "STAPLETOPLEFT", "STAPLEBOTTOMLEFT", "STAPLETOPRIGHT", "STAPLEBOTTOMRIGHT", "STITCHLEFTEDGE", "STITCHTOPEDGE", "STITCHRIGHTEDGE", "STITCHBOTTOMEDGE", "STAPLEDUALLEFT", "STAPLEDUALTOP", "STAPLEDUALRIGHT", "STAPLEDUALBOTTOM", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrintFinishing(v string) (interface{}, error) {
    result := NONE_PRINTFINISHING
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_PRINTFINISHING
        case "STAPLE":
            result = STAPLE_PRINTFINISHING
        case "PUNCH":
            result = PUNCH_PRINTFINISHING
        case "COVER":
            result = COVER_PRINTFINISHING
        case "BIND":
            result = BIND_PRINTFINISHING
        case "SADDLESTITCH":
            result = SADDLESTITCH_PRINTFINISHING
        case "STITCHEDGE":
            result = STITCHEDGE_PRINTFINISHING
        case "STAPLETOPLEFT":
            result = STAPLETOPLEFT_PRINTFINISHING
        case "STAPLEBOTTOMLEFT":
            result = STAPLEBOTTOMLEFT_PRINTFINISHING
        case "STAPLETOPRIGHT":
            result = STAPLETOPRIGHT_PRINTFINISHING
        case "STAPLEBOTTOMRIGHT":
            result = STAPLEBOTTOMRIGHT_PRINTFINISHING
        case "STITCHLEFTEDGE":
            result = STITCHLEFTEDGE_PRINTFINISHING
        case "STITCHTOPEDGE":
            result = STITCHTOPEDGE_PRINTFINISHING
        case "STITCHRIGHTEDGE":
            result = STITCHRIGHTEDGE_PRINTFINISHING
        case "STITCHBOTTOMEDGE":
            result = STITCHBOTTOMEDGE_PRINTFINISHING
        case "STAPLEDUALLEFT":
            result = STAPLEDUALLEFT_PRINTFINISHING
        case "STAPLEDUALTOP":
            result = STAPLEDUALTOP_PRINTFINISHING
        case "STAPLEDUALRIGHT":
            result = STAPLEDUALRIGHT_PRINTFINISHING
        case "STAPLEDUALBOTTOM":
            result = STAPLEDUALBOTTOM_PRINTFINISHING
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PRINTFINISHING
        default:
            return 0, errors.New("Unknown PrintFinishing value: " + v)
    }
    return &result, nil
}
func SerializePrintFinishing(values []PrintFinishing) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
