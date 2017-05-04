package util
import (
	"os"
)

func GetHostName() (string ,error){
	hostname,err:=os.Hostname()
	if  err !=nil{
		return "",err
	}
	return hostname,nil
}