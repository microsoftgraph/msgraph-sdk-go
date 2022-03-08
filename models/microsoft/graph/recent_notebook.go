package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecentNotebook provides operations to call the getRecentNotebooks method.
type RecentNotebook struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The name of the notebook.
    displayName *string;
    // The date and time when the notebook was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    lastAccessedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote client, if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
    links RecentNotebookLinksable;
    // The backend store where the Notebook resides, either OneDriveForBusiness or OneDrive.
    sourceService *OnenoteSourceService;
}
// NewRecentNotebook instantiates a new recentNotebook and sets the default values.
func NewRecentNotebook()(*RecentNotebook) {
    m := &RecentNotebook{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRecentNotebookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecentNotebookFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRecentNotebook(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecentNotebook) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. The name of the notebook.
func (m *RecentNotebook) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecentNotebook) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["lastAccessedTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAccessedTime(val)
        }
        return nil
    }
    res["links"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecentNotebookLinksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(RecentNotebookLinksable))
        }
        return nil
    }
    res["sourceService"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenoteSourceService)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceService(val.(*OnenoteSourceService))
        }
        return nil
    }
    return res
}
// GetLastAccessedTime gets the lastAccessedTime property value. The date and time when the notebook was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *RecentNotebook) GetLastAccessedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastAccessedTime
    }
}
// GetLinks gets the links property value. Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote client, if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
func (m *RecentNotebook) GetLinks()(RecentNotebookLinksable) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// GetSourceService gets the sourceService property value. The backend store where the Notebook resides, either OneDriveForBusiness or OneDrive.
func (m *RecentNotebook) GetSourceService()(*OnenoteSourceService) {
    if m == nil {
        return nil
    } else {
        return m.sourceService
    }
}
func (m *RecentNotebook) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RecentNotebook) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastAccessedTime", m.GetLastAccessedTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("links", m.GetLinks())
        if err != nil {
            return err
        }
    }
    if m.GetSourceService() != nil {
        cast := (*m.GetSourceService()).String()
        err := writer.WriteStringValue("sourceService", &cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecentNotebook) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. The name of the notebook.
func (m *RecentNotebook) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastAccessedTime sets the lastAccessedTime property value. The date and time when the notebook was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *RecentNotebook) SetLastAccessedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastAccessedTime = value
    }
}
// SetLinks sets the links property value. Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote client, if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
func (m *RecentNotebook) SetLinks(value RecentNotebookLinksable)() {
    if m != nil {
        m.links = value
    }
}
// SetSourceService sets the sourceService property value. The backend store where the Notebook resides, either OneDriveForBusiness or OneDrive.
func (m *RecentNotebook) SetSourceService(value *OnenoteSourceService)() {
    if m != nil {
        m.sourceService = value
    }
}
