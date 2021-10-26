package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrinterShare struct {
    PrinterBase
    // If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties.
    allowAllUsers *bool;
    // The groups whose users have access to print using the printer.
    allowedGroups []Group;
    // The users who have access to print using the printer.
    allowedUsers []User;
    // The DateTimeOffset when the printer share was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The printer that this printer share is related to.
    printer *Printer;
}
// Instantiates a new printerShare and sets the default values.
func NewPrinterShare()(*PrinterShare) {
    m := &PrinterShare{
        PrinterBase: *NewPrinterBase(),
    }
    return m
}
// Gets the allowAllUsers property value. If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties.
func (m *PrinterShare) GetAllowAllUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAllUsers
    }
}
// Gets the allowedGroups property value. The groups whose users have access to print using the printer.
func (m *PrinterShare) GetAllowedGroups()([]Group) {
    if m == nil {
        return nil
    } else {
        return m.allowedGroups
    }
}
// Gets the allowedUsers property value. The users who have access to print using the printer.
func (m *PrinterShare) GetAllowedUsers()([]User) {
    if m == nil {
        return nil
    } else {
        return m.allowedUsers
    }
}
// Gets the createdDateTime property value. The DateTimeOffset when the printer share was created. Read-only.
func (m *PrinterShare) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the printer property value. The printer that this printer share is related to.
func (m *PrinterShare) GetPrinter()(*Printer) {
    if m == nil {
        return nil
    } else {
        return m.printer
    }
}
// The deserialization information for the current model
func (m *PrinterShare) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PrinterBase.GetFieldDeserializers()
    res["allowAllUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowAllUsers(val)
        return nil
    }
    res["allowedGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroup() })
        if err != nil {
            return err
        }
        res := make([]Group, len(val))
        for i, v := range val {
            res[i] = *(v.(*Group))
        }
        m.SetAllowedGroups(res)
        return nil
    }
    res["allowedUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUser() })
        if err != nil {
            return err
        }
        res := make([]User, len(val))
        for i, v := range val {
            res[i] = *(v.(*User))
        }
        m.SetAllowedUsers(res)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["printer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinter() })
        if err != nil {
            return err
        }
        m.SetPrinter(val.(*Printer))
        return nil
    }
    return res
}
func (m *PrinterShare) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrinterShare) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PrinterBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAllUsers", m.GetAllowAllUsers())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedGroups()))
        for i, v := range m.GetAllowedGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("allowedGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedUsers()))
        for i, v := range m.GetAllowedUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("allowedUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("printer", m.GetPrinter())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the allowAllUsers property value. If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties.
// Parameters:
//  - value : Value to set for the allowAllUsers property.
func (m *PrinterShare) SetAllowAllUsers(value *bool)() {
    m.allowAllUsers = value
}
// Sets the allowedGroups property value. The groups whose users have access to print using the printer.
// Parameters:
//  - value : Value to set for the allowedGroups property.
func (m *PrinterShare) SetAllowedGroups(value []Group)() {
    m.allowedGroups = value
}
// Sets the allowedUsers property value. The users who have access to print using the printer.
// Parameters:
//  - value : Value to set for the allowedUsers property.
func (m *PrinterShare) SetAllowedUsers(value []User)() {
    m.allowedUsers = value
}
// Sets the createdDateTime property value. The DateTimeOffset when the printer share was created. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *PrinterShare) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the printer property value. The printer that this printer share is related to.
// Parameters:
//  - value : Value to set for the printer property.
func (m *PrinterShare) SetPrinter(value *Printer)() {
    m.printer = value
}
