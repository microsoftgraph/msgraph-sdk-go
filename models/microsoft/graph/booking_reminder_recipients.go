package graph
import (
    "strings"
    "errors"
)
// 
type BookingReminderRecipients int

const (
    ALLATTENDEES_BOOKINGREMINDERRECIPIENTS BookingReminderRecipients = iota
    STAFF_BOOKINGREMINDERRECIPIENTS
    CUSTOMER_BOOKINGREMINDERRECIPIENTS
    UNKNOWNFUTUREVALUE_BOOKINGREMINDERRECIPIENTS
)

func (i BookingReminderRecipients) String() string {
    return []string{"ALLATTENDEES", "STAFF", "CUSTOMER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseBookingReminderRecipients(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ALLATTENDEES":
            return ALLATTENDEES_BOOKINGREMINDERRECIPIENTS, nil
        case "STAFF":
            return STAFF_BOOKINGREMINDERRECIPIENTS, nil
        case "CUSTOMER":
            return CUSTOMER_BOOKINGREMINDERRECIPIENTS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_BOOKINGREMINDERRECIPIENTS, nil
    }
    return 0, errors.New("Unknown BookingReminderRecipients value: " + v)
}
func SerializeBookingReminderRecipients(values []BookingReminderRecipients) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
