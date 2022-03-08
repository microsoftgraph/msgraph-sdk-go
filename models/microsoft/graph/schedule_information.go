package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ScheduleInformation provides operations to call the getSchedule method.
type ScheduleInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Represents a merged view of availability of all the items in scheduleItems. The view consists of time slots. Availability during each time slot is indicated with: 0= free, 1= tentative, 2= busy, 3= out of office, 4= working elsewhere.
    availabilityView *string;
    // Error information from attempting to get the availability of the user, distribution list, or resource.
    error FreeBusyErrorable;
    // An SMTP address of the user, distribution list, or resource, identifying an instance of scheduleInformation.
    scheduleId *string;
    // Contains the items that describe the availability of the user or resource.
    scheduleItems []ScheduleItemable;
    // The days of the week and hours in a specific time zone that the user works. These are set as part of the user's mailboxSettings.
    workingHours WorkingHoursable;
}
// NewScheduleInformation instantiates a new scheduleInformation and sets the default values.
func NewScheduleInformation()(*ScheduleInformation) {
    m := &ScheduleInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateScheduleInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScheduleInformationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewScheduleInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ScheduleInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAvailabilityView gets the availabilityView property value. Represents a merged view of availability of all the items in scheduleItems. The view consists of time slots. Availability during each time slot is indicated with: 0= free, 1= tentative, 2= busy, 3= out of office, 4= working elsewhere.
func (m *ScheduleInformation) GetAvailabilityView()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availabilityView
    }
}
// GetError gets the error property value. Error information from attempting to get the availability of the user, distribution list, or resource.
func (m *ScheduleInformation) GetError()(FreeBusyErrorable) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ScheduleInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["availabilityView"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityView(val)
        }
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateFreeBusyErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(FreeBusyErrorable))
        }
        return nil
    }
    res["scheduleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleId(val)
        }
        return nil
    }
    res["scheduleItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateScheduleItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScheduleItemable, len(val))
            for i, v := range val {
                res[i] = v.(ScheduleItemable)
            }
            m.SetScheduleItems(res)
        }
        return nil
    }
    res["workingHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkingHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkingHours(val.(WorkingHoursable))
        }
        return nil
    }
    return res
}
// GetScheduleId gets the scheduleId property value. An SMTP address of the user, distribution list, or resource, identifying an instance of scheduleInformation.
func (m *ScheduleInformation) GetScheduleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scheduleId
    }
}
// GetScheduleItems gets the scheduleItems property value. Contains the items that describe the availability of the user or resource.
func (m *ScheduleInformation) GetScheduleItems()([]ScheduleItemable) {
    if m == nil {
        return nil
    } else {
        return m.scheduleItems
    }
}
// GetWorkingHours gets the workingHours property value. The days of the week and hours in a specific time zone that the user works. These are set as part of the user's mailboxSettings.
func (m *ScheduleInformation) GetWorkingHours()(WorkingHoursable) {
    if m == nil {
        return nil
    } else {
        return m.workingHours
    }
}
func (m *ScheduleInformation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ScheduleInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("availabilityView", m.GetAvailabilityView())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scheduleId", m.GetScheduleId())
        if err != nil {
            return err
        }
    }
    if m.GetScheduleItems() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScheduleItems()))
        for i, v := range m.GetScheduleItems() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("scheduleItems", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("workingHours", m.GetWorkingHours())
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
func (m *ScheduleInformation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAvailabilityView sets the availabilityView property value. Represents a merged view of availability of all the items in scheduleItems. The view consists of time slots. Availability during each time slot is indicated with: 0= free, 1= tentative, 2= busy, 3= out of office, 4= working elsewhere.
func (m *ScheduleInformation) SetAvailabilityView(value *string)() {
    if m != nil {
        m.availabilityView = value
    }
}
// SetError sets the error property value. Error information from attempting to get the availability of the user, distribution list, or resource.
func (m *ScheduleInformation) SetError(value FreeBusyErrorable)() {
    if m != nil {
        m.error = value
    }
}
// SetScheduleId sets the scheduleId property value. An SMTP address of the user, distribution list, or resource, identifying an instance of scheduleInformation.
func (m *ScheduleInformation) SetScheduleId(value *string)() {
    if m != nil {
        m.scheduleId = value
    }
}
// SetScheduleItems sets the scheduleItems property value. Contains the items that describe the availability of the user or resource.
func (m *ScheduleInformation) SetScheduleItems(value []ScheduleItemable)() {
    if m != nil {
        m.scheduleItems = value
    }
}
// SetWorkingHours sets the workingHours property value. The days of the week and hours in a specific time zone that the user works. These are set as part of the user's mailboxSettings.
func (m *ScheduleInformation) SetWorkingHours(value WorkingHoursable)() {
    if m != nil {
        m.workingHours = value
    }
}
