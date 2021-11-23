package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// printer 
type Printer struct {
    PrinterBase
    // The connectors that are associated with the printer.
    connectors []PrintConnector;
    // True if the printer has a physical device for printing. Read-only.
    hasPhysicalDevice *bool;
    // True if the printer is shared; false otherwise. Read-only.
    isShared *bool;
    // The most recent dateTimeOffset when a printer interacted with Universal Print. Read-only.
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The DateTimeOffset when the printer was registered. Read-only.
    registeredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The list of printerShares that are associated with the printer. Currently, only one printerShare can be associated with the printer. Read-only. Nullable.
    shares []PrinterShare;
    // A list of task triggers that are associated with the printer.
    taskTriggers []PrintTaskTrigger;
}
// NewPrinter instantiates a new printer and sets the default values.
func NewPrinter()(*Printer) {
    m := &Printer{
        PrinterBase: *NewPrinterBase(),
    }
    return m
}
// GetConnectors gets the connectors property value. The connectors that are associated with the printer.
func (m *Printer) GetConnectors()([]PrintConnector) {
    if m == nil {
        return nil
    } else {
        return m.connectors
    }
}
// GetHasPhysicalDevice gets the hasPhysicalDevice property value. True if the printer has a physical device for printing. Read-only.
func (m *Printer) GetHasPhysicalDevice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasPhysicalDevice
    }
}
// GetIsShared gets the isShared property value. True if the printer is shared; false otherwise. Read-only.
func (m *Printer) GetIsShared()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isShared
    }
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The most recent dateTimeOffset when a printer interacted with Universal Print. Read-only.
func (m *Printer) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
// GetRegisteredDateTime gets the registeredDateTime property value. The DateTimeOffset when the printer was registered. Read-only.
func (m *Printer) GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.registeredDateTime
    }
}
// GetShares gets the shares property value. The list of printerShares that are associated with the printer. Currently, only one printerShare can be associated with the printer. Read-only. Nullable.
func (m *Printer) GetShares()([]PrinterShare) {
    if m == nil {
        return nil
    } else {
        return m.shares
    }
}
// GetTaskTriggers gets the taskTriggers property value. A list of task triggers that are associated with the printer.
func (m *Printer) GetTaskTriggers()([]PrintTaskTrigger) {
    if m == nil {
        return nil
    } else {
        return m.taskTriggers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Printer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PrinterBase.GetFieldDeserializers()
    res["connectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintConnector() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintConnector, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintConnector))
            }
            m.SetConnectors(res)
        }
        return nil
    }
    res["hasPhysicalDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasPhysicalDevice(val)
        }
        return nil
    }
    res["isShared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsShared(val)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["registeredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegisteredDateTime(val)
        }
        return nil
    }
    res["shares"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterShare() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterShare, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrinterShare))
            }
            m.SetShares(res)
        }
        return nil
    }
    res["taskTriggers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTaskTrigger() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintTaskTrigger, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintTaskTrigger))
            }
            m.SetTaskTriggers(res)
        }
        return nil
    }
    return res
}
func (m *Printer) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Printer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PrinterBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConnectors()))
        for i, v := range m.GetConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("connectors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasPhysicalDevice", m.GetHasPhysicalDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isShared", m.GetIsShared())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registeredDateTime", m.GetRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetShares()))
        for i, v := range m.GetShares() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("shares", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaskTriggers()))
        for i, v := range m.GetTaskTriggers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("taskTriggers", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectors sets the connectors property value. The connectors that are associated with the printer.
func (m *Printer) SetConnectors(value []PrintConnector)() {
    m.connectors = value
}
// SetHasPhysicalDevice sets the hasPhysicalDevice property value. True if the printer has a physical device for printing. Read-only.
func (m *Printer) SetHasPhysicalDevice(value *bool)() {
    m.hasPhysicalDevice = value
}
// SetIsShared sets the isShared property value. True if the printer is shared; false otherwise. Read-only.
func (m *Printer) SetIsShared(value *bool)() {
    m.isShared = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The most recent dateTimeOffset when a printer interacted with Universal Print. Read-only.
func (m *Printer) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetRegisteredDateTime sets the registeredDateTime property value. The DateTimeOffset when the printer was registered. Read-only.
func (m *Printer) SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registeredDateTime = value
}
// SetShares sets the shares property value. The list of printerShares that are associated with the printer. Currently, only one printerShare can be associated with the printer. Read-only. Nullable.
func (m *Printer) SetShares(value []PrinterShare)() {
    m.shares = value
}
// SetTaskTriggers sets the taskTriggers property value. A list of task triggers that are associated with the printer.
func (m *Printer) SetTaskTriggers(value []PrintTaskTrigger)() {
    m.taskTriggers = value
}
