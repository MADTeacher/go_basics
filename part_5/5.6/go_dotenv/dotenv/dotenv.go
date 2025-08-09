package dotenv

import (
	"bufio"
	"errors"
	"fmt"
	"maps"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type DotEnv struct {
	env map[string]interface{}
}

// Регулярные выражения для int и bool
var (
	intPattern  = regexp.MustCompile(`^-?\d+$`)
	boolPattern = regexp.MustCompile(`^(?i:true|false)$`)
)

// Создаем новый экземпляр DotEnv
// platformEnv указывает, загружать ли переменные из окружения ОС
func NewDotEnv(platformEnv bool) *DotEnv {
	d := &DotEnv{env: make(map[string]interface{})}
	if platformEnv {
		d.env = mergeMaps(d.env, loadPlatformEnv())
	}
	return d
}

// Загружаем переменные из .env файлов
func (d *DotEnv) Load(paths ...string) error {
	if len(paths) == 0 {
		paths = []string{".env"}
	}
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf(".env file not found: %s", path)
		}
		defer file.Close()

		localEnv, err := loadLocalEnv(file)
		if err != nil {
			return err
		}
		d.env = mergeMaps(d.env, localEnv)
	}
	return nil
}

// Возвращает значение переменной окружения по ключу
func (d *DotEnv) GetValue(key string) (interface{}, error) {
	val, ok := d.env[key]
	if !ok {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	return val, nil
}

// Возвращает значение переменной окружения по ключу,
// приведенное к int. Если значение не является int,
// возвращается ошибка.
func (d *DotEnv) GetInt(key string) (int, error) {
	val, ok := d.env[key]
	if !ok {
		return 0, fmt.Errorf("key not found: %s", key)
	}
	i, ok := val.(int)
	if !ok {
		return 0, fmt.Errorf("value for key %s is not an int", key)
	}
	return i, nil
}

// Возвращает значение переменной окружения по ключу,
// приведенное к bool. Если значение не является bool,
// возвращается ошибка.
func (d *DotEnv) GetBool(key string) (bool, error) {
	val, ok := d.env[key]
	if !ok {
		return false, fmt.Errorf("key not found: %s", key)
	}
	b, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("value for key %s is not a bool", key)
	}
	return b, nil
}

// Возвращает значение переменной окружения по ключу,
// приведенное к string.
func (d *DotEnv) GetString(key string) (string, error) {
	val, ok := d.env[key]
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}
	s, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("value for key %s is not a string", key)
	}
	return s, nil
}

// Преобразование таблицы из переменных окружения ОС
func loadPlatformEnv() map[string]interface{} {
	config := make(map[string]interface{})
	for _, value := range os.Environ() {
		parts := strings.SplitN(value, "=", 2)
		if len(parts) != 2 {
			continue
		}
		k := parts[0]
		v := parts[1]

		config[k] = parseValue(v)
	}
	return config
}

// Загрузка из локального .env файла
func loadLocalEnv(file *os.File) (map[string]interface{}, error) {
	scanner := bufio.NewScanner(file)
	config := make(map[string]interface{})

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid line in .env file: " + line)
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		config[key] = parseValue(val)
	}
	return config, nil
}

// Парсинг строки в int, bool или string
func parseValue(val string) interface{} {
	if intPattern.MatchString(val) {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	if boolPattern.MatchString(val) {
		return strings.ToLower(val) == "true"
	}
	return val
}

// Объединение двух map
func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
	maps.Copy(a, b)
	return a
}
