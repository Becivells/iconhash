FROM scratch
COPY iconhash /usr/bin/iconhash
ENTRYPOINT ["/usr/bin/iconhash"]