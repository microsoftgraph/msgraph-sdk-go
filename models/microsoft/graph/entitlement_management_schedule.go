package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EntitlementManagementSchedule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    expiration *ExpirationPattern;
    // 
    recurrence *PatternedRecurrence;
    // 
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new entitlementManagementSchedule and sets the default values.
func NewEntitlementManagementSchedule()(*EntitlementManagementSchedule) {
    m := &EntitlementManagementSchedule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EntitlementManagementSchedule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the expiration property value. 
func (m *EntitlementManagementSchedule) GetExpiration()(*ExpirationPattern) {
    if m == nil {
        return nil
    } else {
        return m.expiration
    }
}
// Gets the recurrence property value. 
func (m *EntitlementManagementSchedule) GetRecurrence()(*PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// Gets the startDateTime property value. 
func (m *EntitlementManagementSchedule) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// The deserialization information for the current model
func (m *EntitlementManagementSchedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expiration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExpirationPattern() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiration(val.(*ExpirationPattern))
        }
        return nil
    }
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPatternedRecurrence() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(*PatternedRecurrence))
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
func (m *EntitlementManagementSchedule) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EntitlementManagementSchedule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("expiration", m.GetExpiration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recurrence", m.GetRecurrence())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *EntitlementManagementSchedule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the expiration property value. 
// Parameters:
//  - value : Value to set for the expiration property.
func (m *EntitlementManagementSchedule) SetExpiration(value *ExpirationPattern)() {
    m.expiration = value
}
// Sets the recurrence property value. 
// Parameters:
//  - value : Value to set for the recurrence property.
func (m *EntitlementManagementSchedule) SetRecurrence(value *PatternedRecurrence)() {
    m.recurrence = value
}
// Sets the startDateTime property value. 
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *EntitlementManagementSchedule) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
