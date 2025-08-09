package jsonstore

import (
	"encoding/json"
	"os"
	"reflect"
)

type JSONStore struct {
	filePath string
	values   map[string]interface{}
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{
		filePath: path,
	}
}

func (js *JSONStore) load() map[string]interface{} {
	if js.values != nil {
		return js.values
	}

	content, err := os.ReadFile(js.filePath)
	if err != nil {
		js.values = make(map[string]interface{})
		return js.values
	}

	var data map[string]interface{}
	if err := json.Unmarshal(content, &data); err != nil {
		js.values = make(map[string]interface{})
		return js.values
	}
	js.values = data
	return js.values
}

func (js *JSONStore) Contains(key string) bool {
	_, ok := js.load()[key]
	return ok
}

func (js *JSONStore) Keys() []string {
	keys := make([]string, 0)
	for k := range js.load() {
		keys = append(keys, k)
	}
	return keys
}

func (js *JSONStore) Values() []interface{} {
	values := make([]interface{}, 0)
	for _, v := range js.load() {
		values = append(values, v)
	}
	return values
}

func (js *JSONStore) GetValue(key string) interface{} {
	return js.load()[key]
}

func (js *JSONStore) GetBool(key string) *bool {
	if v, ok := js.GetValue(key).(bool); ok {
		return &v
	}
	return nil
}

func (js *JSONStore) GetInt(key string) *int {
	// Т.к. все числа по умолчанию декодируются как float64,
	// то сначала приводим к float64
	if f, ok := js.GetValue(key).(float64); ok {
		i := int(f)
		return &i
	}
	return nil
}

func (js *JSONStore) GetFloat(key string) *float64 {
	if v, ok := js.GetValue(key).(float64); ok {
		return &v
	}
	return nil
}

func (js *JSONStore) GetString(key string) *string {
	if v, ok := js.GetValue(key).(string); ok {
		return &v
	}
	return nil
}

// Метод возвращает по ключу срез []interface{}.
// Если значение хранится как []interface{}, возвращаем его напрямую.
// Если значение хранится как []string, []int или []float64,
// преобразуем каждый элемент к interface{} для совместимости
// с Go JSON API. Это обеспечивает универсальный доступ к массивам
// любого типа, сохраненным в JSON.
func (js *JSONStore) GetList(key string) []interface{} {
	// Получаем значение по ключу
	val := js.GetValue(key)
	// Если значение отсутствует, возвращаем nil
	if val == nil {
		return nil
	}
	// Если это уже []interface{} (тип, который возвращает
	// json.Unmarshal), возвращаем напрямую
	if v, ok := val.([]interface{}); ok {
		return v
	}
	// Если это []string, преобразуем каждый элемент к interface{}
	if v, ok := val.([]string); ok {
		// Создаем новый срез с типом []interface{}
		result := make([]interface{}, len(v))
		// Преобразуем каждый элемент []string в interface{}
		for i, s := range v {
			result[i] = s
		}
		return result
	}
	// Если это []int, преобразуем каждый элемент к interface{}
	if v, ok := val.([]int); ok {
		// Создаем новый срез с типом []interface{}
		result := make([]interface{}, len(v))
		// Преобразуем каждый элемент []int в interface{}
		for i, n := range v {
			result[i] = n
		}
		return result
	}
	// Если это []float64 (тип, который возвращает json.Unmarshal),
	// преобразуем каждый элемент к interface{}
	if v, ok := val.([]float64); ok {
		// Создаем новый срез с типом []interface{}
		result := make([]interface{}, len(v))
		// Преобразуем каждый элемент []float64 в interface{}
		for i, n := range v {
			result[i] = n
		}
		return result
	}
	// Если тип не поддерживается, возвращаем nil
	return nil
}

func (js *JSONStore) GetMap(key string) map[string]interface{} {
	val := js.GetValue(key)
	if val == nil {
		return nil
	}
	if m, ok := val.(map[string]interface{}); ok {
		return m
	}
	// Попробуем привести через json, если тип не совпал
	if b, err := json.Marshal(val); err == nil {
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err == nil {
			return m
		}
	}
	return nil
}

func (js *JSONStore) ValueEquals(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func (js *JSONStore) SetValue(key string, value interface{}) {

	values := js.load()
	oldVal := values[key]
	if value == nil {
		delete(values, key)
		js.save(values)
	} else if oldVal == nil || !js.ValueEquals(oldVal, value) {
		values[key] = value
		js.save(values)
	}
}

func (js *JSONStore) ResetValue(key string) {

	values := js.load()
	if _, exists := values[key]; exists {
		delete(values, key)
		js.save(values)
	}
}

func (js *JSONStore) save(data map[string]interface{}) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}
	_ = os.WriteFile(js.filePath, bytes, 0644)
	js.values = data
}
