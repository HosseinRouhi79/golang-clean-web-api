package services

import (
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

func NewOtpService(cfg *config.Config) *OtpService {

	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{cfg: cfg, logger: logger, redisClient: redis}
}

func (optService *OtpService) SetOtp(mobile string, otp string) error {
	//TODO: implement setOtp
	return nil
}

func (optService *OtpService) ValidateOtp(mobile string, otp string) error {
	//TODO: implement validateOtp
	return nil
}
