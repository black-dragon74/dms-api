//
//	This file stores the constants required by this server
//

package utils

import "github.com/black-dragon74/dms-api/types"

const (
	// Base URLs
	DmsURL        = "https://dms.jaipur.manipal.edu/"
	LoginURL      = DmsURL + "Loginform.aspx"
	ProfileURL    = DmsURL + "studentprofile.aspx"
	AttendanceURL = DmsURL + "StudentAttendanceSummary.aspx"
	ResultURL     = DmsURL + "ExamResultDisplay.aspx"
	EventURL      = DmsURL + "studenthomepage.aspx"
	FeeURL        = DmsURL + "feedetailmuj.aspx"

	// Test URLS
	PostManPost = "https://postman-echo.com/post"

	// Cookies
	SessionCookie = "ASP.NET_SessionId"

	// Different HTML selector IDs
	selectorPrefix          = "#ContentPlaceHolder1_"
	IdForName               = selectorPrefix + "txtStudentName"
	IdForPgme               = selectorPrefix + "txtProgramBranch"
	IdForRegno              = selectorPrefix + "txtApplicationNumber"
	IdForAcadYear           = selectorPrefix + "txtAcademicYear"
	IdForSeniorSecMarks     = selectorPrefix + "grdQualification_lblpercent_0"
	IdForSecMarks           = selectorPrefix + "grdQualification_lblpercent_1"
	IdForMarksTable         = selectorPrefix + "grdQualification"
	IdForParentMobile       = selectorPrefix + "txtParentMobile"
	IdForParentEmail        = selectorPrefix + "txtParentPesentEmail"
	IdForParentMother       = selectorPrefix + "txtPMotherName"
	IdForParentFather       = selectorPrefix + "txtPFatherName"
	IdForParentEmergency    = selectorPrefix + "txtEmergencyContactNo"
	IdForAttendanceTable    = selectorPrefix + "grdAttendanceDetails"
	IdForResultTable        = selectorPrefix + "GrdExamResults"
	IdForCgpaTable          = selectorPrefix + "GrdExamResults2"
	IdForInternalsTable     = selectorPrefix + "GrdExamResults"
	IdForAnnouncementsTable = selectorPrefix + "GridAnnounce"
	IdForEventsTable        = selectorPrefix + "grdEventDetails"
	IdForPaidFeeTable       = selectorPrefix + "grdPaidFeesDetails"
	IdForUnpaidFeeTable     = selectorPrefix + "grdFeeDetails"
	IdForCaptcha            = "#imgCaptcha"
	IdForCredentialsError   = "span#labelerror"
	IdForCaptchaError       = "span#lblErrorMsg"

	// Header Values
	HeaderAccept                  = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	HeaderAcceptLanguage          = "en-US,en;q=0.9"
	HeaderCacheControl            = "max-age=0"
	HeaderConnection              = "keep-alive"
	HeaderContentType             = "application/x-www-form-urlencoded"
	HeaderDNT                     = "1"
	HeaderHost                    = "dms.jaipur.manipal.edu"
	HeaderOrigin                  = DmsURL
	HeaderReferer                 = DmsURL + LoginURL
	HeaderSecFetchDest            = "document"
	HeaderSecFetchMode            = "navigate"
	HeaderSecFetchSite            = "same-origin"
	HeaderSecFetchUser            = "?1"
	HeaderUpgradeInsecureRequests = "1"
	HeaderUserAgent               = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"

	// ASP.NET Specific Values
	ASPLoginEventTarget        = "hprlnkStduent"
	ASPLoginEventValidation    = "/wEdAAep4CnDEBO+DK7ogDE9yIJSdR78oILfrSzgm87C/a1IYZxpWckI3qdmfEJVCu2f5cGK8hF2GuqB1EkPPnfRI0IzmX+TFKdoKJU/yfSArg3MIyhcPbspsZuAvIcHzSMoo5oTEvVQ5UbiG8J6VK1P0lg47nw9ow/C86maq2rS+tJpbynFcaYfcgcmP+IGmR/HOn4="
	ASPLoginViewState          = "/wEPDwUKLTQxMTUwMTQyNw9kFgICAw9kFgICAw9kFgQCCQ9kFgJmD2QWAgIBDw8WAh4ISW1hZ2VVcmwFJ0dlbmVyYXRlQ2FwdGNoYS5hc3B4PzYzNzM0MjY4MjUyNjcyOTc5NGRkAhMPDxYCHgdWaXNpYmxlaGRkZFbWHdkndpKMZ2FmRaLaq+f+wL+6ZH//WlWjnAiYNA/W"
	ASPLoginViewStateGenerator = "6ED0046F"

	// Query params, used to avoid usage of magic strings
	VarSessionID = "sessionid"
	VarUserName  = "username"
	VarPassword  = "password"
	VarCaptcha   = "captcha"

	// Various Errors used in the API
	ErrorLoginFailed   = types.APIError("Login failed") // TODO: Must be all lower case
	ErrorTableNotFound = types.APIError("table not found")
	ErrorTableNoBody   = types.APIError("table has no body")
	ErrorTableNoHeader = types.APIError("table has no header")
	ErrorTableNoData   = types.APIError("table has no data")
	ErrorNoAttendance   = types.APIError("attendance not available")
)

// Slices are not thread safe in Go. There is no mutex lock
// on this var as this is only set once. Also, writing to any of the
// slice(s) defined here is a human error and should be avoided at all costs.

// SliceSessionID is a slice containing one element with value VarSessionID
var SliceSessionID = []string{VarSessionID}
