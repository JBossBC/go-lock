package examples

import (
	Mutex "github.com/JBossBC/go-lock"
	"github.com/redis/go-redis/v9"
)

func main() {
	// init the redis server in global config
	Mutex.AssemblyMutex(Mutex.WithStorageClient(redis.NewClient(&redis.Options{Addr: "localhost:6379"})))
	// apply for a new mutex
	mutex := Mutex.NewMutex("example")
	// lock,when it get the lock,it will continue to execute.
	mutex.Lock()
	// unlock
	mutex.UnLock()
	//eventually success,if return false,represent the mutex failed in  delaying
	mutex.Discard()

}
