#!/bin/bash

URL="https://google.com"

echo "Make sure app is running on localhost:8000"

echo "Shortening URL : $URL"

curl --location --request POST 'http://localhost:8000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "'$URL'"
}'

echo "Open the url https://localhost:8000/CODE in the browser to test the redirect, where CODE is given in the above output"
