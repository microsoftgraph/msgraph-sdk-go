package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConditionalAccessFilter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Mode to use for the filter. Possible values are include or exclude.
    mode *FilterMode;
    // Rule syntax is similar to that used for membership rules for groups in Azure Active Directory (Azure AD). For details, see rules with multiple expressions
    rule *string;
}
// Instantiates a new conditionalAccessFilter and sets the default values.
func NewConditionalAccessFilter()(*ConditionalAccessFilter) {
    m := &ConditionalAccessFilter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessFilter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the mode property value. Mode to use for the filter. Possible values are include or exclude.
func (m *ConditionalAccessFilter) GetMode()(*FilterMode) {
    if m == nil {
        return nil
    } else {
        return m.mode
    }
}
// Gets the rule property value. Rule syntax is similar to that used for membership rules for groups in Azure Active Directory (Azure AD). For details, see rules with multiple expressions
func (m *ConditionalAccessFilter) GetRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rule
    }
}
// The deserialization information for the current model
func (m *ConditionalAccessFilter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["mode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFilterMode)
        if err != nil {
            return err
        }
        cast := val.(FilterMode)
        m.SetMode(&cast)
        return nil
    }
    res["rule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRule(val)
        return nil
    }
    return res
}
func (m *ConditionalAccessFilter) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConditionalAccessFilter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetMode() != nil {
        cast := m.GetMode().String()
        err := writer.WriteStringValue("mode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rule", m.GetRule())
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
func (m *ConditionalAccessFilter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the mode property value. Mode to use for the filter. Possible values are include or exclude.
// Parameters:
//  - value : Value to set for the mode property.
func (m *ConditionalAccessFilter) SetMode(value *FilterMode)() {
    m.mode = value
}
// Sets the rule property value. Rule syntax is similar to that used for membership rules for groups in Azure Active Directory (Azure AD). For details, see rules with multiple expressions
// Parameters:
//  - value : Value to set for the rule property.
func (m *ConditionalAccessFilter) SetRule(value *string)() {
    m.rule = value
}
