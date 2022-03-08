package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Fido2AuthenticationMethod provides operations to manage the collection of drive entities.
type Fido2AuthenticationMethod struct {
    AuthenticationMethod
    // Authenticator Attestation GUID, an identifier that indicates the type (e.g. make and model) of the authenticator.
    aaGuid *string;
    // The attestation certificate(s) attached to this security key.
    attestationCertificates []string;
    // The attestation level of this FIDO2 security key. Possible values are: attested, or notAttested.
    attestationLevel *AttestationLevel;
    // The timestamp when this key was registered to the user.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The display name of the key as given by the user.
    displayName *string;
    // The manufacturer-assigned model of the FIDO2 security key.
    model *string;
}
// NewFido2AuthenticationMethod instantiates a new fido2AuthenticationMethod and sets the default values.
func NewFido2AuthenticationMethod()(*Fido2AuthenticationMethod) {
    m := &Fido2AuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// CreateFido2AuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFido2AuthenticationMethodFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFido2AuthenticationMethod(), nil
}
// GetAaGuid gets the aaGuid property value. Authenticator Attestation GUID, an identifier that indicates the type (e.g. make and model) of the authenticator.
func (m *Fido2AuthenticationMethod) GetAaGuid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aaGuid
    }
}
// GetAttestationCertificates gets the attestationCertificates property value. The attestation certificate(s) attached to this security key.
func (m *Fido2AuthenticationMethod) GetAttestationCertificates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.attestationCertificates
    }
}
// GetAttestationLevel gets the attestationLevel property value. The attestation level of this FIDO2 security key. Possible values are: attested, or notAttested.
func (m *Fido2AuthenticationMethod) GetAttestationLevel()(*AttestationLevel) {
    if m == nil {
        return nil
    } else {
        return m.attestationLevel
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The timestamp when this key was registered to the user.
func (m *Fido2AuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. The display name of the key as given by the user.
func (m *Fido2AuthenticationMethod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Fido2AuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["aaGuid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAaGuid(val)
        }
        return nil
    }
    res["attestationCertificates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAttestationCertificates(res)
        }
        return nil
    }
    res["attestationLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttestationLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttestationLevel(val.(*AttestationLevel))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    return res
}
// GetModel gets the model property value. The manufacturer-assigned model of the FIDO2 security key.
func (m *Fido2AuthenticationMethod) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *Fido2AuthenticationMethod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Fido2AuthenticationMethod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("aaGuid", m.GetAaGuid())
        if err != nil {
            return err
        }
    }
    if m.GetAttestationCertificates() != nil {
        err = writer.WriteCollectionOfStringValues("attestationCertificates", m.GetAttestationCertificates())
        if err != nil {
            return err
        }
    }
    if m.GetAttestationLevel() != nil {
        cast := (*m.GetAttestationLevel()).String()
        err = writer.WriteStringValue("attestationLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAaGuid sets the aaGuid property value. Authenticator Attestation GUID, an identifier that indicates the type (e.g. make and model) of the authenticator.
func (m *Fido2AuthenticationMethod) SetAaGuid(value *string)() {
    if m != nil {
        m.aaGuid = value
    }
}
// SetAttestationCertificates sets the attestationCertificates property value. The attestation certificate(s) attached to this security key.
func (m *Fido2AuthenticationMethod) SetAttestationCertificates(value []string)() {
    if m != nil {
        m.attestationCertificates = value
    }
}
// SetAttestationLevel sets the attestationLevel property value. The attestation level of this FIDO2 security key. Possible values are: attested, or notAttested.
func (m *Fido2AuthenticationMethod) SetAttestationLevel(value *AttestationLevel)() {
    if m != nil {
        m.attestationLevel = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The timestamp when this key was registered to the user.
func (m *Fido2AuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the key as given by the user.
func (m *Fido2AuthenticationMethod) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetModel sets the model property value. The manufacturer-assigned model of the FIDO2 security key.
func (m *Fido2AuthenticationMethod) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
