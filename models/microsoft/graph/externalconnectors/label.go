package externalconnectors
import (
    "strings"
    "errors"
)
// 
type Label int

const (
    TITLE_LABEL Label = iota
    URL_LABEL
    CREATEDBY_LABEL
    LASTMODIFIEDBY_LABEL
    AUTHORS_LABEL
    CREATEDDATETIME_LABEL
    LASTMODIFIEDDATETIME_LABEL
    FILENAME_LABEL
    FILEEXTENSION_LABEL
    UNKNOWNFUTUREVALUE_LABEL
)

func (i Label) String() string {
    return []string{"TITLE", "URL", "CREATEDBY", "LASTMODIFIEDBY", "AUTHORS", "CREATEDDATETIME", "LASTMODIFIEDDATETIME", "FILENAME", "FILEEXTENSION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseLabel(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "TITLE":
            return TITLE_LABEL, nil
        case "URL":
            return URL_LABEL, nil
        case "CREATEDBY":
            return CREATEDBY_LABEL, nil
        case "LASTMODIFIEDBY":
            return LASTMODIFIEDBY_LABEL, nil
        case "AUTHORS":
            return AUTHORS_LABEL, nil
        case "CREATEDDATETIME":
            return CREATEDDATETIME_LABEL, nil
        case "LASTMODIFIEDDATETIME":
            return LASTMODIFIEDDATETIME_LABEL, nil
        case "FILENAME":
            return FILENAME_LABEL, nil
        case "FILEEXTENSION":
            return FILEEXTENSION_LABEL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_LABEL, nil
    }
    return 0, errors.New("Unknown Label value: " + v)
}
func SerializeLabel(values []Label) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
