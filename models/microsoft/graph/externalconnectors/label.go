package externalconnectors
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of externalConnection entities.
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
    result := TITLE_LABEL
    switch strings.ToUpper(v) {
        case "TITLE":
            result = TITLE_LABEL
        case "URL":
            result = URL_LABEL
        case "CREATEDBY":
            result = CREATEDBY_LABEL
        case "LASTMODIFIEDBY":
            result = LASTMODIFIEDBY_LABEL
        case "AUTHORS":
            result = AUTHORS_LABEL
        case "CREATEDDATETIME":
            result = CREATEDDATETIME_LABEL
        case "LASTMODIFIEDDATETIME":
            result = LASTMODIFIEDDATETIME_LABEL
        case "FILENAME":
            result = FILENAME_LABEL
        case "FILEEXTENSION":
            result = FILEEXTENSION_LABEL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_LABEL
        default:
            return 0, errors.New("Unknown Label value: " + v)
    }
    return &result, nil
}
func SerializeLabel(values []Label) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
