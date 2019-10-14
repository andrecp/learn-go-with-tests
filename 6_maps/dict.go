package dict

// Dictionary represents a dict
type Dictionary map[string]string

// Errors for dictionary
const (
	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrAlreadyExists = DictionaryErr("cannot add word because it already exists")
)

// DictionaryErr have errors for dicts
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search searches for a value in a dictionary
func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add adds a value in a dictionary, `ErrAlreadyExists` is returned if the key already exists.
func (d Dictionary) Add(key string, value string) error {
	_, ok := d[key]
	if ok {
		return ErrAlreadyExists
	}
	d[key] = value
	return nil
}

// Update updates a value in a dictionary
func (d Dictionary) Update(key string, value string) error {
	d[key] = value
	return nil
}

// Delete updates a value in a dictionary
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
