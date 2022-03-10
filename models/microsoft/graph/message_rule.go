package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MessageRule provides operations to manage the collection of drive entities.
type MessageRule struct {
    Entity
    // Actions to be taken on a message when the corresponding conditions are fulfilled.
    actions MessageRuleActionsable;
    // Conditions that when fulfilled, will trigger the corresponding actions for that rule.
    conditions MessageRulePredicatesable;
    // The display name of the rule.
    displayName *string;
    // Exception conditions for the rule.
    exceptions MessageRulePredicatesable;
    // Indicates whether the rule is in an error condition. Read-only.
    hasError *bool;
    // Indicates whether the rule is enabled to be applied to messages.
    isEnabled *bool;
    // Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
    isReadOnly *bool;
    // Indicates the order in which the rule is executed, among other rules.
    sequence *int32;
}
// NewMessageRule instantiates a new messageRule and sets the default values.
func NewMessageRule()(*MessageRule) {
    m := &MessageRule{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMessageRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageRuleFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMessageRule(), nil
}
// GetActions gets the actions property value. Actions to be taken on a message when the corresponding conditions are fulfilled.
func (m *MessageRule) GetActions()(MessageRuleActionsable) {
    if m == nil {
        return nil
    } else {
        return m.actions
    }
}
// GetConditions gets the conditions property value. Conditions that when fulfilled, will trigger the corresponding actions for that rule.
func (m *MessageRule) GetConditions()(MessageRulePredicatesable) {
    if m == nil {
        return nil
    } else {
        return m.conditions
    }
}
// GetDisplayName gets the displayName property value. The display name of the rule.
func (m *MessageRule) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExceptions gets the exceptions property value. Exception conditions for the rule.
func (m *MessageRule) GetExceptions()(MessageRulePredicatesable) {
    if m == nil {
        return nil
    } else {
        return m.exceptions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateMessageRuleActionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActions(val.(MessageRuleActionsable))
        }
        return nil
    }
    res["conditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateMessageRulePredicatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(MessageRulePredicatesable))
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
        val, err := n.GetObjectValue(CreateMessageRulePredicatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExceptions(val.(MessageRulePredicatesable))
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
// GetHasError gets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
func (m *MessageRule) GetHasError()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasError
    }
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
func (m *MessageRule) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetIsReadOnly gets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
func (m *MessageRule) GetIsReadOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReadOnly
    }
}
// GetSequence gets the sequence property value. Indicates the order in which the rule is executed, among other rules.
func (m *MessageRule) GetSequence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sequence
    }
}
func (m *MessageRule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetActions sets the actions property value. Actions to be taken on a message when the corresponding conditions are fulfilled.
func (m *MessageRule) SetActions(value MessageRuleActionsable)() {
    if m != nil {
        m.actions = value
    }
}
// SetConditions sets the conditions property value. Conditions that when fulfilled, will trigger the corresponding actions for that rule.
func (m *MessageRule) SetConditions(value MessageRulePredicatesable)() {
    if m != nil {
        m.conditions = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the rule.
func (m *MessageRule) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExceptions sets the exceptions property value. Exception conditions for the rule.
func (m *MessageRule) SetExceptions(value MessageRulePredicatesable)() {
    if m != nil {
        m.exceptions = value
    }
}
// SetHasError sets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
func (m *MessageRule) SetHasError(value *bool)() {
    if m != nil {
        m.hasError = value
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
func (m *MessageRule) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetIsReadOnly sets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
func (m *MessageRule) SetIsReadOnly(value *bool)() {
    if m != nil {
        m.isReadOnly = value
    }
}
// SetSequence sets the sequence property value. Indicates the order in which the rule is executed, among other rules.
func (m *MessageRule) SetSequence(value *int32)() {
    if m != nil {
        m.sequence = value
    }
}
