package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// deviceManagementTroubleshootingEvent 
type DeviceManagementTroubleshootingEvent struct {
    Entity
    // Id used for tracing the failure in the service.
    correlationId *string;
    // Time when the event occurred .
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewDeviceManagementTroubleshootingEvent instantiates a new deviceManagementTroubleshootingEvent and sets the default values.
func NewDeviceManagementTroubleshootingEvent()(*DeviceManagementTroubleshootingEvent) {
    m := &DeviceManagementTroubleshootingEvent{
        Entity: *NewEntity(),
    }
    return m
}
// GetCorrelationId gets the correlationId property value. Id used for tracing the failure in the service.
func (m *DeviceManagementTroubleshootingEvent) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// GetEventDateTime gets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementTroubleshootingEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTroubleshootingEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["correlationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
        }
        return nil
    }
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementTroubleshootingEvent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementTroubleshootingEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCorrelationId sets the correlationId property value. Id used for tracing the failure in the service.
func (m *DeviceManagementTroubleshootingEvent) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// SetEventDateTime sets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementTroubleshootingEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
