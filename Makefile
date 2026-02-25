check-dir: src/check-dir.go
	go build -o check-dir src

.PHONY: release
release: check-dir
	mkdir ./phpv
	cp check-dir setup-env switch-version ./phpv
	zip -r ./phpv.zip ./phpv
	rm -rf ./phpv

.PHONY: install
install: check-dir
	mkdir -p ~/.local/opt/phpv
	cp -f check-dir setup-env switch-version ~/.local/opt/phpv/
