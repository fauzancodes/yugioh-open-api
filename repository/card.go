package repository

import (
	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
)

func CreateCard(data models.YOACard) (models.YOACard, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func CreateCardSet(data models.YOACardSet) (models.YOACardSet, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetCardByID(id uint) (response models.YOACard, err error) {
	err = config.DB.Where("id = ?", id).First(&response).Error

	return
}

func GetCardSetsByCardID(id uint) (response []models.YOACardSet, err error) {
	err = config.DB.Where("card_id = ?", id).Find(&response).Error

	return
}

func GetCards(param dto.FindParameter) (responses []models.YOACard, total int64, totalFiltered int64, err error) {
	err = config.DB.Model(responses).Where(param.BaseFilter, param.BaseFilterValues...).Count(&total).Error
	if err != nil {
		return
	}

	err = config.DB.Model(responses).Where(param.Filter, param.FilterValues...).Count(&totalFiltered).Error
	if err != nil {
		return
	}

	if param.Limit == 0 {
		err = config.DB.Debug().Offset(param.Offset).Order(param.Order).Where(param.Filter, param.FilterValues...).Find(&responses).Error
	} else {
		err = config.DB.Debug().Limit(param.Limit).Offset(param.Offset).Order(param.Order).Where(param.Filter, param.FilterValues...).Find(&responses).Error
	}

	return
}

func UpdateCard(data models.YOACard) (models.YOACard, error) {
	err := config.DB.Save(&data).Error

	return data, err
}

func DeleteCard(data models.YOACard) error {
	err := config.DB.Delete(&data).Error

	return err
}

func DeleteCardSet(data models.YOACardSet) error {
	err := config.DB.Delete(&data).Error

	return err
}

func GetCardUtility(field string) (responses []string, err error) {
	nullValues := "''"
	if field == "level" {
		nullValues = "0"
	}
	excludedRaces := `
	 	AND race NOT IN(
			'Thelonious Vi',
			'Pegasus',
			'Jesse Anderso',
			'Tania',
			'Mako',
			'Odion',
			'Dr. Vellian C',
			'Tyranno Hassl',
			'Rex',
			'Yugi',
			'Mai',
			'Camula',
			'Alexis Rhodes',
			'Syrus Truesda',
			'Axel Brodie',
			'Aster Phoenix',
			'Chumley Huffi',
			'Kagemaru',
			'Bastion Misaw',
			'Lumis Umbra',
			'Creator God',
			'Joey',
			'Ishizu',
			'Bonz',
			'Don Zaloog',
			'The Supreme K',
			'Abidos the Th',
			'Lumis and Umb',
			'Amnael',
			'David',
			'Weevil',
			'Adrian Gecko',
			'Yubel',
			'Joey Wheeler',
			'Chazz Princet',
			'Titan',
			'Christine',
			'Espa Roba',
			'Nightshroud',
			'Keith',
			'Tea Gardner',
			'Emma',
			'Yami Bakura',
			'Seto Kaiba',
			'Paradox Broth',
			'Kaiba',
			'Mai Valentine',
			'Jaden Yuki',
			'Yami Marik',
			'Arkana',
			'Zane Truesdal',
			'Andrew',
			'Yami Yugi',
			'Ishizu Ishtar'
		)
	`

	switch field {
	case "card_sets":
		err = config.DB.Raw("SELECT DISTINCT set_name FROM " + models.YOACardSet{}.TableName() + " WHERE set_name != '' AND deleted_at IS NULL").Scan(&responses).Error
	case "rarity":
		err = config.DB.Raw("SELECT DISTINCT set_rarity FROM " + models.YOACardSet{}.TableName() + " WHERE set_rarity != '' AND deleted_at IS NULL").Scan(&responses).Error
	case "rarity_code":
		err = config.DB.Raw("SELECT DISTINCT set_rarity_code FROM " + models.YOACardSet{}.TableName() + " WHERE set_rarity_code != '' AND deleted_at IS NULL").Scan(&responses).Error
	default:
		err = config.DB.Raw("SELECT DISTINCT " + field + " FROM " + models.YOACard{}.TableName() + " WHERE " + field + " != " + nullValues + " AND deleted_at IS NULL" + excludedRaces).Scan(&responses).Error
	}

	return
}
