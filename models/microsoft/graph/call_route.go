package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CallRoute 
type CallRoute struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    final IdentitySetable;
    // 
    original IdentitySetable;
    // Possible values are: forwarded, lookup, selfFork.
    routingType *RoutingType;
}
// NewCallRoute instantiates a new callRoute and sets the default values.
func NewCallRoute()(*CallRoute) {
    m := &CallRoute{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCallRouteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallRouteFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCallRoute(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallRoute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallRoute) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["final"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinal(val.(IdentitySetable))
        }
        return nil
    }
    res["original"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginal(val.(IdentitySetable))
        }
        return nil
    }
    res["routingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoutingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoutingType(val.(*RoutingType))
        }
        return nil
    }
    return res
}
// GetFinal gets the final property value. 
func (m *CallRoute) GetFinal()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.final
    }
}
// GetOriginal gets the original property value. 
func (m *CallRoute) GetOriginal()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.original
    }
}
// GetRoutingType gets the routingType property value. Possible values are: forwarded, lookup, selfFork.
func (m *CallRoute) GetRoutingType()(*RoutingType) {
    if m == nil {
        return nil
    } else {
        return m.routingType
    }
}
// Serialize serializes information the current object
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
        cast := (*m.GetRoutingType()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallRoute) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFinal sets the final property value. 
func (m *CallRoute) SetFinal(value IdentitySetable)() {
    if m != nil {
        m.final = value
    }
}
// SetOriginal sets the original property value. 
func (m *CallRoute) SetOriginal(value IdentitySetable)() {
    if m != nil {
        m.original = value
    }
}
// SetRoutingType sets the routingType property value. Possible values are: forwarded, lookup, selfFork.
func (m *CallRoute) SetRoutingType(value *RoutingType)() {
    if m != nil {
        m.routingType = value
    }
}
