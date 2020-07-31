# Do not run `docker build` manually.
# Please try `make docker` instead.

FROM 476415093293.dkr.ecr.us-west-2.amazonaws.com/goguardian/vault_base

ADD zoneinfo.zip /
ENV ZONEINFO /zoneinfo.zip

ADD go-grpc /

CMD ["/go-grpc"]
