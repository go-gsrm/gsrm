package gsrm

type Sql interface {
	Sql() string
}

type As struct {
	Name string
}

type Where struct {
	Column string
	Value  interface{}
}

type Select struct {
	Columns []string
	Tables  []TableName
}

type TableName struct {
	Name string
	as   *string
}

func (t *TableName) As(name string) *TableName {
	t.as = &name
	return t
}

func (t TableName) Sql() string {
	if t.as != nil {
		return t.Name + " AS " + *t.as
	}
	return t.Name
}
