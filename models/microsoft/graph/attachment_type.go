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
    switch strings.ToUpper(v) {
        case "FILE":
            return FILE_ATTACHMENTTYPE, nil
        case "ITEM":
            return ITEM_ATTACHMENTTYPE, nil
        case "REFERENCE":
            return REFERENCE_ATTACHMENTTYPE, nil
    }
    return 0, errors.New("Unknown AttachmentType value: " + v)
}
func SerializeAttachmentType(values []AttachmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
