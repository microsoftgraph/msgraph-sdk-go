package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DataSubject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Email of the data subject.
    email *string;
    // First name of the data subject.
    firstName *string;
    // Last Name of the data subject.
    lastName *string;
    // The country/region of residency. The residency information is uesed only for internal reporting but not for the content search.
    residency *string;
}
// Instantiates a new dataSubject and sets the default values.
func NewDataSubject()(*DataSubject) {
    m := &DataSubject{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataSubject) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the email property value. Email of the data subject.
func (m *DataSubject) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the firstName property value. First name of the data subject.
func (m *DataSubject) GetFirstName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.firstName
    }
}
// Gets the lastName property value. Last Name of the data subject.
func (m *DataSubject) GetLastName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastName
    }
}
// Gets the residency property value. The country/region of residency. The residency information is uesed only for internal reporting but not for the content search.
func (m *DataSubject) GetResidency()(*string) {
    if m == nil {
        return nil
    } else {
        return m.residency
    }
}
// The deserialization information for the current model
func (m *DataSubject) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    res["firstName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFirstName(val)
        return nil
    }
    res["lastName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastName(val)
        return nil
    }
    res["residency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResidency(val)
        return nil
    }
    return res
}
func (m *DataSubject) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DataSubject) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firstName", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastName", m.GetLastName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("residency", m.GetResidency())
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
func (m *DataSubject) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the email property value. Email of the data subject.
// Parameters:
//  - value : Value to set for the email property.
func (m *DataSubject) SetEmail(value *string)() {
    m.email = value
}
// Sets the firstName property value. First name of the data subject.
// Parameters:
//  - value : Value to set for the firstName property.
func (m *DataSubject) SetFirstName(value *string)() {
    m.firstName = value
}
// Sets the lastName property value. Last Name of the data subject.
// Parameters:
//  - value : Value to set for the lastName property.
func (m *DataSubject) SetLastName(value *string)() {
    m.lastName = value
}
// Sets the residency property value. The country/region of residency. The residency information is uesed only for internal reporting but not for the content search.
// Parameters:
//  - value : Value to set for the residency property.
func (m *DataSubject) SetResidency(value *string)() {
    m.residency = value
}
