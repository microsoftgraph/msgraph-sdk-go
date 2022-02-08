package graph
import (
    "strings"
    "errors"
)
// 
type AttachmentType int

const (
    FILE_ATTACHMENTTYPE AttachmentType = iota
    ITEM_ATTACHMENTTYPE
    REFERENCE_ATTACHMENTTYPE
)

func (i AttachmentType) String() string {
    return []string{"FILE", "ITEM", "REFERENCE"}[i]
}
func ParseAttachmentType(v string) (interface{}, error) {
    result := FILE_ATTACHMENTTYPE
    switch strings.ToUpper(v) {
        case "FILE":
            result = FILE_ATTACHMENTTYPE
        case "ITEM":
            result = ITEM_ATTACHMENTTYPE
        case "REFERENCE":
            result = REFERENCE_ATTACHMENTTYPE
        default:
            return 0, errors.New("Unknown AttachmentType value: " + v)
    }
    return &result, nil
}
func SerializeAttachmentType(values []AttachmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
