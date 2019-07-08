#!/bin/sh
if [ -z $1 ]
then
  echo Set voucher code when you bought something at Cafe
  exit
fi

curl -X POST 'http://10.0.2.1:8002/index.php?zone=cpzone' \
  -F auth_user= \
  -F auth_pass= \
  -F auth_voucher=$1 \
  -F redirurl=http://www.apple.com/library/test/success.html \
  -F accept=Accept
