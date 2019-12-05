# specific node version
FROM node:13.0.1


# specify directory
RUN mkdir -p /user/src/app
WORKDIR /user/src/app
COPY ./package*.json ./

# development env
ENV NODE_ENV="development"

COPY . ./

CMD ["npm","install","--quiet"]
# RUN npm audit fix --quiet

CMD ["npm", "start"]

EXPOSE 3000