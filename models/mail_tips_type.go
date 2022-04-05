package models
import (
    "strings"
    "errors"
)
// Provides operations to call the getMailTips method.
type MailTipsType int

const (
    AUTOMATICREPLIES_MAILTIPSTYPE MailTipsType = iota
    MAILBOXFULLSTATUS_MAILTIPSTYPE
    CUSTOMMAILTIP_MAILTIPSTYPE
    EXTERNALMEMBERCOUNT_MAILTIPSTYPE
    TOTALMEMBERCOUNT_MAILTIPSTYPE
    MAXMESSAGESIZE_MAILTIPSTYPE
    DELIVERYRESTRICTION_MAILTIPSTYPE
    MODERATIONSTATUS_MAILTIPSTYPE
    RECIPIENTSCOPE_MAILTIPSTYPE
    RECIPIENTSUGGESTIONS_MAILTIPSTYPE
)

func (i MailTipsType) String() string {
    return []string{"AUTOMATICREPLIES", "MAILBOXFULLSTATUS", "CUSTOMMAILTIP", "EXTERNALMEMBERCOUNT", "TOTALMEMBERCOUNT", "MAXMESSAGESIZE", "DELIVERYRESTRICTION", "MODERATIONSTATUS", "RECIPIENTSCOPE", "RECIPIENTSUGGESTIONS"}[i]
}
func ParseMailTipsType(v string) (interface{}, error) {
    result := AUTOMATICREPLIES_MAILTIPSTYPE
    switch strings.ToUpper(v) {
        case "AUTOMATICREPLIES":
            result = AUTOMATICREPLIES_MAILTIPSTYPE
        case "MAILBOXFULLSTATUS":
            result = MAILBOXFULLSTATUS_MAILTIPSTYPE
        case "CUSTOMMAILTIP":
            result = CUSTOMMAILTIP_MAILTIPSTYPE
        case "EXTERNALMEMBERCOUNT":
            result = EXTERNALMEMBERCOUNT_MAILTIPSTYPE
        case "TOTALMEMBERCOUNT":
            result = TOTALMEMBERCOUNT_MAILTIPSTYPE
        case "MAXMESSAGESIZE":
            result = MAXMESSAGESIZE_MAILTIPSTYPE
        case "DELIVERYRESTRICTION":
            result = DELIVERYRESTRICTION_MAILTIPSTYPE
        case "MODERATIONSTATUS":
            result = MODERATIONSTATUS_MAILTIPSTYPE
        case "RECIPIENTSCOPE":
            result = RECIPIENTSCOPE_MAILTIPSTYPE
        case "RECIPIENTSUGGESTIONS":
            result = RECIPIENTSUGGESTIONS_MAILTIPSTYPE
        default:
            return 0, errors.New("Unknown MailTipsType value: " + v)
    }
    return &result, nil
}
func SerializeMailTipsType(values []MailTipsType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
