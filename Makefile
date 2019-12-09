all:
	make dependencies
	make core

dependencies:
	go install

core:
	go build -o joueur

clean:
	rm -rf joueur
