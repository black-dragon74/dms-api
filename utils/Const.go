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

	// Cookies
	SessionCookie = "ASP.NET_SessionId"

	// Different HTML selector IDs
	IdForName               = "ContentPlaceHolder1_txtStudentName"
	IdForPgme               = "ContentPlaceHolder1_txtProgramBranch"
	IdForRegno              = "ContentPlaceHolder1_txtApplicationNumber"
	IdForAcadYear           = "ContentPlaceHolder1_txtAcademicYear"
	IdForSeniorSecMarks     = "ContentPlaceHolder1_grdQualification_lblpercent_0"
	IdForSecMarks           = "ContentPlaceHolder1_grdQualification_lblpercent_1"
	IdForMarksTable         = "ContentPlaceHolder1_grdQualification"
	IdForParentMobile       = "ContentPlaceHolder1_txtParentMobile"
	IdForParentEmail        = "ContentPlaceHolder1_txtParentPesentEmail"
	IdForParentMother       = "ContentPlaceHolder1_txtPMotherName"
	IdForParentFather       = "ContentPlaceHolder1_txtPFatherName"
	IdForParentEmergency    = "ContentPlaceHolder1_txtEmergencyContactNo"
	IdForAttendanceTable    = "ContentPlaceHolder1_grdAttendanceDetails"
	IdForResultTable        = "ContentPlaceHolder1_GrdExamResults"
	IdForCgpaTable          = "ContentPlaceHolder1_GrdExamResults2"
	IdForInternalsTable     = "ContentPlaceHolder1_GrdExamResults"
	IdForAnnouncementsTable = "ContentPlaceHolder1_GridAnnounce"
	IdForEventsTable        = "ContentPlaceHolder1_grdEventDetails"
	IdForPaidFeeTable       = "ContentPlaceHolder1_grdPaidFeesDetails"
	IdForUnpaidFeeTable     = "ContentPlaceHolder1_grdFeeDetails"
	IdForCaptcha            = "imgCaptcha"
)
