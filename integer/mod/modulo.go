package mod

import "strconv"

type Mod struct {
	value int
	prime int
}

func New(v, p int) Mod {
	m := Mod{
		value: v,
		prime: p,
	}
	m.fix()
	return m
}

func (m *Mod) fix() {
	if m.value >= m.prime || m.value < 0 {
		m.value = m.value % m.prime
		if m.value < 0 {
			m.value += m.prime
		}
	}
}

func (m *Mod) Inc() *Mod {
	m.value++
	if m.value >= m.prime {
		m.value -= m.prime
	}
	return m
}

func (m *Mod) Dec() *Mod {
	m.value--
	if m.value < 0 {
		m.value += m.prime
	}
	return m
}

func (m *Mod) Add(values ...Mod) *Mod {
	for _, value := range values {
		m.value += value.value
		if m.value >= m.prime {
			m.value -= m.prime
		}
	}
	return m
}

func (m *Mod) Sub(values ...Mod) *Mod {
	for _, value := range values {
		m.value -= value.value
		if m.value < 0 {
			m.value += m.prime
		}
	}
	return m
}

func (m *Mod) Mul(values ...Mod) *Mod {
	for _, value := range values {
		m.value *= value.value
		m.fix()
	}
	return m
}

func (m *Mod) Div(value Mod) *Mod {
	value.Exp(m.prime - 2)
	return m.Mul(value)
}

func (m *Mod) Exp(value int) *Mod {
	if value < 0 {
		m.Exp(-1 * value)
		m.Exp(m.prime - 2)
	} else if value == 0 {
		m.value = 1
	} else if value == 1 {

	} else if value%2 == 0 {
		m.Mul(*m)
		m.Exp(value / 2)
	} else {
		tmp := *m
		m.Mul(tmp)
		m.Exp((value - 1) / 2)
		m.Mul(tmp)
	}
	return m
}

func (m Mod) Value() int {
	return m.value
}

func (m Mod) String() string {
	return strconv.Itoa(m.value)
}
