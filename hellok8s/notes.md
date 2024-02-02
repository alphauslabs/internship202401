docker build --rm -t hellok8s .
docker tag hellok8s asia.gcr.io/mobingi-main/hellok8s:v1
docker push asia.gcr.io/mobingi-main/hellok8s:v1
