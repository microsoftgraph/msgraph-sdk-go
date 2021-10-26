package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IncompleteData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The service does not have source data before the specified time.
    missingDataBeforeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Some data was not recorded due to excessive activity.
    wasThrottled *bool;
}
// Instantiates a new incompleteData and sets the default values.
func NewIncompleteData()(*IncompleteData) {
    m := &IncompleteData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IncompleteData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the missingDataBeforeDateTime property value. The service does not have source data before the specified time.
func (m *IncompleteData) GetMissingDataBeforeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.missingDataBeforeDateTime
    }
}
// Gets the wasThrottled property value. Some data was not recorded due to excessive activity.
func (m *IncompleteData) GetWasThrottled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wasThrottled
    }
}
// The deserialization information for the current model
func (m *IncompleteData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["missingDataBeforeDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetMissingDataBeforeDateTime(val)
        return nil
    }
    res["wasThrottled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetWasThrottled(val)
        return nil
    }
    return res
}
func (m *IncompleteData) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IncompleteData) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("missingDataBeforeDateTime", m.GetMissingDataBeforeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("wasThrottled", m.GetWasThrottled())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *IncompleteData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the missingDataBeforeDateTime property value. The service does not have source data before the specified time.
// Parameters:
//  - value : Value to set for the missingDataBeforeDateTime property.
func (m *IncompleteData) SetMissingDataBeforeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.missingDataBeforeDateTime = value
}
// Sets the wasThrottled property value. Some data was not recorded due to excessive activity.
// Parameters:
//  - value : Value to set for the wasThrottled property.
func (m *IncompleteData) SetWasThrottled(value *bool)() {
    m.wasThrottled = value
}
