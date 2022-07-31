package academy

type RegistRequestBody struct {
	Year        string `json:"year" binding:"required"`
	Month       string `json:"month" binding:"required"`
	Day         string `json:"day" binding:"required"`
	YogaName    string `json:"yogaName" binding:"required"`
	RoomName    string `json:"roomName"`
	StartTime   string `json:"startTime" binding:"required"`
	EndTime     string `json:"endTime" binding:"required"`
	Level       string `json:"level"`
	TeacherName string `json:"teacherName"`
}
