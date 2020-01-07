#!/bin/bash

go build > /dev/null 2>&1
./hkp &

cd web
yarn build > /dev/null 2>&1
yarn start &

cd -

echo "serving backend on port 8080"
echo "serving frontend on port 3000"