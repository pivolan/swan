.PHONY: update update_stopwords

update: update_stopwords

update_stopwords: goose3
	go build
	./update
	go fmt ../...

clean:
	rm -rf goose3
	rm -rf goose3.tar.gz
	go clean

goose3.tar.gz:
	curl -L https://api.github.com/repos/goose3/goose3/tarball > goose3.tar.gz

goose3: goose3.tar.gz
	mkdir goose3
	tar xvf $< --strip-components 1 --directory $@