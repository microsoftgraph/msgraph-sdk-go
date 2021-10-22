package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CallRoute struct {
    additionalData map[string]interface{};
    final *IdentitySet;
    original *IdentitySet;
    routingType *RoutingType;
}
func NewCallRoute()(*CallRoute) {
    m := &CallRoute{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CallRoute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CallRoute) GetFinal()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.final
    }
}
func (m *CallRoute) GetOriginal()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.original
    }
}
func (m *CallRoute) GetRoutingType()(*RoutingType) {
    if m == nil {
        return nil
    } else {
        return m.routingType
    }
}
func (m *CallRoute) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["final"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetFinal(val.(*IdentitySet))
        return nil
    }
    res["original"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetOriginal(val.(*IdentitySet))
        return nil
    }
    res["routingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoutingType)
        if err != nil {
            return err
        }
        cast := val.(RoutingType)
        m.SetRoutingType(&cast)
        return nil
    }
    return res
}
func (m *CallRoute) IsNil()(bool) {
    return m == nil
}
func (m *CallRoute) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("final", m.GetFinal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("original", m.GetOriginal())
        if err != nil {
            return err
        }
    }
    if m.GetRoutingType() != nil {
        cast := m.GetRoutingType().String()
        err := writer.WriteStringValue("routingType", &cast)
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
func (m *CallRoute) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CallRoute) SetFinal(value *IdentitySet)() {
    m.final = value
}
func (m *CallRoute) SetOriginal(value *IdentitySet)() {
    m.original = value
}
func (m *CallRoute) SetRoutingType(value *RoutingType)() {
    m.routingType = value
}
