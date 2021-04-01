package types

type DashboardModel struct {
	AdmDetails struct {
		RegNo    string `json:"reg_no"`
		Name     string `json:"name"`
		Program  string `json:"program"`
		AcadYear string `json:"acad_year"`
	} `json:"adm_details"`

	EduQualifications []map[string]string `json:"edu_qualifications"`

	ParentDetails struct {
		MobileNo         string `json:"mobile_no"`
		Email            string `json:"email"`
		Mother           string `json:"mother"`
		Father           string `json:"father"`
		EmergencyContact string `json:"emergency_contact"`
	} `json:"parent_details"`
}
