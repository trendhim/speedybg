package speedybg

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type PickingValidationErrorType string

const (
	PickingValidationErrorTypeCOMMON_ERROR PickingValidationErrorType = "COMMON_ERROR"

	PickingValidationErrorTypeAMOUNT_INSURANCE_OUT_OF_RANGE PickingValidationErrorType = "AMOUNT_INSURANCE_OUT_OF_RANGE"

	PickingValidationErrorTypeINSURANCE_NOT_ALLOWED_FOR_SERVICE PickingValidationErrorType = "INSURANCE_NOT_ALLOWED_FOR_SERVICE"

	PickingValidationErrorTypeINSURANCE_REQUIRED_FOR_SERVICE PickingValidationErrorType = "INSURANCE_REQUIRED_FOR_SERVICE"

	PickingValidationErrorTypeAMOUNT_CASH_ON_DELIVERY_OUT_OF_RANGE PickingValidationErrorType = "AMOUNT_CASH_ON_DELIVERY_OUT_OF_RANGE"

	PickingValidationErrorTypeAMOUNT_CASH_ON_DELIVERY_NOT_ALLOWED_FOR_SERVICE PickingValidationErrorType = "AMOUNT_CASH_ON_DELIVERY_NOT_ALLOWED_FOR_SERVICE"

	PickingValidationErrorTypeAMOUNT_CASH_ON_DELIVERY_REQUIRED_FOR_SERVICE PickingValidationErrorType = "AMOUNT_CASH_ON_DELIVERY_REQUIRED_FOR_SERVICE"

	PickingValidationErrorTypeAMOUNT_MONEY_TRANSFER_REQ_OUT_OF_RANGE PickingValidationErrorType = "AMOUNT_MONEY_TRANSFER_REQ_OUT_OF_RANGE"

	PickingValidationErrorTypeMONEY_TRANSFER_NOT_ALLOWED PickingValidationErrorType = "MONEY_TRANSFER_NOT_ALLOWED"

	PickingValidationErrorTypeSENDER_HAS_NO_ANNEX_FOR_CASH_ON_DELIVERY PickingValidationErrorType = "SENDER_HAS_NO_ANNEX_FOR_CASH_ON_DELIVERY"

	PickingValidationErrorTypeINSURANCE_CANNOT_BE_LESS_THAN_CASH_ON_DELIVERY PickingValidationErrorType = "INSURANCE_CANNOT_BE_LESS_THAN_CASH_ON_DELIVERY"

	PickingValidationErrorTypeWEIGHT_NOT_IN_RANGE PickingValidationErrorType = "WEIGHT_NOT_IN_RANGE"

	PickingValidationErrorTypeDOCUMENTS_NOT_ALLOWED_FOR_THIS_SERVICE_AND_SITES PickingValidationErrorType = "DOCUMENTS_NOT_ALLOWED_FOR_THIS_SERVICE_AND_SITES"

	PickingValidationErrorTypeDOCUMENT_PICKINGS_CAN_CONSIST_OF_ONLY_ONE_PARCEL PickingValidationErrorType = "DOCUMENT_PICKINGS_CAN_CONSIST_OF_ONLY_ONE_PARCEL"

	PickingValidationErrorTypeDOCUMENT_PICKINGS_ARE_NOT_ALLOWED_TO_HAVE_INSURANCE PickingValidationErrorType = "DOCUMENT_PICKINGS_ARE_NOT_ALLOWED_TO_HAVE_INSURANCE"

	PickingValidationErrorTypeDOCUMENT_PICKINGS_CANNOT_BE_PALLETIZED PickingValidationErrorType = "DOCUMENT_PICKINGS_CANNOT_BE_PALLETIZED"

	PickingValidationErrorTypeFIXED_TIME_DELIVERY_NOT_ALLOWED_FOR_RECEIVER_SITE PickingValidationErrorType = "FIXED_TIME_DELIVERY_NOT_ALLOWED_FOR_RECEIVER_SITE"

	PickingValidationErrorTypeFIXED_TIME_DELIVERY_NOT_ALLOWED_FOR_SERVICE PickingValidationErrorType = "FIXED_TIME_DELIVERY_NOT_ALLOWED_FOR_SERVICE"

	PickingValidationErrorTypeINVALID_FIXED_TIME_FOR_DELIVERY PickingValidationErrorType = "INVALID_FIXED_TIME_FOR_DELIVERY"

	PickingValidationErrorTypeNO_MORE_THAN_3_PHONE_NUMBERS_EXPECTED PickingValidationErrorType = "NO_MORE_THAN_3_PHONE_NUMBERS_EXPECTED"

	PickingValidationErrorTypeINVALID_PHONE_NUMBER PickingValidationErrorType = "INVALID_PHONE_NUMBER"

	PickingValidationErrorTypeLOGGED_CLIENT_OBJECT_MUST_BE_PAYER_OR_SENDER PickingValidationErrorType = "LOGGED_CLIENT_OBJECT_MUST_BE_PAYER_OR_SENDER"

	PickingValidationErrorTypeINVALID_PARCELS_INFO PickingValidationErrorType = "INVALID_PARCELS_INFO"

	PickingValidationErrorTypeINVALID_TAKING_DATE PickingValidationErrorType = "INVALID_TAKING_DATE"

	PickingValidationErrorTypeINVALID_RET_TO_CLIENT_ID PickingValidationErrorType = "INVALID_RET_TO_CLIENT_ID"

	PickingValidationErrorTypeREQUIRED_BRINGING_TO_OFFICE PickingValidationErrorType = "REQUIRED_BRINGING_TO_OFFICE"

	PickingValidationErrorTypeINVALID_PALLET_SIZE PickingValidationErrorType = "INVALID_PALLET_SIZE"

	PickingValidationErrorTypeINVALID_PALLET_WEIGHT PickingValidationErrorType = "INVALID_PALLET_WEIGHT"

	PickingValidationErrorTypeINVALID_PARCEL_ID PickingValidationErrorType = "INVALID_PARCEL_ID"

	PickingValidationErrorTypeINVALID_PARCEL_SIZE PickingValidationErrorType = "INVALID_PARCEL_SIZE"

	PickingValidationErrorTypeINVALID_PARCEL_WEIGHT PickingValidationErrorType = "INVALID_PARCEL_WEIGHT"

	PickingValidationErrorTypeINVALID_EMAIL PickingValidationErrorType = "INVALID_EMAIL"

	PickingValidationErrorTypeINVALID_RETURN_SHIPMENT_REQUEST PickingValidationErrorType = "INVALID_RETURN_SHIPMENT_REQUEST"

	PickingValidationErrorTypeINVALID_RETURN_SERVICE_REQUEST PickingValidationErrorType = "INVALID_RETURN_SERVICE_REQUEST"

	PickingValidationErrorTypeINVALID_BACK_DOCUMENT_REQUEST PickingValidationErrorType = "INVALID_BACK_DOCUMENT_REQUEST"

	PickingValidationErrorTypeINVALID_BACK_RECEIPT_REQUEST PickingValidationErrorType = "INVALID_BACK_RECEIPT_REQUEST"

	PickingValidationErrorTypeINVALID_MONEY_TRANSFER_REQUEST PickingValidationErrorType = "INVALID_MONEY_TRANSFER_REQUEST"

	PickingValidationErrorTypeSERVICE_NOT_ALLOWED_FOR_OFFICE_BTO PickingValidationErrorType = "SERVICE_NOT_ALLOWED_FOR_OFFICE_BTO"

	PickingValidationErrorTypeSERVICE_NOT_ALLOWED_FOR_OFFICE_TBC PickingValidationErrorType = "SERVICE_NOT_ALLOWED_FOR_OFFICE_TBC"

	PickingValidationErrorTypeDELIVERY_FROM_ADDRESS_TO_ADDRESS_NOT_ALLOWED PickingValidationErrorType = "DELIVERY_FROM_ADDRESS_TO_ADDRESS_NOT_ALLOWED"

	PickingValidationErrorTypeDELIVERY_FROM_ADDRESS_TO_OFFICE_NOT_ALLOWED PickingValidationErrorType = "DELIVERY_FROM_ADDRESS_TO_OFFICE_NOT_ALLOWED"

	PickingValidationErrorTypeDELIVERY_FROM_OFFICE_TO_ADDRESS_NOT_ALLOWED PickingValidationErrorType = "DELIVERY_FROM_OFFICE_TO_ADDRESS_NOT_ALLOWED"

	PickingValidationErrorTypeDELIVERY_FROM_OFFICE_TO_OFFICE_NOT_ALLOWED PickingValidationErrorType = "DELIVERY_FROM_OFFICE_TO_OFFICE_NOT_ALLOWED"

	PickingValidationErrorTypeNON_WORKING_DAY_FOR_SELECTED_OFFICE_TBC PickingValidationErrorType = "NON_WORKING_DAY_FOR_SELECTED_OFFICE_TBC"

	PickingValidationErrorTypeNON_WORKING_DAY_FOR_SELECTED_OFFICE_BTO PickingValidationErrorType = "NON_WORKING_DAY_FOR_SELECTED_OFFICE_BTO"

	PickingValidationErrorTypeNON_ACTIVE_OFFICE_TBC PickingValidationErrorType = "NON_ACTIVE_OFFICE_TBC"

	PickingValidationErrorTypeNON_ACTIVE_OFFICE_BTO PickingValidationErrorType = "NON_ACTIVE_OFFICE_BTO"

	PickingValidationErrorTypeFOREIGN_PARCEL_NUMBER_ERROR PickingValidationErrorType = "FOREIGN_PARCEL_NUMBER_ERROR"

	PickingValidationErrorTypeMISSING_REQUIRED_VALUE_FIXED_TIME_FOR_DELIVERY PickingValidationErrorType = "MISSING_REQUIRED_VALUE_FIXED_TIME_FOR_DELIVERY"

	PickingValidationErrorTypeMISSING_REQUIRED_VALUE_CONTENTS PickingValidationErrorType = "MISSING_REQUIRED_VALUE_CONTENTS"

	PickingValidationErrorTypeMISSING_REQUIRED_VALUE_PACKING PickingValidationErrorType = "MISSING_REQUIRED_VALUE_PACKING"

	PickingValidationErrorTypeMISSING_REQUIRED_VALUE_PARTNER_NAME PickingValidationErrorType = "MISSING_REQUIRED_VALUE_PARTNER_NAME"

	PickingValidationErrorTypeMISSING_REQUIRED_VALUE_ADDRESS PickingValidationErrorType = "MISSING_REQUIRED_VALUE_ADDRESS"

	PickingValidationErrorTypeMISSING_REQUIRED_VALUE_PHONE_SENDER PickingValidationErrorType = "MISSING_REQUIRED_VALUE_PHONE_SENDER"

	PickingValidationErrorTypeMISSING_REQUIRED_VALUE_PHONE_RECEIVER PickingValidationErrorType = "MISSING_REQUIRED_VALUE_PHONE_RECEIVER"

	PickingValidationErrorTypeMISSING_REQUIRED_VALUE_BOL PickingValidationErrorType = "MISSING_REQUIRED_VALUE_BOL"

	PickingValidationErrorTypeMISSING_REQUIRED_VALUE_WEIGHT_DECLARED PickingValidationErrorType = "MISSING_REQUIRED_VALUE_WEIGHT_DECLARED"

	PickingValidationErrorTypeMISSING_REQUIRED_VALUE_PARCELS PickingValidationErrorType = "MISSING_REQUIRED_VALUE_PARCELS"

	PickingValidationErrorTypeINVALID_VALUE_PARTNER_NAME PickingValidationErrorType = "INVALID_VALUE_PARTNER_NAME"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_PARCEL_COUNT PickingValidationErrorType = "VALUE_OUT_OF_RANGE_PARCEL_COUNT"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_SIZE PickingValidationErrorType = "VALUE_OUT_OF_RANGE_SIZE"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_PACK_ID PickingValidationErrorType = "VALUE_OUT_OF_RANGE_PACK_ID"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_CONTENTS PickingValidationErrorType = "VALUE_OUT_OF_RANGE_CONTENTS"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_PACKING PickingValidationErrorType = "VALUE_OUT_OF_RANGE_PACKING"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_NOTE_CLIENT PickingValidationErrorType = "VALUE_OUT_OF_RANGE_NOTE_CLIENT"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_CLIENT_FIELD PickingValidationErrorType = "VALUE_OUT_OF_RANGE_CLIENT_FIELD"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_ADDRESS_FIELD PickingValidationErrorType = "VALUE_OUT_OF_RANGE_ADDRESS_FIELD"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_REF1 PickingValidationErrorType = "VALUE_OUT_OF_RANGE_REF1"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_REF2 PickingValidationErrorType = "VALUE_OUT_OF_RANGE_REF2"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_CONSOLIDATION_REF PickingValidationErrorType = "VALUE_OUT_OF_RANGE_CONSOLIDATION_REF"

	PickingValidationErrorTypeVALUE_OUT_OF_RANGE_DEFERRED_DELIVERY_WORK_DAYS PickingValidationErrorType = "VALUE_OUT_OF_RANGE_DEFERRED_DELIVERY_WORK_DAYS"

	PickingValidationErrorTypeOFFICE_BTO_NOT_ALLOWED_FOR_SHIPMENT PickingValidationErrorType = "OFFICE_BTO_NOT_ALLOWED_FOR_SHIPMENT"

	PickingValidationErrorTypeOFFICE_TBC_NOT_ALLOWED_FOR_SHIPMENT PickingValidationErrorType = "OFFICE_TBC_NOT_ALLOWED_FOR_SHIPMENT"

	PickingValidationErrorTypeINVALID_WEIGHT_OR_SIZE PickingValidationErrorType = "INVALID_WEIGHT_OR_SIZE"

	PickingValidationErrorTypeNOT_ALLOWED_ADDITIONAL_TAKING_TIME PickingValidationErrorType = "NOT_ALLOWED_ADDITIONAL_TAKING_TIME"

	PickingValidationErrorTypeBOL_CREATION_LIMIT_VIOLATED PickingValidationErrorType = "BOL_CREATION_LIMIT_VIOLATED"

	PickingValidationErrorTypeBOL_CREATION_DAILY_LIMIT_VIOLATED PickingValidationErrorType = "BOL_CREATION_DAILY_LIMIT_VIOLATED"

	PickingValidationErrorTypeINVALID_DEADLINE_FOR_DELIVERY PickingValidationErrorType = "INVALID_DEADLINE_FOR_DELIVERY"

	PickingValidationErrorTypeINVALID_DECLARED_PARCELS_COUNT PickingValidationErrorType = "INVALID_DECLARED_PARCELS_COUNT"

	PickingValidationErrorTypeBOL_CREATION_TIME_INTERVAL_LIMIT_REACHED PickingValidationErrorType = "BOL_CREATION_TIME_INTERVAL_LIMIT_REACHED"

	PickingValidationErrorTypeBOL_CREATION_TIME_INTERVAL_SERVICE_LIMIT_REACHED PickingValidationErrorType = "BOL_CREATION_TIME_INTERVAL_SERVICE_LIMIT_REACHED"

	PickingValidationErrorTypeBOL_CREATION_DAILY_ADDRESS_LIMIT_REACHED PickingValidationErrorType = "BOL_CREATION_DAILY_ADDRESS_LIMIT_REACHED"

	PickingValidationErrorTypeINVALID_FLOOR_TO_DELIVER_TO PickingValidationErrorType = "INVALID_FLOOR_TO_DELIVER_TO"

	PickingValidationErrorTypeDELIVERY_TO_FLOOR_NOT_ALLOWED_FOR_SERVICE PickingValidationErrorType = "DELIVERY_TO_FLOOR_NOT_ALLOWED_FOR_SERVICE"

	PickingValidationErrorTypeDELIVERY_TO_FLOOR_NOT_ALLOWED_WITHOUT_ANNEX PickingValidationErrorType = "DELIVERY_TO_FLOOR_NOT_ALLOWED_WITHOUT_ANNEX"

	PickingValidationErrorTypeDELIVERY_TO_FLOOR_NOT_ALLOWED_WITHOUT_CONTRACT PickingValidationErrorType = "DELIVERY_TO_FLOOR_NOT_ALLOWED_WITHOUT_CONTRACT"

	PickingValidationErrorTypeINVALID_RETURN_VOUCHER_REQUEST PickingValidationErrorType = "INVALID_RETURN_VOUCHER_REQUEST"

	PickingValidationErrorTypeOPTIONS_BEFORE_PAYMENT_NOT_ALLOWED_FOR_APT PickingValidationErrorType = "OPTIONS_BEFORE_PAYMENT_NOT_ALLOWED_FOR_APT"

	PickingValidationErrorTypeBACK_DOCUMENTS_NOT_ALLOWED_FOR_APT PickingValidationErrorType = "BACK_DOCUMENTS_NOT_ALLOWED_FOR_APT"

	PickingValidationErrorTypeBACK_RECEIPT_NOT_ALLOWED_FOR_APT PickingValidationErrorType = "BACK_RECEIPT_NOT_ALLOWED_FOR_APT"

	PickingValidationErrorTypeBACK_PICKING_NOT_ALLOWED_FOR_APT PickingValidationErrorType = "BACK_PICKING_NOT_ALLOWED_FOR_APT"

	PickingValidationErrorTypeRETURN_SERVICES_NOT_ALLOWED_FOR_APT PickingValidationErrorType = "RETURN_SERVICES_NOT_ALLOWED_FOR_APT"

	PickingValidationErrorTypeSPECIAL_DELIVERY_NOT_ALLOWED_FOR_APT PickingValidationErrorType = "SPECIAL_DELIVERY_NOT_ALLOWED_FOR_APT"

	PickingValidationErrorTypeFIXED_TIME_DELIVERY_NOT_ALLOWED_FOR_APT PickingValidationErrorType = "FIXED_TIME_DELIVERY_NOT_ALLOWED_FOR_APT"

	PickingValidationErrorTypeINSURANCE_NOT_ALLOWED_FOR_APT PickingValidationErrorType = "INSURANCE_NOT_ALLOWED_FOR_APT"

	PickingValidationErrorTypeINVALID_PARCELS_COUNT_FOR_APT PickingValidationErrorType = "INVALID_PARCELS_COUNT_FOR_APT"

	PickingValidationErrorTypeINVALID_PARCELS_COUNT_FOR_PENDING_PARCELS_DESRIPTION_FOR_APT PickingValidationErrorType = "INVALID_PARCELS_COUNT_FOR_PENDING_PARCELS_DESRIPTION_FOR_APT"

	PickingValidationErrorTypeINVALID_SERVICE_TYPE_FOR_APT PickingValidationErrorType = "INVALID_SERVICE_TYPE_FOR_APT"

	PickingValidationErrorTypeOFFICE_OF_TYPE_APT_NOT_ALLOWED_FOR_BTO PickingValidationErrorType = "OFFICE_OF_TYPE_APT_NOT_ALLOWED_FOR_BTO"

	PickingValidationErrorTypeINVALID_RECEIVER_MOBILE_PHONE_NUMBER_FOR_APT_TBC PickingValidationErrorType = "INVALID_RECEIVER_MOBILE_PHONE_NUMBER_FOR_APT_TBC"

	PickingValidationErrorTypeINVALID_OPTIONS_BEFORE_PAYMENT PickingValidationErrorType = "INVALID_OPTIONS_BEFORE_PAYMENT"

	PickingValidationErrorTypeINVALID_OPTIONS_BEFORE_PAYMENT_DETAILS PickingValidationErrorType = "INVALID_OPTIONS_BEFORE_PAYMENT_DETAILS"

	PickingValidationErrorTypePARCEL_SIZE_MUST_BE_DEFINED_FOR_OFFICE_OF_TYPE_APT PickingValidationErrorType = "PARCEL_SIZE_MUST_BE_DEFINED_FOR_OFFICE_OF_TYPE_APT"

	PickingValidationErrorTypePARCEL_PREDEFINED_SIZE_DOES_NOT_EXIST PickingValidationErrorType = "PARCEL_PREDEFINED_SIZE_DOES_NOT_EXIST"

	PickingValidationErrorTypePARCEL_PREDEFINED_SIZE_NOT_ALLOWED PickingValidationErrorType = "PARCEL_PREDEFINED_SIZE_NOT_ALLOWED"

	PickingValidationErrorTypeCONTACT_NAME_REQUIRED PickingValidationErrorType = "CONTACT_NAME_REQUIRED"

	PickingValidationErrorTypeCONTACT_NAME_AND_PARTNER_NAME_CANNOT_BE_EQUALS PickingValidationErrorType = "CONTACT_NAME_AND_PARTNER_NAME_CANNOT_BE_EQUALS"

	PickingValidationErrorTypeCONTACT_NAME_MUST_BE_EMPTY PickingValidationErrorType = "CONTACT_NAME_MUST_BE_EMPTY"

	PickingValidationErrorTypeDECLARED_VALUE_NOT_ALLOWED PickingValidationErrorType = "DECLARED_VALUE_NOT_ALLOWED"
)

