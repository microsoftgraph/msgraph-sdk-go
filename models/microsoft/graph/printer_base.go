package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterBase 
type PrinterBase struct {
    Entity
    // The capabilities of the printer/printerShare.
    capabilities *PrinterCapabilities;
    // The default print settings of printer/printerShare.
    defaults *PrinterDefaults;
    // The name of the printer/printerShare.
    displayName *string;
    // Whether the printer/printerShare is currently accepting new print jobs.
    isAcceptingJobs *bool;
    // The list of jobs that are queued for printing by the printer/printerShare.
    jobs []PrintJob;
    // The physical and/or organizational location of the printer/printerShare.
    location *PrinterLocation;
    // The manufacturer of the printer/printerShare.
    manufacturer *string;
    // The model name of the printer/printerShare.
    model *string;
    // 
    status *PrinterStatus;
}
// NewPrinterBase instantiates a new printerBase and sets the default values.
func NewPrinterBase()(*PrinterBase) {
    m := &PrinterBase{
        Entity: *NewEntity(),
    }
    return m
}
// GetCapabilities gets the capabilities property value. The capabilities of the printer/printerShare.
func (m *PrinterBase) GetCapabilities()(*PrinterCapabilities) {
    if m == nil {
        return nil
    } else {
        return m.capabilities
    }
}
// GetDefaults gets the defaults property value. The default print settings of printer/printerShare.
func (m *PrinterBase) GetDefaults()(*PrinterDefaults) {
    if m == nil {
        return nil
    } else {
        return m.defaults
    }
}
// GetDisplayName gets the displayName property value. The name of the printer/printerShare.
func (m *PrinterBase) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsAcceptingJobs gets the isAcceptingJobs property value. Whether the printer/printerShare is currently accepting new print jobs.
func (m *PrinterBase) GetIsAcceptingJobs()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAcceptingJobs
    }
}
// GetJobs gets the jobs property value. The list of jobs that are queued for printing by the printer/printerShare.
func (m *PrinterBase) GetJobs()([]PrintJob) {
    if m == nil {
        return nil
    } else {
        return m.jobs
    }
}
// GetLocation gets the location property value. The physical and/or organizational location of the printer/printerShare.
func (m *PrinterBase) GetLocation()(*PrinterLocation) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// GetManufacturer gets the manufacturer property value. The manufacturer of the printer/printerShare.
func (m *PrinterBase) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetModel gets the model property value. The model name of the printer/printerShare.
func (m *PrinterBase) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetStatus gets the status property value. 
func (m *PrinterBase) GetStatus()(*PrinterStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["capabilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterCapabilities() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilities(val.(*PrinterCapabilities))
        }
        return nil
    }
    res["defaults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterDefaults() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaults(val.(*PrinterDefaults))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isAcceptingJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAcceptingJobs(val)
        }
        return nil
    }
    res["jobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintJob() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintJob, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintJob))
            }
            m.SetJobs(res)
        }
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(*PrinterLocation))
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*PrinterStatus))
        }
        return nil
    }
    return res
}
func (m *PrinterBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCapabilities sets the capabilities property value. The capabilities of the printer/printerShare.
func (m *PrinterBase) SetCapabilities(value *PrinterCapabilities)() {
    if m != nil {
        m.capabilities = value
    }
}
// SetDefaults sets the defaults property value. The default print settings of printer/printerShare.
func (m *PrinterBase) SetDefaults(value *PrinterDefaults)() {
    if m != nil {
        m.defaults = value
    }
}
// SetDisplayName sets the displayName property value. The name of the printer/printerShare.
func (m *PrinterBase) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsAcceptingJobs sets the isAcceptingJobs property value. Whether the printer/printerShare is currently accepting new print jobs.
func (m *PrinterBase) SetIsAcceptingJobs(value *bool)() {
    if m != nil {
        m.isAcceptingJobs = value
    }
}
// SetJobs sets the jobs property value. The list of jobs that are queued for printing by the printer/printerShare.
func (m *PrinterBase) SetJobs(value []PrintJob)() {
    if m != nil {
        m.jobs = value
    }
}
// SetLocation sets the location property value. The physical and/or organizational location of the printer/printerShare.
func (m *PrinterBase) SetLocation(value *PrinterLocation)() {
    if m != nil {
        m.location = value
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer of the printer/printerShare.
func (m *PrinterBase) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetModel sets the model property value. The model name of the printer/printerShare.
func (m *PrinterBase) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetStatus sets the status property value. 
func (m *PrinterBase) SetStatus(value *PrinterStatus)() {
    if m != nil {
        m.status = value
    }
}
