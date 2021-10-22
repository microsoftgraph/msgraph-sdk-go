package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/externalconnectors"
)

type Property struct {
    additionalData map[string]interface{};
    aliases []string;
    isQueryable *bool;
    isRefinable *bool;
    isRetrievable *bool;
    isSearchable *bool;
    labels []i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label;
    name *string;
    type_escpaped *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.PropertyType;
}
func NewProperty()(*Property) {
    m := &Property{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Property) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Property) GetAliases()([]string) {
    if m == nil {
        return nil
    } else {
        return m.aliases
    }
}
func (m *Property) GetIsQueryable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isQueryable
    }
}
func (m *Property) GetIsRefinable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRefinable
    }
}
func (m *Property) GetIsRetrievable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRetrievable
    }
}
func (m *Property) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
func (m *Property) GetLabels()([]i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label) {
    if m == nil {
        return nil
    } else {
        return m.labels
    }
}
func (m *Property) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *Property) GetType_escpaped()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.PropertyType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *Property) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["aliases"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAliases(res)
        return nil
    }
    res["isQueryable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsQueryable(val)
        return nil
    }
    res["isRefinable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRefinable(val)
        return nil
    }
    res["isRetrievable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRetrievable(val)
        return nil
    }
    res["isSearchable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSearchable(val)
        return nil
    }
    res["labels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParseLabel)
        if err != nil {
            return err
        }
        res := make([]i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label, len(val))
        for i, v := range val {
            res[i] = *(v.(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label))
        }
        m.SetLabels(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParsePropertyType)
        if err != nil {
            return err
        }
        cast := val.(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.PropertyType)
        m.SetType_escpaped(&cast)
        return nil
    }
    return res
}
func (m *Property) IsNil()(bool) {
    return m == nil
}
func (m *Property) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("aliases", m.GetAliases())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isQueryable", m.GetIsQueryable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRefinable", m.GetIsRefinable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRetrievable", m.GetIsRetrievable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("labels", i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.SerializeLabel(m.GetLabels()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Property) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Property) SetAliases(value []string)() {
    m.aliases = value
}
func (m *Property) SetIsQueryable(value *bool)() {
    m.isQueryable = value
}
func (m *Property) SetIsRefinable(value *bool)() {
    m.isRefinable = value
}
func (m *Property) SetIsRetrievable(value *bool)() {
    m.isRetrievable = value
}
func (m *Property) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
func (m *Property) SetLabels(value []i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label)() {
    m.labels = value
}
func (m *Property) SetName(value *string)() {
    m.name = value
}
func (m *Property) SetType_escpaped(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.PropertyType)() {
    m.type_escpaped = value
}
