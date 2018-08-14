docker run -d --name=consul_c1 consul agent -client -node=c1 -join 127.0.0.1
docker run -idt --name=consul_s1 --net=host  -p 8300:8300 -p 8301:8301 -p 8301:8301/udp  -p 8302:8302/udp -p 8302:8302 -p 8400:8400 -p 8500:8500 -p 53:53/udp consul agent -server -bootstrap-expect=1 -node=s1 -bind=127.0.0.1
