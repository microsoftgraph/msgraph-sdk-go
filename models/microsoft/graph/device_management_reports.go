package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementReports 
type DeviceManagementReports struct {
    Entity
    // Entity representing a job to export a report
    exportJobs []DeviceManagementExportJob;
}
// NewDeviceManagementReports instantiates a new deviceManagementReports and sets the default values.
func NewDeviceManagementReports()(*DeviceManagementReports) {
    m := &DeviceManagementReports{
        Entity: *NewEntity(),
    }
    return m
}
// GetExportJobs gets the exportJobs property value. Entity representing a job to export a report
func (m *DeviceManagementReports) GetExportJobs()([]DeviceManagementExportJob) {
    if m == nil {
        return nil
    } else {
        return m.exportJobs
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *DeviceManagementReports) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExportJobs() != nil {
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
// SetExportJobs sets the exportJobs property value. Entity representing a job to export a report
func (m *DeviceManagementReports) SetExportJobs(value []DeviceManagementExportJob)() {
    if m != nil {
        m.exportJobs = value
    }
}
