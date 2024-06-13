#!/usr/bin/bash

ips=(125.71.177.165 111.67.197.25 103.252.4.139 177.250.25.13 88.218.93.92 68.183.48.199 111.231.16.76 223.100.28.112 154.61.76.78 112.192.20.23 164.215.103.234 185.217.1.246 183.81.169.238 122.224.37.86 183.213.92.170 1.192.188.70 36.110.228.254 36.140.104.38 60.255.228.65 117.80.237.70 27.116.53.234 112.192.20.22 47.116.11.110 182.72.90.18 206.189.149.40 183.239.27.18 42.5.104.44 14.29.206.196 179.43.180.108 104.248.153.120)

for element in ${ips[@]}
do
grep -a "$element" fakessh-filter-empty-password.log | awk '{if ( $6 == "123456" ) print $1}' > extracted/ch2/123456-$element.log
echo $element
done