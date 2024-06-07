# ReadMe



```shell
# 过滤空行
$ awk '$5 != "" && $6 != ""' origin.log > fake.log




awk '$6 ~ /^[0-9]+$/ {print $6}' fake.log  | sort | uniq -c | awk '{print $1, ",", $2}' | sort -r -n -k1 > analysis/pureNumber.csv
awk '{print $6}' fake.log | sort | uniq -c | awk '{print $1,",",$2}' | sort -r -n -k1 > analysis/passwordCount.csv
awk '{print $5}' fake.log | sort | uniq -c | awk '{print $1,",",$2}' | sort -r -n -k1 > analysis/usernameCount.csv
awk '{print substr($2, 1, 2)}' fake.log | sort | uniq -c | awk '{print $1,",",$2}' > analysis/timeFrep.csv






```
