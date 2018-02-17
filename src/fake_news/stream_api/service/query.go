package service

import(
	sq "github.com/Masterminds/squirrel"
	"fake_news/stream_api/model"
	"log"
)

func PostToDB(dto model.TrainingDTO) error{
	query := sq.Insert("training").Columns("content", "answer").
		Values(dto.Content, dto.Answer)

	db, err := GetDB()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = query.RunWith(db).Exec()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}