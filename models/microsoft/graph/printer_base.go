package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new printerBase and sets the default values.
func NewPrinterBase()(*PrinterBase) {
    m := &PrinterBase{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the capabilities property value. The capabilities of the printer/printerShare.
func (m *PrinterBase) GetCapabilities()(*PrinterCapabilities) {
    if m == nil {
        return nil
    } else {
        return m.capabilities
    }
}
// Gets the defaults property value. The default print settings of printer/printerShare.
func (m *PrinterBase) GetDefaults()(*PrinterDefaults) {
    if m == nil {
        return nil
    } else {
        return m.defaults
    }
}
// Gets the displayName property value. The name of the printer/printerShare.
func (m *PrinterBase) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isAcceptingJobs property value. Whether the printer/printerShare is currently accepting new print jobs.
func (m *PrinterBase) GetIsAcceptingJobs()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAcceptingJobs
    }
}
// Gets the jobs property value. The list of jobs that are queued for printing by the printer/printerShare.
func (m *PrinterBase) GetJobs()([]PrintJob) {
    if m == nil {
        return nil
    } else {
        return m.jobs
    }
}
// Gets the location property value. The physical and/or organizational location of the printer/printerShare.
func (m *PrinterBase) GetLocation()(*PrinterLocation) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// Gets the manufacturer property value. The manufacturer of the printer/printerShare.
func (m *PrinterBase) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the model property value. The model name of the printer/printerShare.
func (m *PrinterBase) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the status property value. 
func (m *PrinterBase) GetStatus()(*PrinterStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the capabilities property value. The capabilities of the printer/printerShare.
// Parameters:
//  - value : Value to set for the capabilities property.
func (m *PrinterBase) SetCapabilities(value *PrinterCapabilities)() {
    m.capabilities = value
}
// Sets the defaults property value. The default print settings of printer/printerShare.
// Parameters:
//  - value : Value to set for the defaults property.
func (m *PrinterBase) SetDefaults(value *PrinterDefaults)() {
    m.defaults = value
}
// Sets the displayName property value. The name of the printer/printerShare.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PrinterBase) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isAcceptingJobs property value. Whether the printer/printerShare is currently accepting new print jobs.
// Parameters:
//  - value : Value to set for the isAcceptingJobs property.
func (m *PrinterBase) SetIsAcceptingJobs(value *bool)() {
    m.isAcceptingJobs = value
}
// Sets the jobs property value. The list of jobs that are queued for printing by the printer/printerShare.
// Parameters:
//  - value : Value to set for the jobs property.
func (m *PrinterBase) SetJobs(value []PrintJob)() {
    m.jobs = value
}
// Sets the location property value. The physical and/or organizational location of the printer/printerShare.
// Parameters:
//  - value : Value to set for the location property.
func (m *PrinterBase) SetLocation(value *PrinterLocation)() {
    m.location = value
}
// Sets the manufacturer property value. The manufacturer of the printer/printerShare.
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *PrinterBase) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the model property value. The model name of the printer/printerShare.
// Parameters:
//  - value : Value to set for the model property.
func (m *PrinterBase) SetModel(value *string)() {
    m.model = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *PrinterBase) SetStatus(value *PrinterStatus)() {
    m.status = value
}
