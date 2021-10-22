package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Quota struct {
    additionalData map[string]interface{};
    deleted *int64;
    remaining *int64;
    state *string;
    storagePlanInformation *StoragePlanInformation;
    total *int64;
    used *int64;
}
func NewQuota()(*Quota) {
    m := &Quota{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Quota) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Quota) GetDeleted()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deleted
    }
}
func (m *Quota) GetRemaining()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.remaining
    }
}
func (m *Quota) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *Quota) GetStoragePlanInformation()(*StoragePlanInformation) {
    if m == nil {
        return nil
    } else {
        return m.storagePlanInformation
    }
}
func (m *Quota) GetTotal()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
func (m *Quota) GetUsed()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.used
    }
}
func (m *Quota) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDeleted(val)
        return nil
    }
    res["remaining"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetRemaining(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetState(val)
        return nil
    }
    res["storagePlanInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStoragePlanInformation() })
        if err != nil {
            return err
        }
        m.SetStoragePlanInformation(val.(*StoragePlanInformation))
        return nil
    }
    res["total"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotal(val)
        return nil
    }
    res["used"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetUsed(val)
        return nil
    }
    return res
}
func (m *Quota) IsNil()(bool) {
    return m == nil
}
func (m *Quota) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("deleted", m.GetDeleted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("remaining", m.GetRemaining())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storagePlanInformation", m.GetStoragePlanInformation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("used", m.GetUsed())
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
func (m *Quota) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Quota) SetDeleted(value *int64)() {
    m.deleted = value
}
func (m *Quota) SetRemaining(value *int64)() {
    m.remaining = value
}
func (m *Quota) SetState(value *string)() {
    m.state = value
}
func (m *Quota) SetStoragePlanInformation(value *StoragePlanInformation)() {
    m.storagePlanInformation = value
}
func (m *Quota) SetTotal(value *int64)() {
    m.total = value
}
func (m *Quota) SetUsed(value *int64)() {
    m.used = value
}
