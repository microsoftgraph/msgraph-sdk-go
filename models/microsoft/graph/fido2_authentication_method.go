package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new fido2AuthenticationMethod and sets the default values.
func NewFido2AuthenticationMethod()(*Fido2AuthenticationMethod) {
    m := &Fido2AuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// Gets the aaGuid property value. Authenticator Attestation GUID, an identifier that indicates the type (e.g. make and model) of the authenticator.
func (m *Fido2AuthenticationMethod) GetAaGuid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aaGuid
    }
}
// Gets the attestationCertificates property value. The attestation certificate(s) attached to this security key.
func (m *Fido2AuthenticationMethod) GetAttestationCertificates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.attestationCertificates
    }
}
// Gets the attestationLevel property value. The attestation level of this FIDO2 security key. Possible values are: attested, or notAttested.
func (m *Fido2AuthenticationMethod) GetAttestationLevel()(*AttestationLevel) {
    if m == nil {
        return nil
    } else {
        return m.attestationLevel
    }
}
// Gets the createdDateTime property value. The timestamp when this key was registered to the user.
func (m *Fido2AuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the displayName property value. The display name of the key as given by the user.
func (m *Fido2AuthenticationMethod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the model property value. The manufacturer-assigned model of the FIDO2 security key.
func (m *Fido2AuthenticationMethod) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// The deserialization information for the current model
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
            cast := val.(AttestationLevel)
            m.SetAttestationLevel(&cast)
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
func (m *Fido2AuthenticationMethod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
        err = writer.WriteCollectionOfStringValues("attestationCertificates", m.GetAttestationCertificates())
        if err != nil {
            return err
        }
    }
    if m.GetAttestationLevel() != nil {
        cast := m.GetAttestationLevel().String()
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
// Sets the aaGuid property value. Authenticator Attestation GUID, an identifier that indicates the type (e.g. make and model) of the authenticator.
// Parameters:
//  - value : Value to set for the aaGuid property.
func (m *Fido2AuthenticationMethod) SetAaGuid(value *string)() {
    m.aaGuid = value
}
// Sets the attestationCertificates property value. The attestation certificate(s) attached to this security key.
// Parameters:
//  - value : Value to set for the attestationCertificates property.
func (m *Fido2AuthenticationMethod) SetAttestationCertificates(value []string)() {
    m.attestationCertificates = value
}
// Sets the attestationLevel property value. The attestation level of this FIDO2 security key. Possible values are: attested, or notAttested.
// Parameters:
//  - value : Value to set for the attestationLevel property.
func (m *Fido2AuthenticationMethod) SetAttestationLevel(value *AttestationLevel)() {
    m.attestationLevel = value
}
// Sets the createdDateTime property value. The timestamp when this key was registered to the user.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Fido2AuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the displayName property value. The display name of the key as given by the user.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Fido2AuthenticationMethod) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the model property value. The manufacturer-assigned model of the FIDO2 security key.
// Parameters:
//  - value : Value to set for the model property.
func (m *Fido2AuthenticationMethod) SetModel(value *string)() {
    m.model = value
}
