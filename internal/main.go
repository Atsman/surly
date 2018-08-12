package main

func main() {
	config := InitConfig()
	db := InitDB()
	linkRepository := LinkRepository{
		config: config,
		db:     db,
	}
	linkService := LinkService{
		config:         config,
		linkRepository: linkRepository,
	}
	linkCtrl := LinkCtrl{linkService: linkService}
	InitHttp(config, linkCtrl)
}
