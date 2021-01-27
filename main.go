package main

import (
	"errors"
	"fmt"

	//"io"
	"net"
	"os"
	"strconv"
)

var (
	requestMessage, _network, _port string
	ErrInvalidTypeNetwork           = errors.New("invalid type network")
	ErrInvalidPort                  = errors.New("invalid port number")
	ErrInvalidSendToServer          = errors.New("invalid send to server")
	ErrInvalidServerRead            = errors.New("invalid read from server")
)

// запрос
func _textRequest() string {
	var text string
	len := 256
	fmt.Print("Введи запрос:")
	fmt.Println()
	data := make([]byte, len)
	n, err := os.Stdin.Read(data)
	text = string(data[0 : n-1])
	if err != nil {
		return ""
	} else {
		return text
	}
}

// проверка на окончание запросов
func yesNo() int {
	var yesNo string
	len := 4
	data := make([]byte, len)
	n, err := os.Stdin.Read(data)
	yesNo = string(data[0 : n-1])
	if err == nil && (yesNo == "Y" || yesNo == "y" || yesNo == "Н" || yesNo == "н") {
		return 0
	} else {
		return 1
	}
}

// ввод протокола сети
func inpNetwork() (string, int) {
	var typNet string
	len := 256
	data := make([]byte, len)
	n, err := os.Stdin.Read(data)
	typNet = string(data[0 : n-1])
	if err != nil || typNet != "tcp" {
		return typNet, 1
	} else {
		return typNet, 0
	}
}

//ввод номер порта
func inpPort() (string, int) {
	var (
		webPort string
		res     float64
	)
	fmt.Scanf(
		"%s\n",
		&webPort,
	)
	res, err := strconv.ParseFloat(webPort, 16)
	res = res + 1
	if err != nil || len(webPort) != 4 {
		return ":" + webPort, 1
	} else {
		return ":" + webPort, 0
	}
}

// client server
// сервер запущен на локальном компьтере на порте 4545
// клиент подлючается к этому адресу: net.Dial("tcp", ":4545")
// серверу будет отправляться запрос conn.Write([]byte(source))
// с помощью вызова conn.Read выводим полученный ответ на консоль.

func _client() (string, int) {
	answer := ""
	conn, err := net.Dial(_network, _port)
	if err != nil {
		fmt.Println(err)
		return "", 1
	}
	defer conn.Close()
	requestMessage = _textRequest()
	if n, err := conn.Write([]byte(requestMessage)); n == 0 || err != nil {
		fmt.Println(ErrInvalidSendToServer)
		return "", 1
	}
	fmt.Println("Ответ сервера: ")
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println(ErrInvalidServerRead)
		return "", 1
	}
	answer = string(buff[0:n])
	return answer, 0
}

func _beg() {
	fmt.Println("------------------------------------")
	fmt.Println("|          Запуск Go client        |")
	fmt.Println("|     Запрашиваем до посинения!    |")
	fmt.Println("|                                  |")
	fmt.Println("|   (c) jiliaevyp@gmail.com        |")
	fmt.Println("------------------------------------")
}

func main() {
	var _answer string
	err := 1
	_beg() // заголовок
	for err == 1 {
		fmt.Print("Введите тип сети:	")
		_network, err = inpNetwork()
		if err == 1 {
			fmt.Println(ErrInvalidTypeNetwork)
		}
	}
	err = 1
	for err == 1 {
		fmt.Print("Введите номер порта:	")
		_port, err = inpPort()
		if err == 1 {
			fmt.Println(ErrInvalidPort)
		}
	}
	request := 0
	for request == 0 {
		fmt.Print("\n", "Для запроса нажми 'Y'---> ")
		request = yesNo()
		if request == 0 {
			_answer, err = _client()
			if err == 0 {
				fmt.Println(_answer)
			} else {
				fmt.Println(ErrInvalidServerRead)
			}
		} else {
			fmt.Println("Выход")
			fmt.Println("Рад был с Вами пработать!")
			fmt.Print("Обращайтесь в любое время без колебаний!", "\n", "\n")
			return
		}
	}
}
