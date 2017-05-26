set -e
echo "" > coverage.txt
for d in $(go list ./... | grep -v /vendor/); do
    go test -v -coverprofile=coverage.out -covermode=atomic $d
    if [ -f coverage.out ]; then
        cat coverage.out >> coverage.txt
        rm coverage.out
    fi
done
tail -n +2 "coverage.txt" > "coverage.txt.tmp" && mv "coverage.txt.tmp" "coverage.txt"
