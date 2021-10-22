package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CertificateAuthority struct {
    additionalData map[string]interface{};
    certificate []byte;
    certificateRevocationListUrl *string;
    deltaCertificateRevocationListUrl *string;
    isRootAuthority *bool;
    issuer *string;
    issuerSki *string;
}
func NewCertificateAuthority()(*CertificateAuthority) {
    m := &CertificateAuthority{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CertificateAuthority) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CertificateAuthority) GetCertificate()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.certificate
    }
}
func (m *CertificateAuthority) GetCertificateRevocationListUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateRevocationListUrl
    }
}
func (m *CertificateAuthority) GetDeltaCertificateRevocationListUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deltaCertificateRevocationListUrl
    }
}
func (m *CertificateAuthority) GetIsRootAuthority()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRootAuthority
    }
}
func (m *CertificateAuthority) GetIssuer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuer
    }
}
func (m *CertificateAuthority) GetIssuerSki()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuerSki
    }
}
func (m *CertificateAuthority) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetCertificate(val)
        return nil
    }
    res["certificateRevocationListUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificateRevocationListUrl(val)
        return nil
    }
    res["deltaCertificateRevocationListUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeltaCertificateRevocationListUrl(val)
        return nil
    }
    res["isRootAuthority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRootAuthority(val)
        return nil
    }
    res["issuer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIssuer(val)
        return nil
    }
    res["issuerSki"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIssuerSki(val)
        return nil
    }
    return res
}
func (m *CertificateAuthority) IsNil()(bool) {
    return m == nil
}
func (m *CertificateAuthority) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("certificateRevocationListUrl", m.GetCertificateRevocationListUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deltaCertificateRevocationListUrl", m.GetDeltaCertificateRevocationListUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRootAuthority", m.GetIsRootAuthority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issuerSki", m.GetIssuerSki())
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
func (m *CertificateAuthority) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CertificateAuthority) SetCertificate(value []byte)() {
    m.certificate = value
}
func (m *CertificateAuthority) SetCertificateRevocationListUrl(value *string)() {
    m.certificateRevocationListUrl = value
}
func (m *CertificateAuthority) SetDeltaCertificateRevocationListUrl(value *string)() {
    m.deltaCertificateRevocationListUrl = value
}
func (m *CertificateAuthority) SetIsRootAuthority(value *bool)() {
    m.isRootAuthority = value
}
func (m *CertificateAuthority) SetIssuer(value *string)() {
    m.issuer = value
}
func (m *CertificateAuthority) SetIssuerSki(value *string)() {
    m.issuerSki = value
}
