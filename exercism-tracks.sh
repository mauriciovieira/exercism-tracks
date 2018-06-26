alias exercism_build="docker-compose build"
alias exercism_start="docker-compose up -d"
alias exercism_stop="docker-compose stop"
alias exercism_c="docker-compose run --rm exercism-c"
alias exercism_go="docker-compose run --rm exercism-go"
alias exercism_go_test="docker-compose run --rm exercism-go go test ./... -v"
alias exercism_python="docker-compose run --rm exercism-python"
alias exercism_python_test="docker-compose run --rm exercism-python pytest python/**/*_test.py"
alias exercism_ruby="docker-compose run --rm exercism-ruby"
alias exercism_ruby_test="docker-compose run --rm exercism-ruby rake -f ruby/Rakefile"