type ParamLanguage string

const (
	ParamLanguageBG ParamLanguage = "BG"

	ParamLanguageEN ParamLanguage = "EN"
)

type ComplementaryServiceAllowance string

const (
	ComplementaryServiceAllowanceBANNED ComplementaryServiceAllowance = "BANNED"

	ComplementaryServiceAllowanceALLOWED ComplementaryServiceAllowance = "ALLOWED"

	ComplementaryServiceAllowanceREQUIRED ComplementaryServiceAllowance = "REQUIRED"
)

type AddrNomen string

const (
	AddrNomenNO AddrNomen = "NO"

	AddrNomenFULL AddrNomen = "FULL"

	AddrNomenPARTIAL AddrNomen = "PARTIAL"
)

type ConvertToWin1251 struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ convertToWin1251"`

	SessionId string `xml:"sessionId,omitempty"`

	Text string `xml:"text,omitempty"`
}

type ConvertToWin1251Response struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ convertToWin1251Response"`

	Return_ string `xml:"return,omitempty"`
}

type InvalidSessionException struct {
	Message string `xml:"message,omitempty"`
}

type PickingValidationException struct {
	ErrorDescription string `xml:"errorDescription,omitempty"`

	ErrorId string `xml:"errorId,omitempty"`

	ErrorType *PickingValidationErrorType `xml:"errorType,omitempty"`

	Message string `xml:"message,omitempty"`
}

type ListOffices struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listOffices"`

	SessionId string `xml:"sessionId,omitempty"`

	Name string `xml:"name,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`
}

type ListOfficesResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listOfficesResponse"`

	Return_ []*ResultOffice `xml:"return,omitempty"`
}

type ResultOffice struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultOffice"`

	Address *ValueAddress `xml:"address,omitempty"`

	Id int64 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	WorkingTimeFrom int16 `xml:"workingTimeFrom,omitempty"`

	WorkingTimeHalfFrom int16 `xml:"workingTimeHalfFrom,omitempty"`

	WorkingTimeHalfTo int16 `xml:"workingTimeHalfTo,omitempty"`

	WorkingTimeTo int16 `xml:"workingTimeTo,omitempty"`
}

type ValueAddress struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ valueAddress"`

	AddressNote string `xml:"addressNote,omitempty"`

	ApartmentNo string `xml:"apartmentNo,omitempty"`

	BlockNo string `xml:"blockNo,omitempty"`

	CommonObjectId int64 `xml:"commonObjectId,omitempty"`

	CommonObjectName string `xml:"commonObjectName,omitempty"`

	CoordTypeId int32 `xml:"coordTypeId,omitempty"`

	CoordX float64 `xml:"coordX,omitempty"`

	CoordY float64 `xml:"coordY,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`

	Eknm string `xml:"eknm,omitempty"`

	EntranceNo string `xml:"entranceNo,omitempty"`

	FloorNo string `xml:"floorNo,omitempty"`

	FullNomenclature bool `xml:"fullNomenclature,omitempty"`

	MunicipalityName string `xml:"municipalityName,omitempty"`

	PostCode string `xml:"postCode,omitempty"`

	QuarterId int64 `xml:"quarterId,omitempty"`

	QuarterName string `xml:"quarterName,omitempty"`

	QuarterType string `xml:"quarterType,omitempty"`

	RegionName string `xml:"regionName,omitempty"`

	SiteDetails string `xml:"siteDetails,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	SiteName string `xml:"siteName,omitempty"`

	SiteType string `xml:"siteType,omitempty"`

	StateId string `xml:"stateId,omitempty"`

	StreetId int64 `xml:"streetId,omitempty"`

	StreetName string `xml:"streetName,omitempty"`

	StreetNo string `xml:"streetNo,omitempty"`

	StreetType string `xml:"streetType,omitempty"`
}

type CreatePDF struct {
	XMLName xml.Name `xml:"ns1:createPDF"`
	XMLNS   string   `xml:"xmlns:ns1,attr"`

	SessionId string `xml:"sessionId,omitempty"`

	Params *ParamPDF `xml:"params,omitempty"`
}

type ParamPDF struct {
	XMLName xml.Name

	AdditionalBarcodes []*ParamBarcodeInfo `xml:"additionalBarcodes,omitempty"`

	AdditionalBarcodesFormat string `xml:"additionalBarcodesFormat,omitempty"`

	AdditionalCopyForSender bool `xml:"additionalCopyForSender,omitempty"`

	Ids []int64 `xml:"ids,omitempty"`

	IncludeAutoPrintJS bool `xml:"includeAutoPrintJS,omitempty"`

	PrinterName string `xml:"printerName,omitempty"`

	Type_ int32 `xml:"type,omitempty"`
}

type ParamBarcodeInfo struct {
	XMLName xml.Name

	BarcodeLabel string `xml:"barcodeLabel,omitempty"`

	BarcodeValue string `xml:"barcodeValue,omitempty"`
}

type CreatePDFResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ createPDFResponse"`

	Return_ []byte `xml:"return,omitempty"`
}

type ListStreetTypes struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listStreetTypes"`

	SessionId string `xml:"sessionId,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ListStreetTypesResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listStreetTypesResponse"`

	Return_ []string `xml:"return,omitempty"`
}

