package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConditionalAccessGrantControls struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of values of built-in controls required by the policy. Possible values: block, mfa, compliantDevice, domainJoinedDevice, approvedApplication, compliantApplication, passwordChange, unknownFutureValue.
    builtInControls []ConditionalAccessGrantControl;
    // List of custom controls IDs required by the policy. For more information, see Custom controls.
    customAuthenticationFactors []string;
    // Defines the relationship of the grant controls. Possible values: AND, OR.
    operator *string;
    // List of terms of use IDs required by the policy.
    termsOfUse []string;
}
// Instantiates a new conditionalAccessGrantControls and sets the default values.
func NewConditionalAccessGrantControls()(*ConditionalAccessGrantControls) {
    m := &ConditionalAccessGrantControls{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessGrantControls) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the builtInControls property value. List of values of built-in controls required by the policy. Possible values: block, mfa, compliantDevice, domainJoinedDevice, approvedApplication, compliantApplication, passwordChange, unknownFutureValue.
func (m *ConditionalAccessGrantControls) GetBuiltInControls()([]ConditionalAccessGrantControl) {
    if m == nil {
        return nil
    } else {
        return m.builtInControls
    }
}
// Gets the customAuthenticationFactors property value. List of custom controls IDs required by the policy. For more information, see Custom controls.
func (m *ConditionalAccessGrantControls) GetCustomAuthenticationFactors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.customAuthenticationFactors
    }
}
// Gets the operator property value. Defines the relationship of the grant controls. Possible values: AND, OR.
func (m *ConditionalAccessGrantControls) GetOperator()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operator
    }
}
// Gets the termsOfUse property value. List of terms of use IDs required by the policy.
func (m *ConditionalAccessGrantControls) GetTermsOfUse()([]string) {
    if m == nil {
        return nil
    } else {
        return m.termsOfUse
    }
}
// The deserialization information for the current model
func (m *ConditionalAccessGrantControls) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["builtInControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseConditionalAccessGrantControl)
        if err != nil {
            return err
        }
        res := make([]ConditionalAccessGrantControl, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConditionalAccessGrantControl))
        }
        m.SetBuiltInControls(res)
        return nil
    }
    res["customAuthenticationFactors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetCustomAuthenticationFactors(res)
        return nil
    }
    res["operator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperator(val)
        return nil
    }
    res["termsOfUse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTermsOfUse(res)
        return nil
    }
    return res
}
func (m *ConditionalAccessGrantControls) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConditionalAccessGrantControls) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("builtInControls", SerializeConditionalAccessGrantControl(m.GetBuiltInControls()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("customAuthenticationFactors", m.GetCustomAuthenticationFactors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operator", m.GetOperator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("termsOfUse", m.GetTermsOfUse())
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
func (m *ConditionalAccessGrantControls) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the builtInControls property value. List of values of built-in controls required by the policy. Possible values: block, mfa, compliantDevice, domainJoinedDevice, approvedApplication, compliantApplication, passwordChange, unknownFutureValue.
// Parameters:
//  - value : Value to set for the builtInControls property.
func (m *ConditionalAccessGrantControls) SetBuiltInControls(value []ConditionalAccessGrantControl)() {
    m.builtInControls = value
}
// Sets the customAuthenticationFactors property value. List of custom controls IDs required by the policy. For more information, see Custom controls.
// Parameters:
//  - value : Value to set for the customAuthenticationFactors property.
func (m *ConditionalAccessGrantControls) SetCustomAuthenticationFactors(value []string)() {
    m.customAuthenticationFactors = value
}
// Sets the operator property value. Defines the relationship of the grant controls. Possible values: AND, OR.
// Parameters:
//  - value : Value to set for the operator property.
func (m *ConditionalAccessGrantControls) SetOperator(value *string)() {
    m.operator = value
}
// Sets the termsOfUse property value. List of terms of use IDs required by the policy.
// Parameters:
//  - value : Value to set for the termsOfUse property.
func (m *ConditionalAccessGrantControls) SetTermsOfUse(value []string)() {
    m.termsOfUse = value
}
