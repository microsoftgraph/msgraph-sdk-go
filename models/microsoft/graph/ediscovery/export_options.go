package ediscovery
import (
    "strings"
    "errors"
)
// 
type ExportOptions int

const (
    ORIGINALFILES_EXPORTOPTIONS ExportOptions = iota
    TEXT_EXPORTOPTIONS
    PDFREPLACEMENT_EXPORTOPTIONS
    FILEINFO_EXPORTOPTIONS
    TAGS_EXPORTOPTIONS
    UNKNOWNFUTUREVALUE_EXPORTOPTIONS
)

func (i ExportOptions) String() string {
    return []string{"ORIGINALFILES", "TEXT", "PDFREPLACEMENT", "FILEINFO", "TAGS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseExportOptions(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ORIGINALFILES":
            return ORIGINALFILES_EXPORTOPTIONS, nil
        case "TEXT":
            return TEXT_EXPORTOPTIONS, nil
        case "PDFREPLACEMENT":
            return PDFREPLACEMENT_EXPORTOPTIONS, nil
        case "FILEINFO":
            return FILEINFO_EXPORTOPTIONS, nil
        case "TAGS":
            return TAGS_EXPORTOPTIONS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EXPORTOPTIONS, nil
    }
    return 0, errors.New("Unknown ExportOptions value: " + v)
}
func SerializeExportOptions(values []ExportOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
