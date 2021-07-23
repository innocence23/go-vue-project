.PHONY: build dev clean sqlstruct


default: build

build:
		go build -ldflags $(LD_FLAGS) -i -v -o ./bin/$(SERVICE) -race ./$(MAIN)
dev: build
		cp $(CUR_PWD)/conf/conf_dev.ini $(CUR_PWD)/conf/conf.ini && ./bin/$(SERVICE) -v=true
clean:

sqlstruct:
		gormt -H=127.0.0.1 --port=33060 -u=root --password=secret -d=mydb -l=json -F=false


