package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// sharingDetail 
type SharingDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The user who shared the document.
    sharedBy *InsightIdentity;
    // The date and time the file was last shared. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    sharedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    sharingReference *ResourceReference;
    // The subject with which the document was shared.
    sharingSubject *string;
    // Determines the way the document was shared, can be by a 'Link', 'Attachment', 'Group', 'Site'.
    sharingType *string;
}
// NewSharingDetail instantiates a new sharingDetail and sets the default values.
func NewSharingDetail()(*SharingDetail) {
    m := &SharingDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharingDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetSharedBy gets the sharedBy property value. The user who shared the document.
func (m *SharingDetail) GetSharedBy()(*InsightIdentity) {
    if m == nil {
        return nil
    } else {
        return m.sharedBy
    }
}
// GetSharedDateTime gets the sharedDateTime property value. The date and time the file was last shared. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *SharingDetail) GetSharedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.sharedDateTime
    }
}
// GetSharingReference gets the sharingReference property value. 
func (m *SharingDetail) GetSharingReference()(*ResourceReference) {
    if m == nil {
        return nil
    } else {
        return m.sharingReference
    }
}
// GetSharingSubject gets the sharingSubject property value. The subject with which the document was shared.
func (m *SharingDetail) GetSharingSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sharingSubject
    }
}
// GetSharingType gets the sharingType property value. Determines the way the document was shared, can be by a 'Link', 'Attachment', 'Group', 'Site'.
func (m *SharingDetail) GetSharingType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sharingType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharingDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["sharedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInsightIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedBy(val.(*InsightIdentity))
        }
        return nil
    }
    res["sharedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedDateTime(val)
        }
        return nil
    }
    res["sharingReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceReference() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharingReference(val.(*ResourceReference))
        }
        return nil
    }
    res["sharingSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharingSubject(val)
        }
        return nil
    }
    res["sharingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharingType(val)
        }
        return nil
    }
    return res
}
func (m *SharingDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SharingDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("sharedBy", m.GetSharedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("sharedDateTime", m.GetSharedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sharingReference", m.GetSharingReference())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sharingSubject", m.GetSharingSubject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sharingType", m.GetSharingType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharingDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSharedBy sets the sharedBy property value. The user who shared the document.
func (m *SharingDetail) SetSharedBy(value *InsightIdentity)() {
    m.sharedBy = value
}
// SetSharedDateTime sets the sharedDateTime property value. The date and time the file was last shared. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *SharingDetail) SetSharedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sharedDateTime = value
}
// SetSharingReference sets the sharingReference property value. 
func (m *SharingDetail) SetSharingReference(value *ResourceReference)() {
    m.sharingReference = value
}
// SetSharingSubject sets the sharingSubject property value. The subject with which the document was shared.
func (m *SharingDetail) SetSharingSubject(value *string)() {
    m.sharingSubject = value
}
// SetSharingType sets the sharingType property value. Determines the way the document was shared, can be by a 'Link', 'Attachment', 'Group', 'Site'.
func (m *SharingDetail) SetSharingType(value *string)() {
    m.sharingType = value
}
