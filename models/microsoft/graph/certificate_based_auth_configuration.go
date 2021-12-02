package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CertificateBasedAuthConfiguration 
type CertificateBasedAuthConfiguration struct {
    Entity
    // Collection of certificate authorities which creates a trusted certificate chain.
    certificateAuthorities []CertificateAuthority;
}
// NewCertificateBasedAuthConfiguration instantiates a new certificateBasedAuthConfiguration and sets the default values.
func NewCertificateBasedAuthConfiguration()(*CertificateBasedAuthConfiguration) {
    m := &CertificateBasedAuthConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// GetCertificateAuthorities gets the certificateAuthorities property value. Collection of certificate authorities which creates a trusted certificate chain.
func (m *CertificateBasedAuthConfiguration) GetCertificateAuthorities()([]CertificateAuthority) {
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCertificateAuthority() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateAuthority, len(val))
            for i, v := range val {
                res[i] = *(v.(*CertificateAuthority))
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCertificateAuthorities()))
        for i, v := range m.GetCertificateAuthorities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("certificateAuthorities", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateAuthorities sets the certificateAuthorities property value. Collection of certificate authorities which creates a trusted certificate chain.
func (m *CertificateBasedAuthConfiguration) SetCertificateAuthorities(value []CertificateAuthority)() {
    if m != nil {
        m.certificateAuthorities = value
    }
}
