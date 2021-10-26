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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_PRINTFINISHING, nil
        case "STAPLE":
            return STAPLE_PRINTFINISHING, nil
        case "PUNCH":
            return PUNCH_PRINTFINISHING, nil
        case "COVER":
            return COVER_PRINTFINISHING, nil
        case "BIND":
            return BIND_PRINTFINISHING, nil
        case "SADDLESTITCH":
            return SADDLESTITCH_PRINTFINISHING, nil
        case "STITCHEDGE":
            return STITCHEDGE_PRINTFINISHING, nil
        case "STAPLETOPLEFT":
            return STAPLETOPLEFT_PRINTFINISHING, nil
        case "STAPLEBOTTOMLEFT":
            return STAPLEBOTTOMLEFT_PRINTFINISHING, nil
        case "STAPLETOPRIGHT":
            return STAPLETOPRIGHT_PRINTFINISHING, nil
        case "STAPLEBOTTOMRIGHT":
            return STAPLEBOTTOMRIGHT_PRINTFINISHING, nil
        case "STITCHLEFTEDGE":
            return STITCHLEFTEDGE_PRINTFINISHING, nil
        case "STITCHTOPEDGE":
            return STITCHTOPEDGE_PRINTFINISHING, nil
        case "STITCHRIGHTEDGE":
            return STITCHRIGHTEDGE_PRINTFINISHING, nil
        case "STITCHBOTTOMEDGE":
            return STITCHBOTTOMEDGE_PRINTFINISHING, nil
        case "STAPLEDUALLEFT":
            return STAPLEDUALLEFT_PRINTFINISHING, nil
        case "STAPLEDUALTOP":
            return STAPLEDUALTOP_PRINTFINISHING, nil
        case "STAPLEDUALRIGHT":
            return STAPLEDUALRIGHT_PRINTFINISHING, nil
        case "STAPLEDUALBOTTOM":
            return STAPLEDUALBOTTOM_PRINTFINISHING, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTFINISHING, nil
    }
    return 0, errors.New("Unknown PrintFinishing value: " + v)
}
func SerializePrintFinishing(values []PrintFinishing) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
