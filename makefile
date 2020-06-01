## https://opensource.com/article/18/8/what-how-makefile, https://github.com/infracloudio/botkube/blob/develop/Makefile

.PHONY: all say_hello generate clean
#.DEFAULT_GOAL := generate
all: say_hello generate

say_hello: clean   # to call clean 1st than echo hello world
        @echo "Hello World"

generate:
        @echo "Creating empty text files..."
        touch x.txt

clean:   # by default make will not call clean even if it's there in .PHONY
        @echo "Cleaning up..."
        rm *.txt
