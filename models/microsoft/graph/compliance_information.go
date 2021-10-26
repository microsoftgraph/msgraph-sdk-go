package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ComplianceInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Collection of the certification controls associated with certification
    certificationControls []CertificationControl;
    // Compliance certification name (for example, ISO 27018:2014, GDPR, FedRAMP, NIST 800-171)
    certificationName *string;
}
// Instantiates a new complianceInformation and sets the default values.
func NewComplianceInformation()(*ComplianceInformation) {
    m := &ComplianceInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComplianceInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the certificationControls property value. Collection of the certification controls associated with certification
func (m *ComplianceInformation) GetCertificationControls()([]CertificationControl) {
    if m == nil {
        return nil
    } else {
        return m.certificationControls
    }
}
// Gets the certificationName property value. Compliance certification name (for example, ISO 27018:2014, GDPR, FedRAMP, NIST 800-171)
func (m *ComplianceInformation) GetCertificationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificationName
    }
}
// The deserialization information for the current model
func (m *ComplianceInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certificationControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCertificationControl() })
        if err != nil {
            return err
        }
        res := make([]CertificationControl, len(val))
        for i, v := range val {
            res[i] = *(v.(*CertificationControl))
        }
        m.SetCertificationControls(res)
        return nil
    }
    res["certificationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificationName(val)
        return nil
    }
    return res
}
func (m *ComplianceInformation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ComplianceInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCertificationControls()))
        for i, v := range m.GetCertificationControls() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("certificationControls", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("certificationName", m.GetCertificationName())
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
func (m *ComplianceInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the certificationControls property value. Collection of the certification controls associated with certification
// Parameters:
//  - value : Value to set for the certificationControls property.
func (m *ComplianceInformation) SetCertificationControls(value []CertificationControl)() {
    m.certificationControls = value
}
// Sets the certificationName property value. Compliance certification name (for example, ISO 27018:2014, GDPR, FedRAMP, NIST 800-171)
// Parameters:
//  - value : Value to set for the certificationName property.
func (m *ComplianceInformation) SetCertificationName(value *string)() {
    m.certificationName = value
}
