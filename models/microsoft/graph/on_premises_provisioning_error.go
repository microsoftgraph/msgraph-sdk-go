package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnPremisesProvisioningError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Category of the provisioning error. Note: Currently, there is only one possible value. Possible value: PropertyConflict - indicates a property value is not unique. Other objects contain the same value for the property.
    category *string;
    // The date and time at which the error occurred.
    occurredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the directory property causing the error. Current possible values: UserPrincipalName or ProxyAddress
    propertyCausingError *string;
    // Value of the property causing the error.
    value *string;
}
// Instantiates a new onPremisesProvisioningError and sets the default values.
func NewOnPremisesProvisioningError()(*OnPremisesProvisioningError) {
    m := &OnPremisesProvisioningError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesProvisioningError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the category property value. Category of the provisioning error. Note: Currently, there is only one possible value. Possible value: PropertyConflict - indicates a property value is not unique. Other objects contain the same value for the property.
func (m *OnPremisesProvisioningError) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the occurredDateTime property value. The date and time at which the error occurred.
func (m *OnPremisesProvisioningError) GetOccurredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.occurredDateTime
    }
}
// Gets the propertyCausingError property value. Name of the directory property causing the error. Current possible values: UserPrincipalName or ProxyAddress
func (m *OnPremisesProvisioningError) GetPropertyCausingError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.propertyCausingError
    }
}
// Gets the value property value. Value of the property causing the error.
func (m *OnPremisesProvisioningError) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
func (m *OnPremisesProvisioningError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory(val)
        return nil
    }
    res["occurredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetOccurredDateTime(val)
        return nil
    }
    res["propertyCausingError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPropertyCausingError(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *OnPremisesProvisioningError) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnPremisesProvisioningError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("occurredDateTime", m.GetOccurredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("propertyCausingError", m.GetPropertyCausingError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *OnPremisesProvisioningError) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the category property value. Category of the provisioning error. Note: Currently, there is only one possible value. Possible value: PropertyConflict - indicates a property value is not unique. Other objects contain the same value for the property.
// Parameters:
//  - value : Value to set for the category property.
func (m *OnPremisesProvisioningError) SetCategory(value *string)() {
    m.category = value
}
// Sets the occurredDateTime property value. The date and time at which the error occurred.
// Parameters:
//  - value : Value to set for the occurredDateTime property.
func (m *OnPremisesProvisioningError) SetOccurredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.occurredDateTime = value
}
// Sets the propertyCausingError property value. Name of the directory property causing the error. Current possible values: UserPrincipalName or ProxyAddress
// Parameters:
//  - value : Value to set for the propertyCausingError property.
func (m *OnPremisesProvisioningError) SetPropertyCausingError(value *string)() {
    m.propertyCausingError = value
}
// Sets the value property value. Value of the property causing the error.
// Parameters:
//  - value : Value to set for the value property.
func (m *OnPremisesProvisioningError) SetValue(value *string)() {
    m.value = value
}
