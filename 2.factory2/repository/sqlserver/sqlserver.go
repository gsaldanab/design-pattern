package sqlserver

import "fmt"

type SqlServer struct{}

func New() *SqlServer {
	return &SqlServer{}
}

func (c *SqlServer) Find(id int) string {
	return fmt.Sprintf("data from sqlServer -> %d", id)
}

func (c *SqlServer) Save(data string) error {
	fmt.Printf("save %s to sqlServer\n", data)
	return nil
}
