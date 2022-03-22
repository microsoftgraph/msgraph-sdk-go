package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ArchivedPrintJob 
type ArchivedPrintJob struct {
    // True if the job was acquired by a printer; false otherwise. Read-only.
    acquiredByPrinter *bool;
    // The dateTimeOffset when the job was acquired by the printer, if any. Read-only.
    acquiredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The dateTimeOffset when the job was completed, canceled or aborted. Read-only.
    completionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The number of copies that were printed. Read-only.
    copiesPrinted *int32;
    // The user who created the print job. Read-only.
    createdBy UserIdentityable;
    // The dateTimeOffset when the job was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The archived print job's GUID. Read-only.
    id *string;
    // The printer ID that the job was queued for. Read-only.
    printerId *string;
    // The print job's final processing state. Read-only.
    processingState *PrintJobProcessingState;
}
// NewArchivedPrintJob instantiates a new archivedPrintJob and sets the default values.
func NewArchivedPrintJob()(*ArchivedPrintJob) {
    m := &ArchivedPrintJob{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateArchivedPrintJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateArchivedPrintJobFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewArchivedPrintJob(), nil
}
// GetAcquiredByPrinter gets the acquiredByPrinter property value. True if the job was acquired by a printer; false otherwise. Read-only.
func (m *ArchivedPrintJob) GetAcquiredByPrinter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.acquiredByPrinter
    }
}
// GetAcquiredDateTime gets the acquiredDateTime property value. The dateTimeOffset when the job was acquired by the printer, if any. Read-only.
func (m *ArchivedPrintJob) GetAcquiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.acquiredDateTime
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ArchivedPrintJob) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCompletionDateTime gets the completionDateTime property value. The dateTimeOffset when the job was completed, canceled or aborted. Read-only.
func (m *ArchivedPrintJob) GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completionDateTime
    }
}
// GetCopiesPrinted gets the copiesPrinted property value. The number of copies that were printed. Read-only.
func (m *ArchivedPrintJob) GetCopiesPrinted()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.copiesPrinted
    }
}
// GetCreatedBy gets the createdBy property value. The user who created the print job. Read-only.
func (m *ArchivedPrintJob) GetCreatedBy()(UserIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The dateTimeOffset when the job was created. Read-only.
func (m *ArchivedPrintJob) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ArchivedPrintJob) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["acquiredByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcquiredByPrinter(val)
        }
        return nil
    }
    res["acquiredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcquiredDateTime(val)
        }
        return nil
    }
    res["completionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionDateTime(val)
        }
        return nil
    }
    res["copiesPrinted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopiesPrinted(val)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(UserIdentityable))
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
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["printerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterId(val)
        }
        return nil
    }
    res["processingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintJobProcessingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingState(val.(*PrintJobProcessingState))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The archived print job's GUID. Read-only.
func (m *ArchivedPrintJob) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetPrinterId gets the printerId property value. The printer ID that the job was queued for. Read-only.
func (m *ArchivedPrintJob) GetPrinterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.printerId
    }
}
// GetProcessingState gets the processingState property value. The print job's final processing state. Read-only.
func (m *ArchivedPrintJob) GetProcessingState()(*PrintJobProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.processingState
    }
}
// Serialize serializes information the current object
func (m *ArchivedPrintJob) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("acquiredByPrinter", m.GetAcquiredByPrinter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("acquiredDateTime", m.GetAcquiredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("completionDateTime", m.GetCompletionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("copiesPrinted", m.GetCopiesPrinted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("printerId", m.GetPrinterId())
        if err != nil {
            return err
        }
    }
    if m.GetProcessingState() != nil {
        cast := (*m.GetProcessingState()).String()
        err := writer.WriteStringValue("processingState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcquiredByPrinter sets the acquiredByPrinter property value. True if the job was acquired by a printer; false otherwise. Read-only.
func (m *ArchivedPrintJob) SetAcquiredByPrinter(value *bool)() {
    if m != nil {
        m.acquiredByPrinter = value
    }
}
// SetAcquiredDateTime sets the acquiredDateTime property value. The dateTimeOffset when the job was acquired by the printer, if any. Read-only.
func (m *ArchivedPrintJob) SetAcquiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.acquiredDateTime = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ArchivedPrintJob) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCompletionDateTime sets the completionDateTime property value. The dateTimeOffset when the job was completed, canceled or aborted. Read-only.
func (m *ArchivedPrintJob) SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.completionDateTime = value
    }
}
// SetCopiesPrinted sets the copiesPrinted property value. The number of copies that were printed. Read-only.
func (m *ArchivedPrintJob) SetCopiesPrinted(value *int32)() {
    if m != nil {
        m.copiesPrinted = value
    }
}
// SetCreatedBy sets the createdBy property value. The user who created the print job. Read-only.
func (m *ArchivedPrintJob) SetCreatedBy(value UserIdentityable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The dateTimeOffset when the job was created. Read-only.
func (m *ArchivedPrintJob) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetId sets the id property value. The archived print job's GUID. Read-only.
func (m *ArchivedPrintJob) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetPrinterId sets the printerId property value. The printer ID that the job was queued for. Read-only.
func (m *ArchivedPrintJob) SetPrinterId(value *string)() {
    if m != nil {
        m.printerId = value
    }
}
// SetProcessingState sets the processingState property value. The print job's final processing state. Read-only.
func (m *ArchivedPrintJob) SetProcessingState(value *PrintJobProcessingState)() {
    if m != nil {
        m.processingState = value
    }
}
