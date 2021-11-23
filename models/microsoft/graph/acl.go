package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/externalconnectors"
)

// Acl 
type Acl struct {
    // The access granted to the identity. Possible values are: grant, deny, unknownFutureValue.
    accessType *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AccessType;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The type of identity. Possible values are: user, group, everyone, everyoneExceptGuests, externalGroup, unknownFutureValue.
    type_escaped *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AclType;
    // The unique identifer of the identity. In case of Azure Active Directory identities, value is set to the object identifier of the user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. In case of external groups value is set to the ID of the externalGroup
    value *string;
}
// NewAcl instantiates a new acl and sets the default values.
func NewAcl()(*Acl) {
    m := &Acl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAccessType gets the accessType property value. The access granted to the identity. Possible values are: grant, deny, unknownFutureValue.
func (m *Acl) GetAccessType()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AccessType) {
    if m == nil {
        return nil
    } else {
        return m.accessType
    }
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Acl) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetType_escaped gets the type_escaped property value. The type of identity. Possible values are: user, group, everyone, everyoneExceptGuests, externalGroup, unknownFutureValue.
func (m *Acl) GetType_escaped()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AclType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetValue gets the value property value. The unique identifer of the identity. In case of Azure Active Directory identities, value is set to the object identifier of the user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. In case of external groups value is set to the ID of the externalGroup
func (m *Acl) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Acl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParseAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AccessType)
            m.SetAccessType(&cast)
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParseAclType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AclType)
            m.SetType_escaped(&cast)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
func (m *Acl) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Acl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessType() != nil {
        cast := m.GetAccessType().String()
        err := writer.WriteStringValue("accessType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
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
// SetAccessType sets the accessType property value. The access granted to the identity. Possible values are: grant, deny, unknownFutureValue.
func (m *Acl) SetAccessType(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AccessType)() {
    m.accessType = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Acl) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetType_escaped sets the type_escaped property value. The type of identity. Possible values are: user, group, everyone, everyoneExceptGuests, externalGroup, unknownFutureValue.
func (m *Acl) SetType_escaped(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.AclType)() {
    m.type_escaped = value
}
// SetValue sets the value property value. The unique identifer of the identity. In case of Azure Active Directory identities, value is set to the object identifier of the user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. In case of external groups value is set to the ID of the externalGroup
func (m *Acl) SetValue(value *string)() {
    m.value = value
}
