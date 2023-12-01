package work_with_command

import (
	"csv-file/internal/handler"
	"fmt"
)

func WorkWithCSVByCommand() error {
	fmt.Println("Введите команду 'savetodb' для загрузки данных в базу данных")
	var command string
	_, err := fmt.Scan(&command)
	if err != nil {
		return err
	}
	for command != "savetodb" {
		fmt.Println("Команды такой не существует")
		fmt.Println("Введите команду 'savetodb' для загрузки данных в базу данных")
		_, err := fmt.Scan(&command)
		if err != nil {
			return err
		}
	}
	err = handler.DataProcessing()
	if err != nil {
		return err
	}
	return nil
}
