package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/externalconnectors"
)

type Acl struct {
    accessType *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AccessType;
    additionalData map[string]interface{};
    type_escpaped *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AclType;
    value *string;
}
func NewAcl()(*Acl) {
    m := &Acl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Acl) GetAccessType()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AccessType) {
    if m == nil {
        return nil
    } else {
        return m.accessType
    }
}
func (m *Acl) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Acl) GetType_escpaped()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AclType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *Acl) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *Acl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParseAccessType)
        if err != nil {
            return err
        }
        cast := val.(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AccessType)
        m.SetAccessType(&cast)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParseAclType)
        if err != nil {
            return err
        }
        cast := val.(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AclType)
        m.SetType_escpaped(&cast)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *Acl) IsNil()(bool) {
    return m == nil
}
func (m *Acl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessType() != nil {
        cast := m.GetAccessType().String()
        err := writer.WriteStringValue("accessType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetType_escpaped() != nil {
        cast := m.GetType_escpaped().String()
        err := writer.WriteStringValue("type_escpaped", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *Acl) SetAccessType(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AccessType)() {
    m.accessType = value
}
func (m *Acl) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Acl) SetType_escpaped(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AclType)() {
    m.type_escpaped = value
}
func (m *Acl) SetValue(value *string)() {
    m.value = value
}