type ListCountriesEx struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listCountriesEx"`

	SessionId string `xml:"sessionId,omitempty"`

	Filter *ParamFilterCountry `xml:"filter,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ParamFilterCountry struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ paramFilterCountry"`

	CountryId int64 `xml:"countryId,omitempty"`

	IsoAlpha2 string `xml:"isoAlpha2,omitempty"`

	IsoAlpha3 string `xml:"isoAlpha3,omitempty"`

	Name string `xml:"name,omitempty"`

	SearchString string `xml:"searchString,omitempty"`
}

type ListCountriesExResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listCountriesExResponse"`

	Return_ []*ResultCountry `xml:"return,omitempty"`
}

type ResultCountry struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultCountry"`

	CountryId int64 `xml:"countryId,omitempty"`

	Name string `xml:"name,omitempty"`

	IsoAlpha2 string `xml:"isoAlpha2,omitempty"`

	IsoAlpha3 string `xml:"isoAlpha3,omitempty"`

	RequirePostCode bool `xml:"requirePostCode,omitempty"`

	PostCodeFormat string `xml:"postCodeFormat,omitempty"`

	RequireState bool `xml:"requireState,omitempty"`

	SiteNomen int32 `xml:"siteNomen,omitempty"`

	ActiveCurrencyCode string `xml:"activeCurrencyCode,omitempty"`

	AddressTypeParams string `xml:"addressTypeParams,omitempty"`

	AddressTypeParamsCurrent int32 `xml:"addressTypeParamsCurrent,omitempty"`
}

type ListStates struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listStates"`

	SessionId string `xml:"sessionId,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`

	Name string `xml:"name,omitempty"`
}

type ListStatesResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listStatesResponse"`

	Return_ []*ResultState `xml:"return,omitempty"`
}

type ResultState struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultState"`

	CountryId int64 `xml:"countryId,omitempty"`

	Name string `xml:"name,omitempty"`

	StateAlpha string `xml:"stateAlpha,omitempty"`

	StateId string `xml:"stateId,omitempty"`
}

type GetAllowedDaysForTaking struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getAllowedDaysForTaking"`

	SessionId string `xml:"sessionId,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`

	SenderSiteId int64 `xml:"senderSiteId,omitempty"`

	SenderOfficeId int64 `xml:"senderOfficeId,omitempty"`

	MinDate time.Time `xml:"minDate,omitempty"`

	SenderCountryId int64 `xml:"senderCountryId,omitempty"`

	SenderPostCode string `xml:"senderPostCode,omitempty"`

	SenderId int64 `xml:"senderId,omitempty"`
}

type GetAllowedDaysForTakingResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getAllowedDaysForTakingResponse"`

	Return_ []time.Time `xml:"return,omitempty"`
}

type AddParcel struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ addParcel"`

	SessionId string `xml:"sessionId,omitempty"`

	Parcel *ParamParcel `xml:"parcel,omitempty"`
}

type ParamParcel struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ paramParcel"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`

	ForeignParcelNumber string `xml:"foreignParcelNumber,omitempty"`

	PackId int64 `xml:"packId,omitempty"`

	ParcelId int64 `xml:"parcelId,omitempty"`

	PredefinedSize string `xml:"predefinedSize,omitempty"`

	Ref1 string `xml:"ref1,omitempty"`

	Ref2 string `xml:"ref2,omitempty"`

	Size *Size `xml:"size,omitempty"`

	Weight float64 `xml:"weight,omitempty"`
}

type Size struct {
	XMLName xml.Name

	Depth int32 `xml:"depth,omitempty"`

	Height int32 `xml:"height,omitempty"`

	Width int32 `xml:"width,omitempty"`
}

type AddParcelResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ addParcelResponse"`

	Return_ int64 `xml:"return,omitempty"`
}

type GetPickingDeliveryInfo struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getPickingDeliveryInfo"`

	SessionId string `xml:"sessionId,omitempty"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type GetPickingDeliveryInfoResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getPickingDeliveryInfoResponse"`

	Return_ *ResultTrackPickingEx `xml:"return,omitempty"`
}

type ResultTrackPickingEx struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultTrackPickingEx"`

	Consignee string `xml:"consignee,omitempty"`

	Moment time.Time `xml:"moment,omitempty"`

	OperationCode int32 `xml:"operationCode,omitempty"`

	OperationComment string `xml:"operationComment,omitempty"`

	OperationDescription string `xml:"operationDescription,omitempty"`

	RedirectBillOfLading int64 `xml:"redirectBillOfLading,omitempty"`

	ReturnBillOfLading int64 `xml:"returnBillOfLading,omitempty"`

	SiteName string `xml:"siteName,omitempty"`

	SiteType string `xml:"siteType,omitempty"`

	SignatureImage string `xml:"signatureImage,omitempty"`

	Barcode int64 `xml:"barcode,omitempty"`

	ForeignParcelNumber string `xml:"foreignParcelNumber,omitempty"`

	ExceptionCodes []int32 `xml:"exceptionCodes,omitempty"`

	ForeignParcelNumbersList []string `xml:"foreignParcelNumbersList,omitempty"`

	InfoURL string `xml:"infoURL,omitempty"`

	AdditionalInfo string `xml:"additionalInfo,omitempty"`

	UnsuccessfulDeliveryStickerImageURL string `xml:"unsuccessfulDeliveryStickerImageURL,omitempty"`
}

type GetPickingExtendedInfo struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getPickingExtendedInfo"`

	SessionId string `xml:"sessionId,omitempty"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`
}

type GetPickingExtendedInfoResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getPickingExtendedInfoResponse"`

	Return_ *ResultPickingExtendedInfo `xml:"return,omitempty"`
}

type ResultPickingExtendedInfo struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultPickingExtendedInfo"`

	Amounts *ResultAmounts `xml:"amounts,omitempty"`

	BackDocumentsRequest bool `xml:"backDocumentsRequest,omitempty"`

	BackReceiptRequest bool `xml:"backReceiptRequest,omitempty"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`

	CodPayment *CodPayment `xml:"codPayment,omitempty"`

	Contents string `xml:"contents,omitempty"`

	DeadlineDelivery time.Time `xml:"deadlineDelivery,omitempty"`

	DeferredDeliveryWorkDays int32 `xml:"deferredDeliveryWorkDays,omitempty"`

	DeliveryInfo *ResultDeliveryInfo `xml:"deliveryInfo,omitempty"`

	DeliveryToFloorNo int32 `xml:"deliveryToFloorNo,omitempty"`

	DiscCalc *FixedDiscountCardId `xml:"discCalc,omitempty"`

	Documents bool `xml:"documents,omitempty"`

	FixedTimeDelivery int16 `xml:"fixedTimeDelivery,omitempty"`

	Fragile bool `xml:"fragile,omitempty"`

	NoteClient string `xml:"noteClient,omitempty"`

	OfficeToBeCalledId int64 `xml:"officeToBeCalledId,omitempty"`

	OptionsBeforePayment *ResultOptionsBeforePayment `xml:"optionsBeforePayment,omitempty"`

	Packing string `xml:"packing,omitempty"`

	Packings []*ResultPackings `xml:"packings,omitempty"`

	Palletized bool `xml:"palletized,omitempty"`

	PalletsListCalculation string `xml:"palletsListCalculation,omitempty"`

	PalletsListDeclared string `xml:"palletsListDeclared,omitempty"`

	PalletsListMeasured string `xml:"palletsListMeasured,omitempty"`

	Parcels []*ResultParcelInfoEx `xml:"parcels,omitempty"`

	ParcelsCount int32 `xml:"parcelsCount,omitempty"`

	PayCodToThirdParty bool `xml:"payCodToThirdParty,omitempty"`

	PayerRefId int64 `xml:"payerRefId,omitempty"`

	PayerRefInsuranceId int64 `xml:"payerRefInsuranceId,omitempty"`

	PayerRefPackingsId int64 `xml:"payerRefPackingsId,omitempty"`

	PayerType int32 `xml:"payerType,omitempty"`

	PayerTypeInsurance int32 `xml:"payerTypeInsurance,omitempty"`

	PayerTypePackings int32 `xml:"payerTypePackings,omitempty"`

	Receiver *ResultClientInfo `xml:"receiver,omitempty"`

	RedirectBillOfLading int64 `xml:"redirectBillOfLading,omitempty"`

	Ref1 string `xml:"ref1,omitempty"`

	Ref2 string `xml:"ref2,omitempty"`

	RetMoneyTransferReqAmount float64 `xml:"retMoneyTransferReqAmount,omitempty"`

	RetServicesRequest []*ResultReturnServiceRequest `xml:"retServicesRequest,omitempty"`

	RetShipmentRequest *ResultReturnShipmentRequest `xml:"retShipmentRequest,omitempty"`

	RetThirdPartyPayer bool `xml:"retThirdPartyPayer,omitempty"`

	RetToClientId int64 `xml:"retToClientId,omitempty"`

	RetToOfficeId int64 `xml:"retToOfficeId,omitempty"`

	ReturnBillOfLading int64 `xml:"returnBillOfLading,omitempty"`

	ReturnVoucher *ResultReturnVoucher `xml:"returnVoucher,omitempty"`

	Sender *ResultClientInfo `xml:"sender,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`

	SpecialDeliveryId int32 `xml:"specialDeliveryId,omitempty"`

	TakingDate time.Time `xml:"takingDate,omitempty"`

	WeightCalculation float64 `xml:"weightCalculation,omitempty"`

	WeightDeclared float64 `xml:"weightDeclared,omitempty"`

	WeightMeasured float64 `xml:"weightMeasured,omitempty"`

	WillBringToOfficeId int64 `xml:"willBringToOfficeId,omitempty"`

	PickingType int32 `xml:"pickingType,omitempty"`

	PrimaryPickingBOL int64 `xml:"primaryPickingBOL,omitempty"`

	PendingParcelsDescription bool `xml:"pendingParcelsDescription,omitempty"`

	PendingShipmentDescription bool `xml:"pendingShipmentDescription,omitempty"`

	MoneyTransferPayment *MoneyTransferPayment `xml:"moneyTransferPayment,omitempty"`

	ConsolidationRef string `xml:"consolidationRef,omitempty"`
}

type ResultAmounts struct {
	XMLName xml.Name

	CodBase float64 `xml:"codBase,omitempty"`

	CodPremium float64 `xml:"codPremium,omitempty"`

	DiscPcntAdditional float64 `xml:"discPcntAdditional,omitempty"`

	DiscPcntFixed float64 `xml:"discPcntFixed,omitempty"`

	DiscPcntToBeCalled float64 `xml:"discPcntToBeCalled,omitempty"`

	DiscPcntToOffice float64 `xml:"discPcntToOffice,omitempty"`

	DiscountAdditional float64 `xml:"discountAdditional,omitempty"`

	DiscountFixed float64 `xml:"discountFixed,omitempty"`

	DiscountToBeCalled float64 `xml:"discountToBeCalled,omitempty"`

	DiscountToOffice float64 `xml:"discountToOffice,omitempty"`

	FixedTimeDelivery float64 `xml:"fixedTimeDelivery,omitempty"`

	FuelSurcharge float64 `xml:"fuelSurcharge,omitempty"`

	InsuranceBase float64 `xml:"insuranceBase,omitempty"`

	InsurancePremium float64 `xml:"insurancePremium,omitempty"`

	IslandSurcharge float64 `xml:"islandSurcharge,omitempty"`

	Net float64 `xml:"net,omitempty"`

	Packings float64 `xml:"packings,omitempty"`

	PcntFuelSurcharge float64 `xml:"pcntFuelSurcharge,omitempty"`

	Total float64 `xml:"total,omitempty"`

	Tro float64 `xml:"tro,omitempty"`

	Vat float64 `xml:"vat,omitempty"`

	DiscPcntRetShipment float64 `xml:"discPcntRetShipment,omitempty"`

	DiscountRetShipment float64 `xml:"discountRetShipment,omitempty"`

	SpecialDelivery float64 `xml:"specialDelivery,omitempty"`

	MoneyTransferPremium float64 `xml:"moneyTransferPremium,omitempty"`

	TestBeforePayment float64 `xml:"testBeforePayment,omitempty"`

	DeliveryToFloor float64 `xml:"deliveryToFloor,omitempty"`

	HeavyPackageFee float64 `xml:"heavyPackageFee,omitempty"`

	AddrPickupSurcharge float64 `xml:"addrPickupSurcharge,omitempty"`

	AddrDeliverySurcharge float64 `xml:"addrDeliverySurcharge,omitempty"`

	NonStdDeliveryDateSurcharge float64 `xml:"nonStdDeliveryDateSurcharge,omitempty"`

	VoucherDiscount float64 `xml:"voucherDiscount,omitempty"`

	ReturnAmounts *ResultReturnAmounts `xml:"returnAmounts,omitempty"`

	TollSurcharge float64 `xml:"tollSurcharge,omitempty"`

	ProtectiveMeasuresSurcharge float64 `xml:"protectiveMeasuresSurcharge,omitempty"`

	AddrNormSurcharge float64 `xml:"addrNormSurcharge,omitempty"`
}

