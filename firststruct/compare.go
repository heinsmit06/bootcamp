package firststruct

func (u User) Compare(v User) bool {
	if u.Name == v.Name {
		if u.password == v.password {
			return true
		}
	}

	return false
}
