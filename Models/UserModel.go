package Models
type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Dob string `json:"dob" binding:"required""`
	Address string `json:"address" binding:"required"`
	Maths	int	`json:"maths"`
	English int	`json:"english"`
	Science int	`json:"science"`
	Social	int	`json:"social"`
	//Subject []Subject `json:"subject"`

}

//type Subject struct {
//	Id	uint `json:"id"`
//	SubjectName string `json:"subject_name"`
//	Marks string	`json:"marks"`
//}
func (b *User) TableName() string {
	return "user"
}

//func (b *Subject) Table() string{
//	return "subject"
//}
