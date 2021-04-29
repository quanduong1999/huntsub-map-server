package cons

const ENV_HOME_DB = "DID_home_DB"
const ENV_LUCKY_DB = "DID_LUCKY_DB"

const URL_VERIFY = "https://huntsub.com/api/user/verify?id="

const URL_RENEWPASS = "https://huntsub.com/api/user/renew?id="
const M_URL = "https://huntsub.com/v/s?id="
const LINK = "https://huntsub.com/auth/login"

const REDIRECT_VIEW = "http://huntsub.com/api/scan/view?id="

const DIR_DOWNLOAD_EXCEL = "https://huntsub.com/"

// const DIR_DOWNLOAD_EXCEL = "http://localhost:3000/"

//Email Account:
const Email = "digital.identity.did@gmail.com"
const PassWord = "Digital123"

/*Define priority of Calendar*/
const (
	Level_0 = 0 // Status has just creating
	Level_1 = 1 // Status is going to meeting and distance about 30 minues
	Level_2 = 2 // Status has done and dont feedback for each
	Level_3 = 3 // Status has feedback done
)

/*Define kind of Notification*/
const (
	Like            = "like"
	Comment         = "comment"
	Share           = "share"
	Calendar        = "calendar"
	ConfirmCalendar = "confirm_calendar"
)

/**************** DEFINE ALL KEY OF LANGUAGE **************/
const (
	VERIFY_ACCOUNT_HUNTSUB_NETWORK_SYSTEM       = "verify-account-huntsub-network-system"
	GREATING                                    = "greating"
	THANK_YOU_FOR_REGISTER_On_MY_SYSTEM         = "thank-you-for-register-on-my-system"
	PLEASE_VERIFY                               = "please-verify"
	SINGATURE                                   = "signature"
	GET_BACK_PASSWORK                           = "get-back-password"
	NEW_PASSWORD                                = "new-password"
	WE_HAVE_GOT_A_REQUIREMENT_GET_BACK_PASSWORD = "we-have-got-a-requirement-get-back-password"
	PLAESE_VERIFY_BY_LINK_BELOW                 = "please-verify-by-link-below"
	SYSTEM_WILL_BE_SENT_A_NEW_PASSWORD_TO_YOU   = "system-will-be-sent-a-new-password-to-you"
)
