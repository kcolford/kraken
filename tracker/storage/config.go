package storage

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/yaml.v2"
)

// Config defines storage configuration.
type Config struct {
	PeerStore     string      `yaml:"peer_store"`
	TorrentStore  string      `yaml:"torrent_store"`
	ManifestStore string      `yaml:"manifest_store"`
	Redis         RedisConfig `yaml:"redis"`
	MySQL         MySQLConfig `yaml:"mysql"`
}

// MySQLConfig defines configuration for MySQL storage.
type MySQLConfig struct {
	MigrationsDir string `yaml:"migration_dir"`
}

// RedisConfig defines configuration for Redis storage.
type RedisConfig struct {
	Addr              string        `yaml:"addr"`
	DialTimeout       time.Duration `yaml:"dial_timeout"`
	ReadTimeout       time.Duration `yaml:"read_timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	PeerSetWindowSize time.Duration `yaml:"peer_set_window_size"`
	MaxPeerSetWindows int           `yaml:"max_peer_set_windows"`
	TorrentTTL        time.Duration `yaml:"torrent_ttl"`
	MaxIdleConns      int           `yaml:"max_idle_conns"`
	MaxActiveConns    int           `yaml:"max_active_conns"`
	IdleConnTimeout   time.Duration `yaml:"idle_conn_timeout"`
	OriginsTTL        time.Duration `yaml:"origins_ttl"`
}

func (c RedisConfig) String() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("yaml marshal error: %s", err)
	}
	return string(b)
}

func (c RedisConfig) applyDefaults() (RedisConfig, error) {
	if c.Addr == "" {
		return c, errors.New("no addr configured")
	}
	if c.DialTimeout == 0 {
		c.DialTimeout = 5 * time.Second
	}
	if c.ReadTimeout == 0 {
		c.ReadTimeout = 30 * time.Second
	}
	if c.WriteTimeout == 0 {
		c.WriteTimeout = 30 * time.Second
	}
	if c.PeerSetWindowSize == 0 {
		c.PeerSetWindowSize = time.Hour
	}
	if c.MaxPeerSetWindows == 0 {
		c.MaxPeerSetWindows = 10
	}
	if c.TorrentTTL == 0 {
		c.TorrentTTL = time.Hour
	}
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = 10
	}
	if c.MaxActiveConns == 0 {
		c.MaxActiveConns = 50
	}
	if c.IdleConnTimeout == 0 {
		c.IdleConnTimeout = 60 * time.Second
	}
	if c.OriginsTTL == 0 {
		c.OriginsTTL = 10 * time.Hour
	}
	return c, nil
}