package utils

type FieldProperty struct {
	Name          string
	Type          string
	NotNull       bool
	AutoIncrement bool
	Default       string
	PrimaryKey    bool
}

func (f FieldProperty) Define() []string {
	defines := []string{f.Name, f.Type}
	if f.NotNull {
		defines = append(defines, "NOT NULL")
	}
	if f.AutoIncrement {
		defines = append(defines, "AUTO_INCREMENT")
	}
	if f.Default != "" {
		defines = append(defines, "DEFAULT", f.Default)
	}
	return defines
}