type ResultReturnAmounts struct {
	XMLName xml.Name

	MoneyTransferPremium *ReturnAmountDetails `xml:"moneyTransferPremium,omitempty"`
}

type ReturnAmountDetails struct {
	XMLName xml.Name

	Amount float64 `xml:"amount,omitempty"`

	ReturnPayerType int32 `xml:"returnPayerType,omitempty"`
}

type CodPayment struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ codPayment"`

	Date time.Time `xml:"date,omitempty"`

	TotalPayedOutAmount float64 `xml:"totalPayedOutAmount,omitempty"`
}

type ResultDeliveryInfo struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultDeliveryInfo"`

	Consignee string `xml:"consignee,omitempty"`

	DeliveryDate time.Time `xml:"deliveryDate,omitempty"`

	DeliveryNote string `xml:"deliveryNote,omitempty"`
}

type FixedDiscountCardId struct {
	XMLName xml.Name

	AgreementId int64 `xml:"agreementId,omitempty"`

	CardId int64 `xml:"cardId,omitempty"`
}

type ResultOptionsBeforePayment struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultOptionsBeforePayment"`

	Open bool `xml:"open,omitempty"`

	Test bool `xml:"test,omitempty"`

	ReturnServiceTypeId int64 `xml:"returnServiceTypeId,omitempty"`

	ReturnPayerType int32 `xml:"returnPayerType,omitempty"`
}

type ResultPackings struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultPackings"`

	Count int32 `xml:"count,omitempty"`

	PackingId int64 `xml:"packingId,omitempty"`
}

type ResultParcelInfoEx struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultParcelInfoEx"`

	*ResultParcelInfo

	ForeignParcelNumber string `xml:"foreignParcelNumber,omitempty"`

	PackId int64 `xml:"packId,omitempty"`

	SizeDeclared *Size `xml:"sizeDeclared,omitempty"`

	SizeMeasured *Size `xml:"sizeMeasured,omitempty"`

	WeightDeclared float64 `xml:"weightDeclared,omitempty"`

	WeightMeasured float64 `xml:"weightMeasured,omitempty"`

	ForeignParcelNumbersList []string `xml:"foreignParcelNumbersList,omitempty"`
}

type ResultParcelInfo struct {
	XMLName xml.Name

	ParcelId int64 `xml:"parcelId,omitempty"`

	SeqNo int32 `xml:"seqNo,omitempty"`
}

type ResultClientInfo struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultClientInfo"`

	Address *ResultAddressInfo `xml:"address,omitempty"`

	ClientId int64 `xml:"clientId,omitempty"`

	ContactName string `xml:"contactName,omitempty"`

	ObjectName string `xml:"objectName,omitempty"`

	PartnerName string `xml:"partnerName,omitempty"`

	Phones []*ResultPhoneNumber `xml:"phones,omitempty"`
}

type ResultAddressInfo struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultAddressInfo"`

	AddressNote string `xml:"addressNote,omitempty"`

	ApartmentNo string `xml:"apartmentNo,omitempty"`

	BlockNo string `xml:"blockNo,omitempty"`

	CommonObjectId int64 `xml:"commonObjectId,omitempty"`

	CommonObjectName string `xml:"commonObjectName,omitempty"`

	EntranceNo string `xml:"entranceNo,omitempty"`

	FloorNo string `xml:"floorNo,omitempty"`

	MunicipalityName string `xml:"municipalityName,omitempty"`

	PostCode string `xml:"postCode,omitempty"`

	QuarterId int64 `xml:"quarterId,omitempty"`

	QuarterName string `xml:"quarterName,omitempty"`

	QuarterType string `xml:"quarterType,omitempty"`

	RegionName string `xml:"regionName,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	SiteName string `xml:"siteName,omitempty"`

	SiteType string `xml:"siteType,omitempty"`

	StreetId int64 `xml:"streetId,omitempty"`

	StreetName string `xml:"streetName,omitempty"`

	StreetNo string `xml:"streetNo,omitempty"`

	StreetType string `xml:"streetType,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`

	FrnAddressLine1 string `xml:"frnAddressLine1,omitempty"`

	FrnAddressLine2 string `xml:"frnAddressLine2,omitempty"`

	StateId string `xml:"stateId,omitempty"`
}

type ResultPhoneNumber struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultPhoneNumber"`

	Internal string `xml:"internal,omitempty"`

	Number string `xml:"number,omitempty"`
}

type ResultReturnServiceRequest struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultReturnServiceRequest"`

	ParcelsCount int32 `xml:"parcelsCount,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`
}

type ResultReturnShipmentRequest struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultReturnShipmentRequest"`

	AmountInsuranceBase float64 `xml:"amountInsuranceBase,omitempty"`

	Fragile bool `xml:"fragile,omitempty"`

	ParcelsCount int32 `xml:"parcelsCount,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`
}

type ResultReturnVoucher struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultReturnVoucher"`

	PayerType int32 `xml:"payerType,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`
}

type MoneyTransferPayment struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ moneyTransferPayment"`

	Date time.Time `xml:"date,omitempty"`

	TotalPayedOutAmount float64 `xml:"totalPayedOutAmount,omitempty"`
}

type TrackPicking struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ trackPicking"`

	SessionId string `xml:"sessionId,omitempty"`

	BillOfLading string `xml:"billOfLading,omitempty"`
}

type TrackPickingResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ trackPickingResponse"`

	Return_ []*ResultTrackPicking `xml:"return,omitempty"`
}

type ResultTrackPicking struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultTrackPicking"`

	Consignee string `xml:"consignee,omitempty"`

	Moment time.Time `xml:"moment,omitempty"`

	OperationCode int32 `xml:"operationCode,omitempty"`

	OperationComment string `xml:"operationComment,omitempty"`

	OperationDescription string `xml:"operationDescription,omitempty"`

	SiteName string `xml:"siteName,omitempty"`

	SiteType string `xml:"siteType,omitempty"`

	SignatureImage string `xml:"signatureImage,omitempty"`

	Barcode int64 `xml:"barcode,omitempty"`

	ForeignParcelNumber string `xml:"foreignParcelNumber,omitempty"`

	ExceptionCodes []int32 `xml:"exceptionCodes,omitempty"`
}

type ListServices struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listServices"`

	SessionId string `xml:"sessionId,omitempty"`

	Date time.Time `xml:"date,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ListServicesResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listServicesResponse"`

	Return_ []*ResultCourierService `xml:"return,omitempty"`
}

type ResultCourierService struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultCourierService"`

	AllowanceBackDocumentsRequest *ComplementaryServiceAllowance `xml:"allowanceBackDocumentsRequest,omitempty"`

	AllowanceBackReceiptRequest *ComplementaryServiceAllowance `xml:"allowanceBackReceiptRequest,omitempty"`

	AllowanceCashOnDelivery *ComplementaryServiceAllowance `xml:"allowanceCashOnDelivery,omitempty"`

	AllowanceFixedTimeDelivery *ComplementaryServiceAllowance `xml:"allowanceFixedTimeDelivery,omitempty"`

	AllowanceInsurance *ComplementaryServiceAllowance `xml:"allowanceInsurance,omitempty"`

	AllowanceToBeCalled *ComplementaryServiceAllowance `xml:"allowanceToBeCalled,omitempty"`

	Name string `xml:"name,omitempty"`

	TypeId int64 `xml:"typeId,omitempty"`

	CargoType int32 `xml:"cargoType,omitempty"`
}

type ListSites struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listSites"`

	SessionId string `xml:"sessionId,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Name string `xml:"name,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ListSitesResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listSitesResponse"`

	Return_ []*ResultSite `xml:"return,omitempty"`
}

type ResultSite struct {
	XMLName xml.Name

	AddrNomen *AddrNomen `xml:"addrNomen,omitempty"`

	Id int64 `xml:"id,omitempty"`

	Municipality string `xml:"municipality,omitempty"`

	Name string `xml:"name,omitempty"`

	PostCode string `xml:"postCode,omitempty"`

	Region string `xml:"region,omitempty"`

	Type_ string `xml:"type,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`

	ServingOfficeId int64 `xml:"servingOfficeId,omitempty"`

	CoordX float64 `xml:"coordX,omitempty"`

	CoordY float64 `xml:"coordY,omitempty"`

	CoordType int32 `xml:"coordType,omitempty"`

	ServingDays string `xml:"servingDays,omitempty"`
}

type ValidateAddress struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ validateAddress"`

	SessionId string `xml:"sessionId,omitempty"`

	Address *ParamAddress `xml:"address,omitempty"`

	ValidationMode int32 `xml:"validationMode,omitempty"`
}

type ParamAddress struct {
	XMLName xml.Name

	AddressNote string `xml:"addressNote,omitempty"`

	ApartmentNo string `xml:"apartmentNo,omitempty"`

	BlockNo string `xml:"blockNo,omitempty"`

	CommonObjectId int64 `xml:"commonObjectId,omitempty"`

	CoordX float64 `xml:"coordX,omitempty"`

	CoordY float64 `xml:"coordY,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`

	EntranceNo string `xml:"entranceNo,omitempty"`

	FloorNo string `xml:"floorNo,omitempty"`

	FrnAddressLine1 string `xml:"frnAddressLine1,omitempty"`

	FrnAddressLine2 string `xml:"frnAddressLine2,omitempty"`

	PostCode string `xml:"postCode,omitempty"`

	QuarterId int64 `xml:"quarterId,omitempty"`

	QuarterName string `xml:"quarterName,omitempty"`

	QuarterType string `xml:"quarterType,omitempty"`

	SerializedAddress string `xml:"serializedAddress,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	SiteName string `xml:"siteName,omitempty"`

	StateId string `xml:"stateId,omitempty"`

	StreetId int64 `xml:"streetId,omitempty"`

	StreetName string `xml:"streetName,omitempty"`

	StreetNo string `xml:"streetNo,omitempty"`

	StreetType string `xml:"streetType,omitempty"`
}

type ValidateAddressResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ validateAddressResponse"`

	Return_ bool `xml:"return,omitempty"`
}

type GetStateById struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getStateById"`

	SessionId string `xml:"sessionId,omitempty"`

	StateId string `xml:"stateId,omitempty"`
}

type GetStateByIdResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getStateByIdResponse"`

	Return_ *ResultState `xml:"return,omitempty"`
}

type GetClientById struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getClientById"`

	SessionId string `xml:"sessionId,omitempty"`

	ClientId int64 `xml:"clientId,omitempty"`
}

type GetClientByIdResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getClientByIdResponse"`

	Return_ *ResultClientData `xml:"return,omitempty"`
}

type ResultClientData struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultClientData"`

	Address *ResultAddress `xml:"address,omitempty"`

	ClientId int64 `xml:"clientId,omitempty"`

	ContactName string `xml:"contactName,omitempty"`

	ObjectName string `xml:"objectName,omitempty"`

	PartnerName string `xml:"partnerName,omitempty"`

	Phones []*ResultPhoneNumber `xml:"phones,omitempty"`
}

type ResultAddress struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultAddress"`

	AddressNote string `xml:"addressNote,omitempty"`

	ApartmentNo string `xml:"apartmentNo,omitempty"`

	BlockNo string `xml:"blockNo,omitempty"`

	CommonObjectId int64 `xml:"commonObjectId,omitempty"`

	CommonObjectName string `xml:"commonObjectName,omitempty"`

	EntranceNo string `xml:"entranceNo,omitempty"`

	FloorNo string `xml:"floorNo,omitempty"`

	MunicipalityName string `xml:"municipalityName,omitempty"`

	PostCode string `xml:"postCode,omitempty"`

	QuarterId int64 `xml:"quarterId,omitempty"`

	QuarterName string `xml:"quarterName,omitempty"`

	QuarterType string `xml:"quarterType,omitempty"`

	RegionName string `xml:"regionName,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	SiteName string `xml:"siteName,omitempty"`

	SiteType string `xml:"siteType,omitempty"`

	StreetId int64 `xml:"streetId,omitempty"`

	StreetName string `xml:"streetName,omitempty"`

	StreetNo string `xml:"streetNo,omitempty"`

	StreetType string `xml:"streetType,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`

	FrnAddressLine1 string `xml:"frnAddressLine1,omitempty"`

	FrnAddressLine2 string `xml:"frnAddressLine2,omitempty"`

	StateId string `xml:"stateId,omitempty"`
}

type ListOfficesEx struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listOfficesEx"`

	SessionId string `xml:"sessionId,omitempty"`

	Name string `xml:"name,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`
}

type ListOfficesExResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listOfficesExResponse"`

	Return_ []*ResultOfficeEx `xml:"return,omitempty"`
}

