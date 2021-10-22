package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_NOTIFICATIONTEMPLATEBRANDINGOPTIONS, nil
        case "INCLUDECOMPANYLOGO":
            return INCLUDECOMPANYLOGO_NOTIFICATIONTEMPLATEBRANDINGOPTIONS, nil
        case "INCLUDECOMPANYNAME":
            return INCLUDECOMPANYNAME_NOTIFICATIONTEMPLATEBRANDINGOPTIONS, nil
        case "INCLUDECONTACTINFORMATION":
            return INCLUDECONTACTINFORMATION_NOTIFICATIONTEMPLATEBRANDINGOPTIONS, nil
    }
    return 0, errors.New("Unknown NotificationTemplateBrandingOptions value: " + v)
}
func SerializeNotificationTemplateBrandingOptions(values []NotificationTemplateBrandingOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
