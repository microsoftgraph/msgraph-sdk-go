package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementReports struct {
    Entity
    // Entity representing a job to export a report
    exportJobs []DeviceManagementExportJob;
}
// Instantiates a new deviceManagementReports and sets the default values.
func NewDeviceManagementReports()(*DeviceManagementReports) {
    m := &DeviceManagementReports{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the exportJobs property value. Entity representing a job to export a report
func (m *DeviceManagementReports) GetExportJobs()([]DeviceManagementExportJob) {
    if m == nil {
        return nil
    } else {
        return m.exportJobs
    }
}
// The deserialization information for the current model
func (m *DeviceManagementReports) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exportJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExportJob() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExportJob, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementExportJob))
            }
            m.SetExportJobs(res)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementReports) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementReports) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExportJobs()))
        for i, v := range m.GetExportJobs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exportJobs", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the exportJobs property value. Entity representing a job to export a report
// Parameters:
//  - value : Value to set for the exportJobs property.
func (m *DeviceManagementReports) SetExportJobs(value []DeviceManagementExportJob)() {
    m.exportJobs = value
}
