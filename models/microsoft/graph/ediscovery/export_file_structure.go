package ediscovery
import (
    "strings"
    "errors"
)
// 
type ExportFileStructure int

const (
    NONE_EXPORTFILESTRUCTURE ExportFileStructure = iota
    DIRECTORY_EXPORTFILESTRUCTURE
    PST_EXPORTFILESTRUCTURE
    UNKNOWNFUTUREVALUE_EXPORTFILESTRUCTURE
)

func (i ExportFileStructure) String() string {
    return []string{"NONE", "DIRECTORY", "PST", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseExportFileStructure(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_EXPORTFILESTRUCTURE, nil
        case "DIRECTORY":
            return DIRECTORY_EXPORTFILESTRUCTURE, nil
        case "PST":
            return PST_EXPORTFILESTRUCTURE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EXPORTFILESTRUCTURE, nil
    }
    return 0, errors.New("Unknown ExportFileStructure value: " + v)
}
func SerializeExportFileStructure(values []ExportFileStructure) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