type ResultOfficeEx struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultOfficeEx"`

	Address *ResultAddressEx `xml:"address,omitempty"`

	Id int64 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	WorkingTimeFrom int16 `xml:"workingTimeFrom,omitempty"`

	WorkingTimeHalfFrom int16 `xml:"workingTimeHalfFrom,omitempty"`

	WorkingTimeHalfTo int16 `xml:"workingTimeHalfTo,omitempty"`

	WorkingTimeTo int16 `xml:"workingTimeTo,omitempty"`

	MaxParcelDimensions *Size `xml:"maxParcelDimensions,omitempty"`

	MaxParcelWeight float64 `xml:"maxParcelWeight,omitempty"`

	WorkingTimeSchedule []*ResultWorkingTime `xml:"workingTimeSchedule,omitempty"`

	OfficeType int16 `xml:"officeType,omitempty"`

	WorkingTimeDayOffFrom int16 `xml:"workingTimeDayOffFrom,omitempty"`

	WorkingTimeDayOffTo int16 `xml:"workingTimeDayOffTo,omitempty"`

	NearbyOfficeId int64 `xml:"nearbyOfficeId,omitempty"`

	BroughtToAllowed bool `xml:"broughtToAllowed,omitempty"`

	ToBeCalledAllowed bool `xml:"toBeCalledAllowed,omitempty"`
}

type ResultAddressEx struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultAddressEx"`

	AddressNote string `xml:"addressNote,omitempty"`

	ApartmentNo string `xml:"apartmentNo,omitempty"`

	BlockNo string `xml:"blockNo,omitempty"`

	CommonObjectId int64 `xml:"commonObjectId,omitempty"`

	CommonObjectName string `xml:"commonObjectName,omitempty"`

	CoordTypeId int32 `xml:"coordTypeId,omitempty"`

	CoordX float64 `xml:"coordX,omitempty"`

	CoordY float64 `xml:"coordY,omitempty"`

	EntranceNo string `xml:"entranceNo,omitempty"`

	FloorNo string `xml:"floorNo,omitempty"`

	FullAddressString string `xml:"fullAddressString,omitempty"`

	PostCode string `xml:"postCode,omitempty"`

	QuarterId int64 `xml:"quarterId,omitempty"`

	QuarterName string `xml:"quarterName,omitempty"`

	QuarterType string `xml:"quarterType,omitempty"`

	ResultSite *ResultSite `xml:"resultSite,omitempty"`

	StreetId int64 `xml:"streetId,omitempty"`

	StreetName string `xml:"streetName,omitempty"`

	StreetNo string `xml:"streetNo,omitempty"`

	StreetType string `xml:"streetType,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`

	FrnAddressLine1 string `xml:"frnAddressLine1,omitempty"`

	FrnAddressLine2 string `xml:"frnAddressLine2,omitempty"`

	StateId string `xml:"stateId,omitempty"`

	SiteAddressString string `xml:"siteAddressString,omitempty"`

	LocalAddressString string `xml:"localAddressString,omitempty"`
}

type ResultWorkingTime struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultWorkingTime"`

	Date time.Time `xml:"date,omitempty"`

	WorkingTimeFrom int16 `xml:"workingTimeFrom,omitempty"`

	WorkingTimeTo int16 `xml:"workingTimeTo,omitempty"`

	WorkingTimeException bool `xml:"workingTimeException,omitempty"`
}

type CalculateMultipleServices struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ calculateMultipleServices"`

	SessionId string `xml:"sessionId,omitempty"`

	Calculation *ParamCalculation `xml:"calculation,omitempty"`

	ServiceTypeIds []int64 `xml:"serviceTypeIds,omitempty"`
}

type ParamCalculation struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ paramCalculation"`

	AmountCodBase float64 `xml:"amountCodBase,omitempty"`

	AmountInsuranceBase float64 `xml:"amountInsuranceBase,omitempty"`

	AutoAdjustTakingDate bool `xml:"autoAdjustTakingDate,omitempty"`

	BroughtToOffice bool `xml:"broughtToOffice,omitempty"`

	CheckTBCOfficeWorkDay bool `xml:"checkTBCOfficeWorkDay,omitempty"`

	DeferredDeliveryWorkDays int32 `xml:"deferredDeliveryWorkDays,omitempty"`

	DeliveryToFloorNo int32 `xml:"deliveryToFloorNo,omitempty"`

	Documents bool `xml:"documents,omitempty"`

	FixedTimeDelivery int16 `xml:"fixedTimeDelivery,omitempty"`

	Fragile bool `xml:"fragile,omitempty"`

	HalfWorkDayDelivery bool `xml:"halfWorkDayDelivery,omitempty"`

	IgnoreAmountInsuranceBaseIfNotApplicable bool `xml:"ignoreAmountInsuranceBaseIfNotApplicable,omitempty"`

	IncludeShippingPriceInCod bool `xml:"includeShippingPriceInCod,omitempty"`

	OfficeToBeCalledId int64 `xml:"officeToBeCalledId,omitempty"`

	OptionsBeforePayment *ParamOptionsBeforePayment `xml:"optionsBeforePayment,omitempty"`

	Palletized bool `xml:"palletized,omitempty"`

	Parcels []*ParamParcelInfo `xml:"parcels,omitempty"`

	ParcelsCount int32 `xml:"parcelsCount,omitempty"`

	PayCodToThirdParty bool `xml:"payCodToThirdParty,omitempty"`

	PayerRefId int64 `xml:"payerRefId,omitempty"`

	PayerRefInsuranceId int64 `xml:"payerRefInsuranceId,omitempty"`

	PayerRefPackingsId int64 `xml:"payerRefPackingsId,omitempty"`

	PayerType int32 `xml:"payerType,omitempty"`

	PayerTypeInsurance int32 `xml:"payerTypeInsurance,omitempty"`

	PayerTypePackings int32 `xml:"payerTypePackings,omitempty"`

	ReceiverCountryId int64 `xml:"receiverCountryId,omitempty"`

	ReceiverId int64 `xml:"receiverId,omitempty"`

	ReceiverPostCode string `xml:"receiverPostCode,omitempty"`

	ReceiverSiteId int64 `xml:"receiverSiteId,omitempty"`

	RetMoneyTransferReqAmount float64 `xml:"retMoneyTransferReqAmount,omitempty"`

	SenderCountryId int64 `xml:"senderCountryId,omitempty"`

	SenderId int64 `xml:"senderId,omitempty"`

	SenderPostCode string `xml:"senderPostCode,omitempty"`

	SenderSiteId int64 `xml:"senderSiteId,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`

	SpecialDeliveryId int32 `xml:"specialDeliveryId,omitempty"`

	TakingDate time.Time `xml:"takingDate,omitempty"`

	ToBeCalled bool `xml:"toBeCalled,omitempty"`

	WeightDeclared float64 `xml:"weightDeclared,omitempty"`

	WillBringToOfficeId int64 `xml:"willBringToOfficeId,omitempty"`
}

type ParamOptionsBeforePayment struct {
	XMLName xml.Name

	Open bool `xml:"open,omitempty"`

	ReturnPayerType string `xml:"returnPayerType,omitempty"`

	ReturnServiceTypeId int64 `xml:"returnServiceTypeId,omitempty"`

	Test bool `xml:"test,omitempty"`
}

type ParamParcelInfo struct {
	XMLName xml.Name

	ForeignParcelNumber string `xml:"foreignParcelNumber,omitempty"`

	PackId int64 `xml:"packId,omitempty"`

	ParcelId int64 `xml:"parcelId,omitempty"`

	PredefinedSize string `xml:"predefinedSize,omitempty"`

	Ref1 string `xml:"ref1,omitempty"`

	Ref2 string `xml:"ref2,omitempty"`

	SeqNo int32 `xml:"seqNo,omitempty"`

	Size *Size `xml:"size,omitempty"`

	Weight float64 `xml:"weight,omitempty"`
}

type CalculateMultipleServicesResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ calculateMultipleServicesResponse"`

	Return_ []*ResultCalculationMS `xml:"return,omitempty"`
}

type ResultCalculationMS struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultCalculationMS"`

	ErrorDescription string `xml:"errorDescription,omitempty"`

	ResultInfo *ResultCalculation `xml:"resultInfo,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`
}

type ResultCalculation struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultCalculation"`

	Amounts *ResultAmounts `xml:"amounts,omitempty"`

	DeadlineDelivery time.Time `xml:"deadlineDelivery,omitempty"`

	PartialDiscount bool `xml:"partialDiscount,omitempty"`

	TakingDate time.Time `xml:"takingDate,omitempty"`
}

type Login struct {
	XMLName xml.Name `xml:"ns1:login"`
	XMLNS   string   `xml:"xmlns:ns1,attr"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`
}

type LoginResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ loginResponse"`

	Return_ *ResultLogin `xml:"return,omitempty"`
}

type ResultLogin struct {
	XMLName xml.Name

	ClientId int64 `xml:"clientId,omitempty"`

	SessionId string `xml:"sessionId,omitempty"`

	ServerTime time.Time `xml:"serverTime,omitempty"`
}

type InvalidUsernameOrPasswordException struct {
	Message string `xml:"message,omitempty"`
}

type NoUserPermissionsException struct {
	Message string `xml:"message,omitempty"`
}

type AccountLockedException struct {
	Message string `xml:"message,omitempty"`
}

type TrackParcelMultiple struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ trackParcelMultiple"`

	SessionId string `xml:"sessionId,omitempty"`

	Barcodes []string `xml:"barcodes,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`

	ReturnOnlyLastOperation bool `xml:"returnOnlyLastOperation,omitempty"`
}

type TrackParcelMultipleResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ trackParcelMultipleResponse"`

	Return_ []*ResultTrackPickingEx `xml:"return,omitempty"`
}

type ListAllSites struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listAllSites"`

	SessionId string `xml:"sessionId,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`
}

type ListAllSitesResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listAllSitesResponse"`

	Return_ []*ResultSite `xml:"return,omitempty"`
}

type SerializeAddress struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ serializeAddress"`

	SessionId string `xml:"sessionId,omitempty"`

	Address *ParamAddress `xml:"address,omitempty"`
}

type SerializeAddressResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ serializeAddressResponse"`

	Return_ string `xml:"return,omitempty"`
}

type ListStreets struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listStreets"`

	SessionId string `xml:"sessionId,omitempty"`

	Name string `xml:"name,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ListStreetsResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listStreetsResponse"`

	Return_ []*ResultStreet `xml:"return,omitempty"`
}

type ResultStreet struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultStreet"`

	ActualName string `xml:"actualName,omitempty"`

	Id int64 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type CalculatePicking struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ calculatePicking"`

	SessionId string `xml:"sessionId,omitempty"`

	Picking *ParamPicking `xml:"picking,omitempty"`
}

type ParamPicking struct {
	XMLName xml.Name

	AmountCodBase float64 `xml:"amountCodBase,omitempty"`

	AmountInsuranceBase float64 `xml:"amountInsuranceBase,omitempty"`

	BackDocumentsRequest bool `xml:"backDocumentsRequest,omitempty"`

	BackReceiptRequest bool `xml:"backReceiptRequest,omitempty"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`

	ClientSystemId int64 `xml:"clientSystemId,omitempty"`

	ConsolidationRef string `xml:"consolidationRef,omitempty"`

	Contents string `xml:"contents,omitempty"`

	DeferredDeliveryWorkDays int32 `xml:"deferredDeliveryWorkDays,omitempty"`

	DeliveryLimitViolationAutoAdjustment int32 `xml:"deliveryLimitViolationAutoAdjustment,omitempty"`

	DeliveryToFloorNo int32 `xml:"deliveryToFloorNo,omitempty"`

	DiscCalc *FixedDiscountCardId `xml:"discCalc,omitempty"`

	Documents bool `xml:"documents,omitempty"`

	ExciseGoods bool `xml:"exciseGoods,omitempty"`

	FixedTimeDelivery int16 `xml:"fixedTimeDelivery,omitempty"`

	Fragile bool `xml:"fragile,omitempty"`

	HalfWorkDayDelivery bool `xml:"halfWorkDayDelivery,omitempty"`

	IncludeShippingPriceInCod bool `xml:"includeShippingPriceInCod,omitempty"`

	NoteClient string `xml:"noteClient,omitempty"`

	OfficeToBeCalledId int64 `xml:"officeToBeCalledId,omitempty"`

	OptionsBeforePayment *ParamOptionsBeforePayment `xml:"optionsBeforePayment,omitempty"`

	PackId int64 `xml:"packId,omitempty"`

	Packing string `xml:"packing,omitempty"`

	Packings []*ParamPackings `xml:"packings,omitempty"`

	Palletized bool `xml:"palletized,omitempty"`

	Parcels []*ParamParcelInfo `xml:"parcels,omitempty"`

	ParcelsCount int32 `xml:"parcelsCount,omitempty"`

	PayCodToLoggedInClient bool `xml:"payCodToLoggedInClient,omitempty"`

	PayCodToThirdParty bool `xml:"payCodToThirdParty,omitempty"`

	PayerRefId int64 `xml:"payerRefId,omitempty"`

	PayerRefInsuranceId int64 `xml:"payerRefInsuranceId,omitempty"`

	PayerRefPackingsId int64 `xml:"payerRefPackingsId,omitempty"`

	PayerType int32 `xml:"payerType"`

	PayerTypeInsurance int32 `xml:"payerTypeInsurance,omitempty"`

	PayerTypePackings int32 `xml:"payerTypePackings,omitempty"`

	PendingParcelsDescription bool `xml:"pendingParcelsDescription,omitempty"`

	PendingShipmentDescription bool `xml:"pendingShipmentDescription,omitempty"`

	Receiver *ParamClientData `xml:"receiver,omitempty"`

	Ref1 string `xml:"ref1,omitempty"`

	Ref2 string `xml:"ref2,omitempty"`

	RequireUnsuccessfulDeliveryStickerImage bool `xml:"requireUnsuccessfulDeliveryStickerImage,omitempty"`

	RetMoneyTransferReqAmount float64 `xml:"retMoneyTransferReqAmount,omitempty"`

	RetServicesRequest []*ParamReturnServiceRequest `xml:"retServicesRequest,omitempty"`

	RetShipmentRequest *ParamReturnShipmentRequest `xml:"retShipmentRequest,omitempty"`

	RetThirdPartyPayer bool `xml:"retThirdPartyPayer,omitempty"`

	RetToClientId int64 `xml:"retToClientId,omitempty"`

	RetToOfficeId int64 `xml:"retToOfficeId,omitempty"`

	ReturnVoucher *ParamReturnVoucher `xml:"returnVoucher,omitempty"`

	Sender *ParamClientData `xml:"sender,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`

	Size *Size `xml:"size,omitempty"`

	SkipAutomaticParcelsCreation bool `xml:"skipAutomaticParcelsCreation,omitempty"`

	SpecialDeliveryId int32 `xml:"specialDeliveryId,omitempty"`

	TakingDate time.Time `xml:"takingDate,omitempty"`

	WeightDeclared float64 `xml:"weightDeclared,omitempty"`

	WillBringToOffice bool `xml:"willBringToOffice,omitempty"`

	WillBringToOfficeId int64 `xml:"willBringToOfficeId,omitempty"`
}

