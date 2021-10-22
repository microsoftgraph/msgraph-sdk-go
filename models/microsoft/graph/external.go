package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type External struct {
    additionalData map[string]interface{};
    connections []ExternalConnection;
}
func NewExternal()(*External) {
    m := &External{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *External) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *External) GetConnections()([]ExternalConnection) {
    if m == nil {
        return nil
    } else {
        return m.connections
    }
}
func (m *External) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["connections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExternalConnection() })
        if err != nil {
            return err
        }
        res := make([]ExternalConnection, len(val))
        for i, v := range val {
            res[i] = *(v.(*ExternalConnection))
        }
        m.SetConnections(res)
        return nil
    }
    return res
}
func (m *External) IsNil()(bool) {
    return m == nil
}
func (m *External) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConnections()))
        for i, v := range m.GetConnections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("connections", cast)
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
func (m *External) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *External) SetConnections(value []ExternalConnection)() {
    m.connections = value
}
