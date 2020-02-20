#!/bin/sh

echo "\n>>> curl test, POST request"
curl -vk -XPOST \
    -H "Content-type: application/json" \
    -d '{"sender":"MyApp","title":"POST test subject","message":"test body","ishtml":true}' \
    'http://localhost:58725/submit'

echo "\n>>> curl test, GET request"
curl -k 'http://localhost:58725/submit?sender=MyApp&title=GET%20test%20subject&message=test%20body'
