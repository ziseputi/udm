package udm_message

type Event int

const (
	EventGenerateAuthData Event = iota
	EventConfirmAuth
	EventGetAmData
	EventGetIdTranslationResult
	EventInfo
	EventGetSupi
	EventGetSharedData
	EventGetSmData
	EventGetNssai
	EventGetSmfSelectData
	EventGetSmsMngData
	EventGetSmsData
	EventSubscribeToSharedData
	EventSubscribe
	EventUnsubscribeForSharedData
	EventUnsubscribe
	EventModify
	EventModifyForSharedData
	EventGetTraceData
	EventGetUeContextInSmfData
	EventGetUeContextInSmsfData
	EventCreateEeSubscription
	EventDeleteEeSubscription
	EventUpdateEeSubscription
	EventGetAmf3gppAccess
	EventGetAmfNon3gppAccess
	EventRegistrationAmf3gppAccess
	EventRegisterAmfNon3gppAccess
	EventUpdateAmf3gppAccess
	EventUpdateAmfNon3gppAccess
	EventDeregistrationSmfRegistrations
	EventRegistrationSmfRegistrations
	EventGetSmsf3gppAccess
	EventDeregistrationSmsf3gppAccess
	EventDeregistrationSmsfNon3gppAccess
	EventGetSmsfNon3gppAccess
	EventUpdateSMSFReg3GPP
	EventRegistrationSmsfNon3gppAccess
	EventUpdate
	EventDataChangeNotificationToNF
)