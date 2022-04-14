package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateAuthority 
type CertificateAuthority struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Required. The base64 encoded string representing the public certificate.
    certificate []byte
    // The URL of the certificate revocation list.
    certificateRevocationListUrl *string
    // The URL contains the list of all revoked certificates since the last time a full certificate revocaton list was created.
    deltaCertificateRevocationListUrl *string
    // Required. true if the trusted certificate is a root authority, false if the trusted certificate is an intermediate authority.
    isRootAuthority *bool
    // The issuer of the certificate, calculated from the certificate value. Read-only.
    issuer *string
    // The subject key identifier of the certificate, calculated from the certificate value. Read-only.
    issuerSki *string
}
// NewCertificateAuthority instantiates a new certificateAuthority and sets the default values.
func NewCertificateAuthority()(*CertificateAuthority) {
    m := &CertificateAuthority{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCertificateAuthorityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateAuthorityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateAuthority(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateAuthority) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificate gets the certificate property value. Required. The base64 encoded string representing the public certificate.
func (m *CertificateAuthority) GetCertificate()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.certificate
    }
}
// GetCertificateRevocationListUrl gets the certificateRevocationListUrl property value. The URL of the certificate revocation list.
func (m *CertificateAuthority) GetCertificateRevocationListUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateRevocationListUrl
    }
}
// GetDeltaCertificateRevocationListUrl gets the deltaCertificateRevocationListUrl property value. The URL contains the list of all revoked certificates since the last time a full certificate revocaton list was created.
func (m *CertificateAuthority) GetDeltaCertificateRevocationListUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deltaCertificateRevocationListUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateAuthority) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["certificateRevocationListUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateRevocationListUrl(val)
        }
        return nil
    }
    res["deltaCertificateRevocationListUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeltaCertificateRevocationListUrl(val)
        }
        return nil
    }
    res["isRootAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRootAuthority(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["issuerSki"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetIsRootAuthority gets the isRootAuthority property value. Required. true if the trusted certificate is a root authority, false if the trusted certificate is an intermediate authority.
func (m *CertificateAuthority) GetIsRootAuthority()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRootAuthority
    }
}
// GetIssuer gets the issuer property value. The issuer of the certificate, calculated from the certificate value. Read-only.
func (m *CertificateAuthority) GetIssuer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuer
    }
}
// GetIssuerSki gets the issuerSki property value. The subject key identifier of the certificate, calculated from the certificate value. Read-only.
func (m *CertificateAuthority) GetIssuerSki()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuerSki
    }
}
// Serialize serializes information the current object
func (m *CertificateAuthority) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateAuthority) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertificate sets the certificate property value. Required. The base64 encoded string representing the public certificate.
func (m *CertificateAuthority) SetCertificate(value []byte)() {
    if m != nil {
        m.certificate = value
    }
}
// SetCertificateRevocationListUrl sets the certificateRevocationListUrl property value. The URL of the certificate revocation list.
func (m *CertificateAuthority) SetCertificateRevocationListUrl(value *string)() {
    if m != nil {
        m.certificateRevocationListUrl = value
    }
}
// SetDeltaCertificateRevocationListUrl sets the deltaCertificateRevocationListUrl property value. The URL contains the list of all revoked certificates since the last time a full certificate revocaton list was created.
func (m *CertificateAuthority) SetDeltaCertificateRevocationListUrl(value *string)() {
    if m != nil {
        m.deltaCertificateRevocationListUrl = value
    }
}
// SetIsRootAuthority sets the isRootAuthority property value. Required. true if the trusted certificate is a root authority, false if the trusted certificate is an intermediate authority.
func (m *CertificateAuthority) SetIsRootAuthority(value *bool)() {
    if m != nil {
        m.isRootAuthority = value
    }
}
// SetIssuer sets the issuer property value. The issuer of the certificate, calculated from the certificate value. Read-only.
func (m *CertificateAuthority) SetIssuer(value *string)() {
    if m != nil {
        m.issuer = value
    }
}
// SetIssuerSki sets the issuerSki property value. The subject key identifier of the certificate, calculated from the certificate value. Read-only.
func (m *CertificateAuthority) SetIssuerSki(value *string)() {
    if m != nil {
        m.issuerSki = value
    }
}
