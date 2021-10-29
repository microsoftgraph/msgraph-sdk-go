package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsInformationProtectionDataRecoveryCertificate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Data recovery Certificate
    certificate []byte;
    // Data recovery Certificate description
    description *string;
    // Data recovery Certificate expiration datetime
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Data recovery Certificate subject name
    subjectName *string;
}
// Instantiates a new windowsInformationProtectionDataRecoveryCertificate and sets the default values.
func NewWindowsInformationProtectionDataRecoveryCertificate()(*WindowsInformationProtectionDataRecoveryCertificate) {
    m := &WindowsInformationProtectionDataRecoveryCertificate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the certificate property value. Data recovery Certificate
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetCertificate()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.certificate
    }
}
// Gets the description property value. Data recovery Certificate description
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the expirationDateTime property value. Data recovery Certificate expiration datetime
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the subjectName property value. Data recovery Certificate subject name
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetSubjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectName
    }
}
// The deserialization information for the current model
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetCertificate(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["subjectName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubjectName(val)
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionDataRecoveryCertificate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsInformationProtectionDataRecoveryCertificate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subjectName", m.GetSubjectName())
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
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the certificate property value. Data recovery Certificate
// Parameters:
//  - value : Value to set for the certificate property.
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetCertificate(value []byte)() {
    m.certificate = value
}
// Sets the description property value. Data recovery Certificate description
// Parameters:
//  - value : Value to set for the description property.
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetDescription(value *string)() {
    m.description = value
}
// Sets the expirationDateTime property value. Data recovery Certificate expiration datetime
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the subjectName property value. Data recovery Certificate subject name
// Parameters:
//  - value : Value to set for the subjectName property.
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetSubjectName(value *string)() {
    m.subjectName = value
}
