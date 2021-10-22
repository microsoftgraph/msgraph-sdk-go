package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SharingDetail struct {
    additionalData map[string]interface{};
    sharedBy *InsightIdentity;
    sharedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    sharingReference *ResourceReference;
    sharingSubject *string;
    sharingType *string;
}
func NewSharingDetail()(*SharingDetail) {
    m := &SharingDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SharingDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SharingDetail) GetSharedBy()(*InsightIdentity) {
    if m == nil {
        return nil
    } else {
        return m.sharedBy
    }
}
func (m *SharingDetail) GetSharedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.sharedDateTime
    }
}
func (m *SharingDetail) GetSharingReference()(*ResourceReference) {
    if m == nil {
        return nil
    } else {
        return m.sharingReference
    }
}
func (m *SharingDetail) GetSharingSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sharingSubject
    }
}
func (m *SharingDetail) GetSharingType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sharingType
    }
}
func (m *SharingDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["sharedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInsightIdentity() })
        if err != nil {
            return err
        }
        m.SetSharedBy(val.(*InsightIdentity))
        return nil
    }
    res["sharedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSharedDateTime(val)
        return nil
    }
    res["sharingReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceReference() })
        if err != nil {
            return err
        }
        m.SetSharingReference(val.(*ResourceReference))
        return nil
    }
    res["sharingSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSharingSubject(val)
        return nil
    }
    res["sharingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSharingType(val)
        return nil
    }
    return res
}
func (m *SharingDetail) IsNil()(bool) {
    return m == nil
}
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
func (m *SharingDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SharingDetail) SetSharedBy(value *InsightIdentity)() {
    m.sharedBy = value
}
func (m *SharingDetail) SetSharedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sharedDateTime = value
}
func (m *SharingDetail) SetSharingReference(value *ResourceReference)() {
    m.sharingReference = value
}
func (m *SharingDetail) SetSharingSubject(value *string)() {
    m.sharingSubject = value
}
func (m *SharingDetail) SetSharingType(value *string)() {
    m.sharingType = value
}
