package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-business-partner-creates-customer/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) (*General, error) {
	pm := &responses.General{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	general := &General{
		BusinessPartner:               data.BusinessPartner,
		Customer:                      data.Customer,
		Supplier:                      data.Supplier,
		AcademicTitle:                 data.AcademicTitle,
		AuthorizationGroup:            data.AuthorizationGroup,
		BusinessPartnerCategory:       data.BusinessPartnerCategory,
		BusinessPartnerFullName:       data.BusinessPartnerFullName,
		BusinessPartnerGrouping:       data.BusinessPartnerGrouping,
		BusinessPartnerName:           data.BusinessPartnerName,
		CorrespondenceLanguage:        data.CorrespondenceLanguage,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		FirstName:                     data.FirstName,
		Industry:                      data.Industry,
		IsFemale:                      data.IsFemale,
		IsMale:                        data.IsMale,
		IsNaturalPerson:               data.IsNaturalPerson,
		IsSexUnknown:                  data.IsSexUnknown,
		GenderCodeName:                data.GenderCodeName,
		Language:                      data.Language,
		LastChangeDate:                data.LastChangeDate,
		LastChangeTime:                data.LastChangeTime,
		LastName:                      data.LastName,
		OrganizationBPName1:           data.OrganizationBPName1,
		OrganizationBPName2:           data.OrganizationBPName2,
		OrganizationBPName3:           data.OrganizationBPName3,
		OrganizationBPName4:           data.OrganizationBPName4,
		OrganizationFoundationDate:    data.OrganizationFoundationDate,
		OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
		SearchTerm1:                   data.SearchTerm1,
		SearchTerm2:                   data.SearchTerm2,
		AdditionalLastName:            data.AdditionalLastName,
		BirthDate:                     data.BirthDate,
		BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
		BusinessPartnerDeathDate:      data.BusinessPartnerDeathDate,
		BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
		BusinessPartnerType:           data.BusinessPartnerType,
		GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
		IndependentAddressID:          data.IndependentAddressID,
		MiddleName:                    data.MiddleName,
		NameCountry:                   data.NameCountry,
		PersonFullName:                data.PersonFullName,
		PersonNumber:                  data.PersonNumber,
		IsMarkedForArchiving:          data.IsMarkedForArchiving,
		BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
		TradingPartner:                data.TradingPartner,
	}

	return general, nil
}

func ConvertToRole(raw []byte, l *logger.Logger) (*Role, error) {
	pm := &responses.Role{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Role. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	role := &Role{
		BusinessPartner:     data.BusinessPartner,
		BusinessPartnerRole: data.BusinessPartnerRole,
		ValidFrom:           data.ValidFrom,
		ValidTo:             data.ValidTo,
	}

	return role, nil
}

func ConvertToAddress(raw []byte, l *logger.Logger) (*Address, error) {
	pm := &responses.Address{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Address. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	address := &Address{
		BusinessPartner:   data.BusinessPartner,
		AddressID:         data.AddressID,
		ValidityStartDate: data.ValidityStartDate,
		ValidityEndDate:   data.ValidityEndDate,
		Country:           data.Country,
		Region:            data.Region,
		StreetName:        data.StreetName,
		CityName:          data.CityName,
		PostalCode:        data.PostalCode,
		Language:          data.Language,
	}

	return address, nil
}

func ConvertToBank(raw []byte, l *logger.Logger) (*Bank, error) {
	pm := &responses.Bank{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Bank. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	bank := &Bank{
		BusinessPartner:          data.BusinessPartner,
		BankIdentification:       data.BankIdentification,
		BankCountryKey:           data.BankCountryKey,
		BankName:                 data.BankName,
		BankNumber:               data.BankNumber,
		SWIFTCode:                data.SWIFTCode,
		BankControlKey:           data.BankControlKey,
		BankAccountHolderName:    data.BankAccountHolderName,
		BankAccountName:          data.BankAccountName,
		ValidityStartDate:        data.ValidityStartDate,
		ValidityEndDate:          data.ValidityEndDate,
		IBAN:                     data.IBAN,
		IBANValidityStartDate:    data.IBANValidityStartDate,
		BankAccount:              data.BankAccount,
		BankAccountReferenceText: data.BankAccountReferenceText,
		CollectionAuthInd:        data.CollectionAuthInd,
		CityName:                 data.CityName,
		AuthorizationGroup:       data.AuthorizationGroup,
	}

	return bank, nil
}

func ConvertToCustomer(raw []byte, l *logger.Logger) (*Customer, error) {
	pm := &responses.Customer{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Customer. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	customer := &Customer{
		Customer:                    data.Customer,
		AuthorizationGroup:          data.AuthorizationGroup,
		BillingIsBlockedForCustomer: data.BillingIsBlockedForCustomer,
		CreationDate:                data.CreationDate,
		CustomerAccountGroup:        data.CustomerAccountGroup,
		CustomerClassification:      data.CustomerClassification,
		CustomerFullName:            data.CustomerFullName,
		CustomerName:                data.CustomerName,
		DeliveryIsBlocked:           data.DeliveryIsBlocked,
		OrderIsBlockedForCustomer:   data.OrderIsBlockedForCustomer,
		PostingIsBlocked:            data.PostingIsBlocked,
		Supplier:                    data.Supplier,
		CustomerCorporateGroup:      data.CustomerCorporateGroup,
		Industry:                    data.Industry,
		TaxNumber1:                  data.TaxNumber1,
		DeletionIndicator:           data.DeletionIndicator,
		CityCode:                    data.CityCode,
		County:                      data.County,
	}

	return customer, nil
}

func ConvertToSalesArea(raw []byte, l *logger.Logger) (*SalesArea, error) {
	pm := &responses.SalesArea{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesArea. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	salesArea := &SalesArea{
		Customer:                       data.Customer,
		SalesOrganization:              data.SalesOrganization,
		DistributionChannel:            data.DistributionChannel,
		Division:                       data.Division,
		CompleteDeliveryIsDefined:      data.CompleteDeliveryIsDefined,
		Currency:                       data.Currency,
		CustomerAccountAssignmentGroup: data.CustomerAccountAssignmentGroup,
		CustomerPaymentTerms:           data.CustomerPaymentTerms,
		CustomerPriceGroup:             data.CustomerPriceGroup,
		CustomerPricingProcedure:       data.CustomerPricingProcedure,
		DeliveryPriority:               data.DeliveryPriority,
		IncotermsClassification:        data.IncotermsClassification,
		InvoiceDate:                    data.InvoiceDate,
		OrderCombinationIsAllowed:      data.OrderCombinationIsAllowed,
		PartialDeliveryIsAllowed:       data.PartialDeliveryIsAllowed,
		PriceListType:                  data.PriceListType,
		SalesGroup:                     data.SalesGroup,
		SalesOffice:                    data.SalesOffice,
		ShippingCondition:              data.ShippingCondition,
		SupplyingPlant:                 data.SupplyingPlant,
		SalesDistrict:                  data.SalesDistrict,
		InvoiceListSchedule:            data.InvoiceListSchedule,
		ExchangeRateType:               data.ExchangeRateType,
		OrderIsBlockedForCustomer:      data.OrderIsBlockedForCustomer,
		DeliveryIsBlockedForCustomer:   data.DeliveryIsBlockedForCustomer,
		BillingIsBlockedForCustomer:    data.BillingIsBlockedForCustomer,
		DeletionIndicator:              data.DeletionIndicator,
	}

	return salesArea, nil
}

func ConvertToPartnerFunction(raw []byte, l *logger.Logger) (*PartnerFunction, error) {
	pm := &responses.PartnerFunction{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerFunction. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	partnerFunction := &PartnerFunction{
		Customer:                   data.Customer,
		SalesOrganization:          data.SalesOrganization,
		DistributionChannel:        data.DistributionChannel,
		Division:                   data.Division,
		PartnerCounter:             data.PartnerCounter,
		PartnerFunction:            data.PartnerFunction,
		BPCustomerNumber:           data.BPCustomerNumber,
		CustomerPartnerDescription: data.CustomerPartnerDescription,
		DefaultPartner:             data.DefaultPartner,
		Supplier:                   data.Supplier,
		AuthorizationGroup:         data.AuthorizationGroup,
	}

	return partnerFunction, nil
}

func ConvertToCompany(raw []byte, l *logger.Logger) (*Company, error) {
	pm := &responses.Company{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Company. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	company := &Company{
		Customer:                       data.Customer,
		CompanyCode:                    data.CompanyCode,
		APARToleranceGroup:             data.APARToleranceGroup,
		CustomerSupplierClearingIsUsed: data.CustomerSupplierClearingIsUsed,
		HouseBank:                      data.HouseBank,
		PaymentMethodsList:             data.PaymentMethodsList,
		PaymentTerms:                   data.PaymentTerms,
		ReconciliationAccount:          data.ReconciliationAccount,
		DeletionIndicator:              data.DeletionIndicator,
	}

	return company, nil
}
