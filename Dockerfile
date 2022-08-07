FROM node:16.4.2-alpine3.12
ENV NODE_ENV production
RUN mkdir /app && chown -R node:node /app
#to copy manifest files from inventory repo to current working directory
RUN mkdir /app/inventory && chown -R node:node /app/inventory
RUN mkdir /app/inventory/build && chown -R node:node app/inventory/build
RUN ls -l
RUN cd app
RUN ls -l 
WORKDIR /app

USER node
COPY Day3/*.js /app
COPY Day3/inventory/build/*.props /app/inventory/build


EXPOSE 8080
CMD ["node", "getalerts.js"]