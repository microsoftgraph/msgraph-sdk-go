package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CallRoute struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    final *IdentitySet;
    // 
    original *IdentitySet;
    // Possible values are: forwarded, lookup, selfFork.
    routingType *RoutingType;
}
// Instantiates a new callRoute and sets the default values.
func NewCallRoute()(*CallRoute) {
    m := &CallRoute{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallRoute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the final property value. 
func (m *CallRoute) GetFinal()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.final
    }
}
// Gets the original property value. 
func (m *CallRoute) GetOriginal()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.original
    }
}
// Gets the routingType property value. Possible values are: forwarded, lookup, selfFork.
func (m *CallRoute) GetRoutingType()(*RoutingType) {
    if m == nil {
        return nil
    } else {
        return m.routingType
    }
}
// The deserialization information for the current model
func (m *CallRoute) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["final"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinal(val.(*IdentitySet))
        }
        return nil
    }
    res["original"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginal(val.(*IdentitySet))
        }
        return nil
    }
    res["routingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoutingType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RoutingType)
            m.SetRoutingType(&cast)
        }
        return nil
    }
    return res
}
func (m *CallRoute) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CallRoute) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the final property value. 
// Parameters:
//  - value : Value to set for the final property.
func (m *CallRoute) SetFinal(value *IdentitySet)() {
    m.final = value
}
// Sets the original property value. 
// Parameters:
//  - value : Value to set for the original property.
func (m *CallRoute) SetOriginal(value *IdentitySet)() {
    m.original = value
}
// Sets the routingType property value. Possible values are: forwarded, lookup, selfFork.
// Parameters:
//  - value : Value to set for the routingType property.
func (m *CallRoute) SetRoutingType(value *RoutingType)() {
    m.routingType = value
}
