package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Fido2AuthenticationMethod struct {
    AuthenticationMethod
    aaGuid *string;
    attestationCertificates []string;
    attestationLevel *AttestationLevel;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    displayName *string;
    model *string;
}
func NewFido2AuthenticationMethod()(*Fido2AuthenticationMethod) {
    m := &Fido2AuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
func (m *Fido2AuthenticationMethod) GetAaGuid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aaGuid
    }
}
func (m *Fido2AuthenticationMethod) GetAttestationCertificates()([]string) {
    if m == nil {
        return nil
    } else {
        return m.attestationCertificates
    }
}
func (m *Fido2AuthenticationMethod) GetAttestationLevel()(*AttestationLevel) {
    if m == nil {
        return nil
    } else {
        return m.attestationLevel
    }
}
func (m *Fido2AuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *Fido2AuthenticationMethod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Fido2AuthenticationMethod) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *Fido2AuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["aaGuid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAaGuid(val)
        return nil
    }
    res["attestationCertificates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAttestationCertificates(res)
        return nil
    }
    res["attestationLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttestationLevel)
        if err != nil {
            return err
        }
        cast := val.(AttestationLevel)
        m.SetAttestationLevel(&cast)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    return res
}
func (m *Fido2AuthenticationMethod) IsNil()(bool) {
    return m == nil
}
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
func (m *Fido2AuthenticationMethod) SetAaGuid(value *string)() {
    m.aaGuid = value
}
func (m *Fido2AuthenticationMethod) SetAttestationCertificates(value []string)() {
    m.attestationCertificates = value
}
func (m *Fido2AuthenticationMethod) SetAttestationLevel(value *AttestationLevel)() {
    m.attestationLevel = value
}
func (m *Fido2AuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *Fido2AuthenticationMethod) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Fido2AuthenticationMethod) SetModel(value *string)() {
    m.model = value
}
