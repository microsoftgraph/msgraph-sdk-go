package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "SENDTODELEGATEANDINFORMATIONTOPRINCIPAL":
            return SENDTODELEGATEANDINFORMATIONTOPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS, nil
        case "SENDTODELEGATEANDPRINCIPAL":
            return SENDTODELEGATEANDPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS, nil
        case "SENDTODELEGATEONLY":
            return SENDTODELEGATEONLY_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS, nil
    }
    return 0, errors.New("Unknown DelegateMeetingMessageDeliveryOptions value: " + v)
}
func SerializeDelegateMeetingMessageDeliveryOptions(values []DelegateMeetingMessageDeliveryOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
