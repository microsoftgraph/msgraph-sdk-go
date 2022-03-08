package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterShare provides operations to manage the print singleton.
type PrinterShare struct {
    PrinterBase
    // If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties.
    allowAllUsers *bool;
    // The groups whose users have access to print using the printer.
    allowedGroups []Groupable;
    // The users who have access to print using the printer.
    allowedUsers []Userable;
    // The DateTimeOffset when the printer share was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The printer that this printer share is related to.
    printer Printerable;
}
// NewPrinterShare instantiates a new printerShare and sets the default values.
func NewPrinterShare()(*PrinterShare) {
    m := &PrinterShare{
        PrinterBase: *NewPrinterBase(),
    }
    return m
}
// CreatePrinterShareFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterShareFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrinterShare(), nil
}
// GetAllowAllUsers gets the allowAllUsers property value. If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties.
func (m *PrinterShare) GetAllowAllUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAllUsers
    }
}
// GetAllowedGroups gets the allowedGroups property value. The groups whose users have access to print using the printer.
func (m *PrinterShare) GetAllowedGroups()([]Groupable) {
    if m == nil {
        return nil
    } else {
        return m.allowedGroups
    }
}
// GetAllowedUsers gets the allowedUsers property value. The users who have access to print using the printer.
func (m *PrinterShare) GetAllowedUsers()([]Userable) {
    if m == nil {
        return nil
    } else {
        return m.allowedUsers
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The DateTimeOffset when the printer share was created. Read-only.
func (m *PrinterShare) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterShare) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PrinterBase.GetFieldDeserializers()
    res["allowAllUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAllUsers(val)
        }
        return nil
    }
    res["allowedGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Groupable, len(val))
            for i, v := range val {
                res[i] = v.(Groupable)
            }
            m.SetAllowedGroups(res)
        }
        return nil
    }
    res["allowedUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                res[i] = v.(Userable)
            }
            m.SetAllowedUsers(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["printer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinter(val.(Printerable))
        }
        return nil
    }
    return res
}
// GetPrinter gets the printer property value. The printer that this printer share is related to.
func (m *PrinterShare) GetPrinter()(Printerable) {
    if m == nil {
        return nil
    } else {
        return m.printer
    }
}
func (m *PrinterShare) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetAllowedGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedGroups()))
        for i, v := range m.GetAllowedGroups() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("allowedGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedUsers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedUsers()))
        for i, v := range m.GetAllowedUsers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAllowAllUsers sets the allowAllUsers property value. If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties.
func (m *PrinterShare) SetAllowAllUsers(value *bool)() {
    if m != nil {
        m.allowAllUsers = value
    }
}
// SetAllowedGroups sets the allowedGroups property value. The groups whose users have access to print using the printer.
func (m *PrinterShare) SetAllowedGroups(value []Groupable)() {
    if m != nil {
        m.allowedGroups = value
    }
}
// SetAllowedUsers sets the allowedUsers property value. The users who have access to print using the printer.
func (m *PrinterShare) SetAllowedUsers(value []Userable)() {
    if m != nil {
        m.allowedUsers = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The DateTimeOffset when the printer share was created. Read-only.
func (m *PrinterShare) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetPrinter sets the printer property value. The printer that this printer share is related to.
func (m *PrinterShare) SetPrinter(value Printerable)() {
    if m != nil {
        m.printer = value
    }
}
