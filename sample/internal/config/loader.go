package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	configPath := filepath.Join(currentDir, "internal/config/yaml")

	viper.SetConfigName("local")    // 設定ファイルの名前 (拡張子不要)
	viper.SetConfigType("yaml")     // ファイル形式の設定
	viper.AddConfigPath(configPath) // 設定ファイルを探すパスの追加

	// 設定読み込み
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error config file, %s", err)
	}

	// 構造体インスタンスの生成
	cfg := &Config{}

	// 設定の値を構造体にマップ
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return cfg, nil
}
