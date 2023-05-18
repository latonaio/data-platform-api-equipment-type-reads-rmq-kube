package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-equipment-type-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-equipment-type-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var equipmentType *[]dpfm_api_output_formatter.EquipmentType
	var equipmentTypeText *[]dpfm_api_output_formatter.EquipmentTypeText
	for _, fn := range accepter {
		switch fn {
		case "EquipmentType":
			func() {
				equipmentType = c.EquipmentType(mtx, input, output, errs, log)
			}()
		case "EquipmentTypeText":
			func() {
				equipmentTypeText = c.EquipmentTypeText(mtx, input, output, errs, log)
			}()
		case "EquipmentTypeTexts":
			func() {
				equipmentTypeText = c.EquipmentTypeTexts(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		EquipmentType:     equipmentType,
		EquipmentTypeText: equipmentTypeText,
	}

	return data
}

func (c *DPFMAPICaller) EquipmentType(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.EquipmentType {
	equipmentType := input.EquipmentType.EquipmentType

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_equipment_type_equipment_type_data
		WHERE EquipmentType = ?;`, equipmentType,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToEquipmentType(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) EquipmentTypeText(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.EquipmentTypeText {
	var args []interface{}
	equipmentType := input.EquipmentType.EquipmentType
	equipmentTypeText := input.EquipmentType.EquipmentTypeText

	cnt := 0
	for _, v := range equipmentTypeText {
		args = append(args, equipmentType, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_equipment_type_equipment_type_text_data
		WHERE (EquipmentType, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToEquipmentTypeText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) EquipmentTypeTexts(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.EquipmentTypeText {
	where := fmt.Sprintf("WHERE EquipmentType = '%v'", input.EquipmentType.EquipmentType)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_equipment_type_equipment_type_text_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToEquipmentTypeText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
