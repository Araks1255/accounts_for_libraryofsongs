package utils

func ConvertToMap(slice []string) map[int]string { // Функция преобразования слайса в мапу
	result := make(map[int]string)    // Инициализируем мапу
	for i := 0; i < len(slice); i++ { // В цикле
		result[i+1] = slice[i] // Назначем ключу мапы i+1 значение слайса i
	} // И в итоге получается мапа с порядковыми номерами всех значений
	return result // Возвращаем мапу
}
