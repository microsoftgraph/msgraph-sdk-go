package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OptionalClaims 
type OptionalClaims struct {
    // The optional claims returned in the JWT access token.
    accessToken []OptionalClaim;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The optional claims returned in the JWT ID token.
    idToken []OptionalClaim;
    // The optional claims returned in the SAML token.
    saml2Token []OptionalClaim;
}
// NewOptionalClaims instantiates a new optionalClaims and sets the default values.
func NewOptionalClaims()(*OptionalClaims) {
    m := &OptionalClaims{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAccessToken gets the accessToken property value. The optional claims returned in the JWT access token.
func (m *OptionalClaims) GetAccessToken()([]OptionalClaim) {
    if m == nil {
        return nil
    } else {
        return m.accessToken
    }
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaims) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIdToken gets the idToken property value. The optional claims returned in the JWT ID token.
func (m *OptionalClaims) GetIdToken()([]OptionalClaim) {
    if m == nil {
        return nil
    } else {
        return m.idToken
    }
}
// GetSaml2Token gets the saml2Token property value. The optional claims returned in the SAML token.
func (m *OptionalClaims) GetSaml2Token()([]OptionalClaim) {
    if m == nil {
        return nil
    } else {
        return m.saml2Token
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OptionalClaims) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOptionalClaim() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OptionalClaim, len(val))
            for i, v := range val {
                res[i] = *(v.(*OptionalClaim))
            }
            m.SetAccessToken(res)
        }
        return nil
    }
    res["idToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOptionalClaim() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OptionalClaim, len(val))
            for i, v := range val {
                res[i] = *(v.(*OptionalClaim))
            }
            m.SetIdToken(res)
        }
        return nil
    }
    res["saml2Token"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOptionalClaim() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OptionalClaim, len(val))
            for i, v := range val {
                res[i] = *(v.(*OptionalClaim))
            }
            m.SetSaml2Token(res)
        }
        return nil
    }
    return res
}
func (m *OptionalClaims) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OptionalClaims) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessToken()))
        for i, v := range m.GetAccessToken() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("accessToken", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIdToken()))
        for i, v := range m.GetIdToken() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("idToken", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSaml2Token()))
        for i, v := range m.GetSaml2Token() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("saml2Token", cast)
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
// SetAccessToken sets the accessToken property value. The optional claims returned in the JWT access token.
func (m *OptionalClaims) SetAccessToken(value []OptionalClaim)() {
    if m != nil {
        m.accessToken = value
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaims) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIdToken sets the idToken property value. The optional claims returned in the JWT ID token.
func (m *OptionalClaims) SetIdToken(value []OptionalClaim)() {
    if m != nil {
        m.idToken = value
    }
}
// SetSaml2Token sets the saml2Token property value. The optional claims returned in the SAML token.
func (m *OptionalClaims) SetSaml2Token(value []OptionalClaim)() {
    if m != nil {
        m.saml2Token = value
    }
}
