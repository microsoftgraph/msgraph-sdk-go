package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CertificateBasedAuthConfiguration provides operations to manage the collection of certificateBasedAuthConfiguration entities.
type CertificateBasedAuthConfiguration struct {
    Entity
    // Collection of certificate authorities which creates a trusted certificate chain.
    certificateAuthorities []CertificateAuthorityable;
}
// NewCertificateBasedAuthConfiguration instantiates a new certificateBasedAuthConfiguration and sets the default values.
func NewCertificateBasedAuthConfiguration()(*CertificateBasedAuthConfiguration) {
    m := &CertificateBasedAuthConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCertificateBasedAuthConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateBasedAuthConfigurationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCertificateBasedAuthConfiguration(), nil
}
// GetCertificateAuthorities gets the certificateAuthorities property value. Collection of certificate authorities which creates a trusted certificate chain.
func (m *CertificateBasedAuthConfiguration) GetCertificateAuthorities()([]CertificateAuthorityable) {
    if m == nil {
        return nil
    } else {
        return m.certificateAuthorities
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateBasedAuthConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificateAuthorities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificateAuthorityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateAuthorityable, len(val))
            for i, v := range val {
                res[i] = v.(CertificateAuthorityable)
            }
            m.SetCertificateAuthorities(res)
        }
        return nil
    }
    return res
}
func (m *CertificateBasedAuthConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CertificateBasedAuthConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateAuthorities() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCertificateAuthorities()))
        for i, v := range m.GetCertificateAuthorities() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("certificateAuthorities", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateAuthorities sets the certificateAuthorities property value. Collection of certificate authorities which creates a trusted certificate chain.
func (m *CertificateBasedAuthConfiguration) SetCertificateAuthorities(value []CertificateAuthorityable)() {
    if m != nil {
        m.certificateAuthorities = value
    }
}
