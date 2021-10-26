package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceActionResult struct {
    // Action name
    actionName *string;
    // State of the action. Possible values are: none, pending, canceled, active, done, failed, notSupported.
    actionState *ActionState;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Time the action state was last updated
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Time the action was initiated
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new deviceActionResult and sets the default values.
func NewDeviceActionResult()(*DeviceActionResult) {
    m := &DeviceActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the actionName property value. Action name
func (m *DeviceActionResult) GetActionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionName
    }
}
// Gets the actionState property value. State of the action. Possible values are: none, pending, canceled, active, done, failed, notSupported.
func (m *DeviceActionResult) GetActionState()(*ActionState) {
    if m == nil {
        return nil
    } else {
        return m.actionState
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceActionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the lastUpdatedDateTime property value. Time the action state was last updated
func (m *DeviceActionResult) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// Gets the startDateTime property value. Time the action was initiated
func (m *DeviceActionResult) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// The deserialization information for the current model
func (m *DeviceActionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActionName(val)
        return nil
    }
    res["actionState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionState)
        if err != nil {
            return err
        }
        cast := val.(ActionState)
        m.SetActionState(&cast)
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastUpdatedDateTime(val)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    return res
}
func (m *DeviceActionResult) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceActionResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionName", m.GetActionName())
        if err != nil {
            return err
        }
    }
    if m.GetActionState() != nil {
        cast := m.GetActionState().String()
        err := writer.WriteStringValue("actionState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
// Sets the actionName property value. Action name
// Parameters:
//  - value : Value to set for the actionName property.
func (m *DeviceActionResult) SetActionName(value *string)() {
    m.actionName = value
}
// Sets the actionState property value. State of the action. Possible values are: none, pending, canceled, active, done, failed, notSupported.
// Parameters:
//  - value : Value to set for the actionState property.
func (m *DeviceActionResult) SetActionState(value *ActionState)() {
    m.actionState = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceActionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the lastUpdatedDateTime property value. Time the action state was last updated
// Parameters:
//  - value : Value to set for the lastUpdatedDateTime property.
func (m *DeviceActionResult) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// Sets the startDateTime property value. Time the action was initiated
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *DeviceActionResult) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
