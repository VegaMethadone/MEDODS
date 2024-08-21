package guid

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
)

/*
Структура идентификатора:

 GUID STRUCT
     Data1   dd
     Data2   dw
     Data3   dw
     Data4   dw
     Data5   dp
 GUID ENDS
UUID-идентификаторы часто записывают в виде текстовой строки {G4G3G2G1-G6G5-G8G7-G9G10-G11G12G13G14G15G16}, где Gx — значение соответствующего байта структуры в шестнадцатеричном представлении[1]:

Data1 = G4G3G2G1 Data2 = G6G5 Data3 = G8G7 Data4 = G9G10 Data5 = G11G12G13G14G15G16

Например, '22345200-abe8-4f60-90c8-0d43c5f6c0f6' соответствует шестнадцатеричному 128-битному числу 0xF6C0F6C5430DC8904F60ABE822345200

Максимальное значение в GUID соответствует десятичному числу 340 282 366 920 938 463 463 374 607 431 768 211 455 (2128-1).

type guidStruct struct {
	data1 [4]byte
	data2 [2]byte
	data3 [2]byte
	data4 [2]byte
	data5 [6]byte
}
https://ru.wikipedia.org/wiki/GUID
*/

const lettersForHex string = "0123456789abcdef"

func CreateGUID() (string, error) {
	guidArr := make([]string, 5)
	for index := range guidArr {
		guidArr[index] = generateGuidPart(index)
	}

	return strings.Join(guidArr, "-"), nil
}

// В зависимости от указанной позиции GUID, возвращается срез байт соответствующей длины с случайно сгенерированными значениями в шестнадцатеричной системе счисления
func generateGuidPart(pos int) string {
	n := 0
	switch pos {
	case 0:
		n = 4 * 2
	case 1, 2, 3:
		n = 2 * 2
	case 4:
		n = 6 * 2
	}

	guidPart := make([]byte, n)
	for i := range n {
		guidPart[i] = lettersForHex[rand.Intn(len(lettersForHex))]
	}

	return string(guidPart)
}

func ValidateGUID(identifaer string) error {
	str := strings.Split(identifaer, "-")
	// Сначала проверяю, соответствует ли длина части GUID указанной позиции
	for index, guidPart := range str {
		switch index {
		case 0:
			if len(guidPart) != 4*2 {
				return errors.New("invalid guid")
			}
		case 1, 2, 3:
			if len(guidPart) != 2*2 {
				return errors.New("invalid guid")
			}
		case 4:
			if len(guidPart) != 6*2 {
				return errors.New("invalid guid")
			}
		}
	}

	//	Создаю  регулярнео выражения для валидации 16-ричного куска GUID
	pattern := "^[a-fA-F0-9]+"
	parser, err := regexp.Compile(pattern)
	if err != nil {
		return errors.New("faild to make pattern for regexp")
	}

	//	Пробегаюсь по частям  GUID и проверяб  их корректность
	for _, guidPart := range str {
		isOkay := parser.Match([]byte(guidPart))
		if !isOkay {
			return errors.New("invalid guid part")
		}
	}

	return nil
}
