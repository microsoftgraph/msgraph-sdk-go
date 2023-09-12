package security
import (
    "errors"
    "strings"
)
// 
type ExportOptions int

const (
    ORIGINALFILES_EXPORTOPTIONS ExportOptions = iota
    TEXT_EXPORTOPTIONS
    PDFREPLACEMENT_EXPORTOPTIONS
    TAGS_EXPORTOPTIONS
    UNKNOWNFUTUREVALUE_EXPORTOPTIONS
)

func (i ExportOptions) String() string {
    var values []string
    for p := ExportOptions(1); p <= UNKNOWNFUTUREVALUE_EXPORTOPTIONS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"originalFiles", "text", "pdfReplacement", "tags", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseExportOptions(v string) (any, error) {
    var result ExportOptions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "originalFiles":
                result |= ORIGINALFILES_EXPORTOPTIONS
            case "text":
                result |= TEXT_EXPORTOPTIONS
            case "pdfReplacement":
                result |= PDFREPLACEMENT_EXPORTOPTIONS
            case "tags":
                result |= TAGS_EXPORTOPTIONS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_EXPORTOPTIONS
            default:
                return 0, errors.New("Unknown ExportOptions value: " + v)
        }
    }
    return &result, nil
}
func SerializeExportOptions(values []ExportOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ExportOptions) isMultiValue() bool {
    return true
}
