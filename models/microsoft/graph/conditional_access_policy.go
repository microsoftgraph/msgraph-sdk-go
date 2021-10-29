package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConditionalAccessPolicy struct {
    Entity
    // 
    conditions *ConditionalAccessConditionSet;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Not used.
    description *string;
    // Specifies a display name for the conditionalAccessPolicy object.
    displayName *string;
    // Specifies the grant controls that must be fulfilled to pass the policy.
    grantControls *ConditionalAccessGrantControls;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Specifies the session controls that are enforced after sign-in.
    sessionControls *ConditionalAccessSessionControls;
    // Specifies the state of the conditionalAccessPolicy object. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required.
    state *ConditionalAccessPolicyState;
}
// Instantiates a new conditionalAccessPolicy and sets the default values.
func NewConditionalAccessPolicy()(*ConditionalAccessPolicy) {
    m := &ConditionalAccessPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the conditions property value. 
func (m *ConditionalAccessPolicy) GetConditions()(*ConditionalAccessConditionSet) {
    if m == nil {
        return nil
    } else {
        return m.conditions
    }
}
// Gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
func (m *ConditionalAccessPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Not used.
func (m *ConditionalAccessPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Specifies a display name for the conditionalAccessPolicy object.
func (m *ConditionalAccessPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the grantControls property value. Specifies the grant controls that must be fulfilled to pass the policy.
func (m *ConditionalAccessPolicy) GetGrantControls()(*ConditionalAccessGrantControls) {
    if m == nil {
        return nil
    } else {
        return m.grantControls
    }
}
// Gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
func (m *ConditionalAccessPolicy) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// Gets the sessionControls property value. Specifies the session controls that are enforced after sign-in.
func (m *ConditionalAccessPolicy) GetSessionControls()(*ConditionalAccessSessionControls) {
    if m == nil {
        return nil
    } else {
        return m.sessionControls
    }
}
// Gets the state property value. Specifies the state of the conditionalAccessPolicy object. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required.
func (m *ConditionalAccessPolicy) GetState()(*ConditionalAccessPolicyState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *ConditionalAccessPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessConditionSet() })
        if err != nil {
            return err
        }
        m.SetConditions(val.(*ConditionalAccessConditionSet))
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["grantControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessGrantControls() })
        if err != nil {
            return err
        }
        m.SetGrantControls(val.(*ConditionalAccessGrantControls))
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetModifiedDateTime(val)
        return nil
    }
    res["sessionControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessSessionControls() })
        if err != nil {
            return err
        }
        m.SetSessionControls(val.(*ConditionalAccessSessionControls))
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessPolicyState)
        if err != nil {
            return err
        }
        cast := val.(ConditionalAccessPolicyState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *ConditionalAccessPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConditionalAccessPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("conditions", m.GetConditions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("grantControls", m.GetGrantControls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sessionControls", m.GetSessionControls())
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
// Sets the conditions property value. 
// Parameters:
//  - value : Value to set for the conditions property.
func (m *ConditionalAccessPolicy) SetConditions(value *ConditionalAccessConditionSet)() {
    m.conditions = value
}
// Sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ConditionalAccessPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Not used.
// Parameters:
//  - value : Value to set for the description property.
func (m *ConditionalAccessPolicy) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Specifies a display name for the conditionalAccessPolicy object.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ConditionalAccessPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the grantControls property value. Specifies the grant controls that must be fulfilled to pass the policy.
// Parameters:
//  - value : Value to set for the grantControls property.
func (m *ConditionalAccessPolicy) SetGrantControls(value *ConditionalAccessGrantControls)() {
    m.grantControls = value
}
// Sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
// Parameters:
//  - value : Value to set for the modifiedDateTime property.
func (m *ConditionalAccessPolicy) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// Sets the sessionControls property value. Specifies the session controls that are enforced after sign-in.
// Parameters:
//  - value : Value to set for the sessionControls property.
func (m *ConditionalAccessPolicy) SetSessionControls(value *ConditionalAccessSessionControls)() {
    m.sessionControls = value
}
// Sets the state property value. Specifies the state of the conditionalAccessPolicy object. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required.
// Parameters:
//  - value : Value to set for the state property.
func (m *ConditionalAccessPolicy) SetState(value *ConditionalAccessPolicyState)() {
    m.state = value
}
