package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DataSubject provides operations to manage the privacy singleton.
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
// NewDataSubject instantiates a new dataSubject and sets the default values.
func NewDataSubject()(*DataSubject) {
    m := &DataSubject{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDataSubjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataSubjectFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDataSubject(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataSubject) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEmail gets the email property value. Email of the data subject.
func (m *DataSubject) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataSubject) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["firstName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
    res["lastName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
        }
        return nil
    }
    res["residency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResidency(val)
        }
        return nil
    }
    return res
}
// GetFirstName gets the firstName property value. First name of the data subject.
func (m *DataSubject) GetFirstName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.firstName
    }
}
// GetLastName gets the lastName property value. Last Name of the data subject.
func (m *DataSubject) GetLastName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastName
    }
}
// GetResidency gets the residency property value. The country/region of residency. The residency information is uesed only for internal reporting but not for the content search.
func (m *DataSubject) GetResidency()(*string) {
    if m == nil {
        return nil
    } else {
        return m.residency
    }
}
func (m *DataSubject) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataSubject) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEmail sets the email property value. Email of the data subject.
func (m *DataSubject) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetFirstName sets the firstName property value. First name of the data subject.
func (m *DataSubject) SetFirstName(value *string)() {
    if m != nil {
        m.firstName = value
    }
}
// SetLastName sets the lastName property value. Last Name of the data subject.
func (m *DataSubject) SetLastName(value *string)() {
    if m != nil {
        m.lastName = value
    }
}
// SetResidency sets the residency property value. The country/region of residency. The residency information is uesed only for internal reporting but not for the content search.
func (m *DataSubject) SetResidency(value *string)() {
    if m != nil {
        m.residency = value
    }
}
