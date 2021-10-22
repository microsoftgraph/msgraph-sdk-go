package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ObjectIdentity struct {
    additionalData map[string]interface{};
    issuer *string;
    issuerAssignedId *string;
    signInType *string;
}
func NewObjectIdentity()(*ObjectIdentity) {
    m := &ObjectIdentity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ObjectIdentity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ObjectIdentity) GetIssuer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuer
    }
}
func (m *ObjectIdentity) GetIssuerAssignedId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuerAssignedId
    }
}
func (m *ObjectIdentity) GetSignInType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInType
    }
}
func (m *ObjectIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["issuer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIssuer(val)
        return nil
    }
    res["issuerAssignedId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIssuerAssignedId(val)
        return nil
    }
    res["signInType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSignInType(val)
        return nil
    }
    return res
}
func (m *ObjectIdentity) IsNil()(bool) {
    return m == nil
}
func (m *ObjectIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issuerAssignedId", m.GetIssuerAssignedId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signInType", m.GetSignInType())
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
func (m *ObjectIdentity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ObjectIdentity) SetIssuer(value *string)() {
    m.issuer = value
}
func (m *ObjectIdentity) SetIssuerAssignedId(value *string)() {
    m.issuerAssignedId = value
}
func (m *ObjectIdentity) SetSignInType(value *string)() {
    m.signInType = value
}
