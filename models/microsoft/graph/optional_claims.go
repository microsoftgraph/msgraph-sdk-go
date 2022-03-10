package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OptionalClaims provides operations to manage the collection of application entities.
type OptionalClaims struct {
    // The optional claims returned in the JWT access token.
    accessToken []OptionalClaimable;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The optional claims returned in the JWT ID token.
    idToken []OptionalClaimable;
    // The optional claims returned in the SAML token.
    saml2Token []OptionalClaimable;
}
// NewOptionalClaims instantiates a new optionalClaims and sets the default values.
func NewOptionalClaims()(*OptionalClaims) {
    m := &OptionalClaims{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOptionalClaimsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOptionalClaimsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOptionalClaims(), nil
}
// GetAccessToken gets the accessToken property value. The optional claims returned in the JWT access token.
func (m *OptionalClaims) GetAccessToken()([]OptionalClaimable) {
    if m == nil {
        return nil
    } else {
        return m.accessToken
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaims) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OptionalClaims) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOptionalClaimFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OptionalClaimable, len(val))
            for i, v := range val {
                res[i] = v.(OptionalClaimable)
            }
            m.SetAccessToken(res)
        }
        return nil
    }
    res["idToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOptionalClaimFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OptionalClaimable, len(val))
            for i, v := range val {
                res[i] = v.(OptionalClaimable)
            }
            m.SetIdToken(res)
        }
        return nil
    }
    res["saml2Token"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOptionalClaimFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OptionalClaimable, len(val))
            for i, v := range val {
                res[i] = v.(OptionalClaimable)
            }
            m.SetSaml2Token(res)
        }
        return nil
    }
    return res
}
// GetIdToken gets the idToken property value. The optional claims returned in the JWT ID token.
func (m *OptionalClaims) GetIdToken()([]OptionalClaimable) {
    if m == nil {
        return nil
    } else {
        return m.idToken
    }
}
// GetSaml2Token gets the saml2Token property value. The optional claims returned in the SAML token.
func (m *OptionalClaims) GetSaml2Token()([]OptionalClaimable) {
    if m == nil {
        return nil
    } else {
        return m.saml2Token
    }
}
func (m *OptionalClaims) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OptionalClaims) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessToken() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessToken()))
        for i, v := range m.GetAccessToken() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("accessToken", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIdToken() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIdToken()))
        for i, v := range m.GetIdToken() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("idToken", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSaml2Token() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSaml2Token()))
        for i, v := range m.GetSaml2Token() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *OptionalClaims) SetAccessToken(value []OptionalClaimable)() {
    if m != nil {
        m.accessToken = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaims) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIdToken sets the idToken property value. The optional claims returned in the JWT ID token.
func (m *OptionalClaims) SetIdToken(value []OptionalClaimable)() {
    if m != nil {
        m.idToken = value
    }
}
// SetSaml2Token sets the saml2Token property value. The optional claims returned in the SAML token.
func (m *OptionalClaims) SetSaml2Token(value []OptionalClaimable)() {
    if m != nil {
        m.saml2Token = value
    }
}
