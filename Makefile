
SHELL = /usr/bin/env bash

ifeq ($(PROJECT),)
PROJECT = escher
endif

ifeq ($(DOMAIN),)
DOMAIN = benningen.xoreaxeax.de
endif

ifeq ($(SERVER),)
SERVER = root@zfs007.$(DOMAIN)
endif

EXECUTABLE = tmp/$(PROJECT)

all: build install

BUILD = CGO_ENABLED=0 GOOS=freebsd go build -a -installsuffix cgo -o "../$(@)" .

build: $(EXECUTABLE)

$(EXECUTABLE):
	mkdir -pv tmp
	cd src && go get && go fmt && $(BUILD)
	file "$(EXECUTABLE)"
	ls -ali "$(EXECUTABLE)"

INSTALL = chown root:wheel /$(EXECUTABLE) && \
					chmod 0700 /$(EXECUTABLE) && \
					file /$(EXECUTABLE) && \
					ls -ali /$(EXECUTABLE) && \
					mkdir -pv /space/escher/trunk

deploy install: build
	rsync "$(EXECUTABLE)" "$(SERVER):/$(EXECUTABLE)"
	ssh "$(SERVER)" "$(INSTALL)"

run: install
	ssh $(SERVER) -t screen '/$(EXECUTABLE)'

debug: install
	ssh $(SERVER) -t DEBUG=1 screen '/$(EXECUTABLE)'

clean:
	rm -rvf tmp

distclean: clean
	ssh $(SERVER) rm -vf "/$(EXECUTABLE)"
