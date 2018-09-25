all: airports.go airlines.go README.md

%.dat:
	wget https://raw.githubusercontent.com/jpatokal/openflights/master/data/$@

%.csv: %.dat
	cat $< | sed 's/\\N//g' | sed 's/\\\\//g' | sed 's/\\"/""/g' > $@

%.go: %.schema.yaml %.csv
	databundler -pkg openflights -schema $*.schema.yaml -data $*.csv -output $@
	gofmt -w $@

%.doc:
	godoc -src github.com/mmcloughlin/openflights $* | awk '$$1 != "use" && length($$0) > 0' > $@

README.md: README.md.j2 Airport.doc Airline.doc
	j2 $< > $@

deps:
	go get github.com/mmcloughlin/databundler

clean:
	$(RM) *.dat *.csv

.PHONY: all clean deps

.PRECIOUS: %.dat
