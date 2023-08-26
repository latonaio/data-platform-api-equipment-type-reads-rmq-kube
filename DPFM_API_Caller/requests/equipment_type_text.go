package requests

type EquipmentTypeText struct {
	EquipmentType     	string `json:"EquipmentType"`
	Language          	string `json:"Language"`
	EquipmentTypeName 	string `json:"EquipmentTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
