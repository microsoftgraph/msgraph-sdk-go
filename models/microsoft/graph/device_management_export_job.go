package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementExportJob struct {
    Entity
    // Time that the exported report expires
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Filters applied on the report
    filter *string;
    // Format of the exported report. Possible values are: csv, pdf.
    format *DeviceManagementReportFileFormat;
    // Configures how the requested export job is localized. Possible values are: localizedValuesAsAdditionalColumn, replaceLocalizableValues.
    localizationType *DeviceManagementExportJobLocalizationType;
    // Name of the report
    reportName *string;
    // Time that the exported report was requested
    requestDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Columns selected from the report
    select_escaped []string;
    // A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a CachedReportConfiguration id.
    snapshotId *string;
    // Status of the export job. Possible values are: unknown, notStarted, inProgress, completed, failed.
    status *DeviceManagementReportStatus;
    // Temporary location of the exported report
    url *string;
}
// Instantiates a new deviceManagementExportJob and sets the default values.
func NewDeviceManagementExportJob()(*DeviceManagementExportJob) {
    m := &DeviceManagementExportJob{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the expirationDateTime property value. Time that the exported report expires
func (m *DeviceManagementExportJob) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the filter property value. Filters applied on the report
func (m *DeviceManagementExportJob) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// Gets the format property value. Format of the exported report. Possible values are: csv, pdf.
func (m *DeviceManagementExportJob) GetFormat()(*DeviceManagementReportFileFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the localizationType property value. Configures how the requested export job is localized. Possible values are: localizedValuesAsAdditionalColumn, replaceLocalizableValues.
func (m *DeviceManagementExportJob) GetLocalizationType()(*DeviceManagementExportJobLocalizationType) {
    if m == nil {
        return nil
    } else {
        return m.localizationType
    }
}
// Gets the reportName property value. Name of the report
func (m *DeviceManagementExportJob) GetReportName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportName
    }
}
// Gets the requestDateTime property value. Time that the exported report was requested
func (m *DeviceManagementExportJob) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestDateTime
    }
}
// Gets the select_escaped property value. Columns selected from the report
func (m *DeviceManagementExportJob) GetSelect_escaped()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select_escaped
    }
}
// Gets the snapshotId property value. A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a CachedReportConfiguration id.
func (m *DeviceManagementExportJob) GetSnapshotId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.snapshotId
    }
}
// Gets the status property value. Status of the export job. Possible values are: unknown, notStarted, inProgress, completed, failed.
func (m *DeviceManagementExportJob) GetStatus()(*DeviceManagementReportStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the url property value. Temporary location of the exported report
func (m *DeviceManagementExportJob) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// The deserialization information for the current model
func (m *DeviceManagementExportJob) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["filter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFilter(val)
        return nil
    }
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportFileFormat)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementReportFileFormat)
        m.SetFormat(&cast)
        return nil
    }
    res["localizationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExportJobLocalizationType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementExportJobLocalizationType)
        m.SetLocalizationType(&cast)
        return nil
    }
    res["reportName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportName(val)
        return nil
    }
    res["requestDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRequestDateTime(val)
        return nil
    }
    res["select_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSelect_escaped(res)
        return nil
    }
    res["snapshotId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSnapshotId(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportStatus)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementReportStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrl(val)
        return nil
    }
    return res
}
func (m *DeviceManagementExportJob) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementExportJob) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    if m.GetFormat() != nil {
        cast := m.GetFormat().String()
        err = writer.WriteStringValue("format", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalizationType() != nil {
        cast := m.GetLocalizationType().String()
        err = writer.WriteStringValue("localizationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportName", m.GetReportName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestDateTime", m.GetRequestDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("select_escaped", m.GetSelect_escaped())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("snapshotId", m.GetSnapshotId())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the expirationDateTime property value. Time that the exported report expires
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *DeviceManagementExportJob) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the filter property value. Filters applied on the report
// Parameters:
//  - value : Value to set for the filter property.
func (m *DeviceManagementExportJob) SetFilter(value *string)() {
    m.filter = value
}
// Sets the format property value. Format of the exported report. Possible values are: csv, pdf.
// Parameters:
//  - value : Value to set for the format property.
func (m *DeviceManagementExportJob) SetFormat(value *DeviceManagementReportFileFormat)() {
    m.format = value
}
// Sets the localizationType property value. Configures how the requested export job is localized. Possible values are: localizedValuesAsAdditionalColumn, replaceLocalizableValues.
// Parameters:
//  - value : Value to set for the localizationType property.
func (m *DeviceManagementExportJob) SetLocalizationType(value *DeviceManagementExportJobLocalizationType)() {
    m.localizationType = value
}
// Sets the reportName property value. Name of the report
// Parameters:
//  - value : Value to set for the reportName property.
func (m *DeviceManagementExportJob) SetReportName(value *string)() {
    m.reportName = value
}
// Sets the requestDateTime property value. Time that the exported report was requested
// Parameters:
//  - value : Value to set for the requestDateTime property.
func (m *DeviceManagementExportJob) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestDateTime = value
}
// Sets the select_escaped property value. Columns selected from the report
// Parameters:
//  - value : Value to set for the select_escaped property.
func (m *DeviceManagementExportJob) SetSelect_escaped(value []string)() {
    m.select_escaped = value
}
// Sets the snapshotId property value. A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a CachedReportConfiguration id.
// Parameters:
//  - value : Value to set for the snapshotId property.
func (m *DeviceManagementExportJob) SetSnapshotId(value *string)() {
    m.snapshotId = value
}
// Sets the status property value. Status of the export job. Possible values are: unknown, notStarted, inProgress, completed, failed.
// Parameters:
//  - value : Value to set for the status property.
func (m *DeviceManagementExportJob) SetStatus(value *DeviceManagementReportStatus)() {
    m.status = value
}
// Sets the url property value. Temporary location of the exported report
// Parameters:
//  - value : Value to set for the url property.
func (m *DeviceManagementExportJob) SetUrl(value *string)() {
    m.url = value
}
