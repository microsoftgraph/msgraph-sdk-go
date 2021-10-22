package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ComplianceInformation struct {
    additionalData map[string]interface{};
    certificationControls []CertificationControl;
    certificationName *string;
}
func NewComplianceInformation()(*ComplianceInformation) {
    m := &ComplianceInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ComplianceInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ComplianceInformation) GetCertificationControls()([]CertificationControl) {
    if m == nil {
        return nil
    } else {
        return m.certificationControls
    }
}
func (m *ComplianceInformation) GetCertificationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificationName
    }
}
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
func (m *ComplianceInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ComplianceInformation) SetCertificationControls(value []CertificationControl)() {
    m.certificationControls = value
}
func (m *ComplianceInformation) SetCertificationName(value *string)() {
    m.certificationName = value
}
