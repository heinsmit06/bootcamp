package bootcamp

type MyMap struct {
	// Define fields and any necessary structures here
	Data []struct {
		Key   string
		Value interface{}
	}
}

func (m *MyMap) Set(k string, v interface{}) {
	for i := range m.Data {
		if m.Data[i].Key == k {
			m.Data[i].Value = v
			return
		}
	}
	m.Data = append(m.Data, struct {
		Key   string
		Value interface{}
	}{Key: k, Value: v})
}

func (m *MyMap) Get(k string) interface{} {
	for i := range m.Data {
		if m.Data[i].Key == k {
			return m.Data[i].Value
		}
	}
	return nil
}

func (m *MyMap) Has(k string) bool {
	for i := range m.Data {
		if m.Data[i].Key == k {
			return true
		}
	}
	return false
}

func (m *MyMap) Delete(k string) {
	for i := range m.Data {
		if m.Data[i].Key == k && i != len(m.Data)-1 {
			m.Data = append(m.Data[:i], m.Data[i+1:]...)
			return
		}
		if m.Data[i].Key == k {
			m.Data = m.Data[:i]
			return
		}
	}
}

func (m *MyMap) Items() []struct {
	Key   string
	Value interface{}
} {
	return m.Data
}
