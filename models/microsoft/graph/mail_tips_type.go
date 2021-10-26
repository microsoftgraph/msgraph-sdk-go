package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "AUTOMATICREPLIES":
            return AUTOMATICREPLIES_MAILTIPSTYPE, nil
        case "MAILBOXFULLSTATUS":
            return MAILBOXFULLSTATUS_MAILTIPSTYPE, nil
        case "CUSTOMMAILTIP":
            return CUSTOMMAILTIP_MAILTIPSTYPE, nil
        case "EXTERNALMEMBERCOUNT":
            return EXTERNALMEMBERCOUNT_MAILTIPSTYPE, nil
        case "TOTALMEMBERCOUNT":
            return TOTALMEMBERCOUNT_MAILTIPSTYPE, nil
        case "MAXMESSAGESIZE":
            return MAXMESSAGESIZE_MAILTIPSTYPE, nil
        case "DELIVERYRESTRICTION":
            return DELIVERYRESTRICTION_MAILTIPSTYPE, nil
        case "MODERATIONSTATUS":
            return MODERATIONSTATUS_MAILTIPSTYPE, nil
        case "RECIPIENTSCOPE":
            return RECIPIENTSCOPE_MAILTIPSTYPE, nil
        case "RECIPIENTSUGGESTIONS":
            return RECIPIENTSUGGESTIONS_MAILTIPSTYPE, nil
    }
    return 0, errors.New("Unknown MailTipsType value: " + v)
}
func SerializeMailTipsType(values []MailTipsType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
