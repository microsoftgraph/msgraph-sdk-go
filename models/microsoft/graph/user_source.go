package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/ediscovery"
)

// 
type UserSource struct {
    DataSource
    // Email address of the user's mailbox.
    email *string;
    // Specifies which sources are included in this group. Possible values are: mailbox, site.
    includedSources *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.SourceType;
}
// Instantiates a new userSource and sets the default values.
func NewUserSource()(*UserSource) {
    m := &UserSource{
        DataSource: *NewDataSource(),
    }
    return m
}
// Gets the email property value. Email address of the user's mailbox.
func (m *UserSource) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the includedSources property value. Specifies which sources are included in this group. Possible values are: mailbox, site.
func (m *UserSource) GetIncludedSources()(*i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.SourceType) {
    if m == nil {
        return nil
    } else {
        return m.includedSources
    }
}
// The deserialization information for the current model
func (m *UserSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DataSource.GetFieldDeserializers()
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    res["includedSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ParseSourceType)
        if err != nil {
            return err
        }
        cast := val.(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.SourceType)
        m.SetIncludedSources(&cast)
        return nil
    }
    return res
}
func (m *UserSource) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DataSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    if m.GetIncludedSources() != nil {
        cast := m.GetIncludedSources().String()
        err = writer.WriteStringValue("includedSources", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the email property value. Email address of the user's mailbox.
// Parameters:
//  - value : Value to set for the email property.
func (m *UserSource) SetEmail(value *string)() {
    m.email = value
}
// Sets the includedSources property value. Specifies which sources are included in this group. Possible values are: mailbox, site.
// Parameters:
//  - value : Value to set for the includedSources property.
func (m *UserSource) SetIncludedSources(value *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.SourceType)() {
    m.includedSources = value
}
