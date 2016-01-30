package betting

import (
	"errors"
)

var ErrUnknownLoginStatus = errors.New("Unknown loginStatus value")

type LoginStatus int

const (
	LS_SUCCESS                                 LoginStatus = iota
	LS_INVALID_USERNAME_OR_PASSWORD                        // the username or password are invalid
	LS_ACCOUNT_NOW_LOCKED                                  // the account was just locked
	LS_ACCOUNT_ALREADY_LOCKED                              // the account is already locked
	LS_PENDING_AUTH                                        // pending authentication
	LS_TELBET_TERMS_CONDITIONS_NA                          // Telbet terms and conditions rejected
	LS_DUPLICATE_CARDS                                     // duplicate cards
	LS_SECURITY_QUESTION_WRONG_3X                          // the user has entered wrong the security question 3 times
	LS_KYC_SUSPEND                                         // KYC suspended
	LS_SUSPENDED                                           // the account is suspended
	LS_CLOSED                                              // the account is closed
	LS_SELF_EXCLUDED                                       // the account has been self excluded
	LS_INVALID_CONNECTIVITY_TO_REGULATOR_DK                // the DK regulator cannot be accessed due to some internal problems in the system behind or in at regulator; timeout cases included.
	LS_NOT_AUTHORIZED_BY_REGULATOR_DK                      // the user identified by the given credentials is not authorized in the DK's jurisdictions due to the regulators' policies. Ex: the user for which this session should be created is not allowed to act(play, bet) in the DK's jurisdiction.
	LS_INVALID_CONNECTIVITY_TO_REGULATOR_IT                // the IT regulator cannot be accessed due to some internal problems in the system behind or in at regulator; timeout cases included.
	LS_NOT_AUTHORIZED_BY_REGULATOR_IT                      // the user identified by the given credentials is not authorized in the IT's jurisdictions due to the regulators' policies. Ex: the user for which this session should be created is not allowed to act(play, bet) in the IT's jurisdiction.
	LS_SECURITY_RESTRICTED_LOCATION                        // the account is restricted due to security concerns
	LS_BETTING_RESTRICTED_LOCATION                         // the account is accessed from a location where betting is restricted
	LS_TRADING_MASTER                                      // Trading Master Account
	LS_TRADING_MASTER_SUSPENDED                            // Suspended Trading Master Account
	LS_AGENT_CLIENT_MASTER                                 // Agent Client Master
	LS_AGENT_CLIENT_MASTER_SUSPENDED                       // Suspended Agent Client Master
	LS_DANISH_AUTHORIZATION_REQUIRED                       // danish authorization required
	LS_SPAIN_MIGRATION_REQUIRED                            // spain migration required
	LS_DENMARK_MIGRATION_REQUIRED                          // denmark migration required
	LS_SPANISH_TERMS_ACCEPTANCE_REQUIRED                   // The latest spanish terms and conditions version must be accepted
	LS_ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED                // The latest italian contract version must be accepted
	LS_CERT_AUTH_REQUIRED                                  // Certificate required or certificate present but could not authenticate with it
	LS_CHANGE_PASSWORD_REQUIRED                            // change password required
	LS_PERSONAL_MESSAGE_REQUIRED                           // personal message required for the user
	LS_INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED             // The latest international terms and conditions must be accepted prior to logging in.
	LS_EMAIL_LOGIN_NOT_ALLOWED                             // This account has not opted in to log in with the email
	LS_MULTIPLE_USERS_WITH_SAME_CREDENTIAL                 // There is more than one account with the same credential
	LS_ACCOUNT_PENDING_PASSWORD_CHANGE                     // The account must undergo password recovery to reactivate
	LS_TEMPORARY_BAN_TOO_MANY_REQUESTS                     // The limit for successful login requests per minute has been exceeded. New login attempts will be banned for 20 minutes
)

var lsItems = []string{
	"SUCCESS",
	"INVALID_USERNAME_OR_PASSWORD",
	"ACCOUNT_NOW_LOCKED",
	"ACCOUNT_ALREADY_LOCKED",
	"PENDING_AUTH",
	"TELBET_TERMS_CONDITIONS_NA",
	"DUPLICATE_CARDS",
	"SECURITY_QUESTION_WRONG_3X",
	"KYC_SUSPEND",
	"SUSPENDED",
	"CLOSED",
	"SELF_EXCLUDED",
	"INVALID_CONNECTIVITY_TO_REGULATOR_DK",
	"NOT_AUTHORIZED_BY_REGULATOR_DK",
	"INVALID_CONNECTIVITY_TO_REGULATOR_IT",
	"NOT_AUTHORIZED_BY_REGULATOR_IT",
	"SECURITY_RESTRICTED_LOCATION",
	"BETTING_RESTRICTED_LOCATION",
	"TRADING_MASTER",
	"TRADING_MASTER_SUSPENDED",
	"AGENT_CLIENT_MASTER",
	"AGENT_CLIENT_MASTER_SUSPENDED",
	"DANISH_AUTHORIZATION_REQUIRED",
	"SPAIN_MIGRATION_REQUIRED",
	"DENMARK_MIGRATION_REQUIRED",
	"SPANISH_TERMS_ACCEPTANCE_REQUIRED",
	"ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED",
	"CERT_AUTH_REQUIRED",
	"CHANGE_PASSWORD_REQUIRED",
	"PERSONAL_MESSAGE_REQUIRED",
	"INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED",
	"EMAIL_LOGIN_NOT_ALLOWED",
	"MULTIPLE_USERS_WITH_SAME_CREDENTIAL",
	"ACCOUNT_PENDING_PASSWORD_CHANGE",
	"TEMPORARY_BAN_TOO_MANY_REQUESTS",
}

