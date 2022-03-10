package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type NotificationTemplateBrandingOptions int

const (
    NONE_NOTIFICATIONTEMPLATEBRANDINGOPTIONS NotificationTemplateBrandingOptions = iota
    INCLUDECOMPANYLOGO_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
    INCLUDECOMPANYNAME_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
    INCLUDECONTACTINFORMATION_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
)

func (i NotificationTemplateBrandingOptions) String() string {
    return []string{"NONE", "INCLUDECOMPANYLOGO", "INCLUDECOMPANYNAME", "INCLUDECONTACTINFORMATION"}[i]
}
func ParseNotificationTemplateBrandingOptions(v string) (interface{}, error) {
    result := NONE_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
        case "INCLUDECOMPANYLOGO":
            result = INCLUDECOMPANYLOGO_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
        case "INCLUDECOMPANYNAME":
            result = INCLUDECOMPANYNAME_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
        case "INCLUDECONTACTINFORMATION":
            result = INCLUDECONTACTINFORMATION_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
        default:
            return 0, errors.New("Unknown NotificationTemplateBrandingOptions value: " + v)
    }
    return &result, nil
}
func SerializeNotificationTemplateBrandingOptions(values []NotificationTemplateBrandingOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
