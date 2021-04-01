//
//	This file stores the constants required by this server
//

package utils

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
	IdForName               = "#ContentPlaceHolder1_txtStudentName"
	IdForPgme               = "#ContentPlaceHolder1_txtProgramBranch"
	IdForRegno              = "#ContentPlaceHolder1_txtApplicationNumber"
	IdForAcadYear           = "#ContentPlaceHolder1_txtAcademicYear"
	IdForSeniorSecMarks     = "#ContentPlaceHolder1_grdQualification_lblpercent_0"
	IdForSecMarks           = "#ContentPlaceHolder1_grdQualification_lblpercent_1"
	IdForMarksTable         = "#ContentPlaceHolder1_grdQualification"
	IdForParentMobile       = "#ContentPlaceHolder1_txtParentMobile"
	IdForParentEmail        = "#ContentPlaceHolder1_txtParentPesentEmail"
	IdForParentMother       = "#ContentPlaceHolder1_txtPMotherName"
	IdForParentFather       = "#ContentPlaceHolder1_txtPFatherName"
	IdForParentEmergency    = "#ContentPlaceHolder1_txtEmergencyContactNo"
	IdForAttendanceTable    = "#ContentPlaceHolder1_grdAttendanceDetails"
	IdForResultTable        = "#ContentPlaceHolder1_GrdExamResults"
	IdForCgpaTable          = "#ContentPlaceHolder1_GrdExamResults2"
	IdForInternalsTable     = "#ContentPlaceHolder1_GrdExamResults"
	IdForAnnouncementsTable = "#ContentPlaceHolder1_GridAnnounce"
	IdForEventsTable        = "#ContentPlaceHolder1_grdEventDetails"
	IdForPaidFeeTable       = "#ContentPlaceHolder1_grdPaidFeesDetails"
	IdForUnpaidFeeTable     = "#ContentPlaceHolder1_grdFeeDetails"
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
)
