package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	//"strings"
)

var (
	requestMessage, _network, _ip, _port string
	ErrInvalidTypeNetwork                = errors.New("invalid type network")
	ErrInvalidIPaddress                  = errors.New("invalid IP address")
	ErrInvalidPort                       = errors.New("invalid port number")
	ErrInvalidSendToServer               = errors.New("invalid send to server")
	ErrInvalidServerRead                 = errors.New("invalid read from server")
	ErrInvalidProcRequest                = errors.New("invalid procedure of request")
)

const (
	default_net  = "tcp"
	default_IP   = "192.168.0.101"
	default_port = "8181"
)

// запрос
func _textRequest() string {
	var text string
	len := 256
	fmt.Println("Введи текст запроса:")
	data := make([]byte, len)
	n, err := os.Stdin.Read(data)
	text = string(data[0 : n-1])
	if err != nil {
		return ""
	} else {
		return text
	}
}

// проверка на окончание запросов Yes = 1
func yesNo() int {
	var yesNo string
	len := 4
	data := make([]byte, len)
	n, err := os.Stdin.Read(data)
	yesNo = string(data[0 : n-1])
	if err == nil && (yesNo == "Y" || yesNo == "y" || yesNo == "Н" || yesNo == "н") {
		return 1
	} else {
		return 0
	}
}

// ввод протокола сети
func inpNetwork() (string, int) {
	var typNet string
	var err int
	len := 256
	err = 1
	for err == 1 {
		fmt.Print("Тип сети:	", default_net, "\n", " Нажмите (Y) для изменения ")
		yes := yesNo()
		if yes == 0 {
			typNet = default_net
			return typNet, 0
		} else {
			data := make([]byte, len)
			n, err := os.Stdin.Read(data)
			typNet = string(data[0 : n-1])
			if err != nil || typNet != "tcp" {
				fmt.Println(ErrInvalidTypeNetwork)
				return typNet, 1
			}
			return typNet, 0
		}
	}
	return typNet, 0
}

// ввод ip сервера
func inpIP() (string, int) {
	data := ""
	err := 1
	for err == 1 {
		fmt.Print("Введите IP сервера:	", default_IP, "\n", " Нажмите (Y) для изменения ")
		yes := yesNo()
		if yes != 1 {
			data = default_IP
			err = 0
		} else {
			fmt.Scanf(
				"%s\n",
				&data,
			)
			iperr := net.ParseIP(data)
			if iperr == nil {
				fmt.Println(ErrInvalidIPaddress)
				return data, 1
			}
		}
	}
	return data, err
}

//ввод номер порта
func inpPort() (string, int) {
	var (
		webPort string
	)
	err := 1
	for err == 1 {
		fmt.Print("Введите порт:	", default_port, "\n", " Нажмите (Y) для изменения ")
		yes := yesNo()
		if yes != 1 {
			webPort = default_port
			err = 0
		} else {
			fmt.Scanf(
				"%s\n",
				&webPort,
			)
			res, err1 := strconv.ParseFloat(webPort, 16)
			res = res + 1
			err = 0
			if err1 != nil || len(webPort) != 4 {
				fmt.Println(ErrInvalidPort)
				return ":" + webPort, 1
			}
		}
	}
	return ":" + webPort, 0
}

func zikly() int {
	var s string
	var n int
	err1 := 1
	for err1 == 1 {
		fmt.Print("Введите число запросов:   ")
		fmt.Scanf(
			"%s\n",
			&s,
		)
		res, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(ErrInvalidProcRequest)
		} else {
			err1 = 0
			n = res
		}
	}
	return n
}

// client server
// сервер запущен на локальном компьтере на порте 4545
// клиент подлючается к этому адресу: net.Dial("tcp", ":4545")
// серверу будет отправляться запрос conn.Write([]byte(source))
// с помощью вызова conn.Read выводим полученный ответ на консоль.

func _client(_text string) (string, int) {
	answer := ""
	conn, err := net.Dial(_network, _port)
	if err != nil {
		fmt.Println(err)
		return "", 1
	}
	defer conn.Close()
	requestMessage = _text
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
	var err int

	_beg() // заголовок
	err = 1
	for err == 1 {
		_network, err = inpNetwork() // ввод сети сервера
		fmt.Println(_network)
	}
	err = 1
	for err == 1 {
		_ip, err = inpIP() // ввод IP сервера
		fmt.Println(_ip)
	}
	err = 1
	for err == 1 {
		_port, err = inpPort() // ввод порта сервера
		fmt.Println(_port)
	}
	_port = _ip + _port
	zikl := 1
	for zikl == 1 {
		fmt.Println("\n")
		fmt.Println("Сервер:   ", _network, _port)
		fmt.Println("Выберите режим запросов: ручной/автоматический")
		fmt.Println("Нажмите 'Y' для ручного запроса")
		yes := yesNo()
		if yes == 1 {
			request := 1 // ручные запросы
			fmt.Println("Ручной режим")
			for request == 1 {
				if request == 1 {
					text := _textRequest()
					_answer, err = _client(text)
					if err == 0 {
						fmt.Println(_answer)
					} else {
						fmt.Println(ErrInvalidServerRead)
					}
				} else {
					fmt.Println("Выход из режима одиночных запросов")
				}
				fmt.Println("Ручной режим, нажмите 'Y' для нового запроса")
				request = yesNo()
			}
			yes = 0
		} else { // автоматические запросы
			fmt.Println("Режим автоматических запросов")
			n := zikly()
			_text := _textRequest()
			for i := 1; i < n; i++ {
				fmt.Println("Запрос ", i)
				_answer, err = _client(_text)
				if err != 0 {
					fmt.Println(ErrInvalidServerRead)
					break
				}
				fmt.Println(_answer)
			}
			fmt.Println("Конец опросов")
			fmt.Println("Выход из режима автоматических запросов")
		}
		fmt.Println("Нажмите 'Y' для продолжения работы")
		zikl = yesNo()
	}
	fmt.Println("Был счастлив на Вас поработать!")
	fmt.Print("Обращайтесь в любое время без колебаний!", "\n", "\n")
	return
}
