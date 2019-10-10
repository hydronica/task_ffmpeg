# to build this image you must specify the Dockerfile independently of the build context, using -f.
# docker build -f apps/taskmasters/filewatcher/Dockerfile .
# this must be built from the project root directory 

# this is a public image for built task tools
FROM jeremiahz/task-tools:v0.4.3

RUN mkdir /configs
COPY deploy/configs/filewatcher.toml /configs/filewatcher.toml

CMD [ "sh","-c","filewatcher -config /configs/filewatcher.toml" ]