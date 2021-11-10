package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrintUsageByUser struct {
    PrintUsage
    // The UPN of the user represented by these statistics.
    userPrincipalName *string;
}
// Instantiates a new printUsageByUser and sets the default values.
func NewPrintUsageByUser()(*PrintUsageByUser) {
    m := &PrintUsageByUser{
        PrintUsage: *NewPrintUsage(),
    }
    return m
}
// Gets the userPrincipalName property value. The UPN of the user represented by these statistics.
func (m *PrintUsageByUser) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *PrintUsageByUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PrintUsage.GetFieldDeserializers()
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *PrintUsageByUser) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrintUsageByUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PrintUsage.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the userPrincipalName property value. The UPN of the user represented by these statistics.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *PrintUsageByUser) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
