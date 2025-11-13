package tripleseat

import "time"

type BookingResp struct {
	Booking Booking `json:"booking,omitempty"`
}
type Location struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	CustomerID int    `json:"customer_id,omitempty"`
	SiteID     int    `json:"site_id,omitempty"`
}
type EmailAddresses struct {
	ID      int    `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
}
type PhoneNumbers struct {
	ID              int    `json:"id,omitempty"`
	Number          string `json:"number,omitempty"`
	PhoneNumberType string `json:"phone_number_type,omitempty"`
	Extension       string `json:"extension,omitempty"`
}
type Contact struct {
	ID             int              `json:"id,omitempty"`
	FirstName      string           `json:"first_name,omitempty"`
	LastName       string           `json:"last_name,omitempty"`
	AccountID      int              `json:"account_id,omitempty"`
	EmailAddresses []EmailAddresses `json:"email_addresses,omitempty"`
	PhoneNumbers   []PhoneNumbers   `json:"phone_numbers,omitempty"`
}
type Account struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type Owner struct {
	ID           int            `json:"id,omitempty"`
	FirstName    string         `json:"first_name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	Title        string         `json:"title,omitempty"`
	Email        string         `json:"email,omitempty"`
	CreatedAt    string         `json:"created_at,omitempty"`
	UpdatedAt    string         `json:"updated_at,omitempty"`
	LoginCount   int            `json:"login_count,omitempty"`
	LoginAt      string         `json:"login_at,omitempty"`
	CustomerID   int            `json:"customer_id,omitempty"`
	PhoneNumbers []PhoneNumbers `json:"phone_numbers,omitempty"`
}
type Creator struct {
	ID           int            `json:"id,omitempty"`
	FirstName    string         `json:"first_name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	Title        string         `json:"title,omitempty"`
	Email        string         `json:"email,omitempty"`
	CreatedAt    string         `json:"created_at,omitempty"`
	UpdatedAt    string         `json:"updated_at,omitempty"`
	LoginCount   int            `json:"login_count,omitempty"`
	LoginAt      string         `json:"login_at,omitempty"`
	CustomerID   int            `json:"customer_id,omitempty"`
	PhoneNumbers []PhoneNumbers `json:"phone_numbers,omitempty"`
}
type Updator struct {
	ID           int            `json:"id,omitempty"`
	FirstName    string         `json:"first_name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	Title        string         `json:"title,omitempty"`
	Email        string         `json:"email,omitempty"`
	CreatedAt    string         `json:"created_at,omitempty"`
	UpdatedAt    string         `json:"updated_at,omitempty"`
	LoginCount   int            `json:"login_count,omitempty"`
	LoginAt      string         `json:"login_at,omitempty"`
	CustomerID   int            `json:"customer_id,omitempty"`
	PhoneNumbers []PhoneNumbers `json:"phone_numbers,omitempty"`
}
type StatusChanges struct {
	Status         string `json:"status,omitempty"`
	PreviousStatus any    `json:"previous_status,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	CreatedBy      int    `json:"created_by,omitempty"`
}
type OffsiteAddress struct {
	ID          int    `json:"id,omitempty"`
	Address1    string `json:"address1,omitempty"`
	Address2    string `json:"address2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Country     any    `json:"country,omitempty"`
	ZipCode     string `json:"zip_code,omitempty"`
	AddressType string `json:"address_type,omitempty"`
}
type Rooms struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Capacity     int    `json:"capacity,omitempty"`
	Description  string `json:"description,omitempty"`
	SiteID       int    `json:"site_id,omitempty"`
	LocationID   int    `json:"location_id,omitempty"`
	IsUnassigned bool   `json:"is_unassigned,omitempty"`
}
type Views struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Documents struct {
	ID                 int     `json:"id,omitempty"`
	Title              string  `json:"title,omitempty"`
	TemplateDocumentID int     `json:"template_document_id,omitempty"`
	DeletedAt          any     `json:"deleted_at,omitempty"`
	Views              []Views `json:"views,omitempty"`
}
type CustomFields struct {
	ID                  int    `json:"id,omitempty"`
	CustomFieldName     string `json:"custom_field_name,omitempty"`
	CustomFieldID       int    `json:"custom_field_id,omitempty"`
	CustomFieldRequired bool   `json:"custom_field_required,omitempty"`
	CustomFieldSlug     string `json:"custom_field_slug,omitempty"`
	Value               string `json:"value,omitempty"`
}
type CategoryTotals struct {
	Name  string      `json:"name,omitempty"`
	Total StringFloat `json:"total,omitempty"`
}
type BillingTotals struct {
	Name  string      `json:"name,omitempty"`
	Total StringFloat `json:"total,omitempty"`
}
type SelectedLeadSources struct {
	LeadSourceID    any `json:"lead_source_id,omitempty"`
	LeadSourceOther any `json:"lead_source_other,omitempty"`
	LeadSourceName  any `json:"lead_source_name,omitempty"`
}
type Events struct {
	ID                           int                   `json:"id,omitempty"`
	Name                         string                `json:"name,omitempty"`
	EventDate                    string                `json:"event_date,omitempty"`
	EventDateIso8601             string                `json:"event_date_iso8601,omitempty"`
	Status                       string                `json:"status,omitempty"`
	EventStart                   string                `json:"event_start,omitempty"`
	EventEnd                     string                `json:"event_end,omitempty"`
	EventStartUtc                time.Time             `json:"event_start_utc,omitempty"`
	EventEndUtc                  time.Time             `json:"event_end_utc,omitempty"`
	EventStartIso8601            string                `json:"event_start_iso8601,omitempty"`
	EventEndIso8601              string                `json:"event_end_iso8601,omitempty"`
	EventStartTime               string                `json:"event_start_time,omitempty"`
	EventEndTime                 string                `json:"event_end_time,omitempty"`
	StartDate                    string                `json:"start_date,omitempty"`
	EndDate                      string                `json:"end_date,omitempty"`
	SetupTime                    int                   `json:"setup_time,omitempty"`
	TeardownTime                 int                   `json:"teardown_time,omitempty"`
	EventStartTimeWithSetupTime  string                `json:"event_start_time_with_setup_time,omitempty"`
	EventEndTimeWithTeardownTime string                `json:"event_end_time_with_teardown_time,omitempty"`
	EventStartWithSetupIso8601   string                `json:"event_start_with_setup_iso8601,omitempty"`
	EventEndWithTeardownIso8601  string                `json:"event_end_with_teardown_iso8601,omitempty"`
	EventTimezone                string                `json:"event_timezone,omitempty"`
	EventStyle                   string                `json:"event_style,omitempty"`
	GuaranteedGuestCount         any                   `json:"guaranteed_guest_count,omitempty"`
	GuestCount                   int                   `json:"guest_count,omitempty"`
	FoodAndBeverageMin           any                   `json:"food_and_beverage_min,omitempty"`
	PricePerPerson               any                   `json:"price_per_person,omitempty"`
	DepositAmount                StringFloat           `json:"deposit_amount,omitempty"`
	RentalFee                    StringFloat           `json:"rental_fee,omitempty"`
	ActualAmount                 StringFloat           `json:"actual_amount,omitempty"`
	GrandTotal                   StringFloat           `json:"grand_total,omitempty"`
	AmountDue                    StringFloat           `json:"amount_due,omitempty"`
	Description                  string                `json:"description,omitempty"`
	ContactID                    int                   `json:"contact_id,omitempty"`
	AccountID                    int                   `json:"account_id,omitempty"`
	OwnedBy                      int                   `json:"owned_by,omitempty"`
	CreatedAt                    string                `json:"created_at,omitempty"`
	UpdatedAt                    string                `json:"updated_at,omitempty"`
	DeletedAt                    any                   `json:"deleted_at,omitempty"`
	BookingID                    int                   `json:"booking_id,omitempty"`
	UpdatedBy                    int                   `json:"updated_by,omitempty"`
	CreatedBy                    int                   `json:"created_by,omitempty"`
	CustomerID                   int                   `json:"customer_id,omitempty"`
	SiteID                       int                   `json:"site_id,omitempty"`
	LocationID                   int                   `json:"location_id,omitempty"`
	RoomIds                      []int                 `json:"room_ids,omitempty"`
	ManagingUserIds              []any                 `json:"managing_user_ids,omitempty"`
	Unassigned                   bool                  `json:"unassigned,omitempty"`
	EventTypeID                  any                   `json:"event_type_id,omitempty"`
	PostAs                       string                `json:"post_as,omitempty"`
	OffsiteAddress               OffsiteAddress        `json:"offsite_address,omitempty"`
	DocumentIds                  string                `json:"document_ids,omitempty"`
	StartTime                    string                `json:"start_time,omitempty"`
	EndTime                      string                `json:"end_time,omitempty"`
	EventType                    any                   `json:"event_type,omitempty"`
	Rooms                        []Rooms               `json:"rooms,omitempty"`
	Location                     Location              `json:"location,omitempty"`
	Documents                    []Documents           `json:"documents,omitempty"`
	CustomFields                 []CustomFields        `json:"custom_fields,omitempty"`
	CategoryTotals               []CategoryTotals      `json:"category_totals,omitempty"`
	BillingTotals                []BillingTotals       `json:"billing_totals,omitempty"`
	Owner                        Owner                 `json:"owner,omitempty"`
	Creator                      Creator               `json:"creator,omitempty"`
	Updator                      Updator               `json:"updator,omitempty"`
	Managers                     []any                 `json:"managers,omitempty"`
	SelectedLeadSources          []SelectedLeadSources `json:"selected_lead_sources,omitempty"`
	Contact                      Contact               `json:"contact,omitempty"`
	SecondaryContacts            []any                 `json:"secondary_contacts,omitempty"`
	Account                      Account               `json:"account,omitempty"`
	StatusChanges                []StatusChanges       `json:"status_changes,omitempty"`
	Attachments                  []any                 `json:"attachments,omitempty"`
}
type Booking struct {
	ID                     int             `json:"id,omitempty"`
	Name                   string          `json:"name,omitempty"`
	CustomerID             int             `json:"customer_id,omitempty"`
	SiteID                 int             `json:"site_id,omitempty"`
	LocationID             int             `json:"location_id,omitempty"`
	EventIds               []int           `json:"event_ids,omitempty"`
	Status                 string          `json:"status,omitempty"`
	TotalEventActualAmount StringFloat     `json:"total_event_actual_amount,omitempty"`
	TotalActualAmount      StringFloat     `json:"total_actual_amount,omitempty"`
	TotalEventGrandTotal   StringFloat     `json:"total_event_grand_total,omitempty"`
	TotalGrandTotal        StringFloat     `json:"total_grand_total,omitempty"`
	PostAs                 string          `json:"post_as,omitempty"`
	UpdatedBy              int             `json:"updated_by,omitempty"`
	CreatedBy              int             `json:"created_by,omitempty"`
	OwnedBy                int             `json:"owned_by,omitempty"`
	StartDate              string          `json:"start_date,omitempty"`
	EndDate                string          `json:"end_date,omitempty"`
	DefiniteDate           string          `json:"definite_date,omitempty"`
	TentativeDate          any             `json:"tentative_date,omitempty"`
	LostDate               any             `json:"lost_date,omitempty"`
	CreatedAt              string          `json:"created_at,omitempty"`
	UpdatedAt              string          `json:"updated_at,omitempty"`
	DeletedAt              any             `json:"deleted_at,omitempty"`
	Location               Location        `json:"location,omitempty"`
	MarketSegment          any             `json:"market_segment,omitempty"`
	Contact                Contact         `json:"contact,omitempty"`
	SecondaryContacts      []any           `json:"secondary_contacts,omitempty"`
	Account                Account         `json:"account,omitempty"`
	Owner                  Owner           `json:"owner,omitempty"`
	Creator                Creator         `json:"creator,omitempty"`
	Updator                Updator         `json:"updator,omitempty"`
	StatusChanges          []StatusChanges `json:"status_changes,omitempty"`
	SelectedLeadSources    []any           `json:"selected_lead_sources,omitempty"`
	CustomFields           []any           `json:"custom_fields,omitempty"`
	Documents              []any           `json:"documents,omitempty"`
	Events                 []Events        `json:"events,omitempty"`
}

