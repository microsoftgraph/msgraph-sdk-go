package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingAppointment 
type BookingAppointment struct {
    Entity
    // Additional information that is sent to the customer when an appointment is confirmed.
    additionalInformation *string;
    // It lists down the customer properties for an appointment. An appointment will contain a list of customer information and each unit will indicate the properties of a customer who is part of that appointment. Optional.
    customers []BookingCustomerInformationBase;
    // The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
    customerTimeZone *string;
    // The length of the appointment, denoted in ISO8601 format.
    duration *string;
    // 
    endDateTime *DateTimeTimeZone;
    // The current number of customers in the appointment.
    filledAttendeesCount *int32;
    // True indicates that the appointment will be held online. Default value is false.
    isLocationOnline *bool;
    // The URL of the online meeting for the appointment.
    joinWebUrl *string;
    // The maximum number of customers allowed in an appointment.
    maximumAttendeesCount *int32;
    // True indicates that the bookingCustomer for this appointment does not wish to receive a confirmation for this appointment.
    optOutOfCustomerEmail *bool;
    // The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
    postBuffer *string;
    // The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
    preBuffer *string;
    // The regular price for an appointment for the specified bookingService.
    price *float64;
    // A setting to provide flexibility for the pricing structure of services. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet, unknownFutureValue.
    priceType *BookingPriceType;
    // The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
    reminders []BookingReminder;
    // An additional tracking ID for the appointment, if the appointment has been created directly by the customer on the scheduling page, as opposed to by a staff member on the behalf of the customer.
    selfServiceAppointmentId *string;
    // The ID of the bookingService associated with this appointment.
    serviceId *string;
    // The location where the service is delivered.
    serviceLocation *Location;
    // The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it is computed from the service associated with the appointment by the serviceId property.
    serviceName *string;
    // Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
    serviceNotes *string;
    // True indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
    smsNotificationsEnabled *bool;
    // The ID of each bookingStaffMember who is scheduled in this appointment.
    staffMemberIds []string;
    // 
    startDateTime *DateTimeTimeZone;
}
// NewBookingAppointment instantiates a new bookingAppointment and sets the default values.
func NewBookingAppointment()(*BookingAppointment) {
    m := &BookingAppointment{
        Entity: *NewEntity(),
    }
    return m
}
// GetAdditionalInformation gets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
func (m *BookingAppointment) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// GetCustomers gets the customers property value. It lists down the customer properties for an appointment. An appointment will contain a list of customer information and each unit will indicate the properties of a customer who is part of that appointment. Optional.
func (m *BookingAppointment) GetCustomers()([]BookingCustomerInformationBase) {
    if m == nil {
        return nil
    } else {
        return m.customers
    }
}
// GetCustomerTimeZone gets the customerTimeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
func (m *BookingAppointment) GetCustomerTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerTimeZone
    }
}
// GetDuration gets the duration property value. The length of the appointment, denoted in ISO8601 format.
func (m *BookingAppointment) GetDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// GetEndDateTime gets the endDateTime property value. 
func (m *BookingAppointment) GetEndDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetFilledAttendeesCount gets the filledAttendeesCount property value. The current number of customers in the appointment.
func (m *BookingAppointment) GetFilledAttendeesCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.filledAttendeesCount
    }
}
// GetIsLocationOnline gets the isLocationOnline property value. True indicates that the appointment will be held online. Default value is false.
func (m *BookingAppointment) GetIsLocationOnline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLocationOnline
    }
}
// GetJoinWebUrl gets the joinWebUrl property value. The URL of the online meeting for the appointment.
func (m *BookingAppointment) GetJoinWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinWebUrl
    }
}
// GetMaximumAttendeesCount gets the maximumAttendeesCount property value. The maximum number of customers allowed in an appointment.
func (m *BookingAppointment) GetMaximumAttendeesCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumAttendeesCount
    }
}
// GetOptOutOfCustomerEmail gets the optOutOfCustomerEmail property value. True indicates that the bookingCustomer for this appointment does not wish to receive a confirmation for this appointment.
func (m *BookingAppointment) GetOptOutOfCustomerEmail()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.optOutOfCustomerEmail
    }
}
// GetPostBuffer gets the postBuffer property value. The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) GetPostBuffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postBuffer
    }
}
// GetPreBuffer gets the preBuffer property value. The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) GetPreBuffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preBuffer
    }
}
// GetPrice gets the price property value. The regular price for an appointment for the specified bookingService.
func (m *BookingAppointment) GetPrice()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.price
    }
}
// GetPriceType gets the priceType property value. A setting to provide flexibility for the pricing structure of services. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet, unknownFutureValue.
func (m *BookingAppointment) GetPriceType()(*BookingPriceType) {
    if m == nil {
        return nil
    } else {
        return m.priceType
    }
}
// GetReminders gets the reminders property value. The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) GetReminders()([]BookingReminder) {
    if m == nil {
        return nil
    } else {
        return m.reminders
    }
}
// GetSelfServiceAppointmentId gets the selfServiceAppointmentId property value. An additional tracking ID for the appointment, if the appointment has been created directly by the customer on the scheduling page, as opposed to by a staff member on the behalf of the customer.
func (m *BookingAppointment) GetSelfServiceAppointmentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.selfServiceAppointmentId
    }
}
// GetServiceId gets the serviceId property value. The ID of the bookingService associated with this appointment.
func (m *BookingAppointment) GetServiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceId
    }
}
// GetServiceLocation gets the serviceLocation property value. The location where the service is delivered.
func (m *BookingAppointment) GetServiceLocation()(*Location) {
    if m == nil {
        return nil
    } else {
        return m.serviceLocation
    }
}
// GetServiceName gets the serviceName property value. The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it is computed from the service associated with the appointment by the serviceId property.
func (m *BookingAppointment) GetServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceName
    }
}
// GetServiceNotes gets the serviceNotes property value. Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) GetServiceNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceNotes
    }
}
// GetSmsNotificationsEnabled gets the smsNotificationsEnabled property value. True indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
func (m *BookingAppointment) GetSmsNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.smsNotificationsEnabled
    }
}
// GetStaffMemberIds gets the staffMemberIds property value. The ID of each bookingStaffMember who is scheduled in this appointment.
func (m *BookingAppointment) GetStaffMemberIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.staffMemberIds
    }
}
// GetStartDateTime gets the startDateTime property value. 
func (m *BookingAppointment) GetStartDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingAppointment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInformation(val)
        }
        return nil
    }
    res["customers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingCustomerInformationBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCustomerInformationBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingCustomerInformationBase))
            }
            m.SetCustomers(res)
        }
        return nil
    }
    res["customerTimeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerTimeZone(val)
        }
        return nil
    }
    res["duration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["filledAttendeesCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilledAttendeesCount(val)
        }
        return nil
    }
    res["isLocationOnline"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLocationOnline(val)
        }
        return nil
    }
    res["joinWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinWebUrl(val)
        }
        return nil
    }
    res["maximumAttendeesCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAttendeesCount(val)
        }
        return nil
    }
    res["optOutOfCustomerEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptOutOfCustomerEmail(val)
        }
        return nil
    }
    res["postBuffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostBuffer(val)
        }
        return nil
    }
    res["preBuffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreBuffer(val)
        }
        return nil
    }
    res["price"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrice(val)
        }
        return nil
    }
    res["priceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingPriceType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(BookingPriceType)
            m.SetPriceType(&cast)
        }
        return nil
    }
    res["reminders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingReminder() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingReminder, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingReminder))
            }
            m.SetReminders(res)
        }
        return nil
    }
    res["selfServiceAppointmentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelfServiceAppointmentId(val)
        }
        return nil
    }
    res["serviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceId(val)
        }
        return nil
    }
    res["serviceLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceLocation(val.(*Location))
        }
        return nil
    }
    res["serviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceName(val)
        }
        return nil
    }
    res["serviceNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceNotes(val)
        }
        return nil
    }
    res["smsNotificationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsNotificationsEnabled(val)
        }
        return nil
    }
    res["staffMemberIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetStaffMemberIds(res)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    return res
}
func (m *BookingAppointment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingAppointment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("additionalInformation", m.GetAdditionalInformation())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomers()))
        for i, v := range m.GetCustomers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerTimeZone", m.GetCustomerTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("filledAttendeesCount", m.GetFilledAttendeesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLocationOnline", m.GetIsLocationOnline())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumAttendeesCount", m.GetMaximumAttendeesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("optOutOfCustomerEmail", m.GetOptOutOfCustomerEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postBuffer", m.GetPostBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preBuffer", m.GetPreBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("price", m.GetPrice())
        if err != nil {
            return err
        }
    }
    if m.GetPriceType() != nil {
        cast := m.GetPriceType().String()
        err = writer.WriteStringValue("priceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReminders()))
        for i, v := range m.GetReminders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("reminders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("selfServiceAppointmentId", m.GetSelfServiceAppointmentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceId", m.GetServiceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("serviceLocation", m.GetServiceLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceName", m.GetServiceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceNotes", m.GetServiceNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smsNotificationsEnabled", m.GetSmsNotificationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("staffMemberIds", m.GetStaffMemberIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalInformation sets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
func (m *BookingAppointment) SetAdditionalInformation(value *string)() {
    if m != nil {
        m.additionalInformation = value
    }
}
// SetCustomers sets the customers property value. It lists down the customer properties for an appointment. An appointment will contain a list of customer information and each unit will indicate the properties of a customer who is part of that appointment. Optional.
func (m *BookingAppointment) SetCustomers(value []BookingCustomerInformationBase)() {
    if m != nil {
        m.customers = value
    }
}
// SetCustomerTimeZone sets the customerTimeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
func (m *BookingAppointment) SetCustomerTimeZone(value *string)() {
    if m != nil {
        m.customerTimeZone = value
    }
}
// SetDuration sets the duration property value. The length of the appointment, denoted in ISO8601 format.
func (m *BookingAppointment) SetDuration(value *string)() {
    if m != nil {
        m.duration = value
    }
}
// SetEndDateTime sets the endDateTime property value. 
func (m *BookingAppointment) SetEndDateTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetFilledAttendeesCount sets the filledAttendeesCount property value. The current number of customers in the appointment.
func (m *BookingAppointment) SetFilledAttendeesCount(value *int32)() {
    if m != nil {
        m.filledAttendeesCount = value
    }
}
// SetIsLocationOnline sets the isLocationOnline property value. True indicates that the appointment will be held online. Default value is false.
func (m *BookingAppointment) SetIsLocationOnline(value *bool)() {
    if m != nil {
        m.isLocationOnline = value
    }
}
// SetJoinWebUrl sets the joinWebUrl property value. The URL of the online meeting for the appointment.
func (m *BookingAppointment) SetJoinWebUrl(value *string)() {
    if m != nil {
        m.joinWebUrl = value
    }
}
// SetMaximumAttendeesCount sets the maximumAttendeesCount property value. The maximum number of customers allowed in an appointment.
func (m *BookingAppointment) SetMaximumAttendeesCount(value *int32)() {
    if m != nil {
        m.maximumAttendeesCount = value
    }
}
// SetOptOutOfCustomerEmail sets the optOutOfCustomerEmail property value. True indicates that the bookingCustomer for this appointment does not wish to receive a confirmation for this appointment.
func (m *BookingAppointment) SetOptOutOfCustomerEmail(value *bool)() {
    if m != nil {
        m.optOutOfCustomerEmail = value
    }
}
// SetPostBuffer sets the postBuffer property value. The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) SetPostBuffer(value *string)() {
    if m != nil {
        m.postBuffer = value
    }
}
// SetPreBuffer sets the preBuffer property value. The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) SetPreBuffer(value *string)() {
    if m != nil {
        m.preBuffer = value
    }
}
// SetPrice sets the price property value. The regular price for an appointment for the specified bookingService.
func (m *BookingAppointment) SetPrice(value *float64)() {
    if m != nil {
        m.price = value
    }
}
// SetPriceType sets the priceType property value. A setting to provide flexibility for the pricing structure of services. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet, unknownFutureValue.
func (m *BookingAppointment) SetPriceType(value *BookingPriceType)() {
    if m != nil {
        m.priceType = value
    }
}
// SetReminders sets the reminders property value. The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) SetReminders(value []BookingReminder)() {
    if m != nil {
        m.reminders = value
    }
}
// SetSelfServiceAppointmentId sets the selfServiceAppointmentId property value. An additional tracking ID for the appointment, if the appointment has been created directly by the customer on the scheduling page, as opposed to by a staff member on the behalf of the customer.
func (m *BookingAppointment) SetSelfServiceAppointmentId(value *string)() {
    if m != nil {
        m.selfServiceAppointmentId = value
    }
}
// SetServiceId sets the serviceId property value. The ID of the bookingService associated with this appointment.
func (m *BookingAppointment) SetServiceId(value *string)() {
    if m != nil {
        m.serviceId = value
    }
}
// SetServiceLocation sets the serviceLocation property value. The location where the service is delivered.
func (m *BookingAppointment) SetServiceLocation(value *Location)() {
    if m != nil {
        m.serviceLocation = value
    }
}
// SetServiceName sets the serviceName property value. The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it is computed from the service associated with the appointment by the serviceId property.
func (m *BookingAppointment) SetServiceName(value *string)() {
    if m != nil {
        m.serviceName = value
    }
}
// SetServiceNotes sets the serviceNotes property value. Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) SetServiceNotes(value *string)() {
    if m != nil {
        m.serviceNotes = value
    }
}
// SetSmsNotificationsEnabled sets the smsNotificationsEnabled property value. True indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
func (m *BookingAppointment) SetSmsNotificationsEnabled(value *bool)() {
    if m != nil {
        m.smsNotificationsEnabled = value
    }
}
// SetStaffMemberIds sets the staffMemberIds property value. The ID of each bookingStaffMember who is scheduled in this appointment.
func (m *BookingAppointment) SetStaffMemberIds(value []string)() {
    if m != nil {
        m.staffMemberIds = value
    }
}
// SetStartDateTime sets the startDateTime property value. 
func (m *BookingAppointment) SetStartDateTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.startDateTime = value
    }
}
