package postgres

import "api/models"

func UploadText(message models.Message) error {
	tx := Database.MustBegin()
	_, err := tx.NamedExec("INSERT INTO messages (sender, timestamp, content) VALUES (:sender, :timestamp, :content)", message)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