type EventResp struct {
	Event Event `json:"event,omitempty"`
}
type Event struct {
	ID                           int                   `json:"id,omitempty"`
	Name                         string                `json:"name,omitempty"`
	EventDate                    string                `json:"event_date,omitempty"`
	EventDateIso8601             string                `json:"event_date_iso8601,omitempty"`
	Status                       string                `json:"status,omitempty"`
	EventStart                   string                `json:"event_start,omitempty"`
	EventEnd                     string                `json:"event_end,omitempty"`
	EventStartUtc                time.Time             `json:"event_start_utc,omitempty"`
	EventEndUtc                  time.Time             `json:"event_end_utc,omitempty"`
	EventStartIso8601            string                `json:"event_start_iso8601,omitempty"`
	EventEndIso8601              string                `json:"event_end_iso8601,omitempty"`
	EventStartTime               string                `json:"event_start_time,omitempty"`
	EventEndTime                 string                `json:"event_end_time,omitempty"`
	StartDate                    string                `json:"start_date,omitempty"`
	EndDate                      string                `json:"end_date,omitempty"`
	SetupTime                    int                   `json:"setup_time,omitempty"`
	TeardownTime                 int                   `json:"teardown_time,omitempty"`
	EventStartTimeWithSetupTime  string                `json:"event_start_time_with_setup_time,omitempty"`
	EventEndTimeWithTeardownTime string                `json:"event_end_time_with_teardown_time,omitempty"`
	EventStartWithSetupIso8601   string                `json:"event_start_with_setup_iso8601,omitempty"`
	EventEndWithTeardownIso8601  string                `json:"event_end_with_teardown_iso8601,omitempty"`
	EventTimezone                string                `json:"event_timezone,omitempty"`
	EventStyle                   string                `json:"event_style,omitempty"`
	GuaranteedGuestCount         any                   `json:"guaranteed_guest_count,omitempty"`
	GuestCount                   int                   `json:"guest_count,omitempty"`
	FoodAndBeverageMin           any                   `json:"food_and_beverage_min,omitempty"`
	PricePerPerson               any                   `json:"price_per_person,omitempty"`
	DepositAmount                StringFloat           `json:"deposit_amount,omitempty"`
	RentalFee                    string                `json:"rental_fee,omitempty"`
	ActualAmount                 StringFloat           `json:"actual_amount,omitempty"`
	GrandTotal                   StringFloat           `json:"grand_total,omitempty"`
	AmountDue                    StringFloat           `json:"amount_due,omitempty"`
	Description                  string                `json:"description,omitempty"`
	ContactID                    int                   `json:"contact_id,omitempty"`
	AccountID                    int                   `json:"account_id,omitempty"`
	OwnedBy                      int                   `json:"owned_by,omitempty"`
	CreatedAt                    string                `json:"created_at,omitempty"`
	UpdatedAt                    string                `json:"updated_at,omitempty"`
	DeletedAt                    any                   `json:"deleted_at,omitempty"`
	BookingID                    int                   `json:"booking_id,omitempty"`
	UpdatedBy                    int                   `json:"updated_by,omitempty"`
	CreatedBy                    int                   `json:"created_by,omitempty"`
	CustomerID                   int                   `json:"customer_id,omitempty"`
	SiteID                       int                   `json:"site_id,omitempty"`
	LocationID                   int                   `json:"location_id,omitempty"`
	RoomIds                      []int                 `json:"room_ids,omitempty"`
	ManagingUserIds              []any                 `json:"managing_user_ids,omitempty"`
	Unassigned                   bool                  `json:"unassigned,omitempty"`
	EventTypeID                  any                   `json:"event_type_id,omitempty"`
	PostAs                       string                `json:"post_as,omitempty"`
	OffsiteAddress               OffsiteAddress        `json:"offsite_address,omitempty"`
	DocumentIds                  string                `json:"document_ids,omitempty"`
	StartTime                    string                `json:"start_time,omitempty"`
	EndTime                      string                `json:"end_time,omitempty"`
	EventType                    any                   `json:"event_type,omitempty"`
	Rooms                        []Rooms               `json:"rooms,omitempty"`
	Location                     Location              `json:"location,omitempty"`
	Booking                      Booking               `json:"booking,omitempty"`
	Documents                    []Documents           `json:"documents,omitempty"`
	CustomFields                 []CustomFields        `json:"custom_fields,omitempty"`
	CategoryTotals               []CategoryTotals      `json:"category_totals,omitempty"`
	BillingTotals                []BillingTotals       `json:"billing_totals,omitempty"`
	Owner                        Owner                 `json:"owner,omitempty"`
	Creator                      Creator               `json:"creator,omitempty"`
	Updator                      Updator               `json:"updator,omitempty"`
	Managers                     []any                 `json:"managers,omitempty"`
	SelectedLeadSources          []SelectedLeadSources `json:"selected_lead_sources,omitempty"`
	Contact                      Contact               `json:"contact,omitempty"`
	SecondaryContacts            []any                 `json:"secondary_contacts,omitempty"`
	Account                      Account               `json:"account,omitempty"`
	StatusChanges                []StatusChanges       `json:"status_changes,omitempty"`
	Attachments                  []any                 `json:"attachments,omitempty"`
}

