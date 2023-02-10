FROM envoyproxy/envoy:v1.18.4
COPY ./conf/envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml