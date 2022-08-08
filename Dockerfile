FROM node:16.4.2-alpine3.12
ENV NODE_ENV production
# RUN pwd
# RUN mkdir /app && chown -R node:node /app
# #to copy manifest files from inventory repo to current working directory
# RUN mkdir /app/inventory && chown -R node:node /app/inventory
# RUN mkdir /app/inventory/build && chown -R node:node app/inventory/build
RUN ls -l
# RUN cd /home/runner/work && ls -l
# WORKDIR /app
RUN ls -l
USER node
# COPY *.js /app
# COPY inventory/build/*.props /app/inventory/build


EXPOSE 8080
CMD ["node", "getalerts.js"]