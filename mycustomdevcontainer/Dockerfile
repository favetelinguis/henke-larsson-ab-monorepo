# https://www.augmentedmind.de/2022/10/30/container-based-development-envs/

# Contents of Dockerfile
FROM some-language-runtime:major.minor.patch as base
 
FROM base as build
# here are statements that build your application
 
FROM base as vs-code-devcontainer
# Install Git to be able to use it right in the development container
RUN apt-get update && apt-get install git -y
# Optional: install further useful tools, e.g. pre-commit.org or others
 
FROM base as production-app
COPY --from=build /some/artifact.bin /somewhere/
ENTRYPOINT something