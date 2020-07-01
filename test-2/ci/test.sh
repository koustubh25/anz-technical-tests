#!/bin/ash

apk update
apk add curl
wget https://github.com/josephburnett/jd/releases/download/v1.1/jd && chmod a+x jd
curl http://anz-technical-test-2:8000/version -o appVersion.json

# Find diff between the two versions
diff=$(./jd --set appVersion.json version.json)

[ $? -eq 0 ] && [[ -z $diff ]] && echo 'Tests passed' || (echo 'Tests failed' && exit 1)