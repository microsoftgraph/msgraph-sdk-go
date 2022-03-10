package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintUsageByUser provides operations to manage the reportRoot singleton.
type PrintUsageByUser struct {
    PrintUsage
    // The UPN of the user represented by these statistics.
    userPrincipalName *string;
}
// NewPrintUsageByUser instantiates a new printUsageByUser and sets the default values.
func NewPrintUsageByUser()(*PrintUsageByUser) {
    m := &PrintUsageByUser{
        PrintUsage: *NewPrintUsage(),
    }
    return m
}
// CreatePrintUsageByUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintUsageByUserFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrintUsageByUser(), nil
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetUserPrincipalName gets the userPrincipalName property value. The UPN of the user represented by these statistics.
func (m *PrintUsageByUser) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *PrintUsageByUser) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetUserPrincipalName sets the userPrincipalName property value. The UPN of the user represented by these statistics.
func (m *PrintUsageByUser) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
