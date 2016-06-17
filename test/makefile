template_all: build_all template.create template.show template.start_release template.schedule_release template.list template.delete

template_create_all: build_all template.create

template_delete_all: build_all template.delete

template.start_release:
	goxlr template start -n 'test template' -t 'test release'

template.schedule_release:
		goxlr template start -n 'test template' -t 'test release scheduled'

template.delete:
	goxlr template delete -t 'test template'

template.create: tst.json
	goxlr template create -i ./tst.json

template.show:
	goxlr template show  --byTitle 'test template'

template.list:
	goxlr template list

build_all: xlr.build goxlr.build

goxlr.build:
	go build ../
	go install ../

xlr.build:
	go build ../../xlr
	go install ../../xlr
