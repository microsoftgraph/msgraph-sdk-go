package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ScheduleChangeRequest struct {
    ChangeTrackedEntity
    // 
    assignedTo *ScheduleChangeRequestActor;
    // 
    managerActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    managerActionMessage *string;
    // 
    managerUserId *string;
    // 
    senderDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    senderMessage *string;
    // 
    senderUserId *string;
    // 
    state *ScheduleChangeState;
}
// Instantiates a new scheduleChangeRequest and sets the default values.
func NewScheduleChangeRequest()(*ScheduleChangeRequest) {
    m := &ScheduleChangeRequest{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// Gets the assignedTo property value. 
func (m *ScheduleChangeRequest) GetAssignedTo()(*ScheduleChangeRequestActor) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// Gets the managerActionDateTime property value. 
func (m *ScheduleChangeRequest) GetManagerActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.managerActionDateTime
    }
}
// Gets the managerActionMessage property value. 
func (m *ScheduleChangeRequest) GetManagerActionMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managerActionMessage
    }
}
// Gets the managerUserId property value. 
func (m *ScheduleChangeRequest) GetManagerUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managerUserId
    }
}
// Gets the senderDateTime property value. 
func (m *ScheduleChangeRequest) GetSenderDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.senderDateTime
    }
}
// Gets the senderMessage property value. 
func (m *ScheduleChangeRequest) GetSenderMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.senderMessage
    }
}
// Gets the senderUserId property value. 
func (m *ScheduleChangeRequest) GetSenderUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.senderUserId
    }
}
// Gets the state property value. 
func (m *ScheduleChangeRequest) GetState()(*ScheduleChangeState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *ScheduleChangeRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseScheduleChangeRequestActor)
        if err != nil {
            return err
        }
        cast := val.(ScheduleChangeRequestActor)
        m.SetAssignedTo(&cast)
        return nil
    }
    res["managerActionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetManagerActionDateTime(val)
        return nil
    }
    res["managerActionMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagerActionMessage(val)
        return nil
    }
    res["managerUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagerUserId(val)
        return nil
    }
    res["senderDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSenderDateTime(val)
        return nil
    }
    res["senderMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSenderMessage(val)
        return nil
    }
    res["senderUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSenderUserId(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseScheduleChangeState)
        if err != nil {
            return err
        }
        cast := val.(ScheduleChangeState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *ScheduleChangeRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ScheduleChangeRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedTo() != nil {
        cast := m.GetAssignedTo().String()
        err = writer.WriteStringValue("assignedTo", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("managerActionDateTime", m.GetManagerActionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managerActionMessage", m.GetManagerActionMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managerUserId", m.GetManagerUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("senderDateTime", m.GetSenderDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderMessage", m.GetSenderMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderUserId", m.GetSenderUserId())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignedTo property value. 
// Parameters:
//  - value : Value to set for the assignedTo property.
func (m *ScheduleChangeRequest) SetAssignedTo(value *ScheduleChangeRequestActor)() {
    m.assignedTo = value
}
// Sets the managerActionDateTime property value. 
// Parameters:
//  - value : Value to set for the managerActionDateTime property.
func (m *ScheduleChangeRequest) SetManagerActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.managerActionDateTime = value
}
// Sets the managerActionMessage property value. 
// Parameters:
//  - value : Value to set for the managerActionMessage property.
func (m *ScheduleChangeRequest) SetManagerActionMessage(value *string)() {
    m.managerActionMessage = value
}
// Sets the managerUserId property value. 
// Parameters:
//  - value : Value to set for the managerUserId property.
func (m *ScheduleChangeRequest) SetManagerUserId(value *string)() {
    m.managerUserId = value
}
// Sets the senderDateTime property value. 
// Parameters:
//  - value : Value to set for the senderDateTime property.
func (m *ScheduleChangeRequest) SetSenderDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.senderDateTime = value
}
// Sets the senderMessage property value. 
// Parameters:
//  - value : Value to set for the senderMessage property.
func (m *ScheduleChangeRequest) SetSenderMessage(value *string)() {
    m.senderMessage = value
}
// Sets the senderUserId property value. 
// Parameters:
//  - value : Value to set for the senderUserId property.
func (m *ScheduleChangeRequest) SetSenderUserId(value *string)() {
    m.senderUserId = value
}
// Sets the state property value. 
// Parameters:
//  - value : Value to set for the state property.
func (m *ScheduleChangeRequest) SetState(value *ScheduleChangeState)() {
    m.state = value
}
