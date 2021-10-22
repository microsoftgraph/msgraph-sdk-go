package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewReviewerScope struct {
    additionalData map[string]interface{};
    query *string;
    queryRoot *string;
    queryType *string;
}
func NewAccessReviewReviewerScope()(*AccessReviewReviewerScope) {
    m := &AccessReviewReviewerScope{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AccessReviewReviewerScope) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AccessReviewReviewerScope) GetQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.query
    }
}
func (m *AccessReviewReviewerScope) GetQueryRoot()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryRoot
    }
}
func (m *AccessReviewReviewerScope) GetQueryType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryType
    }
}
func (m *AccessReviewReviewerScope) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["query"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQuery(val)
        return nil
    }
    res["queryRoot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQueryRoot(val)
        return nil
    }
    res["queryType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQueryType(val)
        return nil
    }
    return res
}
func (m *AccessReviewReviewerScope) IsNil()(bool) {
    return m == nil
}
func (m *AccessReviewReviewerScope) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryRoot", m.GetQueryRoot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryType", m.GetQueryType())
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
func (m *AccessReviewReviewerScope) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AccessReviewReviewerScope) SetQuery(value *string)() {
    m.query = value
}
func (m *AccessReviewReviewerScope) SetQueryRoot(value *string)() {
    m.queryRoot = value
}
func (m *AccessReviewReviewerScope) SetQueryType(value *string)() {
    m.queryType = value
}
