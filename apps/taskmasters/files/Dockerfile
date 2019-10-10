# to build this image you must specify the Dockerfile independently of the build context, using -f.
# docker build -f apps/taskmasters/files/Dockerfile .
# this must be built from the project root directory 

# this is a public image for built task tools
FROM jeremiahz/task-tools:v0.4.3

RUN mkdir /configs
COPY deploy/configs/files.toml /configs/files.toml

CMD [ "sh","-c","files -config /configs/files.toml" ]