package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SiteCollection struct {
    additionalData map[string]interface{};
    dataLocationCode *string;
    hostname *string;
    root *Root;
}
func NewSiteCollection()(*SiteCollection) {
    m := &SiteCollection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SiteCollection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SiteCollection) GetDataLocationCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataLocationCode
    }
}
func (m *SiteCollection) GetHostname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hostname
    }
}
func (m *SiteCollection) GetRoot()(*Root) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
func (m *SiteCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dataLocationCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDataLocationCode(val)
        return nil
    }
    res["hostname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHostname(val)
        return nil
    }
    res["root"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoot() })
        if err != nil {
            return err
        }
        m.SetRoot(val.(*Root))
        return nil
    }
    return res
}
func (m *SiteCollection) IsNil()(bool) {
    return m == nil
}
func (m *SiteCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dataLocationCode", m.GetDataLocationCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("root", m.GetRoot())
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
func (m *SiteCollection) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SiteCollection) SetDataLocationCode(value *string)() {
    m.dataLocationCode = value
}
func (m *SiteCollection) SetHostname(value *string)() {
    m.hostname = value
}
func (m *SiteCollection) SetRoot(value *Root)() {
    m.root = value
}
