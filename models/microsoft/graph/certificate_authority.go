package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CertificateAuthority struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Required. The base64 encoded string representing the public certificate.
    certificate []byte;
    // The URL of the certificate revocation list.
    certificateRevocationListUrl *string;
    // The URL contains the list of all revoked certificates since the last time a full certificate revocaton list was created.
    deltaCertificateRevocationListUrl *string;
    // Required. true if the trusted certificate is a root authority, false if the trusted certificate is an intermediate authority.
    isRootAuthority *bool;
    // The issuer of the certificate, calculated from the certificate value. Read-only.
    issuer *string;
    // The subject key identifier of the certificate, calculated from the certificate value. Read-only.
    issuerSki *string;
}
// Instantiates a new certificateAuthority and sets the default values.
func NewCertificateAuthority()(*CertificateAuthority) {
    m := &CertificateAuthority{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateAuthority) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the certificate property value. Required. The base64 encoded string representing the public certificate.
func (m *CertificateAuthority) GetCertificate()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.certificate
    }
}
// Gets the certificateRevocationListUrl property value. The URL of the certificate revocation list.
func (m *CertificateAuthority) GetCertificateRevocationListUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateRevocationListUrl
    }
}
// Gets the deltaCertificateRevocationListUrl property value. The URL contains the list of all revoked certificates since the last time a full certificate revocaton list was created.
func (m *CertificateAuthority) GetDeltaCertificateRevocationListUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deltaCertificateRevocationListUrl
    }
}
// Gets the isRootAuthority property value. Required. true if the trusted certificate is a root authority, false if the trusted certificate is an intermediate authority.
func (m *CertificateAuthority) GetIsRootAuthority()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRootAuthority
    }
}
// Gets the issuer property value. The issuer of the certificate, calculated from the certificate value. Read-only.
func (m *CertificateAuthority) GetIssuer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuer
    }
}
// Gets the issuerSki property value. The subject key identifier of the certificate, calculated from the certificate value. Read-only.
func (m *CertificateAuthority) GetIssuerSki()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuerSki
    }
}
// The deserialization information for the current model
func (m *CertificateAuthority) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["certificateRevocationListUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateRevocationListUrl(val)
        }
        return nil
    }
    res["deltaCertificateRevocationListUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeltaCertificateRevocationListUrl(val)
        }
        return nil
    }
    res["isRootAuthority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRootAuthority(val)
        }
        return nil
    }
    res["issuer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["issuerSki"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerSki(val)
        }
        return nil
    }
    return res
}
func (m *CertificateAuthority) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CertificateAuthority) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the certificate property value. Required. The base64 encoded string representing the public certificate.
// Parameters:
//  - value : Value to set for the certificate property.
func (m *CertificateAuthority) SetCertificate(value []byte)() {
    m.certificate = value
}
// Sets the certificateRevocationListUrl property value. The URL of the certificate revocation list.
// Parameters:
//  - value : Value to set for the certificateRevocationListUrl property.
func (m *CertificateAuthority) SetCertificateRevocationListUrl(value *string)() {
    m.certificateRevocationListUrl = value
}
// Sets the deltaCertificateRevocationListUrl property value. The URL contains the list of all revoked certificates since the last time a full certificate revocaton list was created.
// Parameters:
//  - value : Value to set for the deltaCertificateRevocationListUrl property.
func (m *CertificateAuthority) SetDeltaCertificateRevocationListUrl(value *string)() {
    m.deltaCertificateRevocationListUrl = value
}
// Sets the isRootAuthority property value. Required. true if the trusted certificate is a root authority, false if the trusted certificate is an intermediate authority.
// Parameters:
//  - value : Value to set for the isRootAuthority property.
func (m *CertificateAuthority) SetIsRootAuthority(value *bool)() {
    m.isRootAuthority = value
}
// Sets the issuer property value. The issuer of the certificate, calculated from the certificate value. Read-only.
// Parameters:
//  - value : Value to set for the issuer property.
func (m *CertificateAuthority) SetIssuer(value *string)() {
    m.issuer = value
}
// Sets the issuerSki property value. The subject key identifier of the certificate, calculated from the certificate value. Read-only.
// Parameters:
//  - value : Value to set for the issuerSki property.
func (m *CertificateAuthority) SetIssuerSki(value *string)() {
    m.issuerSki = value
}
