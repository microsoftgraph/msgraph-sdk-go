package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ComplianceInformation 
type ComplianceInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Collection of the certification controls associated with certification
    certificationControls []CertificationControl;
    // Compliance certification name (for example, ISO 27018:2014, GDPR, FedRAMP, NIST 800-171)
    certificationName *string;
}
// NewComplianceInformation instantiates a new complianceInformation and sets the default values.
func NewComplianceInformation()(*ComplianceInformation) {
    m := &ComplianceInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComplianceInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificationControls gets the certificationControls property value. Collection of the certification controls associated with certification
func (m *ComplianceInformation) GetCertificationControls()([]CertificationControl) {
    if m == nil {
        return nil
    } else {
        return m.certificationControls
    }
}
// GetCertificationName gets the certificationName property value. Compliance certification name (for example, ISO 27018:2014, GDPR, FedRAMP, NIST 800-171)
func (m *ComplianceInformation) GetCertificationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificationName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComplianceInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certificationControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCertificationControl() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificationControl, len(val))
            for i, v := range val {
                res[i] = *(v.(*CertificationControl))
            }
            m.SetCertificationControls(res)
        }
        return nil
    }
    res["certificationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationName(val)
        }
        return nil
    }
    return res
}
func (m *ComplianceInformation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComplianceInformation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertificationControls sets the certificationControls property value. Collection of the certification controls associated with certification
func (m *ComplianceInformation) SetCertificationControls(value []CertificationControl)() {
    if m != nil {
        m.certificationControls = value
    }
}
// SetCertificationName sets the certificationName property value. Compliance certification name (for example, ISO 27018:2014, GDPR, FedRAMP, NIST 800-171)
func (m *ComplianceInformation) SetCertificationName(value *string)() {
    if m != nil {
        m.certificationName = value
    }
}
