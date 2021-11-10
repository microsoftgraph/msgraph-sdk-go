package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/externalconnectors"
)

// 
type ExternalConnection struct {
    Entity
    // Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
    configuration *Configuration;
    // Description of the connection displayed in the Microsoft 365 admin center. Optional.
    description *string;
    // Read-only. Nullable.
    groups []ExternalGroup;
    // Read-only. Nullable.
    items []ExternalItem;
    // The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
    name *string;
    // Read-only. Nullable.
    operations []ConnectionOperation;
    // Read-only. Nullable.
    schema *Schema;
    // Indicates the current state of the connection. Possible values are: draft, ready, obsolete, limitExceeded, unknownFutureValue.
    state *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ConnectionState;
}
// Instantiates a new externalConnection and sets the default values.
func NewExternalConnection()(*ExternalConnection) {
    m := &ExternalConnection{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the configuration property value. Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
func (m *ExternalConnection) GetConfiguration()(*Configuration) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// Gets the description property value. Description of the connection displayed in the Microsoft 365 admin center. Optional.
func (m *ExternalConnection) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the groups property value. Read-only. Nullable.
func (m *ExternalConnection) GetGroups()([]ExternalGroup) {
    if m == nil {
        return nil
    } else {
        return m.groups
    }
}
// Gets the items property value. Read-only. Nullable.
func (m *ExternalConnection) GetItems()([]ExternalItem) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// Gets the name property value. The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
func (m *ExternalConnection) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the operations property value. Read-only. Nullable.
func (m *ExternalConnection) GetOperations()([]ConnectionOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// Gets the schema property value. Read-only. Nullable.
func (m *ExternalConnection) GetSchema()(*Schema) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
// Gets the state property value. Indicates the current state of the connection. Possible values are: draft, ready, obsolete, limitExceeded, unknownFutureValue.
func (m *ExternalConnection) GetState()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ConnectionState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *ExternalConnection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(*Configuration))
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["groups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExternalGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*ExternalGroup))
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExternalItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*ExternalItem))
            }
            m.SetItems(res)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnectionOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectionOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConnectionOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["schema"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSchema() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val.(*Schema))
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParseConnectionState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ConnectionState)
            m.SetState(&cast)
        }
        return nil
    }
    return res
}
func (m *ExternalConnection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ExternalConnection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schema", m.GetSchema())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the configuration property value. Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
// Parameters:
//  - value : Value to set for the configuration property.
func (m *ExternalConnection) SetConfiguration(value *Configuration)() {
    m.configuration = value
}
// Sets the description property value. Description of the connection displayed in the Microsoft 365 admin center. Optional.
// Parameters:
//  - value : Value to set for the description property.
func (m *ExternalConnection) SetDescription(value *string)() {
    m.description = value
}
// Sets the groups property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the groups property.
func (m *ExternalConnection) SetGroups(value []ExternalGroup)() {
    m.groups = value
}
// Sets the items property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the items property.
func (m *ExternalConnection) SetItems(value []ExternalItem)() {
    m.items = value
}
// Sets the name property value. The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
// Parameters:
//  - value : Value to set for the name property.
func (m *ExternalConnection) SetName(value *string)() {
    m.name = value
}
// Sets the operations property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the operations property.
func (m *ExternalConnection) SetOperations(value []ConnectionOperation)() {
    m.operations = value
}
// Sets the schema property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the schema property.
func (m *ExternalConnection) SetSchema(value *Schema)() {
    m.schema = value
}
// Sets the state property value. Indicates the current state of the connection. Possible values are: draft, ready, obsolete, limitExceeded, unknownFutureValue.
// Parameters:
//  - value : Value to set for the state property.
func (m *ExternalConnection) SetState(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ConnectionState)() {
    m.state = value
}
