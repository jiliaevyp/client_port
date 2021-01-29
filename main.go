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
	defaultNet  = "tcp"
	defaultIp   = "192.168.1.101"
	defaultPort = "8181"
)

// заголовок
func _beg() {
	fmt.Println("-------------------------------------")
	fmt.Println("|        'Go' client-сервер         |")
	fmt.Println("| Запрашиваем серверы до посинения! |")
	fmt.Println("|                                   |")
	fmt.Println("|   (c) jiliaevyp@gmail.com         |")
	fmt.Println("-------------------------------------")
}

// установка конфигурации сервера
func _config() (string, string) {
	var err int
	var ip, port, net string
	err = 1
	for err == 1 {
		net, err = inpNetwork() // ввод сети сервера
		fmt.Println("Тип сети:   ", net, "\n")
	}
	err = 1
	for err == 1 {
		ip, err = inpIP() // ввод IP сервера
		fmt.Println("IP адрес сервера:   ", ip, "\n")
	}
	err = 1
	for err == 1 {
		port, err = inpPort() // ввод порта сервера
		fmt.Println("Порт сервера:   ", port, "\n")
	}
	return net, ip + port
}

// ввод типа протокола сети
func inpNetwork() (string, int) {
	var typNet string
	var err int
	len := 256
	err = 1
	for err == 1 {
		fmt.Print("Тип сети по умолчанию:	", defaultNet, "\n", " Нажмите (Y) для изменения ")
		yes := yesNo()
		if yes == 0 {
			typNet = defaultNet
			return typNet, 0
		} else {
			fmt.Print("Введите тип сети:	")
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

// ввод ip адреса сервера
func inpIP() (string, int) {
	data := ""
	err := 1
	for err == 1 {
		fmt.Print("IP адрес сервера по умолчанию:	", defaultIp, "\n", "Для изменения нажмите 'Y' ")
		yes := yesNo()
		if yes != 1 {
			data = defaultIp
			err = 0
		} else {
			fmt.Print("Введите IP адрес сервера:	")
			fmt.Scanf(
				"%s\n",
				&data,
			)
			iperr := net.ParseIP(data)
			if iperr == nil {
				fmt.Println(ErrInvalidIPaddress)
				return data, 1
			} else {
				err = 0
			}
		}
	}
	return data, err
}

//ввод номера порта
func inpPort() (string, int) {
	var (
		webPort string
	)
	err := 1
	for err == 1 {
		fmt.Print("Порт по умолчанию:	", defaultPort, "\n", "Для изменения нажмите 'Y' ")
		yes := yesNo()
		if yes != 1 {
			webPort = defaultPort
			err = 0
		} else {
			fmt.Print("Введите порт:	")
			fmt.Scanf(
				"%s\n",
				&webPort,
			)
			res, err1 := strconv.ParseFloat(webPort, 16)
			res = res + 1
			err = 0
			if err1 != nil {
				fmt.Println(ErrInvalidPort)
				return ":" + webPort, 1
			}
		}
	}
	return ":" + webPort, 0
}

// текст запроса
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

// проверка на ввод  'Y = 1
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

// ввод числа запросов в автоматическом режиме
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
			return n
		}
	}
	return n
}

// client server
// сервер запущен на локальном компьтере на порте 4545
// клиент подлючается к этому адресу: net.Dial(_network, "_ip:_port")
// серверу будет отправляться запрос '_text' через conn.Write([]byte(source))
// с помощью вызова conn.Read выводим полученный ответ на консоль.

func _client(requestMessage string) (string, int) {
	answer := ""
	conn, err := net.Dial(_network, _port)
	if err != nil {
		fmt.Println(err)
		return "", 1
	}
	defer conn.Close() // закрыть соединение в конце
	if n, err := conn.Write([]byte(requestMessage)); n == 0 || err != nil {
		fmt.Println(ErrInvalidSendToServer)
		return "", 1
	}
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println(ErrInvalidServerRead)
		return "", 1
	}
	answer = string(buff[0:n])
	return answer, 0
}

func main() {
	var _answer, is string
	var err int

	_beg()                      // заголовок
	_network, _port = _config() // конфтигурация сервера
	zikl := 1
	for zikl == 1 {
		fmt.Println("\n")
		fmt.Println("Конфигурация сервера:   ", _network, _port, "\n")
		fmt.Println("Выберите режим запросов: одиночный/автоматический", "\n")
		fmt.Println("Нажмите 'Y' для формирования одного запроса  (для перехода в автоматический режим - любая клавиша)", "\n")
		yes := yesNo()
		if yes == 1 {
			request := 1 // ручные запросы
			fmt.Println("Режим одиночных запросов")
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
				fmt.Println("Режим одиночных запросов ", "\n", "нажмите 'Y' для продолжения (выход из режима - любая клавиша)")
				request = yesNo()
			}
			yes = 0
		} else { // автоматические запросы
			fmt.Println("Режим автоматических запросов")
			n := zikly()
			text := _textRequest()
			for i := 1; i < n; i++ {
				is = strconv.Itoa(i)
				//fmt.Println(i, is)
				_text := is + " " + text
				fmt.Println(is, "запрос: ", _text)
				_answer, err = _client(_text)
				if err != 0 {
					fmt.Println(ErrInvalidServerRead)
					break
				}
				fmt.Println("Ответ сервера:")
				fmt.Println(_answer)
			}
			fmt.Println("Конец запросов")
			fmt.Println("Выход из режима автоматических запросов")
		}
		fmt.Println("Нажмите 'Y' для продолжения работы  (выход - любая клавиша)")
		zikl = yesNo()
	}
	fmt.Println("Был счастлив на Вас поработать!")
	fmt.Print("Обращайтесь в любое время без колебаний!", "\n", "\n")
	return
}
