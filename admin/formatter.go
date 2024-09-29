package admin

type AdminFormatter struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatAdmin(admin *Admin) AdminFormatter {
	formatter := AdminFormatter{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
	}
	return formatter
}
