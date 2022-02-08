package graph
import (
    "strings"
    "errors"
)
// 
type DelegateMeetingMessageDeliveryOptions int

const (
    SENDTODELEGATEANDINFORMATIONTOPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS DelegateMeetingMessageDeliveryOptions = iota
    SENDTODELEGATEANDPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
    SENDTODELEGATEONLY_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
)

func (i DelegateMeetingMessageDeliveryOptions) String() string {
    return []string{"SENDTODELEGATEANDINFORMATIONTOPRINCIPAL", "SENDTODELEGATEANDPRINCIPAL", "SENDTODELEGATEONLY"}[i]
}
func ParseDelegateMeetingMessageDeliveryOptions(v string) (interface{}, error) {
    result := SENDTODELEGATEANDINFORMATIONTOPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
    switch strings.ToUpper(v) {
        case "SENDTODELEGATEANDINFORMATIONTOPRINCIPAL":
            result = SENDTODELEGATEANDINFORMATIONTOPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
        case "SENDTODELEGATEANDPRINCIPAL":
            result = SENDTODELEGATEANDPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
        case "SENDTODELEGATEONLY":
            result = SENDTODELEGATEONLY_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
        default:
            return 0, errors.New("Unknown DelegateMeetingMessageDeliveryOptions value: " + v)
    }
    return &result, nil
}
func SerializeDelegateMeetingMessageDeliveryOptions(values []DelegateMeetingMessageDeliveryOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
