package resource

import (
	"fmt"
	"github.com/jinzhu/inflection"
	"github.com/wokacz/mamba/internal/words"
)

func getModelContent(resourceName string) string {
	className := words.ToPascalCase(resourceName)
	plural := inflection.Plural(resourceName)
	tableName := words.ToSnakeCase(plural)

	return fmt.Sprintf(`

class %s(TimestampModel, BaseModel, table=True):
	__tablename__ = '%s'
	
	# TODO: Define the data model fields here
	pass
`, className, tableName)
}