type ParamPackings struct {
	XMLName xml.Name

	Count int32 `xml:"count,omitempty"`

	PackingId int64 `xml:"packingId,omitempty"`
}

type ParamClientData struct {
	XMLName xml.Name

	Address *ParamAddress `xml:"address,omitempty"`

	ClientId int64 `xml:"clientId,omitempty"`

	ContactName string `xml:"contactName,omitempty"`

	Email string `xml:"email,omitempty"`

	ObjectName string `xml:"objectName,omitempty"`

	PartnerName string `xml:"partnerName,omitempty"`

	Phones []*ParamPhoneNumber `xml:"phones,omitempty"`

	PrivatePersonType int16 `xml:"privatePersonType,omitempty"`
}

type ParamPhoneNumber struct {
	XMLName xml.Name

	Internal string `xml:"internal,omitempty"`

	Number string `xml:"number,omitempty"`
}

type ParamReturnServiceRequest struct {
	XMLName xml.Name

	ParcelsCount int32 `xml:"parcelsCount,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`
}

type ParamReturnShipmentRequest struct {
	XMLName xml.Name

	AmountInsuranceBase float64 `xml:"amountInsuranceBase,omitempty"`

	Fragile bool `xml:"fragile,omitempty"`

	ParcelsCount int32 `xml:"parcelsCount,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`
}

type ParamReturnVoucher struct {
	XMLName xml.Name

	ExternalReturnVoucherId string `xml:"externalReturnVoucherId,omitempty"`

	PayerType int32 `xml:"payerType,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`

	ValidityPeriod int32 `xml:"validityPeriod,omitempty"`
}

type CalculatePickingResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ calculatePickingResponse"`

	Return_ *ResultCalculation `xml:"return,omitempty"`
}

type DeserializeAddress struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ deserializeAddress"`

	SessionId string `xml:"sessionId,omitempty"`

	Address string `xml:"address,omitempty"`
}

type DeserializeAddressResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ deserializeAddressResponse"`

	Return_ *ParamAddress `xml:"return,omitempty"`
}

type CreateBillOfLading struct {
	XMLName xml.Name `xml:"ns1:createBillOfLading"`
	XMLNS   string   `xml:"xmlns:ns1,attr"`

	SessionId string `xml:"sessionId,omitempty"`

	Picking *ParamPicking `xml:"picking,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type CreateBillOfLadingResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ createBillOfLadingResponse"`

	Return_ *ResultBOL `xml:"return,omitempty"`
}

type ResultBOL struct {
	XMLName xml.Name

	Amounts *ResultAmounts `xml:"amounts,omitempty"`

	DeadlineDelivery time.Time `xml:"deadlineDelivery,omitempty"`

	GeneratedParcels []*ResultParcelInfo `xml:"generatedParcels,omitempty"`
}

type ListQuarters struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listQuarters"`

	SessionId string `xml:"sessionId,omitempty"`

	Name string `xml:"name,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ListQuartersResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listQuartersResponse"`

	Return_ []*ResultQuarter `xml:"return,omitempty"`
}

type ResultQuarter struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultQuarter"`

	ActualName string `xml:"actualName,omitempty"`

	Id int64 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type GetSiteById struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getSiteById"`

	SessionId string `xml:"sessionId,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`
}

type GetSiteByIdResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getSiteByIdResponse"`

	Return_ *ResultSite `xml:"return,omitempty"`
}

type AddressSearch struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ addressSearch"`

	SessionId string `xml:"sessionId,omitempty"`

	Address *ParamAddressSearch `xml:"address,omitempty"`
}

type ParamAddressSearch struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ paramAddressSearch"`

	BlockNo string `xml:"blockNo,omitempty"`

	CommonObjectId int64 `xml:"commonObjectId,omitempty"`

	EntranceNo string `xml:"entranceNo,omitempty"`

	QuarterId int64 `xml:"quarterId,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	StreetId int64 `xml:"streetId,omitempty"`

	StreetNo string `xml:"streetNo,omitempty"`

	ReturnCityCenterIfNoAddress bool `xml:"returnCityCenterIfNoAddress,omitempty"`
}

type AddressSearchResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ addressSearchResponse"`

	Return_ []*ResultAddressSearch `xml:"return,omitempty"`
}

type ResultAddressSearch struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultAddressSearch"`

	Actual bool `xml:"actual,omitempty"`

	AdditionalAddressProcessing int32 `xml:"additionalAddressProcessing,omitempty"`

	CoordType int32 `xml:"coordType,omitempty"`

	CoordX float64 `xml:"coordX,omitempty"`

	CoordY float64 `xml:"coordY,omitempty"`

	DistanceToSiteCenter float64 `xml:"distanceToSiteCenter,omitempty"`

	MicroregionId int64 `xml:"microregionId,omitempty"`

	Text string `xml:"text,omitempty"`
}

type ListSitesEx struct {
	XMLName xml.Name `xml:"ns1:listSitesEx"`
	XMLNS   string   `xml:"xmlns:ns1,attr"`

	SessionId string `xml:"sessionId,omitempty"`

	Filter *ParamFilterSite `xml:"filter,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ParamFilterSite struct {
	XMLName xml.Name

	CountryId int64 `xml:"countryId,omitempty"`

	Municipality string `xml:"municipality,omitempty"`

	Name string `xml:"name,omitempty"`

	PostCode string `xml:"postCode,omitempty"`

	Region string `xml:"region,omitempty"`

	SearchString string `xml:"searchString,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type ListSitesExResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listSitesExResponse"`

	Return_ []*ResultSiteEx `xml:"return,omitempty"`
}

type ResultSiteEx struct {
	XMLName xml.Name

	ExactMatch bool `xml:"exactMatch,omitempty"`

	Site *ResultSite `xml:"site,omitempty"`
}

type ListQuarterTypes struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listQuarterTypes"`

	SessionId string `xml:"sessionId,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ListQuarterTypesResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listQuarterTypesResponse"`

	Return_ []string `xml:"return,omitempty"`
}

type CreateCustomTravelLabelPDFType1 struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ createCustomTravelLabelPDFType1"`

	SessionId string `xml:"sessionId,omitempty"`

	ParcelId int64 `xml:"parcelId,omitempty"`
}

type CreateCustomTravelLabelPDFType1Response struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ createCustomTravelLabelPDFType1Response"`

	Return_ []byte `xml:"return,omitempty"`
}

type GetWeightInterval struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getWeightInterval"`

	SessionId string `xml:"sessionId,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`

	SenderSiteId int64 `xml:"senderSiteId,omitempty"`

	ReceiverSiteId int64 `xml:"receiverSiteId,omitempty"`

	Date time.Time `xml:"date,omitempty"`

	Documents bool `xml:"documents,omitempty"`

	SenderCountryId int64 `xml:"senderCountryId,omitempty"`

	SenderPostCode string `xml:"senderPostCode,omitempty"`

	ReceiverCountryId int64 `xml:"receiverCountryId,omitempty"`

	ReceiverPostCode string `xml:"receiverPostCode,omitempty"`

	SenderId int64 `xml:"senderId,omitempty"`

	ReceiverId int64 `xml:"receiverId,omitempty"`

	SenderOfficeId int64 `xml:"senderOfficeId,omitempty"`

	ReceiverOfficeId int64 `xml:"receiverOfficeId,omitempty"`
}

type GetWeightIntervalResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getWeightIntervalResponse"`

	Return_ *ResultMinMaxReal `xml:"return,omitempty"`
}

type ResultMinMaxReal struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultMinMaxReal"`

	MaxValue float64 `xml:"maxValue,omitempty"`

	MinValue float64 `xml:"minValue,omitempty"`
}

type FinalizeBillOfLadingCreation struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ finalizeBillOfLadingCreation"`

	SessionId string `xml:"sessionId,omitempty"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`
}

type FinalizeBillOfLadingCreationResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ finalizeBillOfLadingCreationResponse"`

	Return_ *ResultBOL `xml:"return,omitempty"`
}

type GetSitesByAddrNomenType struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getSitesByAddrNomenType"`

	SessionId string `xml:"sessionId,omitempty"`

	AddrNomen *AddrNomen `xml:"addrNomen,omitempty"`
}

type GetSitesByAddrNomenTypeResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getSitesByAddrNomenTypeResponse"`

	Return_ []*ResultSite `xml:"return,omitempty"`
}

type ListSpecialDeliveryRequirements struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listSpecialDeliveryRequirements"`

	SessionId string `xml:"sessionId,omitempty"`
}

type ListSpecialDeliveryRequirementsResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listSpecialDeliveryRequirementsResponse"`

	Return_ []*ResultSpecialDeliveryRequirement `xml:"return,omitempty"`
}

type ResultSpecialDeliveryRequirement struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultSpecialDeliveryRequirement"`

	SpecialDeliveryId int32 `xml:"specialDeliveryId,omitempty"`

	SpecialDeliveryPrice float64 `xml:"specialDeliveryPrice,omitempty"`

	SpecialDeliveryText string `xml:"specialDeliveryText,omitempty"`
}

type SearchSecondaryPickings struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ searchSecondaryPickings"`

	SessionId string `xml:"sessionId,omitempty"`

	ParamSearchSecondaryPickings *ParamSearchSecondaryPickings `xml:"paramSearchSecondaryPickings,omitempty"`
}

type ParamSearchSecondaryPickings struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ paramSearchSecondaryPickings"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`

	SecondaryPickingType int32 `xml:"secondaryPickingType,omitempty"`
}

type SearchSecondaryPickingsResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ searchSecondaryPickingsResponse"`

	Return_ []*ResultPickingInfo `xml:"return,omitempty"`
}

type ResultPickingInfo struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultPickingInfo"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`

	HasScans bool `xml:"hasScans,omitempty"`

	SecondaryPickingType int32 `xml:"secondaryPickingType,omitempty"`

	ServiceTypeId int64 `xml:"serviceTypeId,omitempty"`

	TakingDate time.Time `xml:"takingDate,omitempty"`
}

type ListContractClients struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listContractClients"`

	SessionId string `xml:"sessionId,omitempty"`
}

type ListContractClientsResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listContractClientsResponse"`

	Return_ []*ResultClientData `xml:"return,omitempty"`
}

type GetPickingParcels struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getPickingParcels"`

	SessionId string `xml:"sessionId,omitempty"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`
}

type GetPickingParcelsResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getPickingParcelsResponse"`

	Return_ []*ResultParcelInfo `xml:"return,omitempty"`
}

type MapForeignBarcode struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ mapForeignBarcode"`

	SessionId string `xml:"sessionId,omitempty"`

	ParcelId int64 `xml:"parcelId,omitempty"`

	ForeignParcelNumber string `xml:"foreignParcelNumber,omitempty"`
}

type MapForeignBarcodeResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ mapForeignBarcodeResponse"`
}

type CreateBillOfLadingPDF struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ createBillOfLadingPDF"`

	SessionId string `xml:"sessionId,omitempty"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`

	IncludeAutoPrintJS bool `xml:"includeAutoPrintJS,omitempty"`
}

type CreateBillOfLadingPDFResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ createBillOfLadingPDFResponse"`

	Return_ []byte `xml:"return,omitempty"`
}

type ListCommonObjects struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listCommonObjects"`

	SessionId string `xml:"sessionId,omitempty"`

	Name string `xml:"name,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ListCommonObjectsResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listCommonObjectsResponse"`

	Return_ []*ResultCommonObject `xml:"return,omitempty"`
}

type ResultCommonObject struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultCommonObject"`

	Address string `xml:"address,omitempty"`

	Id int64 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type UpdateBillOfLading struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ updateBillOfLading"`

	SessionId string `xml:"sessionId,omitempty"`

	Picking *ParamPicking `xml:"picking,omitempty"`
}

type UpdateBillOfLadingResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ updateBillOfLadingResponse"`

	Return_ *ResultBOL `xml:"return,omitempty"`
}

type GetAdditionalUserParams struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getAdditionalUserParams"`

	SessionId string `xml:"sessionId,omitempty"`

	Date time.Time `xml:"date,omitempty"`
}

type GetAdditionalUserParamsResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getAdditionalUserParamsResponse"`

	Return_ []int32 `xml:"return,omitempty"`
}

type InvalidatePicking struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ invalidatePicking"`

	SessionId string `xml:"sessionId,omitempty"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`

	CancelComment string `xml:"cancelComment,omitempty"`
}

type InvalidatePickingResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ invalidatePickingResponse"`
}

type TrackParcel struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ trackParcel"`

	SessionId string `xml:"sessionId,omitempty"`

	ParcelId string `xml:"parcelId,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`

	ReturnOnlyLastOperation bool `xml:"returnOnlyLastOperation,omitempty"`
}

type TrackParcelResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ trackParcelResponse"`

	Return_ []*ResultTrackPickingEx `xml:"return,omitempty"`
}

type ValidatePostCode struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ validatePostCode"`

	SessionId string `xml:"sessionId,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	PostCode string `xml:"postCode,omitempty"`
}

type ValidatePostCodeResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ validatePostCodeResponse"`

	Return_ bool `xml:"return,omitempty"`
}

type IsSessionActive struct {
	XMLName xml.Name `xml:"ns1:isSessionActive"`
	XMLNS   string   `xml:"xmlns:ns1,attr"`

	SessionId string `xml:"sessionId,omitempty"`

	RefreshSession bool `xml:"refreshSession,omitempty"`
}

type IsSessionActiveResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ isSessionActiveResponse"`

	Return_ bool `xml:"return,omitempty"`
}

type GetRoutingLabelInfo struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getRoutingLabelInfo"`

	SessionId string `xml:"sessionId,omitempty"`

	ParcelId int64 `xml:"parcelId,omitempty"`
}

type GetRoutingLabelInfoResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getRoutingLabelInfoResponse"`

	Return_ *ResultRoutingLabelInfo `xml:"return,omitempty"`
}

type ResultRoutingLabelInfo struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultRoutingLabelInfo"`

	HubId int64 `xml:"hubId,omitempty"`

	OfficeId int64 `xml:"officeId,omitempty"`

	DeadlineDay int32 `xml:"deadlineDay,omitempty"`

	DeadlineMonth int32 `xml:"deadlineMonth,omitempty"`

	TourId int64 `xml:"tourId,omitempty"`

	FullBarcode string `xml:"fullBarcode,omitempty"`

	ExportPriority int32 `xml:"exportPriority,omitempty"`
}

type GetAddressNomenclature struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getAddressNomenclature"`

	SessionId string `xml:"sessionId,omitempty"`

	NomenType int32 `xml:"nomenType,omitempty"`

	CountryId int64 `xml:"countryId,omitempty"`
}

type GetAddressNomenclatureResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getAddressNomenclatureResponse"`

	Return_ string `xml:"return,omitempty"`
}

type MakeAddressString struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ makeAddressString"`

	SessionId string `xml:"sessionId,omitempty"`

	Address *ParamAddress `xml:"address,omitempty"`
}

type MakeAddressStringResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ makeAddressStringResponse"`

	Return_ *ResultAddressString `xml:"return,omitempty"`
}

type ResultAddressString struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultAddressString"`

	FullAddress string `xml:"fullAddress,omitempty"`

	LocalAddress string `xml:"localAddress,omitempty"`

	SiteAddress string `xml:"siteAddress,omitempty"`
}

type ListServicesForSites struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listServicesForSites"`

	SessionId string `xml:"sessionId,omitempty"`

	Date time.Time `xml:"date,omitempty"`

	SenderSiteId int64 `xml:"senderSiteId,omitempty"`

	ReceiverSiteId int64 `xml:"receiverSiteId,omitempty"`

	SenderCountryId int64 `xml:"senderCountryId,omitempty"`

	SenderPostCode string `xml:"senderPostCode,omitempty"`

	ReceiverCountryId int64 `xml:"receiverCountryId,omitempty"`

	ReceiverPostCode string `xml:"receiverPostCode,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`

	SenderId int64 `xml:"senderId,omitempty"`

	ReceiverId int64 `xml:"receiverId,omitempty"`

	SenderOfficeId int64 `xml:"senderOfficeId,omitempty"`

	ReceiverOfficeId int64 `xml:"receiverOfficeId,omitempty"`
}

type ListServicesForSitesResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listServicesForSitesResponse"`

	Return_ []*ResultCourierServiceExt `xml:"return,omitempty"`
}

type ResultCourierServiceExt struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultCourierServiceExt"`

	AllowanceBackDocumentsRequest *ComplementaryServiceAllowance `xml:"allowanceBackDocumentsRequest,omitempty"`

	AllowanceBackReceiptRequest *ComplementaryServiceAllowance `xml:"allowanceBackReceiptRequest,omitempty"`

	AllowanceCashOnDelivery *ComplementaryServiceAllowance `xml:"allowanceCashOnDelivery,omitempty"`

	AllowanceFixedTimeDelivery *ComplementaryServiceAllowance `xml:"allowanceFixedTimeDelivery,omitempty"`

	AllowanceInsurance *ComplementaryServiceAllowance `xml:"allowanceInsurance,omitempty"`

	AllowanceToBeCalled *ComplementaryServiceAllowance `xml:"allowanceToBeCalled,omitempty"`

	Name string `xml:"name,omitempty"`

	TypeId int64 `xml:"typeId,omitempty"`

	DeliveryDeadline time.Time `xml:"deliveryDeadline,omitempty"`

	CargoType int32 `xml:"cargoType,omitempty"`

	AllowanceDeliveryToFloor *ComplementaryServiceAllowance `xml:"allowanceDeliveryToFloor,omitempty"`

	AllowanceOptionsBeforePayment *ComplementaryServiceAllowance `xml:"allowanceOptionsBeforePayment,omitempty"`

	AllowanceReturnVoucher *ComplementaryServiceAllowance `xml:"allowanceReturnVoucher,omitempty"`

	RequireParcelsData bool `xml:"requireParcelsData,omitempty"`
}

type Calculate struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ calculate"`

	SessionId string `xml:"sessionId,omitempty"`

	Calculation *ParamCalculation `xml:"calculation,omitempty"`
}

type CalculateResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ calculateResponse"`

	Return_ *ResultCalculation `xml:"return,omitempty"`
}

type CreateOrder struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ createOrder"`

	SessionId string `xml:"sessionId,omitempty"`

	Order *ParamOrder `xml:"order,omitempty"`
}

type ParamOrder struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ paramOrder"`

	BillOfLadingsList []int64 `xml:"billOfLadingsList,omitempty"`

	BillOfLadingsToIncludeType int32 `xml:"billOfLadingsToIncludeType,omitempty"`

	ContactName string `xml:"contactName,omitempty"`

	PhoneNumber *ParamPhoneNumber `xml:"phoneNumber,omitempty"`

	PickupDate time.Time `xml:"pickupDate,omitempty"`

	ReadinessTime int16 `xml:"readinessTime,omitempty"`

	WorkingEndTime int16 `xml:"workingEndTime,omitempty"`
}

type CreateOrderResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ createOrderResponse"`

	Return_ []*ResultOrderPickingInfo `xml:"return,omitempty"`
}

type ResultOrderPickingInfo struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultOrderPickingInfo"`

	BillOfLading int64 `xml:"billOfLading,omitempty"`

	ErrorDescriptions []string `xml:"errorDescriptions,omitempty"`
}

type CreatePDFEx struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ createPDFEx"`

	SessionId string `xml:"sessionId,omitempty"`

	Params *ParamPDF `xml:"params,omitempty"`
}

type CreatePDFExResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ createPDFExResponse"`

	Return_ *ResultCreatePDFEx `xml:"return,omitempty"`
}

type ResultCreatePDFEx struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ resultCreatePDFEx"`

	PdfBytes []byte `xml:"pdfBytes,omitempty"`

	HubId int64 `xml:"hubId,omitempty"`

	OfficeId int64 `xml:"officeId,omitempty"`

	DeadlineDay int32 `xml:"deadlineDay,omitempty"`

	DeadlineMonth int32 `xml:"deadlineMonth,omitempty"`

	TourId int64 `xml:"tourId,omitempty"`

	FullBarcode string `xml:"fullBarcode,omitempty"`

	ExportPriority int32 `xml:"exportPriority,omitempty"`
}

type SearchClients struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ searchClients"`

	SessionId string `xml:"sessionId,omitempty"`

	ClientQuery *ParamClientSearch `xml:"clientQuery,omitempty"`
}

type ParamClientSearch struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ paramClientSearch"`

	ClientId int64 `xml:"clientId,omitempty"`

	ClientName string `xml:"clientName,omitempty"`

	ObjectName string `xml:"objectName,omitempty"`

	Phone string `xml:"phone,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	UserDefTag string `xml:"userDefTag,omitempty"`
}

type SearchClientsResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ searchClientsResponse"`

	Return_ []*ResultClientData `xml:"return,omitempty"`
}

type TrackPickingEx struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ trackPickingEx"`

	SessionId string `xml:"sessionId,omitempty"`

	BillOfLading string `xml:"billOfLading,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`

	ReturnOnlyLastOperation bool `xml:"returnOnlyLastOperation,omitempty"`
}

type TrackPickingExResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ trackPickingExResponse"`

	Return_ []*ResultTrackPickingEx `xml:"return,omitempty"`
}

type ListCountries struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listCountries"`

	SessionId string `xml:"sessionId,omitempty"`

	Name string `xml:"name,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ListCountriesResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listCountriesResponse"`

	Return_ []*ResultCountry `xml:"return,omitempty"`
}

type GetMicroregionId struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getMicroregionId"`

	SessionId string `xml:"sessionId,omitempty"`

	CoordX float64 `xml:"coordX,omitempty"`

	CoordY float64 `xml:"coordY,omitempty"`
}

type GetMicroregionIdResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ getMicroregionIdResponse"`

	Return_ int64 `xml:"return,omitempty"`
}

type ListBlocks struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listBlocks"`

	SessionId string `xml:"sessionId,omitempty"`

	Name string `xml:"name,omitempty"`

	SiteId int64 `xml:"siteId,omitempty"`

	Language *ParamLanguage `xml:"language,omitempty"`
}

type ListBlocksResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ listBlocksResponse"`

	Return_ []string `xml:"return,omitempty"`
}

type SearchPickingsByRefNumber struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ searchPickingsByRefNumber"`

	SessionId string `xml:"sessionId,omitempty"`

	Params *ParamSearchByRefNum `xml:"params,omitempty"`
}

type ParamSearchByRefNum struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ paramSearchByRefNum"`

	DateFrom time.Time `xml:"dateFrom,omitempty"`

	DateTo time.Time `xml:"dateTo,omitempty"`

	ReferenceNumber string `xml:"referenceNumber,omitempty"`

	SearchInField int32 `xml:"searchInField,omitempty"`

	IncludeReturnBols bool `xml:"includeReturnBols,omitempty"`
}

type SearchPickingsByRefNumberResponse struct {
	XMLName xml.Name `xml:"http://ver01.eps.speedy.sirma.com/ searchPickingsByRefNumberResponse"`

	Return_ []int64 `xml:"return,omitempty"`
}

