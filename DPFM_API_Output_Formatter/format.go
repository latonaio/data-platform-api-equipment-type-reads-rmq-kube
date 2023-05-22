package dpfm_api_output_formatter

import (
	"data-platform-api-equipment-type-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToEquipmentType(rows *sql.Rows) (*[]EquipmentType, error) {
	defer rows.Close()
	equipmentType := make([]EquipmentType, 0)

	i := 0
	for rows.Next() {
		pm := &requests.EquipmentType{}
		i++

		err := rows.Scan(
			&pm.EquipmentType,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		equipmentType = append(equipmentType, EquipmentType{
			EquipmentType: data.EquipmentType,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &equipmentType, nil
}

func ConvertToEquipmentTypeText(rows *sql.Rows) (*[]EquipmentTypeText, error) {
	defer rows.Close()
	equipmentTypeText := make([]EquipmentTypeText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := requests.EquipmentTypeText{}

		err := rows.Scan(
			&pm.EquipmentType,
			&pm.Language,
			&pm.EquipmentTypeName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &equipmentTypeText, err
		}

		data := pm
		equipmentTypeText = append(equipmentTypeText, EquipmentTypeText{
			EquipmentType:     data.EquipmentType,
			Language:          data.Language,
			EquipmentTypeName: data.EquipmentTypeName,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &equipmentTypeText, nil
	}

	return &equipmentTypeText, nil
}
