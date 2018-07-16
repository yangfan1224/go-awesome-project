package pilosa
import(
	"github.com/pilosa/go-pilosa"
	"fmt"
)

var pilosaClient *pilosa.Client

func init(){
	pilosaClient = pilosa.DefaultClient()
}

func QueryExample()(error){
	// Retrieve the schema
	schema, err := pilosaClient.Schema()
	if err != nil{
		return err
	}
	// Create an Index object
	myindex, err := schema.Index("test1")

	// Create a Field object
	myfield, err := myindex.Field("f")

	if err != nil{
		return err
	}
	// make sure the index and the field exists on the server
	err = pilosaClient.SyncSchema(schema)
	if err != nil{
		return err
	}

	// Send a Set query. PilosaException is thrown if execution of the query fails.
	response, err := pilosaClient.Query(myfield.SetK("yangfan", "hello"))

	if err != nil{
		return err
	}

	fmt.Println(response)
	return nil
}