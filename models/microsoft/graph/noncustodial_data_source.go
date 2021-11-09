package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type NoncustodialDataSource struct {
    DataSourceContainer
    // Indicates if hold is applied to non-custodial data source (such as mailbox or site).
    applyHoldToSource *bool;
    // User source or SharePoint site data source as non-custodial data source.
    dataSource *DataSource;
}
// Instantiates a new noncustodialDataSource and sets the default values.
func NewNoncustodialDataSource()(*NoncustodialDataSource) {
    m := &NoncustodialDataSource{
        DataSourceContainer: *NewDataSourceContainer(),
    }
    return m
}
// Gets the applyHoldToSource property value. Indicates if hold is applied to non-custodial data source (such as mailbox or site).
func (m *NoncustodialDataSource) GetApplyHoldToSource()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applyHoldToSource
    }
}
// Gets the dataSource property value. User source or SharePoint site data source as non-custodial data source.
func (m *NoncustodialDataSource) GetDataSource()(*DataSource) {
    if m == nil {
        return nil
    } else {
        return m.dataSource
    }
}
// The deserialization information for the current model
func (m *NoncustodialDataSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DataSourceContainer.GetFieldDeserializers()
    res["applyHoldToSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyHoldToSource(val)
        }
        return nil
    }
    res["dataSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDataSource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSource(val.(*DataSource))
        }
        return nil
    }
    return res
}
func (m *NoncustodialDataSource) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *NoncustodialDataSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DataSourceContainer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("applyHoldToSource", m.GetApplyHoldToSource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dataSource", m.GetDataSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the applyHoldToSource property value. Indicates if hold is applied to non-custodial data source (such as mailbox or site).
// Parameters:
//  - value : Value to set for the applyHoldToSource property.
func (m *NoncustodialDataSource) SetApplyHoldToSource(value *bool)() {
    m.applyHoldToSource = value
}
// Sets the dataSource property value. User source or SharePoint site data source as non-custodial data source.
// Parameters:
//  - value : Value to set for the dataSource property.
func (m *NoncustodialDataSource) SetDataSource(value *DataSource)() {
    m.dataSource = value
}
