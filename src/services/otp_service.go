package services

import (
	"fmt"
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/cache"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"github.com/go-redis/redis/v7"
)

type OtpService struct {
	cfg         *config.Config
	logger      logging.Logger
	redisClient *redis.Client
}

type OtpCred struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {

	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{cfg: cfg, logger: logger, redisClient: redis}
}

func (otpService *OtpService) SetOtp(mobile string, otp string) error {
	
	var prefix string = "redis"
	key := fmt.Sprintf("%s:%s", prefix, mobile)
	val := OtpCred{Value: otp, Used: false}
	err := cache.Set[OtpCred](otpService.redisClient, key, val, time.Second*3600)
	if err != nil {
		return err
	}
	return nil
}

func (otpService *OtpService) ValidateOtp(mobile string, otp string) error {
	//TODO: implement validateOtp
	return nil
}