type GenericAllResp[T any] struct {
	Success bool `json:"success,omitempty"`
	Data    []T  `json:"data,omitempty"`
}

type PaymentsAllResp = GenericAllResp[DataPayments]
type RefundsAllResp = GenericAllResp[DataRefunds]

type DataPayments struct {
	Rowid               string      `json:"rowid,omitempty"`
	PaymentMethod       string      `json:"payment_method,omitempty"`
	PayableType         string      `json:"payable_type,omitempty"`
	PayableID           int         `json:"payable_id,omitempty"`
	Payable             string      `json:"payable,omitempty"`
	PayableDate         string      `json:"payable_date,omitempty"`
	PayableStatus       string      `json:"payable_status,omitempty"`
	LocationID          int         `json:"location_id,omitempty"`
	Location            string      `json:"location,omitempty"`
	PaymentID           int         `json:"payment_id,omitempty"`
	PaidAt              Date        `json:"paid_at,omitempty"`
	CardType            string      `json:"card_type,omitempty"`
	CardLast4           string      `json:"card_last4,omitempty"`
	CardHolderName      string      `json:"card_holder_name,omitempty"`
	Amount              StringFloat `json:"amount,omitempty"`
	TransactionID       string      `json:"transaction_id,omitempty"`
	CardID              string      `json:"card_id,omitempty"`
	AuthCode            string      `json:"auth_code,omitempty"`
	Reference           string      `json:"reference,omitempty"`
	FoodSubtotal        StringFloat `json:"food_subtotal,omitempty"`
	BeverageSubtotal    StringFloat `json:"beverage_subtotal,omitempty"`
	AudioVisualSubtotal any         `json:"audio_visual_subtotal,omitempty"`
	LaborSubtotal       any         `json:"labor_subtotal,omitempty"`
	MiscSubtotal        any         `json:"misc_subtotal,omitempty"`
}

