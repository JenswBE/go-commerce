# This script is intended to be sourced like "source lint.sh" or ". lint.sh".
# Config is available at https://raw.githubusercontent.com/JenswBE/setup/main/programming_configs/golang/.golangci.yml
golangci-lint run \
    -c ../setup/programming_configs/golang/.golangci.yml \
    --disable errorlint,goerr113,nestif,noctx,wrapcheck \
    --exclude 'name will be used as .+ by other packages' \
    --exclude 'reassigning variable ErrorStackMarshaler in other package zerolog'