type EPSProvider interface {

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListServices(request *ListServices) (*ListServicesResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	Calculate(request *Calculate) (*CalculateResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	ConvertToWin1251(request *ConvertToWin1251) (*ConvertToWin1251Response, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListAllSites(request *ListAllSites) (*ListAllSitesResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	CreatePDFEx(request *CreatePDFEx) (*CreatePDFExResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListBlocks(request *ListBlocks) (*ListBlocksResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	ValidateAddress(request *ValidateAddress) (*ValidateAddressResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListSites(request *ListSites) (*ListSitesResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListQuarters(request *ListQuarters) (*ListQuartersResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListSitesEx(request *ListSitesEx) (*ListSitesExResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListCountries(request *ListCountries) (*ListCountriesResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	SearchClients(request *SearchClients) (*SearchClientsResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListOffices(request *ListOffices) (*ListOfficesResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	TrackPickingEx(request *TrackPickingEx) (*TrackPickingExResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	AddressSearch(request *AddressSearch) (*AddressSearchResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	TrackPicking(request *TrackPicking) (*TrackPickingResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	SerializeAddress(request *SerializeAddress) (*SerializeAddressResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetClientById(request *GetClientById) (*GetClientByIdResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ValidatePostCode(request *ValidatePostCode) (*ValidatePostCodeResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	AddParcel(request *AddParcel) (*AddParcelResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetSiteById(request *GetSiteById) (*GetSiteByIdResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListCountriesEx(request *ListCountriesEx) (*ListCountriesExResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListQuarterTypes(request *ListQuarterTypes) (*ListQuarterTypesResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListStreets(request *ListStreets) (*ListStreetsResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	CalculatePicking(request *CalculatePicking) (*CalculatePickingResponse, error)

	IsSessionActive(request *IsSessionActive) (*IsSessionActiveResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListStreetTypes(request *ListStreetTypes) (*ListStreetTypesResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListOfficesEx(request *ListOfficesEx) (*ListOfficesExResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	CreateOrder(request *CreateOrder) (*CreateOrderResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	GetWeightInterval(request *GetWeightInterval) (*GetWeightIntervalResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetAddressNomenclature(request *GetAddressNomenclature) (*GetAddressNomenclatureResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	CreatePDF(request *CreatePDF) (*CreatePDFResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetSitesByAddrNomenType(request *GetSitesByAddrNomenType) (*GetSitesByAddrNomenTypeResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListCommonObjects(request *ListCommonObjects) (*ListCommonObjectsResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListContractClients(request *ListContractClients) (*ListContractClientsResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	CreateBillOfLading(request *CreateBillOfLading) (*CreateBillOfLadingResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	CreateBillOfLadingPDF(request *CreateBillOfLadingPDF) (*CreateBillOfLadingPDFResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	UpdateBillOfLading(request *UpdateBillOfLading) (*UpdateBillOfLadingResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	CreateCustomTravelLabelPDFType1(request *CreateCustomTravelLabelPDFType1) (*CreateCustomTravelLabelPDFType1Response, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetRoutingLabelInfo(request *GetRoutingLabelInfo) (*GetRoutingLabelInfoResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	CalculateMultipleServices(request *CalculateMultipleServices) (*CalculateMultipleServicesResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	InvalidatePicking(request *InvalidatePicking) (*InvalidatePickingResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	FinalizeBillOfLadingCreation(request *FinalizeBillOfLadingCreation) (*FinalizeBillOfLadingCreationResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetAllowedDaysForTaking(request *GetAllowedDaysForTaking) (*GetAllowedDaysForTakingResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetPickingParcels(request *GetPickingParcels) (*GetPickingParcelsResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	TrackParcelMultiple(request *TrackParcelMultiple) (*TrackParcelMultipleResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	SearchPickingsByRefNumber(request *SearchPickingsByRefNumber) (*SearchPickingsByRefNumberResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListStates(request *ListStates) (*ListStatesResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	TrackParcel(request *TrackParcel) (*TrackParcelResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetPickingExtendedInfo(request *GetPickingExtendedInfo) (*GetPickingExtendedInfoResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListSpecialDeliveryRequirements(request *ListSpecialDeliveryRequirements) (*ListSpecialDeliveryRequirementsResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	DeserializeAddress(request *DeserializeAddress) (*DeserializeAddressResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException
	//   - PickingValidationException

	MakeAddressString(request *MakeAddressString) (*MakeAddressStringResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetAdditionalUserParams(request *GetAdditionalUserParams) (*GetAdditionalUserParamsResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetPickingDeliveryInfo(request *GetPickingDeliveryInfo) (*GetPickingDeliveryInfoResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	SearchSecondaryPickings(request *SearchSecondaryPickings) (*SearchSecondaryPickingsResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	MapForeignBarcode(request *MapForeignBarcode) (*MapForeignBarcodeResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	ListServicesForSites(request *ListServicesForSites) (*ListServicesForSitesResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetStateById(request *GetStateById) (*GetStateByIdResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidSessionException

	GetMicroregionId(request *GetMicroregionId) (*GetMicroregionIdResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidUsernameOrPasswordException
	//   - NoUserPermissionsException
	//   - AccountLockedException

	Login(request *Login) (*LoginResponse, error)
}

type ePSProvider struct {
	client *Client
}

func NewEPSProvider(client *Client) EPSProvider {
	return &ePSProvider{
		client: client,
	}
}

func (service *ePSProvider) ListServices(request *ListServices) (*ListServicesResponse, error) {
	response := new(ListServicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) Calculate(request *Calculate) (*CalculateResponse, error) {
	response := new(CalculateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ConvertToWin1251(request *ConvertToWin1251) (*ConvertToWin1251Response, error) {
	response := new(ConvertToWin1251Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListAllSites(request *ListAllSites) (*ListAllSitesResponse, error) {
	response := new(ListAllSitesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) CreatePDFEx(request *CreatePDFEx) (*CreatePDFExResponse, error) {
	response := new(CreatePDFExResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListBlocks(request *ListBlocks) (*ListBlocksResponse, error) {
	response := new(ListBlocksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ValidateAddress(request *ValidateAddress) (*ValidateAddressResponse, error) {
	response := new(ValidateAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListSites(request *ListSites) (*ListSitesResponse, error) {
	response := new(ListSitesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListQuarters(request *ListQuarters) (*ListQuartersResponse, error) {
	response := new(ListQuartersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListSitesEx(request *ListSitesEx) (*ListSitesExResponse, error) {
	response := new(ListSitesExResponse)
	request.XMLNS = "http://ver01.eps.speedy.sirma.com/"
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListCountries(request *ListCountries) (*ListCountriesResponse, error) {
	response := new(ListCountriesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) SearchClients(request *SearchClients) (*SearchClientsResponse, error) {
	response := new(SearchClientsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListOffices(request *ListOffices) (*ListOfficesResponse, error) {
	response := new(ListOfficesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) TrackPickingEx(request *TrackPickingEx) (*TrackPickingExResponse, error) {
	response := new(TrackPickingExResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) AddressSearch(request *AddressSearch) (*AddressSearchResponse, error) {
	response := new(AddressSearchResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) TrackPicking(request *TrackPicking) (*TrackPickingResponse, error) {
	response := new(TrackPickingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) SerializeAddress(request *SerializeAddress) (*SerializeAddressResponse, error) {
	response := new(SerializeAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetClientById(request *GetClientById) (*GetClientByIdResponse, error) {
	response := new(GetClientByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ValidatePostCode(request *ValidatePostCode) (*ValidatePostCodeResponse, error) {
	response := new(ValidatePostCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) AddParcel(request *AddParcel) (*AddParcelResponse, error) {
	response := new(AddParcelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetSiteById(request *GetSiteById) (*GetSiteByIdResponse, error) {
	response := new(GetSiteByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListCountriesEx(request *ListCountriesEx) (*ListCountriesExResponse, error) {
	response := new(ListCountriesExResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListQuarterTypes(request *ListQuarterTypes) (*ListQuarterTypesResponse, error) {
	response := new(ListQuarterTypesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListStreets(request *ListStreets) (*ListStreetsResponse, error) {
	response := new(ListStreetsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) CalculatePicking(request *CalculatePicking) (*CalculatePickingResponse, error) {
	response := new(CalculatePickingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) IsSessionActive(request *IsSessionActive) (*IsSessionActiveResponse, error) {
	response := new(IsSessionActiveResponse)
	request.XMLNS = "http://ver01.eps.speedy.sirma.com/"
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListStreetTypes(request *ListStreetTypes) (*ListStreetTypesResponse, error) {
	response := new(ListStreetTypesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListOfficesEx(request *ListOfficesEx) (*ListOfficesExResponse, error) {
	response := new(ListOfficesExResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) CreateOrder(request *CreateOrder) (*CreateOrderResponse, error) {
	response := new(CreateOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetWeightInterval(request *GetWeightInterval) (*GetWeightIntervalResponse, error) {
	response := new(GetWeightIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetAddressNomenclature(request *GetAddressNomenclature) (*GetAddressNomenclatureResponse, error) {
	response := new(GetAddressNomenclatureResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) CreatePDF(request *CreatePDF) (*CreatePDFResponse, error) {
	response := new(CreatePDFResponse)
	request.XMLNS = "http://ver01.eps.speedy.sirma.com/"
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetSitesByAddrNomenType(request *GetSitesByAddrNomenType) (*GetSitesByAddrNomenTypeResponse, error) {
	response := new(GetSitesByAddrNomenTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListCommonObjects(request *ListCommonObjects) (*ListCommonObjectsResponse, error) {
	response := new(ListCommonObjectsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListContractClients(request *ListContractClients) (*ListContractClientsResponse, error) {
	response := new(ListContractClientsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) CreateBillOfLading(request *CreateBillOfLading) (*CreateBillOfLadingResponse, error) {
	response := new(CreateBillOfLadingResponse)
	request.XMLNS = "http://ver01.eps.speedy.sirma.com/"
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) CreateBillOfLadingPDF(request *CreateBillOfLadingPDF) (*CreateBillOfLadingPDFResponse, error) {
	response := new(CreateBillOfLadingPDFResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) UpdateBillOfLading(request *UpdateBillOfLading) (*UpdateBillOfLadingResponse, error) {
	response := new(UpdateBillOfLadingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) CreateCustomTravelLabelPDFType1(request *CreateCustomTravelLabelPDFType1) (*CreateCustomTravelLabelPDFType1Response, error) {
	response := new(CreateCustomTravelLabelPDFType1Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetRoutingLabelInfo(request *GetRoutingLabelInfo) (*GetRoutingLabelInfoResponse, error) {
	response := new(GetRoutingLabelInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) CalculateMultipleServices(request *CalculateMultipleServices) (*CalculateMultipleServicesResponse, error) {
	response := new(CalculateMultipleServicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) InvalidatePicking(request *InvalidatePicking) (*InvalidatePickingResponse, error) {
	response := new(InvalidatePickingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) FinalizeBillOfLadingCreation(request *FinalizeBillOfLadingCreation) (*FinalizeBillOfLadingCreationResponse, error) {
	response := new(FinalizeBillOfLadingCreationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetAllowedDaysForTaking(request *GetAllowedDaysForTaking) (*GetAllowedDaysForTakingResponse, error) {
	response := new(GetAllowedDaysForTakingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetPickingParcels(request *GetPickingParcels) (*GetPickingParcelsResponse, error) {
	response := new(GetPickingParcelsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) TrackParcelMultiple(request *TrackParcelMultiple) (*TrackParcelMultipleResponse, error) {
	response := new(TrackParcelMultipleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) SearchPickingsByRefNumber(request *SearchPickingsByRefNumber) (*SearchPickingsByRefNumberResponse, error) {
	response := new(SearchPickingsByRefNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListStates(request *ListStates) (*ListStatesResponse, error) {
	response := new(ListStatesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) TrackParcel(request *TrackParcel) (*TrackParcelResponse, error) {
	response := new(TrackParcelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetPickingExtendedInfo(request *GetPickingExtendedInfo) (*GetPickingExtendedInfoResponse, error) {
	response := new(GetPickingExtendedInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListSpecialDeliveryRequirements(request *ListSpecialDeliveryRequirements) (*ListSpecialDeliveryRequirementsResponse, error) {
	response := new(ListSpecialDeliveryRequirementsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) DeserializeAddress(request *DeserializeAddress) (*DeserializeAddressResponse, error) {
	response := new(DeserializeAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) MakeAddressString(request *MakeAddressString) (*MakeAddressStringResponse, error) {
	response := new(MakeAddressStringResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetAdditionalUserParams(request *GetAdditionalUserParams) (*GetAdditionalUserParamsResponse, error) {
	response := new(GetAdditionalUserParamsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetPickingDeliveryInfo(request *GetPickingDeliveryInfo) (*GetPickingDeliveryInfoResponse, error) {
	response := new(GetPickingDeliveryInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) SearchSecondaryPickings(request *SearchSecondaryPickings) (*SearchSecondaryPickingsResponse, error) {
	response := new(SearchSecondaryPickingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) MapForeignBarcode(request *MapForeignBarcode) (*MapForeignBarcodeResponse, error) {
	response := new(MapForeignBarcodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) ListServicesForSites(request *ListServicesForSites) (*ListServicesForSitesResponse, error) {
	response := new(ListServicesForSitesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetStateById(request *GetStateById) (*GetStateByIdResponse, error) {
	response := new(GetStateByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) GetMicroregionId(request *GetMicroregionId) (*GetMicroregionIdResponse, error) {
	response := new(GetMicroregionIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ePSProvider) Login(request *Login) (*LoginResponse, error) {
	response := new(LoginResponse)
	request.XMLNS = "http://ver01.eps.speedy.sirma.com/"
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
