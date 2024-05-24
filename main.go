package main

import (
	"fmt"
	"github.com/vonuki/textgame/internal/commands"
	"github.com/vonuki/textgame/internal/items"
	"github.com/vonuki/textgame/internal/locations"
	"github.com/vonuki/textgame/internal/player"
	"log"
	"os"
	"strings"
)

var logger *log.Logger
var world *locations.WorldHandler
var person player.Player
var cmdKeeper *commands.ComKeeper

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
	initGame()
	var command, arg string
	for {
		fmt.Scanf("%s%s\n", &command, &arg)
		if command == "стоп" {
			fmt.Println("Конец игры")
			os.Exit(0)
		}
		result := handleCommand(command + " " + arg)
		fmt.Println(result)
	}
}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все локации
		если что-то было - оно корректно перезатирается
	*/
	logger = log.New(os.Stdout, "Debug: ", 3)

	// Создание Мира (хранилище локаций)
	world = locations.NewWorldHandler()

	// Наполнение Мира и конфигурирование локаций, их взаимосвязей
	kitchen := locations.NewLocation("кухня", "кухня, ничего интересного", "", "ты находишься на кухне, надо собрать рюкзак и идти в универ")
	kitchen.AddLinkedLocation("коридор")
	world.AddLocation(kitchen)

	hall := locations.NewLocation("коридор", "ничего интересного", "", "")
	hall.AddLinkedLocation("кухня")
	hall.AddLinkedLocation("комната")
	hall.AddLinkedLocation("улица")
	world.AddLocation(hall)

	room := locations.NewLocation("комната", "ты в своей комнате", "на столе", "")
	room.AddLinkedLocation("коридор")
	room.PutItem(items.NewItem("ключи"))
	room.PutItem(items.NewItem("конспекты"))
	room.PutItem(items.NewItem("рюкзак"))
	world.AddLocation(room)

	street := locations.NewLocation("улица", "на улице весна", "", "")
	street.AddLinkedLocation("домой")
	world.AddLocation(street)

	// Создание игрока
	person = player.NewPlayer("Дима", "инвентарь", "кухня")

	// Создание хранилища команд и наполнение командами.
	cmdKeeper = commands.NewCommandKeeper()
	cmdKeeper.AddCommand("осмотреться", commands.LookAround)
	cmdKeeper.AddCommand("идти", commands.Move)
	cmdKeeper.AddCommand("взять", commands.Take)

	//logger.Println(fmt.Sprintf("Мир создан: %+v", world))
	//logger.Println(cmdKeeper.GetCommandNames())

}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	//Логирование Служебный вывод
	logger.Println(fmt.Sprintf("Обработка команды: %s", command))

	cmds := strings.Split(command, " ")

	// Вызов обработчика из хранилища команд
	return cmdKeeper.DoCommandByName(cmds[0], world, person, cmds[1:]...)
}