var lsMap = map[string]LoginStatus{
	lsItems[LS_SUCCESS]:                                 LS_SUCCESS,
	lsItems[LS_INVALID_USERNAME_OR_PASSWORD]:            LS_INVALID_USERNAME_OR_PASSWORD,
	lsItems[LS_ACCOUNT_NOW_LOCKED]:                      LS_ACCOUNT_NOW_LOCKED,
	lsItems[LS_ACCOUNT_ALREADY_LOCKED]:                  LS_ACCOUNT_ALREADY_LOCKED,
	lsItems[LS_PENDING_AUTH]:                            LS_PENDING_AUTH,
	lsItems[LS_TELBET_TERMS_CONDITIONS_NA]:              LS_TELBET_TERMS_CONDITIONS_NA,
	lsItems[LS_DUPLICATE_CARDS]:                         LS_DUPLICATE_CARDS,
	lsItems[LS_SECURITY_QUESTION_WRONG_3X]:              LS_SECURITY_QUESTION_WRONG_3X,
	lsItems[LS_KYC_SUSPEND]:                             LS_KYC_SUSPEND,
	lsItems[LS_SUSPENDED]:                               LS_SUSPENDED,
	lsItems[LS_CLOSED]:                                  LS_CLOSED,
	lsItems[LS_SELF_EXCLUDED]:                           LS_SELF_EXCLUDED,
	lsItems[LS_INVALID_CONNECTIVITY_TO_REGULATOR_DK]:    LS_INVALID_CONNECTIVITY_TO_REGULATOR_DK,
	lsItems[LS_NOT_AUTHORIZED_BY_REGULATOR_DK]:          LS_NOT_AUTHORIZED_BY_REGULATOR_DK,
	lsItems[LS_INVALID_CONNECTIVITY_TO_REGULATOR_IT]:    LS_INVALID_CONNECTIVITY_TO_REGULATOR_IT,
	lsItems[LS_NOT_AUTHORIZED_BY_REGULATOR_IT]:          LS_NOT_AUTHORIZED_BY_REGULATOR_IT,
	lsItems[LS_SECURITY_RESTRICTED_LOCATION]:            LS_SECURITY_RESTRICTED_LOCATION,
	lsItems[LS_BETTING_RESTRICTED_LOCATION]:             LS_BETTING_RESTRICTED_LOCATION,
	lsItems[LS_TRADING_MASTER]:                          LS_TRADING_MASTER,
	lsItems[LS_TRADING_MASTER_SUSPENDED]:                LS_TRADING_MASTER_SUSPENDED,
	lsItems[LS_AGENT_CLIENT_MASTER]:                     LS_AGENT_CLIENT_MASTER,
	lsItems[LS_AGENT_CLIENT_MASTER_SUSPENDED]:           LS_AGENT_CLIENT_MASTER_SUSPENDED,
	lsItems[LS_DANISH_AUTHORIZATION_REQUIRED]:           LS_DANISH_AUTHORIZATION_REQUIRED,
	lsItems[LS_SPAIN_MIGRATION_REQUIRED]:                LS_SPAIN_MIGRATION_REQUIRED,
	lsItems[LS_DENMARK_MIGRATION_REQUIRED]:              LS_DENMARK_MIGRATION_REQUIRED,
	lsItems[LS_SPANISH_TERMS_ACCEPTANCE_REQUIRED]:       LS_SPANISH_TERMS_ACCEPTANCE_REQUIRED,
	lsItems[LS_ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED]:    LS_ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED,
	lsItems[LS_CERT_AUTH_REQUIRED]:                      LS_CERT_AUTH_REQUIRED,
	lsItems[LS_CHANGE_PASSWORD_REQUIRED]:                LS_CHANGE_PASSWORD_REQUIRED,
	lsItems[LS_PERSONAL_MESSAGE_REQUIRED]:               LS_PERSONAL_MESSAGE_REQUIRED,
	lsItems[LS_INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED]: LS_INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED,
	lsItems[LS_EMAIL_LOGIN_NOT_ALLOWED]:                 LS_EMAIL_LOGIN_NOT_ALLOWED,
	lsItems[LS_MULTIPLE_USERS_WITH_SAME_CREDENTIAL]:     LS_MULTIPLE_USERS_WITH_SAME_CREDENTIAL,
	lsItems[LS_ACCOUNT_PENDING_PASSWORD_CHANGE]:         LS_ACCOUNT_PENDING_PASSWORD_CHANGE,
	lsItems[LS_TEMPORARY_BAN_TOO_MANY_REQUESTS]:         LS_TEMPORARY_BAN_TOO_MANY_REQUESTS,
}

func (code LoginStatus) String() string {
	return lsItems[code]
}

func (code *LoginStatus) Unmarshal(enum string) error {
	val, ok := lsMap[enum]
	if !ok {
		return ErrUnknownLoginStatus
	}

	*code = val

	return nil
}
