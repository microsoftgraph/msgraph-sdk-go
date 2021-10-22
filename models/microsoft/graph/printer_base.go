package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrinterBase struct {
    Entity
    capabilities *PrinterCapabilities;
    defaults *PrinterDefaults;
    displayName *string;
    isAcceptingJobs *bool;
    jobs []PrintJob;
    location *PrinterLocation;
    manufacturer *string;
    model *string;
    status *PrinterStatus;
}
func NewPrinterBase()(*PrinterBase) {
    m := &PrinterBase{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PrinterBase) GetCapabilities()(*PrinterCapabilities) {
    if m == nil {
        return nil
    } else {
        return m.capabilities
    }
}
func (m *PrinterBase) GetDefaults()(*PrinterDefaults) {
    if m == nil {
        return nil
    } else {
        return m.defaults
    }
}
func (m *PrinterBase) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *PrinterBase) GetIsAcceptingJobs()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAcceptingJobs
    }
}
func (m *PrinterBase) GetJobs()([]PrintJob) {
    if m == nil {
        return nil
    } else {
        return m.jobs
    }
}
func (m *PrinterBase) GetLocation()(*PrinterLocation) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
func (m *PrinterBase) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
func (m *PrinterBase) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *PrinterBase) GetStatus()(*PrinterStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *PrinterBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["capabilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterCapabilities() })
        if err != nil {
            return err
        }
        m.SetCapabilities(val.(*PrinterCapabilities))
        return nil
    }
    res["defaults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterDefaults() })
        if err != nil {
            return err
        }
        m.SetDefaults(val.(*PrinterDefaults))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isAcceptingJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAcceptingJobs(val)
        return nil
    }
    res["jobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintJob() })
        if err != nil {
            return err
        }
        res := make([]PrintJob, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintJob))
        }
        m.SetJobs(res)
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterLocation() })
        if err != nil {
            return err
        }
        m.SetLocation(val.(*PrinterLocation))
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManufacturer(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterStatus() })
        if err != nil {
            return err
        }
        m.SetStatus(val.(*PrinterStatus))
        return nil
    }
    return res
}
func (m *PrinterBase) IsNil()(bool) {
    return m == nil
}
func (m *PrinterBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("capabilities", m.GetCapabilities())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaults", m.GetDefaults())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAcceptingJobs", m.GetIsAcceptingJobs())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetJobs()))
        for i, v := range m.GetJobs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("jobs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PrinterBase) SetCapabilities(value *PrinterCapabilities)() {
    m.capabilities = value
}
func (m *PrinterBase) SetDefaults(value *PrinterDefaults)() {
    m.defaults = value
}
func (m *PrinterBase) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *PrinterBase) SetIsAcceptingJobs(value *bool)() {
    m.isAcceptingJobs = value
}
func (m *PrinterBase) SetJobs(value []PrintJob)() {
    m.jobs = value
}
func (m *PrinterBase) SetLocation(value *PrinterLocation)() {
    m.location = value
}
func (m *PrinterBase) SetManufacturer(value *string)() {
    m.manufacturer = value
}
func (m *PrinterBase) SetModel(value *string)() {
    m.model = value
}
func (m *PrinterBase) SetStatus(value *PrinterStatus)() {
    m.status = value
}
