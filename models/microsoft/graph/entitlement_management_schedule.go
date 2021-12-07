package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EntitlementManagementSchedule 
type EntitlementManagementSchedule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // When the access should expire.
    expiration *ExpirationPattern;
    // For recurring access. Not used at present.
    recurrence *PatternedRecurrence;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewEntitlementManagementSchedule instantiates a new entitlementManagementSchedule and sets the default values.
func NewEntitlementManagementSchedule()(*EntitlementManagementSchedule) {
    m := &EntitlementManagementSchedule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EntitlementManagementSchedule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExpiration gets the expiration property value. When the access should expire.
func (m *EntitlementManagementSchedule) GetExpiration()(*ExpirationPattern) {
    if m == nil {
        return nil
    } else {
        return m.expiration
    }
}
// GetRecurrence gets the recurrence property value. For recurring access. Not used at present.
func (m *EntitlementManagementSchedule) GetRecurrence()(*PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// GetStartDateTime gets the startDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *EntitlementManagementSchedule) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EntitlementManagementSchedule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExpiration sets the expiration property value. When the access should expire.
func (m *EntitlementManagementSchedule) SetExpiration(value *ExpirationPattern)() {
    if m != nil {
        m.expiration = value
    }
}
// SetRecurrence sets the recurrence property value. For recurring access. Not used at present.
func (m *EntitlementManagementSchedule) SetRecurrence(value *PatternedRecurrence)() {
    if m != nil {
        m.recurrence = value
    }
}
// SetStartDateTime sets the startDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *EntitlementManagementSchedule) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
