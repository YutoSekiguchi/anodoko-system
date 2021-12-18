FROM --platform=linux/x86_64 node:15.8

WORKDIR /app
ENV NUXT_HOST 0.0.0.0

COPY package*.json ./
RUN npm install


EXPOSE 3000

CMD [ "npm", "run", "dev" ]