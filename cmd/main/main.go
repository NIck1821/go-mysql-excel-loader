package main

import (
	"flag"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/Nick1821/go-mysql-excel-loader/configs"

	xls_loader "github.com/Nick1821/go-mysql-excel-loader/internal/app/excel_loader"
	"github.com/Nick1821/go-mysql-excel-loader/internal/app/store/mysql_store"
	"github.com/sirupsen/logrus"
)

var (
	configPath                 string
	data_start, data_end, city string
	limit, offset              string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/config.toml", "конфигурация")
	// указываем даты с и по
	flag.StringVar(&data_start, "data_start", "2020-01-01 00:00:00", "дата с (по умолчанию с 2020 года) ")
	flag.StringVar(&data_start, "data_end", "2026-01-01 00:00:00", "дата по (по умолчанию до 2025 года)")
	// указывем город
	flag.StringVar(&city, "city", "Москва", "город(по умолчанию Москва)")
	// указываем лимит и количество пропусков
	flag.StringVar(&limit, "limit", "100", "лимит заявок(по умолчанию 100)")
	flag.StringVar(&offset, "offset", "0", "сколько пропустить(по умолчанию 0)")
}

func main() {

	flag.Parse()

	cfg := configs.NewConfig()
	if _, err := toml.DecodeFile(configPath, cfg); err != nil {
		logrus.Fatal("Config fields doesn't read: ", err)
	}

	// подключение к базе
	db := mysql_store.NewStore()
	if err := db.Open(cfg.DSN); err != nil {
		logrus.Fatalf("%s : %s", "Can't open connect to database", err)
	}

	// конвертация в число limit flag
	limitint, err := strconv.Atoi(limit)
	if err != nil {
		logrus.Fatalf("%s : %s", "Can't convert limit in int", err)
	}

	// конвертация в число offset flag
	offsetint, err := strconv.Atoi(limit)
	if err != nil {
		logrus.Fatalf("%s : %s", "Can't convert offset in int", err)
	}

	// получение лидов
	leads, err := db.GetPromoRep().GetLead(data_start, data_end, city, limitint, offsetint)
	if err != nil {
		logrus.Fatalf("%s : %s", "Can't get lead from database", err)
	}

	// запись лидов в xlsx файл
	err = xls_loader.Loader_XLS(leads)
	if err != nil {
		logrus.Fatalf("%s : %s", "Problem with xls file", err)
	}
}
