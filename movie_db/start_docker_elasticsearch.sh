docker run -d --name es01 --net elastic -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -it -t --rm -v $PWD/elastic_data:/usr/share/elasticsearch/data elasticsearch:7.9.3
