PHONY: gen

gen:
  docker run -v $(pwd):/defs namely/protoc-all -f ./data/messages.proto -l go