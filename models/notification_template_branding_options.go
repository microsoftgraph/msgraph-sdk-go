package models
import (
    "errors"
    "strings"
)
// Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
type NotificationTemplateBrandingOptions int

const (
    // Indicates that no branding options are set in the message template.
    NONE_NOTIFICATIONTEMPLATEBRANDINGOPTIONS NotificationTemplateBrandingOptions = iota
    // Indicates to include company logo in the message template.
    INCLUDECOMPANYLOGO_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
    // Indicates to include company name in the message template.
    INCLUDECOMPANYNAME_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
    // Indicates to include contact information in the message template.
    INCLUDECONTACTINFORMATION_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
    // Indicates to include company portal website link in the message template.
    INCLUDECOMPANYPORTALLINK_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
    // Indicates to include device details in the message template.
    INCLUDEDEVICEDETAILS_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
)

func (i NotificationTemplateBrandingOptions) String() string {
    var values []string
    for p := NotificationTemplateBrandingOptions(1); p <= UNKNOWNFUTUREVALUE_NOTIFICATIONTEMPLATEBRANDINGOPTIONS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "includeCompanyLogo", "includeCompanyName", "includeContactInformation", "includeCompanyPortalLink", "includeDeviceDetails", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseNotificationTemplateBrandingOptions(v string) (any, error) {
    var result NotificationTemplateBrandingOptions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
            case "includeCompanyLogo":
                result |= INCLUDECOMPANYLOGO_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
            case "includeCompanyName":
                result |= INCLUDECOMPANYNAME_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
            case "includeContactInformation":
                result |= INCLUDECONTACTINFORMATION_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
            case "includeCompanyPortalLink":
                result |= INCLUDECOMPANYPORTALLINK_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
            case "includeDeviceDetails":
                result |= INCLUDEDEVICEDETAILS_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_NOTIFICATIONTEMPLATEBRANDINGOPTIONS
            default:
                return 0, errors.New("Unknown NotificationTemplateBrandingOptions value: " + v)
        }
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
func (i NotificationTemplateBrandingOptions) isMultiValue() bool {
    return true
}
