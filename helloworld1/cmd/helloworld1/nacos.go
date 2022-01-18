package main

import (
	"encoding/json"

	"github.com/bytectl/gopkg/crypto/gconfig"
	"helloworld1/internal/conf"
	nacosV2 "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
)

var (
	key = []byte("xxxxxx")
)

func newConfig(confPath string, logger log.Logger) config.Config {
	log := log.NewHelper(log.With(logger, "module", "nacos/config"))
	c := config.New(
		config.WithSource(
			env.NewSource("CONFIG_"),
			file.NewSource(confPath),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}
	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}
	bs, _ := json.Marshal(rc)
	log.Debugf("%v\n", string(bs))
	client := newNacosConfigClient(&rc)

	source := []config.Source{}
	// add file config
	source = append(source, file.NewSource(confPath))
	// add nacos config
	for _, dataId := range rc.Nacos.DataIds {
		source = append(source, nacosV2.NewConfigSource(client,
			nacosV2.WithGroup(rc.Nacos.Group),
			nacosV2.WithDataID(dataId),
			nacosV2.WithCacheDir("/tmp/nacos/cache"),
		))
	}
	cc := config.New(
		config.WithSource(source...),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			err := yaml.Unmarshal(kv.Value, v)
			if err != nil {
				return err
			}
			gconfig.DecryptConfigMap(v, key)
			return nil
		}),
	)
	return cc
}

func newNacosConfigClient(conf *conf.Registry) config_client.IConfigClient {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.Address, conf.Nacos.Port),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         conf.Nacos.NamespaceId, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          conf.Nacos.RotateTime,
		MaxAge:              conf.Nacos.MaxAge,
		LogLevel:            conf.Nacos.LogLevel,
	}

	// a more graceful way to create naming client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	return client
}