type DataRefunds struct {
	Rowid                 string      `json:"rowid,omitempty"`
	PaymentMethod         string      `json:"payment_method,omitempty"`
	PayableType           string      `json:"payable_type,omitempty"`
	PayableID             int         `json:"payable_id,omitempty"`
	Payable               string      `json:"payable,omitempty"`
	PayableDate           string      `json:"payable_date,omitempty"`
	PayableStatus         string      `json:"payable_status,omitempty"`
	LocationID            int         `json:"location_id,omitempty"`
	Location              string      `json:"location,omitempty"`
	PaymentID             int         `json:"payment_id,omitempty"`
	PaidAt                string      `json:"paid_at,omitempty"`
	CardType              string      `json:"card_type,omitempty"`
	CardLast4             any         `json:"card_last4,omitempty"`
	CardHolderName        string      `json:"card_holder_name,omitempty"`
	OriginalAmount        StringFloat `json:"original_amount,omitempty"`
	OriginalTransactionID string      `json:"original_transaction_id,omitempty"`
	CardID                string      `json:"card_id,omitempty"`
	AuthCode              string      `json:"auth_code,omitempty"`
	RefundedAt            string      `json:"refunded_at,omitempty"`
	RefundAmount          StringFloat `json:"refund_amount,omitempty"`
	RefundReason          string      `json:"refund_reason,omitempty"`
	UserID                string      `json:"user_id,omitempty"`
	Reference             string      `json:"reference,omitempty"`
}
