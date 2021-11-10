package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MessageRule struct {
    Entity
    // Actions to be taken on a message when the corresponding conditions are fulfilled.
    actions *MessageRuleActions;
    // Conditions that when fulfilled, will trigger the corresponding actions for that rule.
    conditions *MessageRulePredicates;
    // The display name of the rule.
    displayName *string;
    // Exception conditions for the rule.
    exceptions *MessageRulePredicates;
    // Indicates whether the rule is in an error condition. Read-only.
    hasError *bool;
    // Indicates whether the rule is enabled to be applied to messages.
    isEnabled *bool;
    // Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
    isReadOnly *bool;
    // Indicates the order in which the rule is executed, among other rules.
    sequence *int32;
}
// Instantiates a new messageRule and sets the default values.
func NewMessageRule()(*MessageRule) {
    m := &MessageRule{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the actions property value. Actions to be taken on a message when the corresponding conditions are fulfilled.
func (m *MessageRule) GetActions()(*MessageRuleActions) {
    if m == nil {
        return nil
    } else {
        return m.actions
    }
}
// Gets the conditions property value. Conditions that when fulfilled, will trigger the corresponding actions for that rule.
func (m *MessageRule) GetConditions()(*MessageRulePredicates) {
    if m == nil {
        return nil
    } else {
        return m.conditions
    }
}
// Gets the displayName property value. The display name of the rule.
func (m *MessageRule) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the exceptions property value. Exception conditions for the rule.
func (m *MessageRule) GetExceptions()(*MessageRulePredicates) {
    if m == nil {
        return nil
    } else {
        return m.exceptions
    }
}
// Gets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
func (m *MessageRule) GetHasError()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasError
    }
}
// Gets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
func (m *MessageRule) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Gets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
func (m *MessageRule) GetIsReadOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReadOnly
    }
}
// Gets the sequence property value. Indicates the order in which the rule is executed, among other rules.
func (m *MessageRule) GetSequence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sequence
    }
}
// The deserialization information for the current model
func (m *MessageRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessageRuleActions() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActions(val.(*MessageRuleActions))
        }
        return nil
    }
    res["conditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessageRulePredicates() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(*MessageRulePredicates))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["exceptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessageRulePredicates() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExceptions(val.(*MessageRulePredicates))
        }
        return nil
    }
    res["hasError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasError(val)
        }
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["isReadOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReadOnly(val)
        }
        return nil
    }
    res["sequence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequence(val)
        }
        return nil
    }
    return res
}
func (m *MessageRule) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MessageRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("actions", m.GetActions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditions", m.GetConditions())
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
        err = writer.WriteObjectValue("exceptions", m.GetExceptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasError", m.GetHasError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReadOnly", m.GetIsReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sequence", m.GetSequence())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the actions property value. Actions to be taken on a message when the corresponding conditions are fulfilled.
// Parameters:
//  - value : Value to set for the actions property.
func (m *MessageRule) SetActions(value *MessageRuleActions)() {
    m.actions = value
}
// Sets the conditions property value. Conditions that when fulfilled, will trigger the corresponding actions for that rule.
// Parameters:
//  - value : Value to set for the conditions property.
func (m *MessageRule) SetConditions(value *MessageRulePredicates)() {
    m.conditions = value
}
// Sets the displayName property value. The display name of the rule.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *MessageRule) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the exceptions property value. Exception conditions for the rule.
// Parameters:
//  - value : Value to set for the exceptions property.
func (m *MessageRule) SetExceptions(value *MessageRulePredicates)() {
    m.exceptions = value
}
// Sets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
// Parameters:
//  - value : Value to set for the hasError property.
func (m *MessageRule) SetHasError(value *bool)() {
    m.hasError = value
}
// Sets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *MessageRule) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// Sets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
// Parameters:
//  - value : Value to set for the isReadOnly property.
func (m *MessageRule) SetIsReadOnly(value *bool)() {
    m.isReadOnly = value
}
// Sets the sequence property value. Indicates the order in which the rule is executed, among other rules.
// Parameters:
//  - value : Value to set for the sequence property.
func (m *MessageRule) SetSequence(value *int32)() {
    m.sequence = value
}
