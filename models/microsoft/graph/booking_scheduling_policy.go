package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingSchedulingPolicy provides operations to manage the solutionsRoot singleton.
type BookingSchedulingPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // True if to allow customers to choose a specific person for the booking.
    allowStaffSelection *bool;
    // Maximum number of days in advance that a booking can be made. It follows the ISO 8601 format.
    maximumAdvance *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The minimum amount of time before which bookings and cancellations must be made. It follows the ISO 8601 format.
    minimumLeadTime *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // True to notify the business via email when a booking is created or changed. Use the email address specified in the email property of the bookingBusiness entity for the business.
    sendConfirmationsToOwner *bool;
    // Duration of each time slot, denoted in ISO 8601 format.
    timeSlotInterval *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
}
// NewBookingSchedulingPolicy instantiates a new bookingSchedulingPolicy and sets the default values.
func NewBookingSchedulingPolicy()(*BookingSchedulingPolicy) {
    m := &BookingSchedulingPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBookingSchedulingPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingSchedulingPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBookingSchedulingPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingSchedulingPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowStaffSelection gets the allowStaffSelection property value. True if to allow customers to choose a specific person for the booking.
func (m *BookingSchedulingPolicy) GetAllowStaffSelection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowStaffSelection
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingSchedulingPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowStaffSelection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowStaffSelection(val)
        }
        return nil
    }
    res["maximumAdvance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAdvance(val)
        }
        return nil
    }
    res["minimumLeadTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumLeadTime(val)
        }
        return nil
    }
    res["sendConfirmationsToOwner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendConfirmationsToOwner(val)
        }
        return nil
    }
    res["timeSlotInterval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeSlotInterval(val)
        }
        return nil
    }
    return res
}
// GetMaximumAdvance gets the maximumAdvance property value. Maximum number of days in advance that a booking can be made. It follows the ISO 8601 format.
func (m *BookingSchedulingPolicy) GetMaximumAdvance()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.maximumAdvance
    }
}
// GetMinimumLeadTime gets the minimumLeadTime property value. The minimum amount of time before which bookings and cancellations must be made. It follows the ISO 8601 format.
func (m *BookingSchedulingPolicy) GetMinimumLeadTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.minimumLeadTime
    }
}
// GetSendConfirmationsToOwner gets the sendConfirmationsToOwner property value. True to notify the business via email when a booking is created or changed. Use the email address specified in the email property of the bookingBusiness entity for the business.
func (m *BookingSchedulingPolicy) GetSendConfirmationsToOwner()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendConfirmationsToOwner
    }
}
// GetTimeSlotInterval gets the timeSlotInterval property value. Duration of each time slot, denoted in ISO 8601 format.
func (m *BookingSchedulingPolicy) GetTimeSlotInterval()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.timeSlotInterval
    }
}
func (m *BookingSchedulingPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingSchedulingPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowStaffSelection", m.GetAllowStaffSelection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("maximumAdvance", m.GetMaximumAdvance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("minimumLeadTime", m.GetMinimumLeadTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendConfirmationsToOwner", m.GetSendConfirmationsToOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("timeSlotInterval", m.GetTimeSlotInterval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingSchedulingPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowStaffSelection sets the allowStaffSelection property value. True if to allow customers to choose a specific person for the booking.
func (m *BookingSchedulingPolicy) SetAllowStaffSelection(value *bool)() {
    if m != nil {
        m.allowStaffSelection = value
    }
}
// SetMaximumAdvance sets the maximumAdvance property value. Maximum number of days in advance that a booking can be made. It follows the ISO 8601 format.
func (m *BookingSchedulingPolicy) SetMaximumAdvance(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.maximumAdvance = value
    }
}
// SetMinimumLeadTime sets the minimumLeadTime property value. The minimum amount of time before which bookings and cancellations must be made. It follows the ISO 8601 format.
func (m *BookingSchedulingPolicy) SetMinimumLeadTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.minimumLeadTime = value
    }
}
// SetSendConfirmationsToOwner sets the sendConfirmationsToOwner property value. True to notify the business via email when a booking is created or changed. Use the email address specified in the email property of the bookingBusiness entity for the business.
func (m *BookingSchedulingPolicy) SetSendConfirmationsToOwner(value *bool)() {
    if m != nil {
        m.sendConfirmationsToOwner = value
    }
}
// SetTimeSlotInterval sets the timeSlotInterval property value. Duration of each time slot, denoted in ISO 8601 format.
func (m *BookingSchedulingPolicy) SetTimeSlotInterval(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.timeSlotInterval = value
    }
}
