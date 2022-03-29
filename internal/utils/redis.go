package utils

// "github.com/go-redis/redis"

//var client *redis.Client

// InitRedis connection
func InitRedis() {
	//Initializing redis
	// dsn := os.Getenv("REDIS_DSN")
	// if len(dsn) == 0 {
	// 	dsn = "localhost:6379"
	// }
	// client = redis.NewClient(&redis.Options{
	// 	Addr: dsn,
	// })
	// _, err := client.Ping().Result()
	// if err != nil {
	// 	panic(err)
	// }
}

//RedisTokenDetails structure
type RedisTokenDetails struct {
	AccessToken string
	AccessUUID  string
	AtExpires   int64
	UserID      uint64
}

//CreateRedisAuth  redis document
func CreateRedisAuth(userid uint64, td *RedisTokenDetails) error {
	// at := time.Unix(td.AtExpires, 0)
	// now := time.Now()

	// errAccess := client.Set(td.AccessUUID, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	// if errAccess != nil {
	// 	return errAccess
	// }
	return nil
}
