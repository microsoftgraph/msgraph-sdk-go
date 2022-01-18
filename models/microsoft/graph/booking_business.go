package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingBusiness 
type BookingBusiness struct {
    Entity
    // The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page.
    address *PhysicalAddress;
    // All the appointments of this business. Read-only. Nullable.
    appointments []BookingAppointment;
    // The hours of operation for the business.
    businessHours []BookingWorkHours;
    // The type of business.
    businessType *string;
    // The set of appointments of this business in a specified date range. Read-only. Nullable.
    calendarView []BookingAppointment;
    // All the customers of this business. Read-only. Nullable.
    customers []BookingCustomerBase;
    // All the custom questions of this business. Read-only. Nullable.
    customQuestions []BookingCustomQuestion;
    // The code for the currency that the business operates in on Microsoft Bookings.
    defaultCurrencyIso *string;
    // The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
    displayName *string;
    // The email address for the business.
    email *string;
    // The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
    isPublished *bool;
    // The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
    phone *string;
    // The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
    publicUrl *string;
    // Specifies how bookings can be created for this business.
    schedulingPolicy *BookingSchedulingPolicy;
    // All the services offered by this business. Read-only. Nullable.
    services []BookingService;
    // All the staff members that provide services in this business. Read-only. Nullable.
    staffMembers []BookingStaffMemberBase;
    // The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
    webSiteUrl *string;
}
// NewBookingBusiness instantiates a new bookingBusiness and sets the default values.
func NewBookingBusiness()(*BookingBusiness) {
    m := &BookingBusiness{
        Entity: *NewEntity(),
    }
    return m
}
// GetAddress gets the address property value. The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page.
func (m *BookingBusiness) GetAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetAppointments gets the appointments property value. All the appointments of this business. Read-only. Nullable.
func (m *BookingBusiness) GetAppointments()([]BookingAppointment) {
    if m == nil {
        return nil
    } else {
        return m.appointments
    }
}
// GetBusinessHours gets the businessHours property value. The hours of operation for the business.
func (m *BookingBusiness) GetBusinessHours()([]BookingWorkHours) {
    if m == nil {
        return nil
    } else {
        return m.businessHours
    }
}
// GetBusinessType gets the businessType property value. The type of business.
func (m *BookingBusiness) GetBusinessType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.businessType
    }
}
// GetCalendarView gets the calendarView property value. The set of appointments of this business in a specified date range. Read-only. Nullable.
func (m *BookingBusiness) GetCalendarView()([]BookingAppointment) {
    if m == nil {
        return nil
    } else {
        return m.calendarView
    }
}
// GetCustomers gets the customers property value. All the customers of this business. Read-only. Nullable.
func (m *BookingBusiness) GetCustomers()([]BookingCustomerBase) {
    if m == nil {
        return nil
    } else {
        return m.customers
    }
}
// GetCustomQuestions gets the customQuestions property value. All the custom questions of this business. Read-only. Nullable.
func (m *BookingBusiness) GetCustomQuestions()([]BookingCustomQuestion) {
    if m == nil {
        return nil
    } else {
        return m.customQuestions
    }
}
// GetDefaultCurrencyIso gets the defaultCurrencyIso property value. The code for the currency that the business operates in on Microsoft Bookings.
func (m *BookingBusiness) GetDefaultCurrencyIso()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultCurrencyIso
    }
}
// GetDisplayName gets the displayName property value. The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
func (m *BookingBusiness) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmail gets the email property value. The email address for the business.
func (m *BookingBusiness) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetIsPublished gets the isPublished property value. The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
func (m *BookingBusiness) GetIsPublished()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPublished
    }
}
// GetPhone gets the phone property value. The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
func (m *BookingBusiness) GetPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phone
    }
}
// GetPublicUrl gets the publicUrl property value. The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
func (m *BookingBusiness) GetPublicUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicUrl
    }
}
// GetSchedulingPolicy gets the schedulingPolicy property value. Specifies how bookings can be created for this business.
func (m *BookingBusiness) GetSchedulingPolicy()(*BookingSchedulingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.schedulingPolicy
    }
}
// GetServices gets the services property value. All the services offered by this business. Read-only. Nullable.
func (m *BookingBusiness) GetServices()([]BookingService) {
    if m == nil {
        return nil
    } else {
        return m.services
    }
}
// GetStaffMembers gets the staffMembers property value. All the staff members that provide services in this business. Read-only. Nullable.
func (m *BookingBusiness) GetStaffMembers()([]BookingStaffMemberBase) {
    if m == nil {
        return nil
    } else {
        return m.staffMembers
    }
}
// GetWebSiteUrl gets the webSiteUrl property value. The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
func (m *BookingBusiness) GetWebSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webSiteUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingBusiness) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(*PhysicalAddress))
        }
        return nil
    }
    res["appointments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingAppointment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingAppointment, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingAppointment))
            }
            m.SetAppointments(res)
        }
        return nil
    }
    res["businessHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingWorkHours() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkHours, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingWorkHours))
            }
            m.SetBusinessHours(res)
        }
        return nil
    }
    res["businessType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessType(val)
        }
        return nil
    }
    res["calendarView"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingAppointment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingAppointment, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingAppointment))
            }
            m.SetCalendarView(res)
        }
        return nil
    }
    res["customers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingCustomerBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCustomerBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingCustomerBase))
            }
            m.SetCustomers(res)
        }
        return nil
    }
    res["customQuestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingCustomQuestion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCustomQuestion, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingCustomQuestion))
            }
            m.SetCustomQuestions(res)
        }
        return nil
    }
    res["defaultCurrencyIso"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultCurrencyIso(val)
        }
        return nil
    }
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
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["isPublished"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPublished(val)
        }
        return nil
    }
    res["phone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    res["publicUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicUrl(val)
        }
        return nil
    }
    res["schedulingPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingSchedulingPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingPolicy(val.(*BookingSchedulingPolicy))
        }
        return nil
    }
    res["services"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingService() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingService, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingService))
            }
            m.SetServices(res)
        }
        return nil
    }
    res["staffMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingStaffMemberBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingStaffMemberBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingStaffMemberBase))
            }
            m.SetStaffMembers(res)
        }
        return nil
    }
    res["webSiteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebSiteUrl(val)
        }
        return nil
    }
    return res
}
func (m *BookingBusiness) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingBusiness) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppointments()))
        for i, v := range m.GetAppointments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appointments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBusinessHours()))
        for i, v := range m.GetBusinessHours() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("businessHours", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("businessType", m.GetBusinessType())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomers()))
        for i, v := range m.GetCustomers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomQuestions()))
        for i, v := range m.GetCustomQuestions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customQuestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultCurrencyIso", m.GetDefaultCurrencyIso())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPublished", m.GetIsPublished())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publicUrl", m.GetPublicUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedulingPolicy", m.GetSchedulingPolicy())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStaffMembers()))
        for i, v := range m.GetStaffMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("staffMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webSiteUrl", m.GetWebSiteUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page.
func (m *BookingBusiness) SetAddress(value *PhysicalAddress)() {
    if m != nil {
        m.address = value
    }
}
// SetAppointments sets the appointments property value. All the appointments of this business. Read-only. Nullable.
func (m *BookingBusiness) SetAppointments(value []BookingAppointment)() {
    if m != nil {
        m.appointments = value
    }
}
// SetBusinessHours sets the businessHours property value. The hours of operation for the business.
func (m *BookingBusiness) SetBusinessHours(value []BookingWorkHours)() {
    if m != nil {
        m.businessHours = value
    }
}
// SetBusinessType sets the businessType property value. The type of business.
func (m *BookingBusiness) SetBusinessType(value *string)() {
    if m != nil {
        m.businessType = value
    }
}
// SetCalendarView sets the calendarView property value. The set of appointments of this business in a specified date range. Read-only. Nullable.
func (m *BookingBusiness) SetCalendarView(value []BookingAppointment)() {
    if m != nil {
        m.calendarView = value
    }
}
// SetCustomers sets the customers property value. All the customers of this business. Read-only. Nullable.
func (m *BookingBusiness) SetCustomers(value []BookingCustomerBase)() {
    if m != nil {
        m.customers = value
    }
}
// SetCustomQuestions sets the customQuestions property value. All the custom questions of this business. Read-only. Nullable.
func (m *BookingBusiness) SetCustomQuestions(value []BookingCustomQuestion)() {
    if m != nil {
        m.customQuestions = value
    }
}
// SetDefaultCurrencyIso sets the defaultCurrencyIso property value. The code for the currency that the business operates in on Microsoft Bookings.
func (m *BookingBusiness) SetDefaultCurrencyIso(value *string)() {
    if m != nil {
        m.defaultCurrencyIso = value
    }
}
// SetDisplayName sets the displayName property value. The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
func (m *BookingBusiness) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmail sets the email property value. The email address for the business.
func (m *BookingBusiness) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetIsPublished sets the isPublished property value. The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
func (m *BookingBusiness) SetIsPublished(value *bool)() {
    if m != nil {
        m.isPublished = value
    }
}
// SetPhone sets the phone property value. The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
func (m *BookingBusiness) SetPhone(value *string)() {
    if m != nil {
        m.phone = value
    }
}
// SetPublicUrl sets the publicUrl property value. The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
func (m *BookingBusiness) SetPublicUrl(value *string)() {
    if m != nil {
        m.publicUrl = value
    }
}
// SetSchedulingPolicy sets the schedulingPolicy property value. Specifies how bookings can be created for this business.
func (m *BookingBusiness) SetSchedulingPolicy(value *BookingSchedulingPolicy)() {
    if m != nil {
        m.schedulingPolicy = value
    }
}
// SetServices sets the services property value. All the services offered by this business. Read-only. Nullable.
func (m *BookingBusiness) SetServices(value []BookingService)() {
    if m != nil {
        m.services = value
    }
}
// SetStaffMembers sets the staffMembers property value. All the staff members that provide services in this business. Read-only. Nullable.
func (m *BookingBusiness) SetStaffMembers(value []BookingStaffMemberBase)() {
    if m != nil {
        m.staffMembers = value
    }
}
// SetWebSiteUrl sets the webSiteUrl property value. The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
func (m *BookingBusiness) SetWebSiteUrl(value *string)() {
    if m != nil {
        m.webSiteUrl = value
    }
}
