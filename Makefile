# Based on https://medium.com/@olebedev/live-code-reloading-for-golang-web-projects-in-19-lines-8b2e8777b1ea
PID = /tmp/risuto.pid

live-reload: restart
	@fswatch -e '.*/tmp/.*' -e '.*/public/.*' -o . | xargs -n1 -I{}  make restart || make kill

kill:
	@kill `cat $(PID)` || true

before:
	# @echo "Generate assets"
	# @go generate
	@echo "Build binary"
	@go build -o /tmp/risuto risuto.go

restart: kill before
	 @/tmp/risuto server -b localhost -p 5000 & echo $$! > $(PID)

.PHONY: serve restart kill before # let's go to reserve rules names

# fswatch --batch-marker=EOF -xn . | while read file event; do
#    echo $file $event
#    if [ $file = "EOF" ]; then
#       echo TRIGGER
#    fi
# done
