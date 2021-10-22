package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AppConsentApprovalRoute struct {
    Entity
    appConsentRequests []AppConsentRequest;
}
func NewAppConsentApprovalRoute()(*AppConsentApprovalRoute) {
    m := &AppConsentApprovalRoute{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AppConsentApprovalRoute) GetAppConsentRequests()([]AppConsentRequest) {
    if m == nil {
        return nil
    } else {
        return m.appConsentRequests
    }
}
func (m *AppConsentApprovalRoute) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appConsentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppConsentRequest() })
        if err != nil {
            return err
        }
        res := make([]AppConsentRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppConsentRequest))
        }
        m.SetAppConsentRequests(res)
        return nil
    }
    return res
}
func (m *AppConsentApprovalRoute) IsNil()(bool) {
    return m == nil
}
func (m *AppConsentApprovalRoute) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppConsentRequests()))
        for i, v := range m.GetAppConsentRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appConsentRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AppConsentApprovalRoute) SetAppConsentRequests(value []AppConsentRequest)() {
    m.appConsentRequests = value
}
